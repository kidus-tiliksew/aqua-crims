package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kidus-tiliksew/aqua-crims/application"
	"github.com/kidus-tiliksew/aqua-crims/controllers"
	"github.com/kidus-tiliksew/aqua-crims/rabbitmq"
	"github.com/kidus-tiliksew/aqua-crims/postgres"

	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const NotificationExchange = "notifications"

func main() {
	// load envs
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Panic("DATABASE_DSN environment variable is not set")
	}
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("failed to connect to database: %v", err)
	}

	// migrate schema
	db.AutoMigrate(&postgres.Customer{}, &postgres.CloudResource{}, &postgres.Notification{})

	// initialize repositories
	customers := postgres.NewCustomerRepository(db)
	resources := postgres.NewCloudResourceRepository(db)
	notifications := postgres.NewNotificationRepository(db)

	// initialize notification service
	ampDsn := os.Getenv("AMPQ_DSN")
	if dsn == "" {
		log.Panic("AMPQ_DSN environment variable is not set")
	}
	ns, err := rabbitmq.NewNotificationService(notifications, ampDsn, NotificationExchange)
	if err != nil {
		log.Panicf("failed to initialize notification service: %v", err)
	}

	// initialize application
	app := application.New(customers, resources, notifications, ns)

	// initialize controllers
	customerController := controllers.NewCustomerController(app)
	cloudResourceController := controllers.NewCloudResourceController(app)
	notificationController := controllers.NewNotificationController(app)


	r := gin.Default()
	r.POST("/customers", customerController.CustomerCreate)
	r.POST("/customers/:id/cloud-resources", customerController.CustomerCreateCloudResources)
	r.GET("/customers/:id/cloud-resources", cloudResourceController.CloudResourceFindByCustomer)
	r.GET("/customers/:id/notifications", notificationController.NotificationGetByUser)
	r.DELETE("/customers/:id/notifications", notificationController.DeleteNotificationByUser)

	r.POST("/cloud-resources", cloudResourceController.CloudResourceCreate)
	r.PUT("/cloud-resources/:id", cloudResourceController.CloudResourceUpdate)
	r.DELETE("/cloud-resources/:id", cloudResourceController.CloudResourceDelete)
	r.DELETE("/notifications/:id", notificationController.DeleteNotification)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
