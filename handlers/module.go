package handlers

import "github.com/brandenc40/wuphf.com/controllers"

type Handlers struct {
	*controllers.Controllers
}

const successMessage = "OK"

func New() *Handlers {
	return &Handlers{
		controllers.New(),
	}
}
