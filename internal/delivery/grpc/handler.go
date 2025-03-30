package grpc

import (
	"task-manager/internal/domain"
)

// TaskServer 實現 gRPC 服務器
type TaskServer struct {
	taskService domain.TaskService
}

// NewTaskServer 創建新的 gRPC 服務器
func NewTaskServer(service domain.TaskService) *TaskServer {
	return &TaskServer{
		taskService: service,
	}
}

// 這裡將實現 protobuf 定義的服務方法
// 例如：
// service TaskService {
//   rpc CreateTask(CreateTaskRequest) returns (Task);
//   rpc GetTask(GetTaskRequest) returns (Task);
//   rpc ListTasks(ListTasksRequest) returns (ListTasksResponse);
//   rpc UpdateTask(UpdateTaskRequest) returns (Task);
//   rpc DeleteTask(DeleteTaskRequest) returns (DeleteTaskResponse);
// }
