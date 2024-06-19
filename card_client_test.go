package gofall_test

import (
	"context"
	"testing"

	"github.com/SethCurry/gofall"
)

func Test_Client_Card_Named(t *testing.T) {
	t.Parallel()

	client := gofall.NewClient(nil)

	card, err := client.Card.Named(context.Background(), "Black Lotus")
	if err != nil {
		t.Fatalf("failed to query for Black Lotus by named: %v", err)
	}

	if card.ID != "bd8fa327-dd41-4737-8f19-2cf5eb1f7cdd" {
		t.Errorf("returned card has ID %q instead of expected value", card.ID)
	}

	if card.Name != "Black Lotus" {
		t.Errorf("returned card is named %q instead of \"Black Lotus\"", card.Name)
	}
}

func Test_Client_Card_Search(t *testing.T) {
	t.Parallel()

	client := gofall.NewClient(nil)

	cardPager, err := client.Card.Search(context.Background(), "Black Lotus", gofall.CardSearchOptions{})
	if err != nil {
		t.Fatalf("failed to search for Black Lotus: %v", err)
	}

	cards, err := cardPager.Next(context.Background())
	if err != nil {
		t.Fatalf("failed to retrieve next page of results: %v", err)
	}

	if len(cards) != 2 {
		t.Errorf("unexpected number of cards returned: %d", len(cards))
	}
}

func Test_Client_Card_Autocomplete(t *testing.T) {
	t.Parallel()

	client := gofall.NewClient(nil)

	autocomplete, err := client.Card.Autocomplete(context.Background(), "Black Lotu")
	if err != nil {
		t.Fatalf("failed to get autocomplete suggestions: %v", err)
	}

	if len(autocomplete) != 1 {
		t.Fatalf("expected one suggestion but got %d", len(autocomplete))
	}
}
