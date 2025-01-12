package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kidus-tiliksew/aqua-crims/application"
	"github.com/kidus-tiliksew/aqua-crims/controllers"
	"github.com/kidus-tiliksew/aqua-crims/domain"
	memorydb "github.com/kidus-tiliksew/aqua-crims/memory_db"
	"github.com/stretchr/testify/assert"
)

func setupNotificationRouter(app application.App) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	notificationController := controllers.NewNotificationController(app)
	router.DELETE("/notifications/:id", notificationController.DeleteNotification)
	router.DELETE("/customers/:id/notifications", notificationController.DeleteNotificationByUser)
	router.GET("/customers/:id/notifications", notificationController.NotificationGetByUser)
	return router
}

func TestNotificationController_DeleteNotification(t *testing.T) {
	repo := &memorydb.InMemoryNotificationRepository{
		Notifications: []domain.Notification{
			{ID: 1, UserID: "user1", Message: "Test Notification"},
		},
	}
	app := application.New(nil, nil, repo, nil)
	router := setupNotificationRouter(app)

	req, _ := http.NewRequest("DELETE", "/notifications/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
	assert.Len(t, repo.Notifications, 0)
}

func TestNotificationController_DeleteNotificationByUser(t *testing.T) {
	repo := &memorydb.InMemoryNotificationRepository{
		Notifications: []domain.Notification{
			{ID: 1, UserID: "user1", Message: "Test Notification 1"},
			{ID: 2, UserID: "user1", Message: "Test Notification 2"},
		},
	}
	app := application.New(nil, nil, repo, nil)
	router := setupNotificationRouter(app)

	req, _ := http.NewRequest("DELETE", "/customers/user1/notifications", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)
	assert.Len(t, repo.Notifications, 0)
}

func TestNotificationController_NotificationGetByUser(t *testing.T) {
	repo := &memorydb.InMemoryNotificationRepository{
		Notifications: []domain.Notification{
			{ID: 1, UserID: "user1", Message: "Test Notification 1"},
			{ID: 2, UserID: "user1", Message: "Test Notification 2"},
		},
	}
	app := application.New(nil, nil, repo, nil)
	router := setupNotificationRouter(app)

	req, _ := http.NewRequest("GET", "/customers/user1/notifications", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var notifications []domain.Notification
	err := json.Unmarshal(resp.Body.Bytes(), &notifications)
	assert.NoError(t, err)
	assert.Len(t, notifications, 2)
	assert.Equal(t, "Test Notification 1", notifications[0].Message)
	assert.Equal(t, "Test Notification 2", notifications[1].Message)
}
