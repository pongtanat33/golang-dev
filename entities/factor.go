package entities

type FactorKey struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

type Factor struct {
	Factor string `json:"factor" `
}

type CreateFactorReq struct {
	Name string  `json:"name" db:"name"`
	Factor float64 `json:"factor" db:"factor"`
	UnitId string  `json:"unit_id" db:"unit_id"`
	User   string  `json:"user" db:"user"`
	Type   *string  `json:"type" db:"type"`
}

type FactorTypeRes struct {
	Id string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	GWP float64 `json:"gwp" db:"gwp"`
	GWP_Id string `json:"gwp_id" db:"gwp_id"`
	GWP_Name string `json:"gwp_name" db:"gwp_name"`
	GWP_Type string `json:"gwp_type" db:"gwp_type"`
	GWP_Source string `json:"gwp_source" db:"gwp_source"`

}