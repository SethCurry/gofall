package gofall

type Card struct {
	// Scryfall's ID unique ID for the card.
	ID string `json:"id"`

	// The name of the card.  For dual-faced or split cards, it will look like:
	// "$first_side_name // $second_side_name".  E.g. "Fire // Ice"
	Name string `json:"name"`

	// The Oracle ID for the card, from Wizards.
	OracleID string `json:"oracle_id"`

	// The rarity of the card as a string, such as "Common", "Uncommon", "Rare", etc.
	Rarity string `json:"rarity"`

	// The flavor text as printed on the card.
	FlavorText string `json:"flavor_text"`

	// The ID of the card back design.
	CardBackID string `json:"card_back_id"`

	// The name of the artist who created the card's illustration.
	Artist string `json:"artist"`

	// The ID of the illustration that was created for the card.
	IllustrationID string `json:"illustration_id"`

	// The border color of the card as a string.  Un-sets will have silver borders.
	BorderColor string `json:"border_color"`

	// Frame is the edition of the card frame used for the printing.
	Frame string `json:"frame"`

	// The language of the card printing.  This will influence fields like the
	// flavor text, type line, and oracle text.
	Language string `json:"lang"`

	// SetID is Scryfall's UUID for the set.
	SetID string `json:"set_id"`

	// SetCode is the 3 character code for the set.
	SetCode string `json:"set"`

	// SetName is the full name of the set.
	SetName string `json:"set_name"`

	// SetType is the type of set this printing is in.
	SetType string `json:"set_type"`

	// SetURI is a link to the card's set object on Scryfall.
	SetURI          string `json:"set_uri"`
	SetSearchURI    string `json:"set_search_uri"`
	ScryfallSetURI  string `json:"scryfall_set_uri"`
	RulingsURI      string `json:"rulings_uri"`
	PrintsSearchURI string `json:"prints_search_uri"`
	CollectorNumber string `json:"collector_number"`
	URI             string `json:"uri"`
	ScryfallURI     string `json:"scryfall_uri"`
	Layout          string `json:"layout"`
	ManaCost        string `json:"mana_cost"`
	TypeLine        string `json:"type_line"`
	OracleText      string `json:"oracle_text"`
	Power           string `json:"power"`
	Toughness       string `json:"toughness"`
	Loyalty         string `json:"loyalty"`

	HighResImage   bool `json:"highres_image"`
	HighResScan    bool `json:"highres_scan"`
	Reserved       bool `json:"reserved"`
	Foil           bool `json:"foil"`
	NonFoil        bool `json:"nonfoil"`
	Oversized      bool `json:"oversized"`
	Promo          bool `json:"promo"`
	Reprint        bool `json:"reprint"`
	Variation      bool `json:"variation"`
	Digital        bool `json:"digital"`
	FullArt        bool `json:"full_art"`
	Textless       bool `json:"textless"`
	Booster        bool `json:"booster"`
	StorySpotlight bool `json:"story_spotlight"`

	// The converted mana cost of the card.  This is only a float because of
	// Un-set cards that have silly mana costs including half a mana.
	CMC float32 `json:"cmc"`

	// The MTG: Online ID for the card.
	MTGOID int `json:"mtgo_id"`

	// The MTG: Online ID for the card's foil variant.
	MTGOFoilID int `json:"mtgo_foil_id"`

	// The TCGPlayer ID for the card.
	TCGPlayerID int `json:"tcgplayer_id"`

	// The CardMarket ID for the card.
	CardmarketID int `json:"cardmarket_id"`

	// The EDHREC rank of the card.
	EDHRecRank int `json:"edhrec_rank"`

	// The PennyRank of the card.
	PennyRank int `json:"penny_rank"`

	// The multiverse IDs for the card.
	MultiverseIDs []int        `json:"multiverse_ids"`
	Colors        []string     `json:"colors"`
	ColorIdentity []string     `json:"color_identity"`
	Keywords      []string     `json:"keywords"`
	Games         []string     `json:"games"`
	Finishes      []string     `json:"finishes"`
	ArtistIDs     []string     `json:"artist_ids"`
	ReleasedAt    Date         `json:"released_at"`
	ImageURIs     ImageURIs    `json:"image_uris"`
	RelatedURIs   RelatedURIs  `json:"related_uris"`
	Prices        Prices       `json:"prices"`
	Legality      CardLegality `json:"legalities"`
	AllParts      []Part       `json:"all_parts"`
}

type Part struct {
	Object    string    `json:"object"`
	ID        string    `json:"id"`
	Component Component `json:"component"`
	Name      string    `json:"name"`
	TypeLine  string    `json:"type_line"`
	URI       string    `json:"uri"`
}

type Prices struct {
	USD       string  `json:"usd"`
	USDFoil   *string `json:"usd_foil"`
	USDEtched *string `json:"usd_etched"`
	EUR       string  `json:"eur"`
	EURFoil   *string `json:"eur_foil"`
	Tix       string  `json:"tix"`
}

type RelatedURIs struct {
	Gatherer                  string `json:"gatherer"`
	TCGPlayerInfiniteArticles string `json:"tcgplayer_infinite_articles"`
	TCGPlayerInfiniteDecks    string `json:"tcgplayer_infinite_decks"`
	EDHRec                    string `json:"edhrec"`
}
