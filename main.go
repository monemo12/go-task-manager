package main

import (
	"log"

	"task-manager/internal/delivery/rest"
	"task-manager/internal/domain"
	repo_task_memory "task-manager/internal/repository/task/memory"
	service_task "task-manager/internal/service/task"
	"task-manager/pkg/validator"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化依賴
	repo := repo_task_memory.NewRepository()
	taskValidator := validator.NewTaskValidator()

	// TODO: 實現一個簡單的通知器
	notifier := &SimpleNotifier{}

	// 初始化服務
	taskService := service_task.NewService(repo, taskValidator, notifier)

	// 初始化 HTTP handler
	taskHandler := rest.NewTaskHandler(taskService)

	// 設置 Gin router
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.UseRawPath = true
	r.UnescapePathValues = false

	// 使用自定義的 JSON 序列化
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	// 註冊路由
	tasks := r.Group("/tasks")
	{
		tasks.GET("/", taskHandler.ListTasks)
		tasks.POST("/", taskHandler.CreateTask)
		tasks.GET("/:id", taskHandler.GetTask)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
	}

	// 啟動服務器
	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// SimpleNotifier 是一個簡單的通知實現
type SimpleNotifier struct{}

func (n *SimpleNotifier) NotifyTaskCreated(task *domain.Task) {
	log.Printf("Task created: %s\n", task.Title)
}

func (n *SimpleNotifier) NotifyTaskCompleted(task *domain.Task) {
	log.Printf("Task completed: %s\n", task.Title)
}

func (n *SimpleNotifier) NotifyTaskDueSoon(task *domain.Task) {
	log.Printf("Task due soon: %s\n", task.Title)
}
