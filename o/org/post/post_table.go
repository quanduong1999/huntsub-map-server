package post

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TablePost = model.NewTable("huntsub-server", "post", "po")

func NewPostID() string {
	return TablePost.Next()
}

func (s *Post) MakeID() string {
	return TablePost.IdMaker.Next()
}

func (b *Post) Create() error {
	var _ = TablePost.EnsureAddressIndex()
	return TablePost.Create(b)
}

func MarkDelete(id string) error {
	return TablePost.MarkDelete(id)
}

func (v *Post) Update(newValue *Post) (*Post, error) {
	var values = map[string]interface{}{}

	if !reflect.DeepEqual(newValue.GetImages(), v.GetImages()) {
		if len(newValue.GetImages()) > 0 {
			values["images"] = newValue.GetImages()
		}
	}
	if newValue.GetText() != v.GetText() && newValue.GetText() != "" {
		values["text"] = newValue.GetText()
	}
	if !reflect.DeepEqual(newValue.GetVideos(), v.GetVideos()) {
		if len(newValue.GetVideos()) > 0 {
			values["videos"] = newValue.GetVideos()
		}
	}
	if newValue.GetLike() != v.GetLike() {
		if newValue.GetLike() != 0 {
			values["like"] = newValue.GetLike()
		}
	}
	if newValue.GetComment() != v.GetComment() {
		if newValue.GetComment() != 0 {
			values["comment"] = newValue.GetComment()
		}
	}
	if newValue.GetShare() != v.GetShare() {
		if newValue.GetShare() != 0 {
			values["share"] = newValue.GetShare()
		}
	}
	if newValue.GetPrice() != v.GetPrice() {
		if newValue.GetPrice() != "" {
			values["price"] = newValue.GetPrice()
		}
	}
	if newValue.GetType() != v.GetType() {
		if newValue.GetType() != "" {
			values["type"] = newValue.GetType()
		}
	}
	if newValue.GetJob() != v.GetJob() {
		if newValue.GetJob() != "" {
			values["job"] = newValue.GetJob()
		}
	}
	if newValue.GetName() != v.GetName() {
		if newValue.GetName() != "" {
			values["name"] = newValue.GetName()
		}
	}
	if newValue.GetRankName() != v.GetRankName() {
		if newValue.GetRankName() != "" {
			values["rankname"] = newValue.GetRankName()
		}
	}
	if newValue.GetAvatar() != v.GetAvatar() {
		if newValue.GetAvatar() != "" {
			values["avatar"] = newValue.GetAvatar()
		}
	}

	if newValue.GetDeal() != v.GetDeal() {
		values["deal"] = newValue.GetDeal()
	}

	if newValue.GetCategory() != v.GetCategory() {
		if newValue.GetCategory() != "" {
			values["category"] = newValue.GetCategory()
		}
	}

	err := TablePost.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

var _ = TablePost.EnsureIndex(mgo.Index{
	Key:        []string{"ctime"},
	Background: true,
})
var _ = TablePost.EnsureAddressIndex()
