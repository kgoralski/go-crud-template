package domain

// Bank domain model
type Bank struct {
	ID   int    `json:"id" DB:"id"`
	Name string `json:"name" DB:"name"`
}

//Service instance to manage Banks
type Service struct {
	s Store
}

//NewService creates new Bank Service
func NewService(store Store) *Service {
	return &Service{s: store}
}

//GetBanks returns all banks
func (svc Service) GetBanks() ([]Bank, error) {
	return svc.s.getAll()
}

//GetBank returns single Bank
func (svc Service) GetBank(id int) (*Bank, error) {
	return svc.s.get(id)
}

//Create a Bank
func (svc Service) Create(bank Bank) (int, error) {
	return svc.s.create(bank)
}

//DeleteBanks all Banks
func (svc Service) DeleteBanks() error {
	return svc.s.deleteAll()
}

//Update single Bank
func (svc Service) Update(bank Bank) (*Bank, error) {
	return svc.s.update(bank)
}

//Delete single Bank
func (svc Service) Delete(id int) error {
	return svc.s.delete(id)
}
