package domain

type Card struct {
	Number  string  `json:"Number"`
	Name    string  `json:"Name"`
	Type    string  `json:"Type"`
	Level   string  `json:"Level"`
	Country string  `json:"Country"`
}
