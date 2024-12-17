package account

import (
	"context"
	"time"

	"github.com/blog/poc/pkg/middleware"
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

	user := AccountDao{
		Email:     req.Email,
		UserId:    req.UserId,
		AccountId: req.AccountId,
		Password:  middleware.CreatePasswordHash(req.Password),
		CreatedAt: time.Now(),
		CreatedBy: req.UserId,
	}

	err := s.DAO.CreateAccount(ctx, user)
	if err != nil {
		return nil, err
	}

	payload := middleware.TokenPayload{
		UserID:    user.UserId,
		AccountID: user.AccountId,
		Scopes:    []string{"all"},
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
		UserId:       user.UserId,
		AccountId:    user.AccountId,
	}, nil

}
