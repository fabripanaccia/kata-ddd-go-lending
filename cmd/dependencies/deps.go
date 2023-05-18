package dependencies

import (
	"context"
)

type Deps struct {
}

const (
	service    = "service"
	hoursInDay = 24
)

func NewDeps(ctx context.Context) *Deps {
	return &Deps{}
}
