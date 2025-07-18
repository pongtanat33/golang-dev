package entities

type UnitsKey struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

type UnitsListReq struct {
	IsCalculated bool   `json:"is_calculated" db:"is_calculated"`
	Lang         string `db:"lang"`
}

type Unit struct {
	Id   string `json:"id" `
	Name string `json:"name" `
}
