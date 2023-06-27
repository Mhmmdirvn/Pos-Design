package register

type UseCase struct {
	Repo Repository
}

func (usecase UseCase) Register(register *Admin) error {
	err := usecase.Repo.Register(register)
	return err
}