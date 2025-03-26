package gofall

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// CardClient contains methods for querying Scryfall for cards by
// a search query, name, etc.
type CardClient struct {
	client *http.Client
}

// CardNamedRequest contains the parameters for a named card search.
// You must provide either Exact or Fuzzy, but not both.
type CardNamedRequest struct {
	// Exact performs an exact match search for the provided card name.
	Exact string

	// Fuzzy performs a fuzzy match search for the provided card name.
	Fuzzy string

	// Set is the set code of the card to search for.
	// Optional.
	Set *string

	// Face is the face of a double-faced card to search for.
	// Optional.
	Face *string

	// Version is the image version to return.
	// Optional.
	Version *ImageType
}

// validate ensures that either Exact or Fuzzy is provided, but not both.
func (c CardNamedRequest) validate() error {
	if c.Exact != "" && c.Fuzzy != "" {
		return errors.New("cannot provide both exact and fuzzy search parameters")
	}

	if c.Exact == "" && c.Fuzzy == "" {
		return errors.New("must provide either exact or fuzzy search parameters")
	}

	return nil
}

// addToQuery adds the parameters to the provided URL query.
// Unset parameters are not added to the query, and will use
// Scryfall's default behavior.
func (c CardNamedRequest) addToQuery(query url.Values) {
	if c.Exact != "" {
		query.Add("exact", c.Exact)
	}

	if c.Fuzzy != "" {
		query.Add("fuzzy", c.Fuzzy)
	}

	if c.Set != nil {
		query.Add("set", *c.Set)
	}

	if c.Face != nil {
		query.Add("face", *c.Face)
	}

	if c.Version != nil {
		query.Add("version", string(*c.Version))
	}
}

// Named searches for a card by its name.
// Returns an error if more than one card is found.
func (c *CardClient) Named(ctx context.Context, options CardNamedRequest) (*Card, error) {
	if err := options.validate(); err != nil {
		return nil, err
	}

	var card Card

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/named", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	query := req.URL.Query()
	options.addToQuery(query)
	req.URL.RawQuery = query.Encode()

	err = doRequest(c.client, req, &card)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return &card, nil
}

type listContainer[T any] struct {
	Object     Object   `json:"object"`
	Data       []T      `json:"data"`
	HasMore    bool     `json:"has_more"`
	NextPage   string   `json:"next_page"`
	TotalCards int      `json:"total_cards"`
	Warnings   []string `json:"warnings"`
}

// CardSearchPager allows reading through several pages of card search results.
type CardSearchPager struct {
	client   *http.Client
	nextPage string
	done     bool
}

// HasMore returns true if there are more pages of results left, or false
// if it has already read all the pages.
func (c *CardSearchPager) HasMore() bool {
	return !c.done
}

// Next reads the next page of results from the search and returns the cards on it.
// This does require an HTTP request, so it will incur latency.  This method is not safe
// to call in a goroutine.
func (c *CardSearchPager) Next(ctx context.Context) ([]Card, error) {
	if c.done {
		return nil, io.EOF
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.nextPage, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var lst listContainer[Card]

	err = doRequest(c.client, req, &lst)
	if err != nil {
		apiErr := APIError{}
		if errors.As(err, &apiErr) {
			if apiErr.Status == 422 {
				return nil, ErrNoBackFace
			}
		}
	}

	c.nextPage = lst.NextPage
	c.done = !lst.HasMore

	return lst.Data, nil
}

// UniqueMode defines how card results are made unique, eg by card, art, or print.
type UniqueMode string

const (
	// UniqueCards returns a single result per card name.
	UniqueCards UniqueMode = "cards"

	// UniqueArt returns a single result per card art.
	UniqueArt UniqueMode = "art"

	// UniquePrint returns a single result per card print,
	// effectively disabling the unique feature.
	UniquePrint UniqueMode = "print"
)

// OrderDirection defines the direction of the sorting, eg ascending or descending.
type OrderDirection string

const (
	// OrderDirectionAuto sorts the cards in the default direction for the given order.
	OrderDirectionAuto OrderDirection = "auto"

	// OrderDirectionAscending sorts the cards in ascending order.
	OrderDirectionAscending OrderDirection = "asc"

	// OrderDirectionDescending sorts the cards in descending order.
	OrderDirectionDescending OrderDirection = "desc"
)

// CardSearchOptions provide optional parameters for searching for cards.
type CardSearchOptions struct {
	Unique            *UniqueMode
	Order             *Order
	Direction         *OrderDirection
	IncludeExtras     bool
	IncludeVariations bool
}

func (c CardSearchOptions) addToQuery(query url.Values) {
	if c.Unique != nil {
		query.Add("unique", string(*c.Unique))
	}

	if c.Order != nil {
		query.Add("order", string(*c.Order))
	}

	if c.Direction != nil {
		query.Add("dir", string(*c.Direction))
	}
}

// Search queries Scryfall for cards matching the provided query.
// See https://scryfall.com/docs/syntax for more information on the query syntax.
func (c *CardClient) Search(ctx context.Context, query string, opts CardSearchOptions) (*CardSearchPager, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/search", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	q := req.URL.Query()
	q.Add("q", query)
	opts.addToQuery(q)

	req.URL.RawQuery = q.Encode()

	pager := &CardSearchPager{
		client:   c.client,
		nextPage: req.URL.String(),
	}

	return pager, nil
}

type autocompleteResponse struct {
	Data []string `json:"data"`
}

// Autocomplete returns a list of cards that start with the provided string.
// Useful for providing autocomplete suggestions to users.
func (c *CardClient) Autocomplete(ctx context.Context, query string) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/autocomplete", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	q := req.URL.Query()
	q.Add("q", query)
	req.URL.RawQuery = q.Encode()

	var autocomplete autocompleteResponse

	err = doRequest(c.client, req, &autocomplete)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return autocomplete.Data, nil
}

type RandomCardOptions struct {
	Face    string
	Version *ImageType
}

func (r RandomCardOptions) addToQuery(query url.Values) {
	if r.Face != "" {
		query.Add("face", r.Face)
	}

	if r.Version != nil {
		query.Add("version", string(*r.Version))
	}
}

// Random returns a random card from the provided query.
func (c *CardClient) Random(ctx context.Context, query string, opts RandomCardOptions) (*Card, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/random", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	q := req.URL.Query()
	opts.addToQuery(q)
	q.Add("q", query)

	req.URL.RawQuery = q.Encode()

	var card Card

	err = doRequest(c.client, req, &card)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return &card, nil
}

// CardIdentifier allows providing one of several identifiers for a card.
// At least one identifier is required to function.
type CardIdentifier struct {
	ID              string `json:"id"`
	MTGOID          int    `json:"mtgo_id"`
	MultiverseID    int    `json:"multiverse_id"`
	OracleID        string `json:"oracle_id"`
	IllustrationID  string `json:"illustration_id"`
	Name            string `json:"name"`
	Set             string `json:"set"`
	CollectorNumber string `json:"collector_number"`
}

type collectionRequest struct {
	Identifiers []CardIdentifier `json:"identifiers"`
}

func (c *CardClient) Collection(ctx context.Context, identifiers []CardIdentifier) ([]Card, error) {
	marshalled, err := json.Marshal(collectionRequest{Identifiers: identifiers})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal identifiers: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/random", bytes.NewBuffer(marshalled))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var list listContainer[Card]

	err = doRequest(c.client, req, &list)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return list.Data, nil
}
