package db

type Region struct {
	Model
	Name     string     `json:"name"`
	Slug     string     `json:"slug"`
	Articles []*Article `json:"articles"`
}
