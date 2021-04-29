package oev

type ObjectActionName string

const (
	ObjectActionCreate     = ObjectActionName("create")
	ObjectActionUpdate     = ObjectActionName("update")
	ObjectActionMarkDelete = ObjectActionName("mark_delete")
	ObjectActionDelete     = ObjectActionName("delete")
	ObjectActionConfirm    = ObjectActionName("confirm")

	ObjectActionLike     = ObjectActionName("like")
	ObjectActionShare    = ObjectActionName("share")
	ObjectActionReport   = ObjectActionName("report")
	ObjectActionBlock    = ObjectActionName("block")
	ObjectActionUnFollow = ObjectActionName("unfollow")
	ObjectActionFollow   = ObjectActionName("follow")

	ObjectActionFeedBack = ObjectActionName("feedback")
	ObjectActionLevelUp  = ObjectActionName("level_up")
	ObjectActionMeeting  = ObjectActionName("meeting")
)
