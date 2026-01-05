package main

import (
	"backend/controller"
	"backend/models"
	"backend/repository"
	"backend/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("banco.db"), &gorm.Config{})
	if err != nil {
		panic("erro ao se conectar com o banco de dados: " + err.Error())
	}

	db.AutoMigrate(&models.Task{})

	return db
}

func main() {
	db := initDb()
	server := gin.Default()

	server.Use(cors.Default())
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"}, AllowCredentials: true}))

	repository := repository.NewTaskRepository(db)
	service := service.NewTaskService(repository)
	controller := controller.NewTaskController(service)

	server.POST("/task", controller.CreateTask)
	server.GET("task", controller.ListTasks)
	server.GET("task/:id", controller.FindOneTask)
	server.PUT("task/:id", controller.EditTask)
	server.PUT("task/:id/status", controller.EditStatusTask)
	server.DELETE("task/:id", controller.DeleteOneTask)

	server.Run(":8888")
}
