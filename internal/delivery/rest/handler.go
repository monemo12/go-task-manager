package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/monemo12/task-manager/internal/domain"
)

// TaskHandler 處理 RESTful API 請求
type TaskHandler struct {
	taskService domain.TaskService
}

// NewTaskHandler 創建新的 REST API 處理器
func NewTaskHandler(service domain.TaskService) *TaskHandler {
	return &TaskHandler{
		taskService: service,
	}
}

// 這裡將實現具體的 HTTP 處理方法
// 例如：
// - POST /tasks
// - GET /tasks
// - GET /tasks/:id
// - PUT /tasks/:id
// - DELETE /tasks/:id
// - GET /tasks

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := h.taskService.CreateTask(c.Request.Context(), task.Title, task.Description, task.Priority, task.DueDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTask)
}

func (h *TaskHandler) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := h.taskService.GetTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.taskService.UpdateTaskStatus(c.Request.Context(), id, task.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = h.taskService.UpdateTaskPriority(c.Request.Context(), id, task.Priority)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.taskService.DeleteTask(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func (h *TaskHandler) ListTasks(c *gin.Context) {
	tasks, err := h.taskService.ListTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// TestPathHandler 用於測試路徑處理
func (h *TaskHandler) TestPathHandler(c *gin.Context) {
	name := c.Param("name")
	path := c.Request.URL.Path
	rawPath := c.Request.URL.RawPath

	c.JSON(http.StatusOK, gin.H{
		"name":     name,    // 解析後的參數
		"path":     path,    // 當前路徑
		"raw_path": rawPath, // 原始路徑
	})
}
