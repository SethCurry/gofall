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

type CardClient struct {
	client *http.Client
}

// ImageType defines the types of images such as small, normal, large, etc.
type ImageType string

const (
	// ImageTypeSmall is the smallest image size.
	ImageTypeSmall ImageType = "small"

	// ImageTypeNormal is the normal image size.
	ImageTypeNormal ImageType = "normal"

	// ImageTypeLarge is the largest image size.
	ImageTypeLarge ImageType = "large"

	// ImageTypePng is the PNG image type.
	ImageTypePng ImageType = "png"

	// ImageTypeArtCrop is the art crop image type.
	ImageTypeArtCrop ImageType = "art_crop"

	// ImageTypeBorderCrop is the border crop image type.
	ImageTypeBorderCrop ImageType = "border_crop"
)

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
func (c CardNamedRequest) addToQuery(q url.Values) {
	if c.Exact != "" {
		q.Add("exact", c.Exact)
	}

	if c.Fuzzy != "" {
		q.Add("fuzzy", c.Fuzzy)
	}

	if c.Set != nil {
		q.Add("set", *c.Set)
	}

	if c.Face != nil {
		q.Add("face", *c.Face)
	}

	if c.Version != nil {
		q.Add("version", string(*c.Version))
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

	q := req.URL.Query()
	options.addToQuery(q)
	req.URL.RawQuery = q.Encode()

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

type CardSearchPager struct {
	client   *http.Client
	nextPage string
	done     bool
}

func (c *CardSearchPager) HasMore() bool {
	return !c.done
}

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

// Order defines how the returned cards are sorted.
type Order string

const (
	// OrderName sorts cards by name, A → Z
	OrderName Order = "name"

	// OrderSet sorts cards by their set and collector number: AAA/#1 → ZZZ/#999
	OrderSet Order = "set"

	// OrderReleased sorts cards by their release date: Newest → Oldest
	OrderReleased Order = "released"

	// OrderRarity sorts cards by their rarity: Common → Mythic
	OrderRarity Order = "rarity"

	// OrderColor sort cards by their color and color identity:
	// WUBRG → multicolor → colorless
	OrderColor Order = "color"

	// OrderUSD sorts cards by their lowest known U.S. Dollar price:
	// 0.01 → highest, null last
	OrderUSD Order = "usd"

	// OrderTix sorts cards by their lowest known
	// TIX price: 0.01 → highest, null last
	OrderTix Order = "tix"

	// OrderEur sorts cards by their lowest known Euro price:
	// 0.01 → highest, null last
	OrderEur Order = "eur"

	// OrderCMC sorts cards by their mana value: 0 → highest
	OrderCMC Order = "cmc"

	// OrderPower sorts cards by their power: null → highest
	OrderPower Order = "power"

	// OrderToughness sorts cards by their toughness: null → highest
	OrderToughness Order = "toughness"

	// OrderEDHREC sorts cards by their EDHREC ranking: lowest → highest
	OrderEDHREC Order = "edhrec"

	// OrderPenny sorts cards by their Penny Dreadful ranking: lowest → highest
	OrderPenny Order = "penny"

	// OrderArtist sorts cards by their front-side artist name: A → Z
	OrderArtist Order = "artist"

	// OrderReview sorts cards how podcasts review sets,
	// usually color & CMC, lowest → highest, with Booster Fun cards at the end
	OrderReview Order = "review"
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

func (c *CardClient) Search(ctx context.Context, query string, opts CardSearchOptions) (*CardSearchPager, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/search", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	q := req.URL.Query()
	q.Add("q", query)
	if opts.Unique != nil {
		q.Add("unique", string(*opts.Unique))
	}
	if opts.Order != nil {
		q.Add("order", string(*opts.Order))
	}
	if opts.Direction != nil {
		q.Add("dir", string(*opts.Direction))
	}

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
	Query   string
	Face    string
	Version *ImageType
	Pretty  bool
}

func (c *CardClient) Random(ctx context.Context, opts RandomCardOptions) (*Card, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/random", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	q := req.URL.Query()

	if opts.Query != "" {
		q.Add("q", opts.Query)
	}

	if opts.Face != "" {
		q.Add("face", opts.Face)
	}

	if opts.Version != nil {
		q.Add("version", string(*opts.Version))
	}

	if opts.Pretty {
		q.Add("pretty", "true")
	}

	req.URL.RawQuery = q.Encode()

	var card Card

	err = doRequest(c.client, req, &card)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return &card, nil
}

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
