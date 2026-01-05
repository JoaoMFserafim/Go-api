package controller

import (
	"backend/models"
	"backend/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	Service *service.TaskService
}

func NewTaskController(service *service.TaskService) *TaskController {
	return &TaskController{Service: service}
}

func (controller *TaskController) CreateTask(context *gin.Context) {
	var task models.Task

	if err := context.ShouldBindJSON(&task); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.Service.CreateTask(&task); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, task)
}

func (controller *TaskController) ListTasks(context *gin.Context) {
	var tasks []models.Task

	tasks, err := controller.Service.FindAllTasks()

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "not found the tasks"})
		return
	}

	context.JSON(http.StatusOK, tasks)
}

func (controller *TaskController) FindOneTask(context *gin.Context) {
	var task models.Task
	id := context.Param("id")
	idCnv, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id invalid"})
		return
	}
	task, err = controller.Service.FindOneTask(idCnv)
	context.JSON(http.StatusOK, task)
}

func (controller *TaskController) DeleteOneTask(context *gin.Context) {
	id := context.Param("id")
	idcvs, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadGateway, err.Error())
		return
	}
	controller.Service.DeleteOneTask(idcvs)

	context.JSON(http.StatusOK, "Deleted")
}

func (controller *TaskController) EditTask(context *gin.Context) {
	id := context.Param("id")
	idcvs, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input models.Task

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task, err := controller.Service.EditOne(idcvs, input)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	context.JSON(http.StatusOK, task)

}

func (controller *TaskController) EditStatusTask(context *gin.Context) {
	id := context.Param("id")
	idcvs, err := strconv.Atoi(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	controller.Service.EditStatusTask(idcvs)

	context.JSON(http.StatusOK, gin.H{"sucess": "status changed"})
}
