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

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestCloudResourceController_CloudResourceCreate(t *testing.T) {
	repo := &memorydb.InMemoryCloudResourceRepository{}
	ns := &memorydb.InMemoryNotificationReceiver{}
	app := application.New(nil, repo, nil, ns)
	controller := controllers.NewCloudResourceController(app)

	router := setupRouter()
	router.POST("/cloud-resources", controller.CloudResourceCreate)

	payload := commands.CloudResourceCreate{
		CustomerID: 1,
		Name:       "Test Resource",
		Type:       "S3",
		Region:     "us-west-2",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/cloud-resources", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var resource domain.CloudResource
	err := json.Unmarshal(resp.Body.Bytes(), &resource)
	assert.NoError(t, err)
	assert.Equal(t, payload.Name, resource.Name)
	assert.Equal(t, payload.Type, resource.Type)
	assert.Equal(t, payload.Region, resource.Region)
	assert.Equal(t, payload.CustomerID, resource.CustomerID)
}

func TestCloudResourceController_CloudResourceUpdate(t *testing.T) {
	repo := &memorydb.InMemoryCloudResourceRepository{
		Resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource", Type: "S3", Region: "us-west-2"},
		},
	}
	customers := &memorydb.InMemoryCustomerRepository{
		Customers: []domain.Customer{
			{ID: 1, Name: "Test Customer"},
		},
	}
	ns := &memorydb.InMemoryNotificationReceiver{}
	app := application.New(customers, repo, nil, ns)
	controller := controllers.NewCloudResourceController(app)

	router := setupRouter()
	router.PUT("/cloud-resources/:id", controller.CloudResourceUpdate)

	payload := commands.CloudResourceUpdate{
		ID:         1,
		CustomerID: 1,
		Name:       "Updated Resource",
		Type:       "S3",
		Region:     "us-west-2",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("PUT", "/cloud-resources/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var resource domain.CloudResource
	err := json.Unmarshal(resp.Body.Bytes(), &resource)
	assert.NoError(t, err)
	assert.Equal(t, payload.Name, resource.Name)
}

func TestCloudResourceController_CloudResourceDelete(t *testing.T) {
	repo := &memorydb.InMemoryCloudResourceRepository{
		Resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource", Type: "S3", Region: "us-west-2"},
		},
	}
	ns := &memorydb.InMemoryNotificationReceiver{}
	app := application.New(nil, repo, nil, ns)
	controller := controllers.NewCloudResourceController(app)

	router := setupRouter()
	router.DELETE("/cloud-resources/:id", controller.CloudResourceDelete)

	req, _ := http.NewRequest("DELETE", "/cloud-resources/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Len(t, repo.Resources, 0)
}

func TestCloudResourceController_CloudResourceGet(t *testing.T) {
	repo := &memorydb.InMemoryCloudResourceRepository{
		Resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource", Type: "S3", Region: "us-west-2"},
		},
	}
	ns := &memorydb.InMemoryNotificationReceiver{}
	app := application.New(nil, repo, nil, ns)
	controller := controllers.NewCloudResourceController(app)

	router := setupRouter()
	router.GET("/cloud-resources/:id", controller.CloudResourceGet)

	req, _ := http.NewRequest("GET", "/cloud-resources/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var resource domain.CloudResource
	err := json.Unmarshal(resp.Body.Bytes(), &resource)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), resource.ID)
}

func TestCloudResourceController_CloudResourceFindByCustomer(t *testing.T) {
	repo := &memorydb.InMemoryCloudResourceRepository{
		Resources: []domain.CloudResource{
			{ID: 1, CustomerID: 1, Name: "Test Resource 1", Type: "S3", Region: "us-west-2"},
			{ID: 2, CustomerID: 1, Name: "Test Resource 2", Type: "EC2", Region: "us-east-1"},
		},
	}
	ns := &memorydb.InMemoryNotificationReceiver{}
	app := application.New(nil, repo, nil, ns)
	controller := controllers.NewCloudResourceController(app)

	router := setupRouter()
	router.GET("/customers/:id/cloud-resources", controller.CloudResourceFindByCustomer)

	req, _ := http.NewRequest("GET", "/customers/1/cloud-resources", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var resources []domain.CloudResource
	err := json.Unmarshal(resp.Body.Bytes(), &resources)
	assert.NoError(t, err)
	assert.Len(t, resources, 2)
}
