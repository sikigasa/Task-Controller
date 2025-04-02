package usecase

import (
	"github.com/sikigasa/task-controller/internal/infra"
	tag "github.com/sikigasa/task-controller/proto/v1"
)

type TagService struct {
	tag.UnimplementedTaskServiceServer
	tagRepo infra.TagRepo
}
