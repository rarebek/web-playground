package main

import (
	"database/sql"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rarebek/web-playground/config"
	_ "github.com/rarebek/web-playground/docs" // Import generated docs
	"github.com/rarebek/web-playground/handlers"
	"github.com/rarebek/web-playground/repo"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title           Web Playground API
// @version         1.0
// @description     A beautiful API documentation for Web Playground
// @contact.name   Nodirbek
// @contact.email  nomonovn2@gmail.com
// @host      localhost:6655
// @BasePath  /api/v1
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err.Error())
	}
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cfg.PgUsername, cfg.PgPass, cfg.PgHost, cfg.PgPort, cfg.PgDb)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(dsn)

	newRepo := repo.NewRepo(conn)
	newHandlers := handlers.NewHandlers(newRepo)

	defer func(apiClient *client.Client) {
		err := apiClient.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}(apiClient)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
	}))

	routerGroup := router.Group("/api/v1")

	// API routes
	routerGroup.POST("/register", newHandlers.Register)

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err = router.Run(":6655")
	if err != nil {
		log.Fatal(err.Error())
	}
}
