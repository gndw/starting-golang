package httpmiddlewarelog

import (
	"github.com/gndw/starting-golang/internals/constants"
)

type Service interface {
	LogMiddleware(f constants.HttpFunction) constants.HttpFunction
}
