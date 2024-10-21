package httpmiddleware

import (
	"github.com/gndw/starting-golang/internals/constants"
)

//go:generate mockery --name Service
type Service interface {
	LogMiddleware(f constants.HttpFunction) constants.HttpFunction
}
