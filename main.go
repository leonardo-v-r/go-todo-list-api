package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/leonardo-v-r/golang-to-do-list-api/handler"
	"github.com/leonardo-v-r/golang-to-do-list-api/service"
	storage "github.com/leonardo-v-r/golang-to-do-list-api/storage/sqlite"
	"github.com/leonardo-v-r/golang-to-do-list-api/storage/sqlite/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	db, err := storage.New(os.Getenv("DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Migrate(); err != nil {
		log.Fatal(err)
	}

	taskRepo := repository.NewTaskRepository(db.DB)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	router := handler.NewRouter(*taskHandler)
	router.Serve(":5000")
}
