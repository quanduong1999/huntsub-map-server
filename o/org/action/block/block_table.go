package block

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableBlock = model.NewTable("huntsub-server", "block", "blo")

func NewBlockID() string {
	return TableBlock.Next()
}

func (s *Block) MakeID() string {
	return TableBlock.IdMaker.Next()
}

func (b *Block) Create() error {
	return TableBlock.Create(b)
}

func MarkDelete(id string) error {
	return TableBlock.MarkDelete(id)
}

func (v *Block) Update(newValue *Block) (*Block, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetData(), v.GetData()) {
		if newValue.Data.Job != "" {
			values["data"] = newValue.GetData()
		}
	}

	err := TableBlock.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableBlock.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})
