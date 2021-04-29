package oev

import (
	"huntsub/huntsub-map-server/o/org/action/block"
	"huntsub/huntsub-map-server/o/org/action/follow"
	"huntsub/huntsub-map-server/o/org/action/like"
	"huntsub/huntsub-map-server/o/org/action/report"
	"huntsub/huntsub-map-server/o/org/action/share"
	"huntsub/huntsub-map-server/o/org/calendar"
	"huntsub/huntsub-map-server/o/org/channel"
	"huntsub/huntsub-map-server/o/org/chatroom"
	"huntsub/huntsub-map-server/o/org/comment"
	"huntsub/huntsub-map-server/o/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/o/org/feedback"
	"huntsub/huntsub-map-server/o/org/library/photo"
	"huntsub/huntsub-map-server/o/org/library/video"
	"huntsub/huntsub-map-server/o/org/meeting"
	"huntsub/huntsub-map-server/o/org/notification"
	"huntsub/huntsub-map-server/o/org/post"
	"huntsub/huntsub-map-server/o/org/user"
	userReport "huntsub/huntsub-map-server/o/report/user"
)

type ObjectCategoryName string

const (
	ObjectCategoryUser         = ObjectCategoryName("user")
	ObjectCategoryPost         = ObjectCategoryName("post")
	ObjectCategoryMeeting      = ObjectCategoryName("meeting")
	ObjectCategoryPhoto        = ObjectCategoryName("photo")
	ObjectCategoryVideo        = ObjectCategoryName("video")
	ObjectCategoryFeedback     = ObjectCategoryName("feedback")
	ObjectCategoryWHFeedback   = ObjectCategoryName("whfeedback")
	ObjectCategoryComment      = ObjectCategoryName("comment")
	ObjectCategoryChannel      = ObjectCategoryName("channel")
	ObjectCategoryChatroom     = ObjectCategoryName("chatroom")
	ObjectCategoryCalendar     = ObjectCategoryName("calendar")
	ObjectCategoryNotification = ObjectCategoryName("notification")
	ObjectCategoryUnknow       = ObjectCategoryName("unknow")

	//Report
	ObjectCategoryUserReport = ObjectCategoryName("user_report")
	//Action
	ObjectCategoryActionLike    = ObjectCategoryName("action_like")
	ObjectCategoryActionComment = ObjectCategoryName("action_comment")
	ObjectCategoryActionShare   = ObjectCategoryName("action_share")
	ObjectCategoryActionReport  = ObjectCategoryName("action_report")
	ObjectCategoryActionBlock   = ObjectCategoryName("action_block")
	ObjectCategoryActionFollow  = ObjectCategoryName("action_follow")
)

func GetCategory(data interface{}) ObjectCategoryName {
	switch data.(type) {
	case *user.User:
		return ObjectCategoryUser
	case *post.Post:
		return ObjectCategoryPost
	case *meeting.Meeting:
		return ObjectCategoryMeeting
	case *photo.Photo:
		return ObjectCategoryPhoto
	case *video.Video:
		return ObjectCategoryVideo
	case *feedback.Feedback:
		return ObjectCategoryFeedback
	case *whfeedback.WHFeedback:
		return ObjectCategoryWHFeedback
	case *comment.Comment:
		return ObjectCategoryComment
	case *channel.Channel:
		return ObjectCategoryChannel
	case *chatroom.ChatRoom:
		return ObjectCategoryChatroom
	case *calendar.Calendar:
		return ObjectCategoryCalendar
	case *like.Like:
		return ObjectCategoryActionLike
	case *share.Share:
		return ObjectCategoryActionShare
	case *follow.Follow:
		return ObjectCategoryActionFollow
	case *report.Report:
		return ObjectCategoryActionReport
	case *block.Block:
		return ObjectCategoryActionBlock
	case *notification.NotifiationManagement:
		return ObjectCategoryNotification
	case *userReport.UserReport:
		return ObjectCategoryUserReport
	default:
		return ObjectCategoryUnknow
	}
}
