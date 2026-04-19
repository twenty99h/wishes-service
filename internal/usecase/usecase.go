package usecase

type Postgres interface {
}

type UseCase struct {
	postgres Postgres
}

func New(postgres Postgres) *UseCase {
	return &UseCase{postgres: postgres}
}
