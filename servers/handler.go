package servers

import (
	"fmt"

	_usersHttp "bufferbox_backend_go/modules/users/controllers"
	_usersRepository "bufferbox_backend_go/modules/users/repositories"
	_usersUsecase "bufferbox_backend_go/modules/users/usecases"
	"bufferbox_backend_go/pkg/utils"

	_addressDataHttp "bufferbox_backend_go/modules/address_data/controllers"
	_addressDataRepository "bufferbox_backend_go/modules/address_data/repositories"
	_addressDataUsecase "bufferbox_backend_go/modules/address_data/usecases"

	_industryTypeHttp "bufferbox_backend_go/modules/industry_type/controllers"
	_industryTypeRepository "bufferbox_backend_go/modules/industry_type/repositories"
	_industryTypeUsecase "bufferbox_backend_go/modules/industry_type/usecases"

	_companyHttp "bufferbox_backend_go/modules/company/controllers"
	_companyRepository "bufferbox_backend_go/modules/company/repositories"
	_companyUsecase "bufferbox_backend_go/modules/company/usecases"

	_subscriptionHttp "bufferbox_backend_go/modules/subscription/controllers"
	_subscriptionRepository "bufferbox_backend_go/modules/subscription/repositories"
	_subscriptionUsecase "bufferbox_backend_go/modules/subscription/usecases"

	_roleHttp "bufferbox_backend_go/modules/role/controllers"
	_roleRepository "bufferbox_backend_go/modules/role/repositories"
	_roleUsecase "bufferbox_backend_go/modules/role/usecases"

	_groupTagHttp "bufferbox_backend_go/modules/group_tag/controllers"
	_groupTagRepository "bufferbox_backend_go/modules/group_tag/repositories"
	_groupTagUsecase "bufferbox_backend_go/modules/group_tag/usecases"

	_siteHttp "bufferbox_backend_go/modules/sites/controllers"
	_siteRepository "bufferbox_backend_go/modules/sites/repositories"
	_siteUsecase "bufferbox_backend_go/modules/sites/usecases"

	_activityHttp "bufferbox_backend_go/modules/activity/controllers"
	_activityRepository "bufferbox_backend_go/modules/activity/repositories"
	_activityUsecase "bufferbox_backend_go/modules/activity/usecases"

	_activityTaskHttp "bufferbox_backend_go/modules/activity_task/controllers"
	_activityTaskRepository "bufferbox_backend_go/modules/activity_task/repositories"
	_activityTaskUsecase "bufferbox_backend_go/modules/activity_task/usecases"

	_emissionHttp "bufferbox_backend_go/modules/emission/controllers"
	_emissionRepository "bufferbox_backend_go/modules/emission/repositories"
	_emissionUsecase "bufferbox_backend_go/modules/emission/usecases"

	_attachmentHttp "bufferbox_backend_go/modules/attachment/controllers"
	_attachmentRepository "bufferbox_backend_go/modules/attachment/repositories"
	_attachmentUsecase "bufferbox_backend_go/modules/attachment/usecases"

	_notificationHttp "bufferbox_backend_go/modules/notification/controllers"
	_notificationRepository "bufferbox_backend_go/modules/notification/repositories"
	_notificationUsecase "bufferbox_backend_go/modules/notification/usecases"

	ws "bufferbox_backend_go/pkg/ws"
	"github.com/gofiber/websocket/v2"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/api/v1")

	notificationGroup, ok := v1.Group("/noti").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, notificationGroup should be of type *fiber.Group")
	}

	notificationRepository := _notificationRepository.NewNotificationRepository(s.Db)
	notificationUsecase := _notificationUsecase.NewNotificationUsecase(notificationRepository)
	_notificationHttp.NewNotificationController(notificationGroup, notificationUsecase)

	subscriptionGroup, ok := v1.Group("/subscription").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, subscriptionGroup should be of type *fiber.Group")
	}

	subscriptionRepository := _subscriptionRepository.NewSubscriptionRepository(s.Db)
	subscriptionUsecase := _subscriptionUsecase.NewSubscriptionUsecase(subscriptionRepository)
	_subscriptionHttp.NewSubscriptionController(subscriptionGroup, subscriptionUsecase)

	emissionGroup, ok := v1.Group("/emission").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, emissionGroup should be of type *fiber.Group")
	}

	emissionRepository := _emissionRepository.NewEmissionRepository(s.Db)
	emissionUsecase := _emissionUsecase.NewEmissionUsecase(emissionRepository)
	_emissionHttp.NewEmissionController(emissionGroup, emissionUsecase)

	activityGroup, ok := v1.Group("/activity").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, activityGroup should be of type *fiber.Group")
	}

	activityRepository := _activityRepository.NewActivityRepository(s.Db)
	activityUsecase := _activityUsecase.NewActivityUsecase(activityRepository)
	_activityHttp.NewActivityController(activityGroup, activityUsecase)

	attachmentGroup, ok := v1.Group("/attachment").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, attachmentGroup should be of type *fiber.Group")
	}
	attachmentRepository := _attachmentRepository.NewAttachmentRepository(s.Db)
	attachmentUsecase := _attachmentUsecase.NewAttachmentUsecase(attachmentRepository)
	_attachmentHttp.NewAttachmentController(attachmentGroup, attachmentUsecase)

	activityTaskGroup, ok := v1.Group("/activity-task").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, activityTaskGroup should be of type *fiber.Group")
	}

	activityTaskRepository := _activityTaskRepository.NewActivityTaskRepository(s.Db)
	activityTaskUsecase := _activityTaskUsecase.NewActivityTaskUsecase(activityTaskRepository, attachmentRepository)
	_activityTaskHttp.NewActivityTaskController(activityTaskGroup, activityTaskUsecase)

	groupTagGroup, ok := v1.Group("/group-tag").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, group_tag_Group should be of type *fiber.Group")
	}

	groupTagRepository := _groupTagRepository.NewGroupTagRepository(s.Db)
	groupTagUsecase := _groupTagUsecase.NewGroupTagUsecase(groupTagRepository)
	_groupTagHttp.NewGroupTagController(groupTagGroup, groupTagUsecase)

	roleGroup, ok := v1.Group("/role").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, roleGroup should be of type *fiber.Group")
	}

	roleRepository := _roleRepository.NewRoleRepository(s.Db)
	roleUsecase := _roleUsecase.NewRoleUsecase(roleRepository)
	_roleHttp.NewRoleController(roleGroup, roleUsecase)

	companyGroup, ok := v1.Group("/company").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, companyGroup should be of type *fiber.Group")
	}

	companyRepository := _companyRepository.NewCompanyRepository(s.Db)
	companyUsecase := _companyUsecase.NewCompanyUsecase(companyRepository, roleRepository)
	_companyHttp.NewCompanyController(companyGroup, companyUsecase)

	siteGroup, ok := v1.Group("/site").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, siteGroup should be of type *fiber.Group")
	}

	siteRepository := _siteRepository.NewSiteRepository(s.Db)
	siteUsecase := _siteUsecase.NewSiteUsecase(siteRepository, roleRepository)
	_siteHttp.NewSiteController(siteGroup, siteUsecase)

	industryTypeGroup, ok := v1.Group("/industry-type").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, industryGroup should be of type *fiber.Group")
	}

	industryTypeRepository := _industryTypeRepository.NewIndustryTypeRepository(s.Db)
	industryTypeUsecase := _industryTypeUsecase.NewIndustryTypeUsecase(industryTypeRepository)
	_industryTypeHttp.NewIndustryTypeController(industryTypeGroup, industryTypeUsecase)

	addressDataGroup, ok := v1.Group("/address").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, addressGroup should be of type *fiber.Group")
	}

	addressDataRepository := _addressDataRepository.NewAddressDataRepository(s.Db)
	addressDataUsecase := _addressDataUsecase.NewAddressDataUsecase(addressDataRepository)
	_addressDataHttp.NewAddressDataController(addressDataGroup, addressDataUsecase)

	usersGroup, ok := v1.Group("/users").(*fiber.Group)
	if !ok {
		return fmt.Errorf("error, usersGroup should be of type *fiber.Group")
	}
	usersRepository := _usersRepository.NewUsersRepository(s.Db)
	usersUsecase := _usersUsecase.NewUsersUsecase(usersRepository)
	_usersHttp.NewUsersController(usersGroup, usersUsecase)

	// --- Custom notify endpoint (simple example) ---
	v1.Get("/notify/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong"})
	})
	// --- End custom notify endpoint ---

	// --- WebSocket notify endpoint ---
	v1.Use("/notify/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	v1.Get("/notify/ws", websocket.New(ws.ConnHandler))
	// --- End WebSocket notify endpoint ---

	s.App.Use(func(c *fiber.Ctx) error {
		return utils.HandleResponse(c, fiber.ErrInternalServerError.Code, "error, end point not found", nil)
	})
	return nil
}
