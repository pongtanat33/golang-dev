package usecases

import ()

type addressDataUse struct {
	AddressDataRepo entities.AddressDataRepository
}

func NewAddressDataUsecase(addressDataRepo entities.AddressDataRepository) entities.AddressDataUsecase {
	return &addressDataUse{
		AddressDataRepo: addressDataRepo,
	}
}

func (a *addressDataUse) CountryMasterData() ([]entities.CountryRes, error) {
	countrys, err := a.AddressDataRepo.CountryMasterData()
	if err != nil {
		return nil, err
	}
	return countrys, nil
}

func (a *addressDataUse) ProvinceMasterData(req *entities.ProvinceReq) ([]entities.ProvinceRes, error) {
	provinces, err := a.AddressDataRepo.ProvinceMasterData(req)
	if err != nil {
		return nil, err
	}
	return provinces, nil
}
