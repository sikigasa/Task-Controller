package usecase

// import (
// 	"context"
// 	"testing"

// 	task "github.com/sikigasa/task-controller/proto/v1"
// 	"google.golang.org/protobuf/types/known/timestamppb"
// )

// func TestTask(t *testing.T) {
// 	t.Run("CreateTask", testCreateTask)
// }

// func testCreateTask(t *testing.T) {

// 	t.Run("正常系", func(t *testing.T) {

// 		req := &task.CreateTaskRequest{
// 			Title:       "test",
// 			Description: "test",
// 			LimitedAt:   timestamppb.Now(),
// 			TagIds:      []string{"tag1", "tag2"},
// 		}

// 		res, err := taskService.CreateTask(context.Background(), req)
// 		if err != nil {
// 			t.Errorf("expected no error, got %v", err)
// 		}

// 		if res == nil {
// 			t.Errorf("expected response, got nil")
// 		}
// 	})
// }
