package tests

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/kidus-tiliksew/aqua-crims/application"
	"github.com/kidus-tiliksew/aqua-crims/application/commands"
	"github.com/kidus-tiliksew/aqua-crims/domain"
	"github.com/kidus-tiliksew/aqua-crims/postgres"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type notificationReceiver struct {
}

func (n *notificationReceiver) SendStructuredMessage(userID, message string) error {
	return nil
}

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "user",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "aqua-crims",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	postgresContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("failed to start container: %v", err)
	}
	defer postgresContainer.Terminate(ctx)

	host, err := postgresContainer.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get container host: %v", err)
	}

	port, err := postgresContainer.MappedPort(ctx, "5432")
	if err != nil {
		log.Fatalf("failed to get container port: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=user password=password dbname=aqua-crims sslmode=disable", host, port.Port())
	db, err = gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&postgres.Customer{}, &postgres.CloudResource{}, &postgres.Notification{})

	code := m.Run()
	os.Exit(code)
}

func TestIntegration_CreateCustomer(t *testing.T) {
	repo := postgres.NewCustomerRepository(db)
	app := application.New(repo, nil, nil, &notificationReceiver{})

	_, err := app.CustomerCreate(context.Background(), commands.CustomerCreate{
		Name:  "Test Customer",
		Email: "kidus.tiliksew@gmail.com",
	})
	assert.NoError(t, err)

	var result domain.Customer
	err = db.First(&result, "name = ?", "Test Customer").Error
	assert.NoError(t, err)
	assert.Equal(t, "Test Customer", result.Name)
}

func TestIntegration_CreateCloudResource(t *testing.T) {
	customerRepo := postgres.NewCustomerRepository(db)
	resourceRepo := postgres.NewCloudResourceRepository(db)
	notificationRepo := postgres.NewNotificationRepository(db)
	ns := &notificationReceiver{}

	app := application.New(customerRepo, resourceRepo, notificationRepo, ns)

	// Create a customer first
	customer, err := app.CustomerCreate(context.Background(), commands.CustomerCreate{
		Name:  "Test Customer",
		Email: "kidus@gmail.com",
	})
	assert.NoError(t, err)

	// Create a cloud resource for the customer
	_, err = app.CloudResourceCreate(context.Background(), commands.CloudResourceCreate{
		CustomerID: customer.ID,
		Name:       "Test Resource",
		Type:       "S3",
		Region:     "us-west-2",
	})
	assert.NoError(t, err)

	var result domain.CloudResource
	err = db.First(&result, "name = ?", "Test Resource").Error
	assert.NoError(t, err)
	assert.Equal(t, "Test Resource", result.Name)
	assert.Equal(t, customer.ID, result.CustomerID)
}
