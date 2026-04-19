package v1

import (
	"github.com/twenty99h/wishes-service/internal/usecase"
)

type Handlers struct {
	usecase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{usecase: uc}
}
