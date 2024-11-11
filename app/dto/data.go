package dto

type Data struct {
	Data []Card `json:"data"`
}

type Card struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	Description string    `json:"desc"`
	Race        string    `json:"race"`
	Archetype   string    `json:"archetype"`
	Attack      int       `json:"atk"`
	Defense     int       `json:"def"`
	Level       int       `json:"level"`
	Attribute   string    `json:"attribute"`
	CardSets    []CardSet `json:"card_sets"`
	CardImages  []Image   `json:"card_images"`
}

type CardSet struct {
	SetName       string `json:"set_name"`
	SetRarity     string `json:"set_rarity"`
	SetRarityCode string `json:"set_rarity_code"`
}

type Image struct {
	ImageURL string `json:"image_url"`
}

type CardImage struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}
