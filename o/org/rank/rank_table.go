package rank

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableRank = model.NewTable("huntsub-server", "rank", "ran")

func NewPostID() string {
	return TableRank.Next()
}

func (s *Rank) MakeID() string {
	return TableRank.IdMaker.Next()
}

func (b *Rank) Create() error {
	var _ = TableRank.EnsureAddressIndex()
	return TableRank.Create(b)
}

func MarkDelete(id string) error {
	return TableRank.MarkDelete(id)
}

func (v *Rank) Update(newValue *Rank) (*Rank, error) {
	var values = map[string]interface{}{}

	if newValue.GetRankName() != v.GetRankName() {
		if newValue.GetRankName() != "" {
			values["rankname"] = newValue.GetRankName()
		}
	}
	if newValue.GetRadius() != v.GetRadius() {
		if newValue.GetRadius() != 0 {
			values["radius"] = newValue.GetRadius()
		}
	}
	if newValue.GetPostNumber() != v.GetPostNumber() {
		if newValue.GetPostNumber() != 0 {
			values["postnumber"] = newValue.GetPostNumber()
		}
	}
	if newValue.GetExperienceNumber() != v.GetExperienceNumber() {
		if newValue.GetExperienceNumber() != 0 {
			values["experiencenumber"] = newValue.GetExperienceNumber()
		}
	}

	if !reflect.DeepEqual(newValue.GetLocation(), v.GetLocation()) {
		if newValue.Location.Type != "" {
			values["location"] = newValue.GetLocation()
		}
	}

	if newValue.GetLikeNumber() != v.GetLikeNumber() {
		if newValue.GetLikeNumber() != 0 {
			values["likenumber"] = newValue.GetLikeNumber()
		}
	}

	if newValue.GetCommentNumber() != v.GetCommentNumber() {
		if newValue.GetCommentNumber() != 0 {
			values["commentnumber"] = newValue.GetCommentNumber()
		}
	}

	if newValue.GetShareNumber() != v.GetShareNumber() {
		if newValue.GetShareNumber() != 0 {
			values["sharenumber"] = newValue.GetShareNumber()
		}
	}

	err := TableRank.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TableRank.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})
var _ = TableRank.EnsureAddressIndex()
