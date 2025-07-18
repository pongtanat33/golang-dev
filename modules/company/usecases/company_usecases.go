package usecases

import (
	"bufferbox_backend_go/constants"
	"bufferbox_backend_go/entities"
	"bufferbox_backend_go/logs"
	"bufferbox_backend_go/pkg/utils"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type companyUse struct {
	CompanyRepo entities.CompanyRepository
	RoleRepo    entities.RoleRepository
}

func NewCompanyUsecase(companyRepo entities.CompanyRepository, roleRepo entities.RoleRepository) entities.CompanyUsecase {
	return &companyUse{
		CompanyRepo: companyRepo,
		RoleRepo:    roleRepo,
	}
}

func (c *companyUse) CompanyMemberDelete(req *entities.MemberDeleteReq) error {

	chkOwner, err := utils.CheckUserOwner(req.Email, req.TablePermissionId, "Owner")
	if err != nil {
		logs.Error(err)
		return err
	}

	if chkOwner {
		logs.Error(errors.New("cannot delete owner"))
		return fmt.Errorf("cannot delete owner")
	}
	err = utils.DeleteMember(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *companyUse) CompanyMemberChange(req *entities.UpsertUserPermissionReq) error {
	err := utils.UpsertUserPermission(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *companyUse) CompanyMember(req *entities.CompanyMemberReq) ([]entities.TeamMemberRes, error) {
	req.RoleType = constants.C_ROLE_TYPE
	companyMember, err := utils.GetTeamMembers(req.CompanyId, req.RoleType)
	if err != nil {
		return nil, err
	}
	return companyMember, nil
}

func (c *companyUse) CompanyGroupsByName(req *entities.CompanyGroupByNameReq) ([]entities.CompanyGroupRes, int, error) {
	companyGroupByName, totalRows, err := c.CompanyRepo.CompanyGroupsByName(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	return companyGroupByName, totalRows, nil
}

func (c *companyUse) CompanyGroups(req *entities.CompanyGroupReq) ([]entities.CompanyGroupRes, int, error) {
	companyGroup, totalRows, err := c.CompanyRepo.CompanyGroups(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	return companyGroup, totalRows, nil
}

func (c *companyUse) CompanySites(req *entities.CompanySiteReq) ([]entities.CompanySiteRes, int, error) {
	companySite, totalRows, err := c.CompanyRepo.CompanySites(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	return companySite, totalRows, nil
}

func (c *companyUse) CompanySelectedSites(req *entities.CompanySiteReq) ([]entities.CompanySiteRes, int, error) {
	companySite, totalRows, err := c.CompanyRepo.CompanySelectedSites(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	return companySite, totalRows, nil
}

func (c *companyUse) CompanyValue(company_id string, user_id string) (*entities.CompanyRes, error) {

	companyChan := make(chan *entities.CompanyRes)
	siteQtyChan := make(chan int)
	errChan := make(chan error)

	// Start GetCompanyValue Goroutine
	go func() {
		company, err := utils.GetCompanyValue(company_id, user_id)
		if err != nil {
			errChan <- err
			return
		}
		companyChan <- company
	}()

	// Start CompanyCountSite Goroutine
	go func() {
		res, err := utils.CompanyCountSite(company_id)
		if err != nil {
			errChan <- err
			return
		}
		siteQtyChan <- res
	}()

	// Wait for both Goroutines to complete
	var company *entities.CompanyRes
	var siteQty int
	for i := 0; i < 2; i++ {
		select {
		case comp := <-companyChan:
			company = comp
		case qty := <-siteQtyChan:
			siteQty = qty
		case err := <-errChan:
			return nil, err
		}
	}

	// Assign the siteQty to the company object
	company.SiteQty = siteQty

	// Update user last visited company
	err := utils.UpdateUserLastVisitedCompany(user_id, company_id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// elapsed := time.Since(start) // เวลาที่ใช้ทั้งหมด
	// fmt.Println("Time elapsed:", elapsed)
	return company, nil
}

func (c *companyUse) CompanyLists(user_id string) ([]entities.CompanyList, *entities.CompanyRes, error) {
	companyListChan := make(chan []entities.CompanyList)

	// Start GetCompanyValue Goroutine
	go func() {
		defer close(companyListChan) // ปิด channel หลังจากที่ goroutine เสร็จสิ้น
		companyList, err := utils.CompanyLists(user_id, constants.C_ROLE_TYPE, "")
		if err != nil {
			return
		}
		companyListChan <- companyList
	}()

	user, err := utils.GetUserValue(user_id)
	if err != nil {
		logs.Error(err)
		return nil, nil, err
	}

	if user.LastVisitedCompany != nil {
		lastVisitedCompany, err := utils.GetCompanyValue(*user.LastVisitedCompany, user_id)
		//wait for GetCompanyValue Goroutine
		companyList := <-companyListChan

		if err != nil {
			logs.Error(err)
			//wait for GetCompanyValue Goroutine
			lastVisitedCompany, err = utils.GetCompanyValue(companyList[0].Id, user_id)

			if err != nil {
				return nil, nil, err
			}
		}

		if !lastVisitedCompany.IsActive {
			lastVisitedCompany = nil

			return companyList, lastVisitedCompany, nil
		}

		return companyList, lastVisitedCompany, nil

	} else {

		//wait for GetCompanyValue Goroutine
		companyList := <-companyListChan

		lastVisitedCompany, err := utils.GetCompanyValue(companyList[0].Id, user_id)

		if err != nil {
			logs.Error(err)
			return nil, nil, err
		}

		return companyList, lastVisitedCompany, nil
	}

}

func (c *companyUse) CompanyCreate(req *entities.CompanyReq) (*entities.CompanyRes, error) {
	// Define a channel for errors
	errChan := make(chan error)
	var wg sync.WaitGroup

	company, err := c.CompanyRepo.CompanyCreate(req)
	if err != nil {
		return nil, err
	}

	// Create role in a goroutine
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	rowsStandard, err := c.RoleRepo.GetRoleStandard(constants.C_ROLE_TYPE)
	// 	if err != nil {
	// 		errChan <- err
	// 		return
	// 	}
	// 	if len(rowsStandard) == 0 {
	// 		return
	// 	}

	// 	for _, row := range rowsStandard {
	// 		roleFeature, err := c.RoleRepo.GetRoleFeatureStandard(row.Id)
	// 		if err != nil {
	// 			errChan <- err
	// 			continue
	// 		}

	// 		if roleFeature != nil {
	// 			createRoleReq := &entities.RoleCreateReq{
	// 				RoleName: row.Name,
	// 				RoleType: row.RoleType,
	// 				User:     req.User,
	// 			}

	// 			roleID, err := c.RoleRepo.CreateRole(createRoleReq)
	// 			if err != nil {
	// 				errChan <- err
	// 				continue
	// 			}

	// 			errPermission := utils.CreateUserPermissionByRoleId(company.Id, req.User, roleID, req.User)
	// 			if errPermission != nil {
	// 				errChan <- errPermission
	// 				continue
	// 			}

	// 			for _, feature := range roleFeature {
	// 				roleFeatureReq := &entities.RoleFeatureReq{
	// 					RoleId:    roleID,
	// 					FeatureId: feature.FeatureId,
	// 					Create:    feature.Create,
	// 					Read:      feature.Read,
	// 					Update:    feature.Update,
	// 					Delete:    feature.Delete,
	// 				}
	// 				if err := c.RoleRepo.UpsertRoleFeature(roleFeatureReq); err != nil {
	// 					errChan <- err
	// 				}
	// 			}
	// 		}
	// 	}
	// }()

	//Other goroutines (create user permission, update last visited company)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := utils.CreateUserPermissionStandardRole(company.Id, req.User, constants.Owner, constants.C_ROLE_TYPE, req.User)
		if err != nil {
			errChan <- err
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := utils.UpdateUserLastVisitedCompany(req.User, company.Id)
		if err != nil {
			errChan <- err
		}
	}()

	// Wait for goroutines and handle errors
	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}

	// subscription
	packReq := new(entities.PackageUserReq)
	packReq.PackageId = constants.FreeTrial
	packReq.User = req.User
	existTrial, err := utils.GetPackageFreeTrialByUser(packReq)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// check what package to subscription
	subReq := &entities.TicketReq{
		CompanyId: company.Id,
		TicketKey: utils.GenTicketKey16Digit(),
		User:      req.User,
		Type:      constants.TrialType,
	}

	// user already free trial is issued
	if existTrial.TicketId == "" {

		subReq.PackageId = constants.FreeTrial

		// if not expired, then use free trial
	} else if !existTrial.IsExpired && !existTrial.IsUsed {

		subReq.PackageId = constants.FreeTrial

	} else {

		// freemium subscription
		subReq.PackageId = constants.Freemium
		subReq.Type = constants.FreeType
	}

	_, errSubscription := utils.NewSubscription(subReq)
	if errSubscription != nil {
		return nil, errSubscription
	}

	return company, nil
}

func (c *companyUse) CompanyUpdate(req *entities.CompanyReq) error {
	result, err := c.CompanyRepo.CompanyUpdate(req)
	if err != nil {
		return err
	}
	if result.Result == "duplicate_name" {
		return errors.New("duplicated_site_name")
	}
	return nil
}

func (c *companyUse) CompanyDelete(req *entities.CompanyReq) error {
	// // get ticket id
	subReq := new(entities.SubscriptionReq)
	subReq.CompanyId = req.Id
	subscription, _ := utils.GetSubscription(subReq)
	// del unsubscription
	delReq := new(entities.UnsubscriptionReq)
	delReq.CompanyId = req.Id
	delReq.TicketId = subscription.TicketId
	delReq.User = req.User
	_ = utils.Unsubscription(delReq)

	// do delete company
	err := c.CompanyRepo.CompanyDelete(req)
	if err != nil {
		return err
	}

	// Last visite company
	var companyList []entities.CompanyList
	companyList, err = utils.CompanyLists(req.User, constants.C_ROLE_TYPE, "")
	if err != nil {
		return nil
	}

	if companyList != nil {
		err = utils.UpdateUserLastVisitedCompany(req.User, companyList[0].Id)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return nil
}

func (c *companyUse) CompanyTargetBaseAll(req *entities.CompanyTargetBaseReq) ([]entities.CompanyTargetBaseRes, error) {
	result, err := c.CompanyRepo.CompanyTargetBaseAll(req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *companyUse) CompanyTargetBaseById(req *entities.CompanyTargetBaseByIdReq) (*entities.CompanyTargetBaseRes, error) {
	result, err := c.CompanyRepo.CompanyTargetBaseById(req)
	if err != nil {
		return nil, err
	}
	if result.Id == "" {
		return nil, nil
	}
	return result, nil
}

func (c *companyUse) CompanyTargetAll(req *entities.CompanyTargetBaseReq) ([]entities.CompanyTargetRes, error) {
	result, err := c.CompanyRepo.CompanyTargetAll(req)
	if err != nil {
		return nil, err
	}
	var companyTargetResult []entities.CompanyTargetRes
	for _, res := range result {
		var resCompanyTarget entities.CompanyTargetRes
		resCompanyTarget.Id = res.Id
		resCompanyTarget.Year = res.Year
		resCompanyTarget.CompanyId = res.CompanyId
		resCompanyTarget.Scope1Max = res.Scope1Max
		resCompanyTarget.Scope2Max = res.Scope2Max
		resCompanyTarget.Scope3Max = res.Scope3Max
		resCompanyTarget.Scope1ReductionPercent = res.Scope1ReductionPercent
		resCompanyTarget.Scope2ReductionPercent = res.Scope2ReductionPercent
		resCompanyTarget.Scope3ReductionPercent = res.Scope3ReductionPercent
		resCompanyTarget.CreatedAt = res.CreatedAt
		resCompanyTarget.MainTarget = res.MainTarget
		if res.Base != nil {
			tgb := new(entities.CompanyTargetBaseByIdReq)
			tgb.CompanyId = res.CompanyId
			tgb.BaseId = *res.Base
			targetBase, err := c.CompanyRepo.CompanyTargetBaseById(tgb)
			if err != nil {
				return nil, err
			}
			resCompanyTarget.Base = targetBase
		} else {
			resCompanyTarget.Base = nil
		}
		companyTargetResult = append(companyTargetResult, resCompanyTarget)
	}
	return companyTargetResult, nil
}

func (c *companyUse) CompanyTargetById(req *entities.CompanyTargetByIdReq) (*entities.CompanyTargetRes, error) {
	result, err := c.CompanyRepo.CompanyTargetById(req)
	if err != nil {
		return nil, err
	}
	var resCompanyTarget entities.CompanyTargetRes
	resCompanyTarget.Id = result.Id
	resCompanyTarget.Year = result.Year
	resCompanyTarget.CompanyId = result.CompanyId
	resCompanyTarget.Scope1Max = result.Scope1Max
	resCompanyTarget.Scope2Max = result.Scope2Max
	resCompanyTarget.Scope3Max = result.Scope3Max
	resCompanyTarget.Scope1ReductionPercent = result.Scope1ReductionPercent
	resCompanyTarget.Scope2ReductionPercent = result.Scope2ReductionPercent
	resCompanyTarget.Scope3ReductionPercent = result.Scope3ReductionPercent
	resCompanyTarget.CreatedAt = result.CreatedAt
	if result.Base != nil {
		tgb := new(entities.CompanyTargetBaseByIdReq)
		tgb.CompanyId = result.CompanyId
		tgb.BaseId = *result.Base
		targetBase, err := c.CompanyRepo.CompanyTargetBaseById(tgb)
		if err != nil {
			return nil, err
		}
		resCompanyTarget.Base = targetBase
	} else {
		resCompanyTarget.Base = nil
	}
	return &resCompanyTarget, nil
}

func (c *companyUse) CompanyTargetDelete(req *entities.CompanyTargetDeleteReq) error {
	err := c.CompanyRepo.CompanyTargetDelete(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *companyUse) CompanyMainTargetUpdate(req *entities.CompanyMainTargetUpdateReq) error {
	err := c.CompanyRepo.CompanyMainTargetUpdate(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *companyUse) CompanyTargetBaseDelete(req *entities.CompanyTargetBaseDeleteReq) error {
	result, err := c.CompanyRepo.CompanyTargetBaseDelete(req)
	if err != nil {
		return err
	}
	if !result.Result {
		return errors.New("company_target_of_this_base_is_active")
	}
	return nil
}

func (c *companyUse) CompanyTargetBaseUpsert(req *entities.CompanyTargetBaseUpsertReq) (*entities.CompanyTargetBaseUpsertRes, error) {
	if req.Id == nil {
		newId := uuid.New().String()
		req.Id = &newId
	}
	err := c.CompanyRepo.CompanyTargetBaseUpsert(req)
	if err != nil {
		return nil, err
	}
	var resCompanyTarget entities.CompanyTargetBaseUpsertRes
	resCompanyTarget.Id = req.Id
	resCompanyTarget.Year = req.Year
	resCompanyTarget.CompanyId = req.CompanyId
	resCompanyTarget.Scope1 = req.Scope1
	resCompanyTarget.Scope2 = req.Scope2
	resCompanyTarget.Scope3 = req.Scope3

	return &resCompanyTarget, nil
}

func (c *companyUse) CompanyTargetUpsert(req *entities.CompanyTargetUpsertReq) (*entities.CompanyTargetUpsertRes, error) {
	if req.Id == nil {
		newId := uuid.New().String()
		req.Id = &newId
	}
	err := c.CompanyRepo.CompanyTargetUpsert(req)
	if err != nil {
		return nil, err
	}
	var resCompanyTarget entities.CompanyTargetUpsertRes
	resCompanyTarget.Id = req.Id
	resCompanyTarget.Year = req.Year
	resCompanyTarget.CompanyId = req.CompanyId
	resCompanyTarget.Scope1Max = req.Scope1Max
	resCompanyTarget.Scope2Max = req.Scope2Max
	resCompanyTarget.Scope3Max = req.Scope3Max
	resCompanyTarget.Scope1ReductionPercent = req.Scope1ReductionPercent
	resCompanyTarget.Scope2ReductionPercent = req.Scope2ReductionPercent
	resCompanyTarget.Scope3ReductionPercent = req.Scope3ReductionPercent
	if req.Base != nil {
		tgb := new(entities.CompanyTargetBaseByIdReq)
		tgb.CompanyId = req.CompanyId
		tgb.BaseId = *req.Base
		targetBase, err := c.CompanyRepo.CompanyTargetBaseById(tgb)
		if err != nil {
			return nil, err
		}
		resCompanyTarget.Base = targetBase
	} else {
		resCompanyTarget.Base = nil
	}

	// set main target
	companyTargetReq := new(entities.CompanyTargetBaseReq)
	companyTargetReq.CompanyId = req.CompanyId
	companyTarget, _ := c.CompanyTargetAll(companyTargetReq)
	if len(companyTarget) == 1 {
		for i := 0; i < len(companyTarget); i++ {
			if companyTarget[i].MainTarget != true {
				setMainTargetReq := new(entities.CompanyMainTargetUpdateReq)
				setMainTargetReq.Id = *req.Id
				setMainTargetReq.User = req.User
				_ = c.CompanyMainTargetUpdate(setMainTargetReq)
				break
			}
		}
	}

	return &resCompanyTarget, nil
}

func (c *companyUse) CompanyGroupSites(req *entities.CompanySiteReq) ([]entities.CompanyGroupSiteRes, int, error) {
	companySite, totalRows, err := c.CompanyRepo.CompanyGroupSites(req)
	if err != nil {
		logs.Error(err)
		return nil, 0, err
	}
	return companySite, totalRows, nil
}
