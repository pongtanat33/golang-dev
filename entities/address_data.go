package entities

type AddressDataUsecase interface {
	CountryMasterData() ([]CountryRes, error)
	ProvinceMasterData(*ProvinceReq) ([]ProvinceRes, error)
}

type AddressDataRepository interface {
	CountryMasterData() ([]CountryRes, error)
	ProvinceMasterData(*ProvinceReq) ([]ProvinceRes, error)
}

type CountryRes struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Iso2 string `json:"iso2" db:"iso2"`
}

type ProvinceReq struct {
	PostCode string `json:"postcode" db:"postcode"`
}

type ProvinceRes struct {
	ProvinceId      string `json:"province_id" db:"province_id"`
	ProvinceName    string `json:"province_name" db:"province_name"`
	DistrictId      string `json:"district_id" db:"district_id"`
	DistrictName    string `json:"district_name" db:"district_name"`
	SubDistrcitId   string `json:"subdistrict_id" db:"subdistrict_id"`
	SubDistrictName string `json:"subdistrict_name" db:"subdistrict_name"`
}