package post

import (
	"context"
	"errors"
	"time"
)

type Service struct {
	DAO *DAO
}

func NewService(dao *DAO) *Service {
	return &Service{
		DAO: dao,
	}
}

func (s *Service) CreatePost(ctx context.Context, Id string, req BlogPost) error {

	dto := BlogPostDao{
		AccountId:   Id,
		ID:          req.ID,
		Title:       req.Title,
		CoverImage:  req.CoverImage,
		Content:     req.Content,
		Tags:        req.Tags,
		AuthorID:    req.AuthorID,
		Author:      req.Author,
		PublishedAt: req.PublishedAt,
		ReadTime:    req.ReadTime,
		Visibility:  req.Visibility,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		CreatedBy:   Id,
	}

	err := s.DAO.CreatePost(ctx, dto)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdatePost(ctx context.Context, Id string, req UpdateBlogPost, PostId string) error {

	post, err := s.DAO.GetPostById(ctx, Id, PostId)
	if err != nil {
		return errors.New("post not found")
	}

	dto := BlogPostDao{
		AccountId:   Id,
		ID:          PostId,
		Title:       req.Title,
		CoverImage:  req.CoverImage,
		Content:     req.Content,
		Tags:        req.Tags,
		AuthorID:    req.AuthorID,
		Author:      req.Author,
		PublishedAt: req.PublishedAt,
		ReadTime:    req.ReadTime,
		Visibility:  req.Visibility,
		Status:      req.Status,
		CreatedAt:   post.CreatedAt,
		CreatedBy:   post.CreatedBy,
		UpdatedAt:   time.Now(),
		UpdatedBy:   Id,
	}

	err = s.DAO.CreatePost(ctx, dto)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetPostById(ctx context.Context, Id, PostId string) (*GetPost, error) {

	post, err := s.DAO.GetPostById(ctx, Id, PostId)
	if err != nil {
		return nil, errors.New("post not found")
	}

	return post, nil

}

func (s *Service) GetAllPost(ctx context.Context, Id string) ([]GetPost, int64, error) {

	post, count, err := s.DAO.GetAllPost(ctx, Id)
	if err != nil {
		return nil, 0, errors.New("post not found")
	}

	return post, count, nil

}

func (s *Service) DeletePost(ctx context.Context, Id, PostId string) error {

	err := s.DAO.DeletePost(ctx, Id, PostId)
	if err != nil {
		return errors.New("post not found")
	}

	return nil

}
