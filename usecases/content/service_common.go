package content

var _ Usecase = &Service{}

type Service struct {
	db DatabaseRepository
}

func NewService(db DatabaseRepository) *Service {
	return &Service{
		db: db,
	}
}
