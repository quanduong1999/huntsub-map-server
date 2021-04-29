package photo

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TablePhoto = model.NewTable("huntsub-server", "photos", "pt")

func NewPhotoID() string {
	return TablePhoto.Next()
}

func (s *Photo) MakeID() string {
	return TablePhoto.IdMaker.Next()
}

func (b *Photo) Create() error {
	return TablePhoto.Create(b)
}

func MarkDelete(id string) error {
	return TablePhoto.MarkDelete(id)
}

func (v *Photo) Update(newValue *Photo) (*Photo, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetImages(), v.GetImages()) {
		if len(newValue.GetImages()) > 0 {
			values["images"] = newValue.GetImages()
		}
	}

	err := TablePhoto.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TablePhoto.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})
