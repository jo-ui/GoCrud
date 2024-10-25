package main

import (
	"fmt"
	"log"
	"time"

	"go_crud/docs"
	"go_crud/internal/domain"
	"go_crud/internal/handler"
	"go_crud/internal/repository"
	"go_crud/internal/routes"
	"go_crud/internal/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	files "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

// @title Person CRUD API
// @version 1.0
// @description This is a sample CRUD API for managing persons.
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// cross-site resource sharing
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // specify allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize in-memory SQLite database
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&domain.Person{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Initialize components
	personRepo := repository.NewPersonRepository(db)
	personUsecase := usecase.NewPersonUsecase(personRepo)
	personHandler := handler.NewPersonHandler(personUsecase)

	routes.RegisterRoutes(r, personHandler)

	// Swagger setup
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	fmt.Println("Server running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
