package entities


type GWPKey struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}


type GWPRes struct {
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Source string `json:"source" db:"source"`
	GWP float64 `json:"gwp" db:"gwp"`
	Type string `json:"type"`
}
