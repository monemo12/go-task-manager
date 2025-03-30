package graphql

import (
	"task-manager/internal/domain"
)

// TaskResolver GraphQL 解析器
type TaskResolver struct {
	taskService domain.TaskService
}

// NewTaskResolver 創建新的 GraphQL 解析器
func NewTaskResolver(service domain.TaskService) *TaskResolver {
	return &TaskResolver{
		taskService: service,
	}
}

// 這裡將實現 GraphQL 查詢和變更
// 例如：
// type Query {
//   task(id: ID!): Task
//   tasks: [Task]
// }
//
// type Mutation {
//   createTask(input: CreateTaskInput!): Task
//   updateTask(id: ID!, input: UpdateTaskInput!): Task
//   deleteTask(id: ID!): Boolean
// }
