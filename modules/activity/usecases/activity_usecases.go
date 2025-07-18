package usecases

type activityUse struct {
	ActivityRepo entities.ActivityRepository
}

func NewActivityUsecase(activityRepo entities.ActivityRepository) entities.ActivityUsecase {
	return &activityUse{
		ActivityRepo: activityRepo,
	}
}

func (a *activityUse) GetActivityAvailableSyncBySite(req *entities.ActivityListsReq) ([]entities.ActivityListRes, error) {
	return a.ActivityRepo.GetActivityAvailableSyncBySite(req)
}

func (a *activityUse) GetActivitySyncBySite(req *entities.ActivitySyncBySiteReq) ([]entities.ActivitySyncBySiteRes, error) {
	return a.ActivityRepo.GetActivitySyncBySite(req)
}

func (a *activityUse) CheckActivitySyncConnection(req *entities.ActivitySyncReq) (bool, error) {
	return a.ActivityRepo.CheckActivitySyncConnection(req)
}

func (a *activityUse) CreateActivitySync(req *entities.ActivitySyncReq) error {
	return a.ActivityRepo.CreateActivitySync(req)
}

func (a *activityUse) DeleteActivitySync(req *entities.ActivitySyncReq) error {
	return a.ActivityRepo.DeleteActivitySync(req)
}

func (a *activityUse) DeleteActivitySyncBySource(req *entities.ActivitySyncReq) error {
	return a.ActivityRepo.DeleteActivitySyncBySource(req)
}

func (a *activityUse) DuplicateActivity(req *entities.DuplicateActivityReq) error {
	activity, err := utils.GetActivityValue(req.ActivityId, "en")
	if err != nil {
		return err
	}

	if activity.ActivityName == req.ActivityName {
		return errors.New("activity_name_already_exist")
	}

	activityCreateReq := new(entities.ActivityCreateReq)
	activityCreateReq.Name = req.ActivityName
	activityCreateReq.Description = activity.Description
	activityCreateReq.SiteId = activity.SiteId
	activityCreateReq.EmissionFactorId = activity.EmissionFactorId
	activityCreateReq.User = req.User
	activityCreateReq.CompanyId = req.CompanyId
	activityCreateReq.FrequencyType = activity.ActivityFrequency
	activityCreateReq.FrequencySelectedDate = activity.FrequencyDate

	new_activity_id, err := a.ActivityRepo.CreateActivity(activityCreateReq)
	if err != nil {
		return err
	}

	activityFrequency := &entities.ActivityFrequency{
		ActivityId:   new_activity_id,
		Type:         *activity.ActivityFrequency,
		SelectedDate: activity.FrequencyDate,
		CreatedBy:    req.User,
	}
	_, errFrequency := utils.CreateActivityFrequency(activityFrequency)
	if errFrequency != nil {
		return nil
	}

	return nil
}

func (a *activityUse) UpdateActivityAutoScope2(req *entities.UpdateActivityAutoScope2Req) error {

	// CHECK IF ALREADY USED OR NOT
	row, errCheck := utils.CountBill(&req.BillingInfo)
	if errCheck != nil {
		return errCheck
	}

	if row != nil && *row > 0 {
		return errors.New("bill_already_used")
	}

	// update activity name
	err := a.ActivityRepo.UpdateActivityAutoScope2(req)
	if err != nil {
		return err
	}

	// Create Bill
	upsertBillReq := new(entities.UpsertBillingReq)
	upsertBillReq.ActivityId = req.Id
	upsertBillReq.CustomerNumber = req.BillingInfo.CustomerNumber
	upsertBillReq.PeaNumber = req.BillingInfo.PeaNumber
	upsertBillReq.SiteId = req.SiteId
	upsertBillReq.TaxId = req.BillingInfo.TaxId
	upsertBillReq.User = req.User

	_, errUpsertBill := utils.UpsertBill(upsertBillReq)
	if errUpsertBill != nil {
		return errUpsertBill
	}

	return nil

}

func (a *activityUse) CreateActivityAutoScope2(req *entities.CreateActivityAutoScope2Req) error {

	// create channel for go routine
	countBillChan := make(chan int, 1)
	if req.IsAuto {
		// CHECK IF ALREADY USED OR NOT
		go func() {
			row, errCheck := utils.CountBill(&req.BillingInfo)
			if errCheck != nil {
				logs.Error(errCheck)
			}
			countBillChan <- *row
		}()

		select {
		case count := <-countBillChan:
			if count > 0 {
				return errors.New("bill_already_used")
			}
		default:
			// Do nothing if no value is received from countBillChan
		}
	}

	// Create Activity & Activity Transaction
	peaEmissionID, err := utils.GetPeaEmissionId()
	if err != nil {
		return err
	}
	activityCreateReq := new(entities.ActivityCreateReq)
	activityCreateReq.EmissionFactorId = peaEmissionID
	activityCreateReq.FrequencyType = req.FrequencyType
	activityCreateReq.FrequencySelectedDate = req.FrequencySelectedDate
	activityCreateReq.Name = req.Name
	activityCreateReq.Description = req.ActivityDescription
	activityCreateReq.SiteId = req.SiteId
	activityCreateReq.CompanyId = req.CompanyId

	activityCreateReq.User = req.User

	activityId, err := a.CreateActivity(activityCreateReq)
	if err != nil {
		return err
	}

	if req.IsAuto {

		// Create Bill
		upsertBillReq := new(entities.UpsertBillingReq)
		upsertBillReq.ActivityId = activityId
		upsertBillReq.CustomerNumber = req.BillingInfo.CustomerNumber
		upsertBillReq.PeaNumber = req.BillingInfo.PeaNumber
		upsertBillReq.SiteId = req.SiteId
		upsertBillReq.TaxId = req.BillingInfo.TaxId
		upsertBillReq.User = req.User

		_, errUpsertBill := utils.UpsertBill(upsertBillReq)
		if errUpsertBill != nil {
			return errUpsertBill
		}

		getBillingTransactionReq := new(entities.GetBillingTransactionReq)
		getBillingTransactionReq.CustomerNumber = req.BillingInfo.CustomerNumber
		getBillingTransactionReq.PeaNumber = ""

		// Get the current date and time in the local timezone
		now := time.Now().UTC()

		// Get the year and month as integers
		year := now.Year()
		month := int(now.Month())

		// Format the year and month as "YYYYMM"
		period := year*100 + month

		getBillingTransactionReq.Period = period

		// result, err := utils.GetBillingTransaction(getBillingTransactionReq)
		// if err != nil {
		// 	return err
		// }

		// Calculate the period 36 months ago
		oneYearAgo := now.AddDate(-3, 0, 0)
		yearAgo := oneYearAgo.Year()
		monthAgo := int(oneYearAgo.Month())
		periodAgo := yearAgo*100 + monthAgo

		//combinedResults := result
		combinedResults := []entities.GetBillingTransactionRes{}

		for current := periodAgo; current <= period; {
			year := current / 100
			month := current % 100

			getBillingTransactionReq.Period = current
			resultAgo, err := utils.GetBillingTransaction(getBillingTransactionReq)
			if err != nil {
				return err
			}
			if len(resultAgo) > 0 {
				if resultAgo[0].Period == current {
					combinedResults = append(combinedResults, resultAgo[0])
				}
			}

			// ถ้าเดือนเป็น 1 ให้ย้อนกลับไปปีที่แล้ว และตั้งเดือนเป็น 12
			if month == 12 {
				year++
				month = 1
			} else {
				month++
			}
			//fmt.Println(current)
			// อัพเดต current ด้วยค่าปีและเดือนใหม่
			current = year*100 + month
		}

		for _, billingTransaction := range combinedResults {

			if billingTransaction.KwhTotal > 0 {
				year := billingTransaction.Period / 100
				month := billingTransaction.Period % 100

				// Get the last day of the month
				lastDay := getLastDayOfMonth(year, month)
				//periodString := fmt.Sprintf("%04d-%02d-01T00:00:00.000Z", year, month)
				periodString := fmt.Sprintf("%04d-%02d-%2dT00:00:00.000Z", year, month, lastDay)

				billingDate, err := time.Parse("2006-01-02T15:04:05.000Z", periodString)
				if err != nil {
					logs.Error("Error parsing date:" + err.Error())
					return err
				}

				createActivityTransactionReq := new(entities.CreateActivityTransactionReq)
				createActivityTransactionReq.ActionDate = billingDate.Format("2006-01-02 15:04:05") // Example format: "YYYY-MM-DD HH:MM:SS"
				createActivityTransactionReq.ActivityId = activityId
				createActivityTransactionReq.Amount = billingTransaction.KwhTotal
				createActivityTransactionReq.CompanyId = req.CompanyId
				createActivityTransactionReq.User = req.User
				createActivityTransactionReq.Evidence = billingTransaction.PeaNumber

				_, errCreateActivityTransaction := a.CreateActivityTransaction(createActivityTransactionReq)
				if errCreateActivityTransaction != nil {
					logs.Error(errCreateActivityTransaction)
					return errCreateActivityTransaction
				}
			}
		}
	}

	return nil

}

func (a *activityUse) CheckBillingInfo(req *entities.CustomerInfoReq) (*entities.CheckBillingInfoRes, error) {

	connectBillingInfoRes := new(entities.CheckBillingInfoRes)
	// CHECK IF ALREADY USED OR NOT

	row, err := utils.CountBill(req)
	if err != nil {
		return nil, err
	}

	if row != nil && *row > 0 {
		connectBillingInfoRes.BillNotUsed = true
	} else {
		connectBillingInfoRes.BillNotUsed = false
	}

	// CHECK IF CA VALID
	result, err := utils.CheckCarbonRegister(req)
	if err != nil {
		return nil, err
	}

	// true
	if result.Data.Data {
		connectBillingInfoRes.BillInfoValid = true
	} else {
		connectBillingInfoRes.BillInfoValid = false
	}

	return connectBillingInfoRes, nil
}

func (a *activityUse) GetActivityValueById(activity_id string, lang string) (*entities.ActivityValueRes, error) {
	activity, err := utils.GetActivityValue(activity_id, lang)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	return activity, nil
}

func (a *activityUse) GetActivityPeaEmissionValueById(activity_id string, lang string) (*entities.ActivityPeaEmissionValueRes, error) {

	activity, err := utils.GetActivityPeaEmissionValue(activity_id, lang)
	if err != nil {
		logs.Error(err)
		return nil, err
	}

	if activity.IsConnected != nil {
		activity.IsAuto = true
	} else {
		activity.IsAuto = false
	}

	return activity, nil
}

func (a *activityUse) ActivityTransaction(req *entities.ActivityTransactionReq) ([]entities.ActivityTransactionRes, int, error) {
	activityTransactions, err := a.ActivityRepo.ActivityTransaction(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}

	activityTransactionRowsCount, errRowCount := a.ActivityRepo.ActivityTransactionRowCount(req.Id)
	if errRowCount != nil {
		fmt.Println(errRowCount)
		return nil, 0, errRowCount
	}

	return activityTransactions, activityTransactionRowsCount, nil
}

func (a *activityUse) ActivityTransactionByDate(req *entities.ActivityTransactionByDateReq) ([]entities.ActivityTransactionRes, int, error) {
	activityTransactions, err := a.ActivityRepo.ActivityTransactionByDate(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}

	rowsCount, errRowCount := a.ActivityRepo.ActivityTransactionByDateRowCount(req.Id, req.Date)
	if errRowCount != nil {

		return nil, 0, errRowCount
	}

	return activityTransactions, rowsCount, err
}

func (a *activityUse) ActivityByScope(req *entities.ActivityByScopeReq) ([]entities.ActivityByScopeRes, int, error) {
	activities, err := a.ActivityRepo.ActivityByScope(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}

	return activities, len(activities), err
}

func (a *activityUse) ActivityByScopeEmission(req *entities.ActivityByScopeEmissionReq) ([]entities.ActivityByScopeRes, int, error) {
	activities, err := a.ActivityRepo.ActivityByScopeEmission(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}

	activityByScopeEmissionRowsCount, errRowCount := a.ActivityRepo.ActivityByScopeEmissionRowCount(req)
	if errRowCount != nil {
		return nil, 0, errRowCount
	}

	return activities, activityByScopeEmissionRowsCount, err
}

func (a *activityUse) ActivityLists(req *entities.ActivityListsReq) ([]entities.ActivityListRes, error) {
	return a.ActivityRepo.ActivityLists(req)
}

func (a *activityUse) CreateActivity(req *entities.ActivityCreateReq) (string, error) {
	activity_id, err := a.ActivityRepo.CreateActivity(req)
	if err != nil {
		return "", err
	}
	if req.FrequencyType == nil && req.FrequencySelectedDate == nil {
		logs.Error("create_activity_not_create_frequency_case")
	} else {
		activityFrequency := &entities.ActivityFrequency{
			ActivityId:   activity_id,
			Type:         *req.FrequencyType,
			SelectedDate: req.FrequencySelectedDate,
			CreatedBy:    req.User,
		}
		_, errFrequency := utils.CreateActivityFrequency(activityFrequency)
		if errFrequency != nil {
			return "", errFrequency
		}
	}
	return activity_id, err
}

func (a *activityUse) CreateActivityTransaction(req *entities.CreateActivityTransactionReq) (string, error) {
	id, err := a.ActivityRepo.CreateActivityTransaction(req)
	if err != nil {
		return "", err
	}

	// check activity source is synced
	syncList := a.ActivityRepo.GetActivitySyncBySourceId(req.ActivityId)

	if len(syncList) > 0 {
		for _, target := range syncList {
			if target.StartSyncAt.Before(time.Now().UTC().UTC()) {
				syncReq := entities.CreateActivityTransactionReq{
					ActivityId:             target.TargetActivityId,
					Amount:                 req.Amount * target.ConversionRate,
					ActionDate:             req.ActionDate,
					Evidence:               req.Evidence,
					Additional:             req.Additional,
					User:                   req.User,
					CompanyId:              req.CompanyId,
					ReferenceTransactionId: id,
				}
				a.ActivityRepo.CreateActivityTransaction(&syncReq)
			}
		}
	}

	return id, nil
}

func (a *activityUse) UpdateActivityTransaction(req *entities.UpdateActivityTransactionReq) error {
	_, err := a.ActivityRepo.UpdateActivityTransaction(req)

	//fmt.Println(err)

	if err != nil {
		return err
	}

	// check reference transaction
	referenceTransaction, err := a.ActivityRepo.GetReferenceTransactionById(req.Id)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		return nil
	}

	//fmt.Println(referenceTransaction)

	if referenceTransaction.ReferenceTransactionId == nil {
		return nil
	}

	// get activity sync by target
	syncList := a.ActivityRepo.GetActivitySyncByTargetId(referenceTransaction.ActivityId)
	//fmt.Println(syncList)

	// Find the item in the list
	for _, item := range syncList {
		// check if source is equal to target
		if item.SourceActivityId == req.ActivityId {
			// Do something with the found item

			syncReq := entities.UpdateActivityTransactionReq{
				Id:         referenceTransaction.Id,
				ActivityId: referenceTransaction.ActivityId,
				Amount:     req.Amount * item.ConversionRate,
				ActionDate: req.ActionDate,
				Evidence:   req.Evidence,
				Additional: req.Additional,
				User:       req.User,
				CompanyId:  req.CompanyId,
			}
			a.ActivityRepo.UpdateActivityTransaction(&syncReq)

			break
		}
	}

	return nil
}

func (a *activityUse) UpdateActivity(req *entities.UpdateActivityReq) error {
	err := a.ActivityRepo.UpdateActivity(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *activityUse) DeleteActivity(req *entities.DeleteActivityReq) error {
	// delete activity
	if err := a.ActivityRepo.DeleteActivity(req); err != nil {
		return err
	}

	// delete activity sync (not return error)
	if err := a.ActivityRepo.DeleteAllActivitySyncById(req); err != nil {
		return nil
	}

	return nil
}

func (a *activityUse) DeleteActivityTransaction(req *entities.DeleteActivityTransactionReq) error {
	err := a.ActivityRepo.DeleteActivityTransaction(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *activityUse) CreateActivityWithCustomizedEmission(req *entities.CreateActivityWithCustomizedEmissionReq) error {

	siteValue, err := utils.GetSiteValue(req.SiteId)
	if err != nil {
		return err
	}
	emission := new(entities.CreateEmissionReq)
	emission.CountryId = siteValue.CountryID
	emission.User = req.User
	emission.Name = req.EmissionName
	emission.Source = req.Source
	emission.TypeId = req.TypeId
	emission.Version = "1"
	emission.StartDate = time.Now().UTC()
	endDate, err := time.Parse("2006-01-02T15:04:05.000Z", constants.EMISSION_END_DATE)
	if err != nil {
		return err
	}
	emission.EndDate = endDate

	emissionId, err := utils.CreateEmission(emission)
	if err != nil {
		logs.Error(err)
		return err
	}

	factor := new(entities.CreateFactorReq)
	if req.IsCalculated {
		factor.Factor = constants.IS_CALCULATED_FACTOR_VALUE
		factor.UnitId = *req.FactorUnitId
		factor.Name = "TOTAL"
	} else {
		factor.Factor = *req.Factor
		factor.UnitId = *req.FactorUnitId
		factor.Name = "TOTAL"
	}
	factor.User = req.User
	factorId, err := utils.CreateFactor(factor)
	if err != nil {
		return err
	}

	emissionFactor := new(entities.CreateEmissionFactorReq)
	emissionFactor.User = req.User
	emissionFactor.EmissionId = emissionId
	emissionFactor.FactorId = factorId
	emissionFactor.SiteId = req.SiteId
	emissionFactor.CompanyId = req.CompanyId

	emissionFactorId, err := utils.CreateEmissionFactor(emissionFactor)
	if err != nil {
		return err
	}
	activityCreate := new(entities.ActivityCreateReq)
	activityCreate.User = req.User
	activityCreate.EmissionFactorId = emissionFactorId
	activityCreate.FrequencySelectedDate = req.FrequencySelectedDate
	activityCreate.FrequencyType = req.FrequencyType
	activityCreate.Name = req.ActivityName
	activityCreate.Description = req.ActivityDescription
	activityCreate.SiteId = req.SiteId
	activity_id, err := a.ActivityRepo.CreateActivity(activityCreate)
	if err != nil {
		return err
	}
	if req.FrequencyType == nil && req.FrequencySelectedDate == nil {
		logs.Error("create_activity_not_create_frequency_case")
	} else {
		activityFrequency := &entities.ActivityFrequency{
			ActivityId:   activity_id,
			Type:         *req.FrequencyType,
			SelectedDate: req.FrequencySelectedDate,
			CreatedBy:    req.User,
		}
		_, errFrequency := utils.CreateActivityFrequency(activityFrequency)
		if errFrequency != nil {
			return errFrequency
		}
	}
	return nil
}

// UpdateActivityWithCustomizedEmission implements entities.ActivityUsecase.
func (a *activityUse) UpdateActivityWithCustomizedEmission(req *entities.UpdateActivityWithCustomizedEmissionReq) error {
	siteValue, err := utils.GetSiteValue(req.SiteId)
	if err != nil {
		return err
	}

	/* Create Emission */
	emission := new(entities.CreateEmissionReq)
	emission.CountryId = siteValue.CountryID
	emission.User = req.User
	emission.Name = req.EmissionName
	emission.Source = req.Source
	emission.TypeId = req.TypeId
	emission.Version = "1"
	emission.StartDate = time.Now().UTC()
	endDate, err := time.Parse("2006-01-02T15:04:05.000Z", constants.EMISSION_END_DATE)
	if err != nil {
		return err
	}
	emission.EndDate = endDate

	emissionId, err := utils.CreateEmission(emission)
	if err != nil {
		logs.Error(err)
		return err
	}
	/* Create Factor */
	factor := new(entities.CreateFactorReq)
	if req.IsCalculated {
		factor.Factor = constants.IS_CALCULATED_FACTOR_VALUE
		factor.UnitId = *req.FactorUnitId
		factor.Name = "TOTAL"
	} else {
		factor.Factor = *req.Factor
		factor.UnitId = *req.FactorUnitId
		factor.Name = "TOTAL"
	}
	factor.User = req.User
	factorId, err := utils.CreateFactor(factor)
	if err != nil {
		return err
	}

	/* Create Emission Factor */
	emissionFactor := new(entities.CreateEmissionFactorReq)
	emissionFactor.User = req.User
	emissionFactor.EmissionId = emissionId
	emissionFactor.FactorId = factorId
	emissionFactor.SiteId = req.SiteId
	emissionFactor.CompanyId = req.CompanyId

	emissionFactorId, err := utils.CreateEmissionFactor(emissionFactor)
	if err != nil {
		return err
	}

	/* update Activity */
	activityUpdate := new(entities.UpdateActivityReq)
	activityUpdate.Id = req.Id
	activityUpdate.UpdateUser = req.User
	activityUpdate.EmissionFactorId = emissionFactorId
	activityUpdate.FrequencySelectedDate = req.FrequencySelectedDate
	activityUpdate.FrequencyType = *req.FrequencyType
	activityUpdate.Name = req.ActivityName
	activityUpdate.Description = req.ActivityDescription
	activityUpdate.SiteId = req.SiteId
	activityUpdate.CompanyId = req.CompanyId

	err = a.ActivityRepo.UpdateActivity(activityUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (a *activityUse) UpdateActivityStatus(req *entities.UpdateActivityStatusReq) error {

	err := a.ActivityRepo.UpdateActivityStatus(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *activityUse) DisconnectBill(req *entities.DisConnectBillReq) error {
	err := utils.DisconnectBill(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *activityUse) UpdateBillConnect(req *entities.UpsertBillingConnectReq) error {
	err := utils.UpsertBillConnect(req)
	if err != nil {
		return err
	}
	return nil
}

func (a *activityUse) CreateActivityWithRecepiEmission(req *entities.CreateActivityWithRecepiEmissionReq) error {

	siteValue, err := utils.GetSiteValue(req.SiteId)
	if err != nil {
		return err
	}

	// create emission
	emission := new(entities.CreateEmissionReq)
	emission.CountryId = siteValue.CountryID
	emission.User = req.User
	emission.Name = req.EmissionName
	emission.Source = req.Source
	emission.TypeId = req.TypeId
	emission.Version = "1"
	emission.StartDate = time.Now().UTC()
	endDate, err := time.Parse("2006-01-02T15:04:05.000Z", constants.EMISSION_END_DATE)
	if err != nil {
		return err
	}
	emission.EndDate = endDate

	emissionId, err := utils.CreateEmission(emission)
	if err != nil {
		logs.Error(err)
		return err
	}

	var total = 0.00
	// loop for create factor,emision factor
	for _, factor_type := range req.FactorType {

		factor := new(entities.CreateFactorReq)
		factor.Name = factor_type.Name
		factor.Factor = factor_type.Factor
		factor.UnitId = *req.FactorUnitId
		factor.User = req.User
		factor.Type = factor_type.Id

		factorId, err := utils.CreateFactor(factor)
		if err != nil {
			return err
		}

		emissionFactor := new(entities.CreateEmissionFactorReq)
		emissionFactor.User = req.User
		emissionFactor.EmissionId = emissionId
		emissionFactor.FactorId = factorId
		emissionFactor.CompanyId = req.CompanyId
		emissionFactor.SiteId = req.SiteId

		_, err = utils.CreateEmissionFactor(emissionFactor)
		if err != nil {
			return err
		}

		gwp, err := utils.GetGWPValueById(factor_type.GWP_Id)
		if err != nil {
			return err
		}

		total += (factor_type.Factor * gwp.GWP)

	} // end for

	factor := new(entities.CreateFactorReq)
	factor.Name = "TOTAL"
	factor.Factor = total
	factor.UnitId = *req.FactorUnitId
	factor.User = req.User

	factorId, err := utils.CreateFactor(factor)
	if err != nil {
		return err
	}

	emissionFactor := new(entities.CreateEmissionFactorReq)
	emissionFactor.User = req.User
	emissionFactor.EmissionId = emissionId
	emissionFactor.FactorId = factorId
	emissionFactor.CompanyId = req.CompanyId
	emissionFactor.SiteId = req.SiteId

	emissionFactorId, err := utils.CreateEmissionFactor(emissionFactor)
	if err != nil {
		return err
	}

	activityCreate := new(entities.ActivityCreateReq)
	activityCreate.User = req.User
	activityCreate.EmissionFactorId = emissionFactorId
	activityCreate.FrequencySelectedDate = req.FrequencySelectedDate
	activityCreate.FrequencyType = req.FrequencyType
	activityCreate.Name = req.ActivityName
	activityCreate.Description = req.ActivityDescription
	activityCreate.SiteId = req.SiteId
	activity_id, err := a.ActivityRepo.CreateActivity(activityCreate)
	if err != nil {
		return err
	}

	if req.FrequencyType == nil && req.FrequencySelectedDate == nil {
		logs.Error("create_activity_not_create_frequency_case")
	} else {
		activityFrequency := &entities.ActivityFrequency{
			ActivityId:   activity_id,
			Type:         *req.FrequencyType,
			SelectedDate: req.FrequencySelectedDate,
			CreatedBy:    req.User,
		}
		_, errFrequency := utils.CreateActivityFrequency(activityFrequency)
		if errFrequency != nil {
			return errFrequency
		}
	}

	return nil
}

// UpdateActivityWithRecepiEmission implements entities.ActivityUsecase.
func (a *activityUse) UpdateActivityWithRecepiEmission(req *entities.UpdateActivityWithRecepiEmissionReq) error {

	siteValue, err := utils.GetSiteValue(req.SiteId)
	if err != nil {
		return err
	}

	/* create emission */ // create emission
	emission := new(entities.CreateEmissionReq)
	emission.CountryId = siteValue.CountryID
	emission.User = req.User
	emission.Name = req.EmissionName
	emission.Source = req.Source
	emission.TypeId = req.TypeId
	emission.Version = "1"
	emission.StartDate = time.Now().UTC()
	endDate, err := time.Parse("2006-01-02T15:04:05.000Z", constants.EMISSION_END_DATE)
	if err != nil {
		return err
	}
	emission.EndDate = endDate

	emissionId, err := utils.CreateEmission(emission)
	if err != nil {
		logs.Error(err)
		return err
	}

	/* create factor */
	var total = 0.00
	// loop for create factor,emision factor
	for _, factor_type := range req.FactorType {

		factor := new(entities.CreateFactorReq)
		factor.Name = factor_type.Name
		factor.Factor = factor_type.Factor
		factor.UnitId = *req.FactorUnitId
		factor.User = req.User
		factor.Type = factor_type.Id

		factorId, err := utils.CreateFactor(factor)
		if err != nil {
			return err
		}

		emissionFactor := new(entities.CreateEmissionFactorReq)
		emissionFactor.User = req.User
		emissionFactor.EmissionId = emissionId
		emissionFactor.FactorId = factorId
		emissionFactor.CompanyId = req.CompanyId
		emissionFactor.SiteId = req.SiteId

		_, err = utils.CreateEmissionFactor(emissionFactor)
		if err != nil {
			return err
		}

		gwp, err := utils.GetGWPValueById(factor_type.GWP_Id)
		if err != nil {
			return err
		}

		total += (factor_type.Factor * gwp.GWP)

	} // end for

	factor := new(entities.CreateFactorReq)
	factor.Name = "TOTAL"
	factor.Factor = total
	factor.UnitId = *req.FactorUnitId
	factor.User = req.User

	factorId, err := utils.CreateFactor(factor)
	if err != nil {
		return err
	}

	/* create emission factor */
	emissionFactor := new(entities.CreateEmissionFactorReq)
	emissionFactor.User = req.User
	emissionFactor.EmissionId = emissionId
	emissionFactor.FactorId = factorId
	emissionFactor.CompanyId = req.CompanyId
	emissionFactor.SiteId = req.SiteId

	emissionFactorId, err := utils.CreateEmissionFactor(emissionFactor)
	if err != nil {
		return err
	}

	/* update activity */
	activityUpdate := new(entities.UpdateActivityReq)
	activityUpdate.Id = req.Id
	activityUpdate.UpdateUser = req.User
	activityUpdate.EmissionFactorId = emissionFactorId
	activityUpdate.Name = req.ActivityName
	activityUpdate.Description = req.ActivityDescription
	activityUpdate.SiteId = req.SiteId
	activityUpdate.CompanyId = req.CompanyId
	err = a.ActivityRepo.UpdateActivity(activityUpdate)
	if err != nil {
		return err
	}

	return nil
}

func getLastDayOfMonth(year, month int) int {
	// Get the first day of the next month
	firstDayOfNextMonth := time.Date(year, time.Month(month)+1, 1, 0, 0, 0, 0, time.UTC)

	// Subtract one day from the first day of the next month to get the last day of the current month
	lastDayOfMonth := firstDayOfNextMonth.AddDate(0, 0, -1)

	return lastDayOfMonth.Day()
}
