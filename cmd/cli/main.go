package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kidus-tiliksew/aqua-crims/domain"
	"github.com/kidus-tiliksew/aqua-crims/postgres"
	"github.com/spf13/cobra"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "CLI for aqua-crims",
}

var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed cloud resources",
	Run: func(cmd *cobra.Command, args []string) {
		dsn := "host=0.0.0.0 user=postgres password=password dbname=aqua-crims"
		db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}
		resourceRepo := postgres.NewCloudResourceRepository(db)
		err = resourceRepo.Create(context.Background(), &domain.CloudResource{
			Name:   "S3 Bucket",
			Region: "us-west-2",
		})
		if err != nil {
			log.Fatalf("failed to seed resources: %v", err)
		}
		fmt.Println("Successfully seeded cloud resources.")
	},
}

func init() {
	rootCmd.AddCommand(SeedCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("command failed: %v", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("command failed: %v", err)
	}
}
