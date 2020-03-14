package handlers

import (
	"github.com/brandenc40/wuphf.com/common"
	"github.com/brandenc40/wuphf.com/controllers"
)

type Handlers struct {
	context *common.AppContext

	controllers *controllers.Controllers
}

const successMessage = "OK"

func New(appContext *common.AppContext) *Handlers {
	return &Handlers{
		context:     appContext,
		controllers: controllers.New(),
	}
}
