package transformer

type ProjectorItem struct {
	Template string
	OutKey   string
}
type Projector []ProjectorItem

type DataObject map[string]string
