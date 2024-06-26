package main

import (
	"log"
	"net/http"

	"go_backend/api"
	"go_backend/repositories"
	"go_backend/services"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// 連接數據庫
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdbname port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// 初始化存儲庫和服務
	projectRepo := &repositories.ProjectRepository{DB: db}
	projectService := services.NewProjectService(projectRepo)
	projectHandler := api.NewProjectHandler(projectService)

	// 設置路由
	r := mux.NewRouter()
	r.HandleFunc("/api/projects", projectHandler.GetAllProjects).Methods("GET")
	r.HandleFunc("/api/projects/{id}", projectHandler.GetProject).Methods("GET")
	r.HandleFunc("/api/projects", projectHandler.CreateProject).Methods("POST")
	r.HandleFunc("/api/projects/{id}", projectHandler.UpdateProject).Methods("PUT")
	r.HandleFunc("/api/projects/{id}", projectHandler.DeleteProject).Methods("DELETE")

	// 使用 CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	})

	handler := c.Handler(r)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
