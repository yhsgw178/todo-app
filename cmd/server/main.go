package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yhsgw178/todoapp/internal/config"
	"github.com/yhsgw178/todoapp/internal/database"
	"github.com/yhsgw178/todoapp/pkg/handler"
	"github.com/yhsgw178/todoapp/pkg/repository"
	"github.com/yhsgw178/todoapp/pkg/service"
)

func main() {
	cfg := config.LoadConfig()
	database.InitDB(cfg.DB)

	repo := repository.NewPostgresTaskRepository(database.DB)
	service := service.NewTaskService(repo)
	handler := handler.NewTaskHandler(service)

	r := mux.NewRouter()
	r.HandleFunc("/task", handler.CreateTask).Methods("POST")

	log.Println("Starting server on: 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
