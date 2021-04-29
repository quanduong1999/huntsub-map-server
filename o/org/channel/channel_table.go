package channel

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableChannel = model.NewTable("huntsub-server", "channel", "ch")

func NewChannelID() string {
	return TableChannel.Next()
}

func (s *Channel) MakeID() string {
	return TableChannel.IdMaker.Next()
}

func (b *Channel) Create() error {
	var _ = TableChannel.EnsureAddressIndex()
	return TableChannel.Create(b)
}

func MarkDelete(id string) error {
	return TableChannel.MarkDelete(id)
}

func (v *Channel) Update(newValue *Channel) (*Channel, error) {
	var values = map[string]interface{}{}

	if newValue.GetJob() != v.GetJob() {
		if newValue.GetJob() != "" {
			values["job"] = newValue.GetJob()
		}
	}

	if newValue.GetIsActive() != v.GetIsActive() {
		values["isactive"] = newValue.GetIsActive()
	}

	if !reflect.DeepEqual(newValue.GetDescribeJob(), v.GetDescribeJob()) {
		if len(newValue.GetDescribeJob()) > 0 {
			values["describejob"] = newValue.GetDescribeJob()
		}
	}

	if newValue.GetIntroduction() != v.GetIntroduction() {
		if newValue.GetIntroduction() != "" {
			values["introduction"] = newValue.GetIntroduction()
		}
	}

	if !reflect.DeepEqual(newValue.GetLocation(), v.GetLocation()) {
		if newValue.Location.Type != "" {
			values["location"] = newValue.GetLocation()
		}
	}
	if !reflect.DeepEqual(newValue.GetWorkTime(), v.GetWorkTime()) {
		if newValue.WorkTime.Start != "" {
			values["worktime"] = newValue.GetWorkTime()
		}
	}

	err := TableChannel.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableChannel.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})

var _ = TableChannel.EnsureAddressIndex()
