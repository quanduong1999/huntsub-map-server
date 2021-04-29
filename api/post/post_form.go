package post

import (
	"huntsub/huntsub-map-server/o/org/action/like"
	"huntsub/huntsub-map-server/o/org/post"
	"huntsub/huntsub-map-server/o/org/user"
)

type PostForm struct {
	ID       string        `json:"id"`
	MTime    int64         `json:"mtime"`
	UserID   string        `json:"userid"`
	Title    string        `json:"title"`
	Text     string        `json:"text" `
	Images   []string      `json:"images"`
	Videos   []string      `json:"videos"`
	Tags     []post.Tag    `json:"tags"`
	Like     int           `json:"like"`
	Comment  int           `json:"comment"`
	Share    int           `json:"share"`
	Price    string        `json:"price"`
	Type     post.ModePost `json:"type"`
	Name     string        `json:"name"`
	RankName string        `json:"rankname"`
	Job      string        `json:"job"`
	Avatar   string        `json:"avatar"`
	Location user.Location `json:"location"`
	Deal     bool          `json:"deal"`
	Heart    bool          `json:"heart"`
}

func NewPostForm(p *post.Post, userid string) *PostForm {
	var s = &PostForm{}
	s.ID = p.ID
	s.MTime = p.MTime
	s.UserID = p.UserID
	s.Title = p.Title
	s.Text = p.Text
	s.Images = p.Images
	s.Videos = p.Videos
	s.Like = p.Like
	s.Comment = p.Comment
	s.Share = p.Share
	s.Price = p.Price
	s.Type = p.Type
	s.Name = p.Name
	s.RankName = p.RankName
	s.Job = p.Job
	s.Avatar = p.Avatar
	s.Location = p.Location
	s.Deal = p.Deal
	_, err := like.GetLike(map[string]interface{}{
		"dtime":  0,
		"postid": p.ID,
		"userid": userid,
	})
	if err != nil {
		s.Heart = false
	} else {
		s.Heart = true
	}

	return s
}
