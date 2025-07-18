package entities

type AdminAuthenticationRes struct {
	IsValid bool `db:"is_valid"`
	ResultMessage    *string    `db:"result_message"`

}