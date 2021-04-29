package post

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/org/user"
)

type ModePost string

const (
	ModeFinder     = ModePost("mode_finder")
	ModeSeller     = ModePost("mode_seller")
	ModeExperience = ModePost("mode_experience")
)

type Post struct {
	mgo.BaseModel `bson:",inline"`
	UserID        string   `json:"userid" bson:"userid"`
	Title         string   `json:"title" bson:"title"`
	Text          string   `json:"text" bson:"text"`
	Images        []string `json:"images" bson:"images"`
	Videos        []string `json:"videos" bson:"videos"`
	// Tags          []Tag         `json:"tags" bson:"tags"`
	Like     int           `json:"like" bson:"like"`
	Comment  int           `json:"comment" bson:"comment"`
	Share    int           `json:"share" bson:"share"`
	Price    string        `json:"price" bson:"price"`
	Type     ModePost      `json:"type" bson:"type"`
	Name     string        `json:"name" bson:"name"`
	RankName string        `json:"rankname" bson:"rankname"`
	Job      string        `json:"job" bson:"job"`
	Avatar   string        `json:"avatar" bson:"avatar"`
	Location user.Location `json:"location" bson:"location"`
	Deal     bool          `json:"deal" bson:"deal"`
	Category string        `json:"category" bson:"category"`
}

type Tag struct {
	Name   string `json:"name" bson:"name"`
	UserID string `json:"userid" bson:"userid"`
}

func (s *Post) GetImages() []string {
	return s.Images
}

func (s *Post) GetVideos() []string {
	return s.Videos
}

func (s *Post) GetText() string {
	return s.Text
}

func (s *Post) GetLike() int {
	return s.Like
}

func (s *Post) GetComment() int {
	return s.Comment
}

func (s *Post) GetShare() int {
	return s.Share
}

func (s *Post) GetPrice() string {
	return s.Price
}

func (s *Post) GetType() ModePost {
	return s.Type
}

func (s *Post) GetName() string {
	return s.Name
}

func (s *Post) GetRankName() string {
	return s.RankName
}

func (s *Post) GetJob() string {
	return s.Job
}

func (s *Post) GetAvatar() string {
	return s.Avatar
}

func (s *Post) GetDeal() bool {
	return s.Deal
}

func (s *Post) GetCategory() string {
	return s.Category
}
