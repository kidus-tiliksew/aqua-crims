package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kidus-tiliksew/aqua-crims/application"
	"github.com/kidus-tiliksew/aqua-crims/application/commands"
	"github.com/kidus-tiliksew/aqua-crims/controllers"
	"github.com/kidus-tiliksew/aqua-crims/domain"
	memorydb "github.com/kidus-tiliksew/aqua-crims/memory_db"
	"github.com/stretchr/testify/assert"
)

func setupCustomerRouter(app application.App) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	customerController := controllers.NewCustomerController(app)
	router.POST("/customers", customerController.CustomerCreate)
	router.POST("/customers/:id/cloud-resources", customerController.CustomerCreateCloudResources)
	return router
}

func TestCustomerController_CustomerCreate(t *testing.T) {
	customers := &memorydb.InMemoryCustomerRepository{}
	notificationReceiver := &memorydb.InMemoryNotificationReceiver{}

	app := application.New(customers, nil, nil, notificationReceiver)
	router := setupCustomerRouter(app)

	payload := commands.CustomerCreate{
		Name:  "Test Customer",
		Email: "test@example.com",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/customers", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var customer domain.Customer
	err := json.Unmarshal(resp.Body.Bytes(), &customer)
	assert.NoError(t, err)
	assert.Equal(t, payload.Name, customer.Name)
	assert.Equal(t, payload.Email, customer.Email)
}

func TestCustomerController_CustomerCreateCloudResources(t *testing.T) {
	customerRepo := &memorydb.InMemoryCustomerRepository{
		Customers: []domain.Customer{
			{ID: 1, Name: "Test Customer", Email: "test@example.com"},
		},
	}
	resourceRepo := &memorydb.InMemoryCloudResourceRepository{
		Resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource", Type: "S3", Region: "us-west-2"},
		},
	}
	notificationReceiver := &memorydb.InMemoryNotificationReceiver{}
	app := application.New(customerRepo, resourceRepo, nil, notificationReceiver)
	router := setupCustomerRouter(app)

	payload := commands.CustomerCreateCloudResources{
		CustomerID: 1,
		Names:      []string{"Test Resource"},
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/customers/1/cloud-resources", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var resource []*domain.CloudResource
	err := json.Unmarshal(resp.Body.Bytes(), &resource)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(resource))
}
