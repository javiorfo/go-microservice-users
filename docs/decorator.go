package docs

import (
	"github.com/javiorfo/go-microservice-users/config"
	"github.com/swaggo/swag"
)

type SwaggerInfoWrapper struct {
	swag.Spec
}

func (i *SwaggerInfoWrapper) ReadDoc() string {
	i.BasePath = config.AppContextPath
	i.Title = config.AppName
	return i.Spec.ReadDoc()
}
