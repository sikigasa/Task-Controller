package usecase

import (
	"context"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	tag "github.com/sikigasa/task-controller/gen"
	tagConnect "github.com/sikigasa/task-controller/gen/protov1connect"
	"github.com/sikigasa/task-controller/internal/domain"
	"github.com/sikigasa/task-controller/internal/infra"
)

type TagService struct {
	tagConnect.UnimplementedTagServiceHandler
	tagRepo infra.TagRepo
}

func NewTagService(tagRepo infra.TagRepo) tagConnect.TagServiceClient {
	return &TagService{
		tagRepo: tagRepo,
	}
}

func (t *TagService) CreateTag(ctx context.Context, req *connect.Request[tag.CreateTagRequest]) (*connect.Response[tag.CreateTagResponse], error) {
	uuid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	param := domain.CreateTagParam{
		ID:   uuid.String(),
		Name: req.Msg.Name,
	}

	if err := t.tagRepo.CreateTag(ctx, param); err != nil {
		return nil, err
	}

	return connect.NewResponse(&tag.CreateTagResponse{
		Id: param.ID,
	}), nil
}

func (t *TagService) ListTag(ctx context.Context, req *connect.Request[tag.ListTagRequest]) (*connect.Response[tag.ListTagResponse], error) {
	param := domain.ListTagParam{
		Limit:  req.Msg.Limit,
		Offset: req.Msg.Offset,
	}

	tags, err := t.tagRepo.ListTag(ctx, param)
	if err != nil {
		return nil, err
	}

	var tagList []*tag.Tag
	for _, t := range tags {
		tagList = append(tagList, &tag.Tag{
			Id:   t.ID,
			Name: t.Name,
		})
	}

	return connect.NewResponse(&tag.ListTagResponse{
		Tags: tagList,
	}), nil
}

func (t *TagService) DeleteTag(ctx context.Context, req *connect.Request[tag.DeleteTagRequest]) (*connect.Response[tag.DeleteTagResponse], error) {
	param := domain.DeleteTagParam{
		ID: req.Msg.Id,
	}

	if err := t.tagRepo.DeleteTag(ctx, param); err != nil {
		return nil, err
	}

	return connect.NewResponse(&tag.DeleteTagResponse{
		Success: true,
	}), nil
}
