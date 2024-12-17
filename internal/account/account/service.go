package account

type Service struct {
	DAO *DAO
}

func NewService(dao *DAO) *Service {
	return &Service{
		DAO: dao,
	}
}
