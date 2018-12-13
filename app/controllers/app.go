package controllers

import (
	"Transformer/model"
	"Transformer/services/transformer"
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Transform(body model.ProjectedData) revel.Result {
	result := make([]transformer.DataObject, len(body.Data))
	for i, item := range body.Data {
		result[i] = transformer.Project(item, body.Proj)
	}
	return c.RenderJSON(result)
}
