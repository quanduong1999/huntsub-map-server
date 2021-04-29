package video

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableVideo = model.NewTable("huntsub-server", "videos", "vd")

func NewVideoID() string {
	return TableVideo.Next()
}

func (s *Video) MakeID() string {
	return TableVideo.IdMaker.Next()
}

func (b *Video) Create() error {
	return TableVideo.Create(b)
}

func MarkDelete(id string) error {
	return TableVideo.MarkDelete(id)
}

func (v *Video) Update(newValue *Video) (*Video, error) {
	var values = map[string]interface{}{}

	if reflect.DeepEqual(newValue.GetVideos(), v.GetVideos()) {
		if len(newValue.GetVideos()) > 0 {
			values["videos"] = newValue.GetVideos()
		}
	}

	err := TableVideo.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableVideo.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})
