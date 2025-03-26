package gofall

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
