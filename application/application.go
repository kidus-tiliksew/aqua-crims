package application

import (
	"context"

	"github.com/kidus-tiliksew/aqua-crims/application/commands"
	"github.com/kidus-tiliksew/aqua-crims/application/queries"
	"github.com/kidus-tiliksew/aqua-crims/domain"
)

type (
	App interface {
		Commands
		Queries
	}

	Commands interface {
		CustomerCreate(ctx context.Context, cmd commands.CustomerCreate) (*domain.Customer, error)
		CustomerCreateCloudResources(ctx context.Context, cmd commands.CustomerCreateCloudResources) ([]*domain.CloudResource, error)
		CloudResourceCreate(ctx context.Context, cmd commands.CloudResourceCreate) (*domain.CloudResource, error)
		CloudResourceUpdate(ctx context.Context, cmd commands.CloudResourceUpdate) error
		CloudResourceDelete(ctx context.Context, cmd commands.CloudResourceDelete) error
		NotificationCreate(ctx context.Context, cmd commands.NotificationCreate) (*domain.Notification, error)
		NotificationDelete(ctx context.Context, id int64) error
		NotificationDeleteByUser(ctx context.Context, userID string) error
	}

	Queries interface {
		CloudResourceGet(ctx context.Context, query queries.CloudResourceGet) (*domain.CloudResource, error)
		CloudResourceGetByCustomer(ctx context.Context, query queries.CloudResourceGetByCustomer) ([]domain.CloudResource, error)
		CloudResourceGetByName(ctx context.Context, query queries.CloudResourceGetByName) (*domain.CloudResource, error)
		NotificationGetByUser(ctx context.Context, userID string) ([]*domain.Notification, error)
	}

	Application struct {
		appCommands
		appQueries
	}

	appCommands struct {
		commands.CustomerCreateHandler
		commands.CustomerCreateCloudResourcesHandler
		commands.CloudResourceCreateHandler
		commands.CloudResourceUpdateHandler
		commands.CloudResourceDeleteHandler
		commands.NotificationCreateHandler
		commands.NotificationDeleteHandler
		commands.NotificationDeleteByUserHandler
	}

	appQueries struct {
		queries.CloudResourceGetHandler
		queries.CloudResourceGetByCustomerHandler
		queries.CloudResourceGetByNameHandler
		queries.NotificationGetByUserHandler
	}
)

var _ App = (*Application)(nil)

func New(
	customers domain.CustomerRepository,
	resources domain.CloudResourceRepository,
	notifications domain.NotificationRepository,
	notificationReceiver domain.NotificationReceiver,
) *Application {
	return &Application{
		appCommands: appCommands{
			CustomerCreateHandler:               commands.NewCustomerCreateHandler(customers, notificationReceiver),
			CloudResourceCreateHandler:          commands.NewCloudResourceCreateHandler(resources, notificationReceiver),
			CloudResourceUpdateHandler:          commands.NewCloudResourceUpdateHandler(resources, customers, notificationReceiver),
			CloudResourceDeleteHandler:          commands.NewCloudResourceDeleteHandler(resources),
			CustomerCreateCloudResourcesHandler: commands.NewCustomerCreateCloudResourcesHandler(resources),
			NotificationCreateHandler:           commands.NewNotificationCreateHandler(notifications),
			NotificationDeleteHandler:           commands.NewNotificationDeleteHandler(notifications),
			NotificationDeleteByUserHandler:     commands.NewNotificationDeleteByUserHandler(notifications),
		},
		appQueries: appQueries{
			CloudResourceGetHandler:           queries.NewCloudResourceGetHandler(resources),
			CloudResourceGetByNameHandler:     queries.NewCloudResourceGetByNameHandler(resources),
			CloudResourceGetByCustomerHandler: queries.NewCloudResourceGetByCustomerHandler(resources),
			NotificationGetByUserHandler:      queries.NewNotificationGetByUserHandler(notifications),
		},
	}
}
