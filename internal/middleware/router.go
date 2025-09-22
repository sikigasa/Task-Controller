package middleware

import (
	"net/http"

	connect "github.com/sikigasa/task-controller/gen/protov1connect"
)

// HTTPルーティングを管理するインターフェース
type Router interface {
	SetupRoutes() *http.ServeMux
}

type router struct {
	taskService connect.TaskServiceHandler
	tagService  connect.TagServiceHandler
}

// 新しいRouterインスタンスを作成
func NewRouter(taskService connect.TaskServiceHandler, tagService connect.TagServiceHandler) Router {
	return &router{
		taskService: taskService,
		tagService:  tagService,
	}
}

// Connect RPCサービスのルートを設定
func (r *router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	taskPath, taskHandler := connect.NewTaskServiceHandler(r.taskService)
	mux.Handle(taskPath, taskHandler)

	tagPath, tagHandler := connect.NewTagServiceHandler(r.tagService)
	mux.Handle(tagPath, tagHandler)

	return mux
}
