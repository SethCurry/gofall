package gofall

import (
	"context"
	"fmt"
	"net/http"
)

// RulingClient contains methods for interacting with rulings.
type RulingClient struct {
	client *http.Client
}

// ByScryfallID fetches all of the rulings for a card by its Scryfall ID.
func (r *RulingClient) ByScryfallID(ctx context.Context, id string) ([]Ruling, error) {
	// https://scryfall.com/docs/api/rulings/id
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/"+id+"/rulings", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var lst listContainer[Ruling]

	err = doRequest(r.client, req, &lst)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return lst.Data, nil
}

// ByMultiverseID fetches all of the rulings for a card by its Multiverse ID.
func (r *RulingClient) ByMultiverseID(ctx context.Context, id int) ([]Ruling, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/multiverse/"+fmt.Sprint(id)+"/rulings", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var list listContainer[Ruling]

	err = doRequest(r.client, req, &list)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return list.Data, nil
}

// ByMTGOID fetches all of the rulings for a card by its MTGO ID.
func (r *RulingClient) ByMTGOID(ctx context.Context, id int) ([]Ruling, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/mtgo/"+fmt.Sprint(id)+"/rulings", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var list listContainer[Ruling]

	err = doRequest(r.client, req, &list)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return list.Data, nil
}

// ByArenaID fetches all of the rulings for a card by its Arena ID.
func (r *RulingClient) ByArenaID(ctx context.Context, id int) ([]Ruling, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/arena/"+fmt.Sprint(id)+"/rulings", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var list listContainer[Ruling]

	err = doRequest(r.client, req, &list)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return list.Data, nil
}

// ByCodeAndNumber fetches all of the rulings for a card by its set code and collector number.
func (r *RulingClient) ByCodeAndNumber(ctx context.Context, code string, number string) ([]Ruling, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.scryfall.com/cards/"+code+"/"+number+"/rulings", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}

	var list listContainer[Ruling]

	err = doRequest(r.client, req, &list)
	if err != nil {
		return nil, fmt.Errorf("failed to perform HTTP request: %w", err)
	}

	return list.Data, nil
}
