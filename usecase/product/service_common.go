package product

type Service struct {
	db      DatabaseRepository
	storage StorageRepository
}

func NewService(db DatabaseRepository, storage StorageRepository) *Service {
	return &Service{
		db:      db,
		storage: storage,
	}
}
