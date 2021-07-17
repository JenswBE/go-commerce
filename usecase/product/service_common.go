package product

type Service struct {
	db           DatabaseRepository
	imageStorage StorageRepository
}

func NewService(db DatabaseRepository, imageStorage StorageRepository) *Service {
	return &Service{
		db:           db,
		imageStorage: imageStorage,
	}
}
