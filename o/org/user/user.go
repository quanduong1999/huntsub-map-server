package user

import (
	"db/mgo"
	"http/web"
	"huntsub/huntsub-map-server/x/mlog"
)

var objectUserLog = mlog.NewTagLog("object_user")

//User : Employee
type User struct {
	mgo.BaseModel   `bson:",inline"`
	Username        string       `bson:"username" json:"username,omitempty"`
	Name            string       `bson:"name" json:"name"`
	Password        string       `bson:"password"`
	PrimaryPassword string       `bson:"primary_password"`
	Email           string       `bson:"email" json:"email,omitempty"`
	IDCard          string       `bson:"idcard" json:"idcard,omitempty"` //CMTND
	Phone           string       `bson:"phone" json:"phone"`
	Sex             string       `bson:"sex" json:"sex"`
	Birthday        string       `bson:"birthday" json:"birthday"`
	Nationality     string       `bson:"nationality" json:"nationality"`
	Avatar          string       `bson:"avatar" json:"avatar"`
	BackGround      string       `json:"background" bson:"background"`
	ReadedLaw       bool         `bson:"readed_law" json:"readed_law"`
	Verify          bool         `bson:"verify" json:"verify"`
	Home            Location     `json:"home" bson:"home"`
	Active          bool         `json:"active" bson:"active"`
	ExpoToken       string       `json:"expotoken" bson:"expotoken"`
	StatusActive    StatusActive `json:"statusactive" bson:"statusactive"`
	Location        Location     `json:"location" bson:"location"`
	Language        string       `json:"language" bson:"language"`
	Job             string       `json:"job" bson:"job"`
}

type StatusActive struct {
	Online  bool   `json:"online" bson:"online"`
	RoomID  string `json:"roomid" bson:"roomid"`
	TimeOut int64  `json:"timeout" bson:"timeout"`
	Device  string `json:"device" bson:"device"`
}

type Location struct {
	Address     string    `json:"address" bson:"address"`
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
}

func (v *User) ensureUniqueUsername() error {
	if len(v.Username) < 3 {
		return web.BadRequest("Username must be at least 6 characters")
	}
	if err := TableUser.NotExist(map[string]interface{}{
		"username": v.Username,
	}); err != nil {
		return web.BadRequest("Username was exist")
	}
	return nil
}

func (v *User) GetName() string {
	return v.Name
}

func (v *User) GetEmail() string {
	return v.Email
}

func (v *User) GetPhone() string {
	return v.Phone
}

func (v *User) GetSex() string {
	return v.Sex
}

func (v *User) GetBirthday() string {
	return v.Birthday
}

func (v *User) GetNationality() string {
	return v.Nationality
}

func (v *User) GetAvatar() string {
	return v.Avatar
}

func (v *User) GetBackGround() string {
	return v.BackGround
}

func (v *User) GetReadedLaw() bool {
	return v.ReadedLaw
}

func (v *User) GetIDCard() string {
	return v.IDCard
}

func (v *User) GetVerify() bool {
	return v.Verify
}

func (v *User) GetActive() bool {
	return v.Active
}

func (v *User) GetExpoToken() string {
	return v.ExpoToken
}

func (v *User) GetStatusActive() StatusActive {
	return v.StatusActive
}

func (v *User) Getlanguage() string {
	return v.Language
}

func (v *User) GetLocation() Location {
	return v.Location
}

func (v *User) GetHome() Location {
	return v.Home
}
