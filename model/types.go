package model

import "Transformer/services/transformer"

type ProjectedData struct {
	Data []transformer.DataObject
	Proj []transformer.ProjectorItem
}
