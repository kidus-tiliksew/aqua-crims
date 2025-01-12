package main

import (
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"github.com/kidus-tiliksew/aqua-crims/application"
	grpcservice "github.com/kidus-tiliksew/aqua-crims/grpc"
	"github.com/kidus-tiliksew/aqua-crims/grpc/proto"
	"github.com/kidus-tiliksew/aqua-crims/notification"
	"github.com/kidus-tiliksew/aqua-crims/postgres"
	"google.golang.org/grpc"
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
	db.AutoMigrate(&postgres.Notification{})

	customers := postgres.NewCustomerRepository(db)
	resources := postgres.NewCloudResourceRepository(db)
	notifications := postgres.NewNotificationRepository(db)

	// initialize notification service
	ampDsn := os.Getenv("AMPQ_DSN")
	if dsn == "" {
		log.Panic("AMPQ_DSN environment variable is not set")
	}
	ns, err := notification.NewNotificationService(notifications, ampDsn, NotificationExchange)
	if err != nil {
		log.Panicf("failed to initialize notification service: %v", err)
	}
	go ns.Subscribe()

	// initialize application
	app := application.New(customers, resources, notifications, ns)

	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "9090"
	}
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterNotificationServiceServer(grpcServer, grpcservice.NewNotificationGRPCServer(app))
	log.Printf("gRPC server listening on %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
