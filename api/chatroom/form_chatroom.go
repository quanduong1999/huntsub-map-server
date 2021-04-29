package chatroom

import (
	"http/web"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/o/org/chatroom"
)

func NewChatRoomForm(arr []*chatroom.ChatRoom, userID string) []*chatroom.ChatRoomForm {
	var s = []*chatroom.ChatRoomForm{}
	for _, c := range arr {
		x := &chatroom.ChatRoomForm{}
		x.BackGround = c.BackGround
		x.ColorText = c.ColorText
		x.RoomName = c.RoomName
		x.LastMessage = c.LastMessage
		x.IsGroup = c.IsGroup
		x.RoomID = c.ID
		if c.IsGroup {

		} else {
			id, num := checkValue(userID, c.Users)
			u, err := cache_user.Get(id)
			web.AssertNil(err)
			x.Person = c.Users[num]
			x.Person.TimeOut = u.StatusActive.TimeOut
		}
		s = append(s, x)
	}

	return s
}

func checkValue(id string, persons []chatroom.Person) (string, int) {
	if id == persons[0].UserID {
		return persons[1].UserID, 1
	} else {
		return persons[0].UserID, 0
	}
}
