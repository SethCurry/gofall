package gofall

// CardLegality stores the legality of a card in various formats.
type CardLegality struct {
	Standard        Legality `json:"standard"`
	Future          Legality `json:"future"`
	Historic        Legality `json:"historic"`
	Gladiator       Legality `json:"gladiator"`
	Pioneer         Legality `json:"pioneer"`
	Explorer        Legality `json:"explorer"`
	Modern          Legality `json:"modern"`
	Legacy          Legality `json:"legacy"`
	Pauper          Legality `json:"pauper"`
	Vintage         Legality `json:"vintage"`
	Penny           Legality `json:"penny"`
	Commander       Legality `json:"commander"`
	Oathbreaker     Legality `json:"oathbreaker"`
	Brawl           Legality `json:"brawl"`
	HistoricBrawl   Legality `json:"historicbrawl"`
	Alchemy         Legality `json:"alchemy"`
	PauperCommander Legality `json:"paupercommander"`
	Duel            Legality `json:"duel"`
	OldSchool       Legality `json:"oldschool"`
	PreModern       Legality `json:"premodern"`
	PrEDH           Legality `json:"predh"`
}
