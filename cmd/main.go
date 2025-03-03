package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rarebek/web-playground/config"
	"github.com/rarebek/web-playground/handlers"
	"github.com/rarebek/web-playground/repo"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	config, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.PG_USERNAME, config.PG_PASS, config.PG_HOST, config.PG_PORT, config.PG_DB)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := repo.NewRepo(conn)

	handlers := handlers.NewHandlers(repo)

	defer apiClient.Close()

	router := gin.Default()

	router.POST("/user", handlers.Register)

	router.Run(":6655")

	// containers, err := apiClient.ContainerList(context.Background(), container.ListOptions{All: true})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// for _, ctr := range containers {
	// 	fmt.Printf("%s %s (status: %s)\n", ctr.ID, ctr.Image, ctr.Status)
	// }
}
