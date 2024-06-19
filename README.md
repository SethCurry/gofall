# gofall

This project is not affiliated with Scryfall in any way.

## Examples

### Cards

Find a card by name:

```go
client := gofall.NewClient(nil)

card, err := client.Card.Named("Lightning Bolt")
if err != nil {
    panic(err)
}

fmt.Println(card.ID)
```

Search for cards:

```go
client := gofall.NewClient(nil)

cardPager, err := client.Card.Search(context.Background(), "Black Lotus", gofall.CardSearchOptions{})
if err != nil {
    panic(err)
}

for cardPager.HasMore() {
    cards, err := cardPager.NextPage()
    if err != nil {
        panic(err)
    }

    for _, card := range cards {
        fmt.Println(card.Name)
    }
}
```

Try to autocomplete a card name:

```go
client := gofall.NewClient(nil)

cardNames, err := client.Card.Autocomplete("Black Lot")
if err != nil {
    panic(err)
}

for _, cardName := range cardNames {
    fmt.Println(cardName)
}
```
