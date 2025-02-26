package account

import (
	"context"
	"fmt"
	"time"

	"github.com/blog/poc/pkg/middleware"
	"github.com/blog/poc/pkg/utils"
)

type Service struct {
	DAO *DAO
}

func NewService(dao *DAO) *Service {
	return &Service{
		DAO: dao,
	}
}

func (s *Service) Signup(ctx context.Context, req SignUp) (*SignUpRes, error) {

	firstname, lastname := utils.ExtractNameFromEmail(req.Email)

	fullname := fmt.Sprintf("%s %s", firstname, lastname)

	user := AccountDao{
		Email:      req.Email,
		Id:         req.Id,
		FullName:   fullname,
		Password:   middleware.CreatePasswordHash(req.Password),
		CreatedAt:  time.Now(),
		CreatedBy:  req.Id,
		JoinedAt:   time.Now(),
		IsVerified: true,
	}

	err := s.DAO.CreateAccount(ctx, user)
	if err != nil {
		return nil, err
	}

	payload := middleware.TokenPayload{
		Id:     user.Id,
		Scopes: []string{"all"},
	}
	// go func() {
	// 	_, err = utils.SendMail(req.Email)
	// 	if err != nil {
	// 		return
	// 	}
	// }()

	accessToken, refreshToken, err := middleware.GenerateJWTTokens(payload)
	if err != nil {
		return nil, err
	}
	return &SignUpRes{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Id:           user.Id,
	}, nil

}

func (s *Service) Login(ctx context.Context, req Login) (*LoginResponse, error) {

	user, err := s.DAO.FindAccount(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	verify := middleware.VerifyPasswordHash(req.Password, user.Password)

	if !verify {
		return nil, fmt.Errorf("Invaild email or password", err)
	}

	payload := middleware.TokenPayload{
		Id:     user.Id,
		Scopes: []string{"all"},
	}

	accessToken, refreshToken, err := middleware.GenerateJWTTokens(payload)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Id:           user.Id,
	}, nil
}

func (s *Service) GetAccount(ctx context.Context, Id string) (*Account, error) {

	data, err := s.DAO.GetAccountById(ctx, Id)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (s *Service) UpdateAccount(ctx context.Context, Id string, req UpdateAccount) error {

	data, err := s.DAO.GetAccountForUpdate(ctx, Id)
	if err != nil {
		return err
	}

	dto := AccountDao{
		Id:             Id,
		Email:          req.Email,
		Username:       req.Username,
		FullName:       req.FullName,
		Password:       data.Password,
		Bio:            req.Bio,
		ProfileImage:   req.ProfileImage,
		CoverImage:     req.CoverImage,
		Website:        req.Website,
		Location:       req.Location,
		SocialLinks:    req.SocialLinks,
		FollowerCount:  data.FollowerCount,
		FollowingCount: data.FollowingCount,
		JoinedAt:       data.JoinedAt,
		IsVerified:     data.IsVerified,
		IsPremium:      data.IsPremium,
		LastActive:     data.LastActive,
		CreatedAt:      data.CreatedAt,
		CreatedBy:      data.CreatedBy,
		UpdatedAt:      time.Now(),
		UpdatedBy:      Id,
	}

	err = s.DAO.CreateAccount(ctx, dto)
	if err != nil {
		return err
	}

	return nil

}
