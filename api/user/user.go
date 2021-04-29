package user

import (
	"fmt"
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_user "huntsub/huntsub-map-server/cache/org/user"
	"huntsub/huntsub-map-server/config/cons"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/channel"
	"huntsub/huntsub-map-server/o/org/datawarehouse/whfeedback"
	"huntsub/huntsub-map-server/o/org/rank"
	"huntsub/huntsub-map-server/o/org/user"
	"huntsub/huntsub-map-server/x/language"
	"huntsub/huntsub-map-server/x/mail"
	"huntsub/huntsub-map-server/x/math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var obj = event.ObjectEventSource

type UserServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewUserServer() *UserServer {
	var s = &UserServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)
	s.HandleFunc("/verify", s.HandleVerify)
	s.HandleFunc("/near", s.HandleGetAllUserNear)
	s.HandleFunc("/info", s.HandleGetAllInforUser)
	s.HandleFunc("/renew", s.HandleGetBackAccount)
	s.HandleFunc("/confirm", s.HandleConfirmAccount)
	s.HandleFunc("/checklistactive", s.HandleCheckListActive)

	return s
}

func (s *UserServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var u = &user.User{}
	s.MustDecodeBody(r, &u)
	u.Username = strings.ToLower(u.Username)
	u.Location = u.Home
	if err := u.Insert(); err != nil {
		s.SendError(w, err)
		return
	}
	//Create folder stogare in firebase2
	//End

	//Create folder stogare in local
	os.MkdirAll(makePath("video", u), os.ModeSticky|0755)
	os.MkdirAll(makePath("img", u), os.ModeSticky|0755)
	os.MkdirAll(makePath("img_msg", u), os.ModeSticky|0755)
	os.MkdirAll(makePath("audio", u), os.ModeSticky|0755)
	/*Sending email verifitation for new user*/
	to := []string{}
	to = append(to, u.Username)
	ok := mail.SendSingleMail(cons.Email, to, language.Translation(u.Language, cons.VERIFY_ACCOUNT_HUNTSUB_NETWORK_SYSTEM), u.GetID(), func() string {
		link := cons.URL_VERIFY + u.GetID()
		body := ""
		body += fmt.Sprintf(language.Translation(u.Language, cons.GREATING) + "\r\n\n" + u.Name)
		body += fmt.Sprintf(
			language.Translation(u.Language, cons.THANK_YOU_FOR_REGISTER_On_MY_SYSTEM) +
				language.Translation(u.Language, cons.PLEASE_VERIFY) +
				"\r\n\n" + link + "\r\n\n")
		body += fmt.Sprintf(language.Translation(u.Language, cons.SINGATURE))
		return body
	})
	web.AssertNil(ok)
	obj.EmitCreate(u)
	s.SendData(w, u)
}

func makePath(mode string, u *user.User) string {
	var t = time.Now()
	switch mode {
	case "video":
		return fmt.Sprintf("static/upload/user/%04d-%02d-%02d/%s/video", t.Year(), t.Month(), t.Day(), u.ID)
	case "img":
		return fmt.Sprintf("static/upload/user/%04d-%02d-%02d/%s/img", t.Year(), t.Month(), t.Day(), u.ID)
	case "img_msg":
		return fmt.Sprintf("static/upload/user/%04d-%02d-%02d/%s/img_msg", t.Year(), t.Month(), t.Day(), u.ID)
	case "audio":
		return fmt.Sprintf("static/upload/user/%04d-%02d-%02d/%s/audio", t.Year(), t.Month(), t.Day(), u.ID)
	}
	return ""
}

func (s *UserServer) mustGetUser(r *http.Request) *user.User {
	var id = r.URL.Query().Get("id")
	var u, err = cache_user.Get(id)
	web.AssertNil(err)
	return u
}

func (s *UserServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetUser(r)
	s.SendData(w, u)
}

func (s *UserServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newUser = &user.User{}
	s.MustDecodeBody(r, &newUser)
	newUser.Username = strings.ToLower(newUser.Username)
	sess := session.MustAuthScope(r)
	u, err := cache_user.Get(sess.UserID)
	web.AssertNil(err)
	res, err := u.Update(newUser)
	web.AssertNil(err)
	s.SendData(w, res)
	obj.EmitUpdate(res)
}

func (s *UserServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := user.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = user.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *UserServer) HandleVerify(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	var u, err = user.GetByID(id)
	web.AssertNil(err)
	var newUser = &user.User{}
	newUser.Verify = true
	c, err := u.Update(newUser)
	web.AssertNil(err)
	obj.EmitUpdate(c)
	s.Success(w)

}

func (s *UserServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetUser(r)
	web.AssertNil(user.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *UserServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	u, err := cache_user.Get(sess.UserID)
	web.AssertNil(err)
	var value = r.URL.Query().Get("value")
	skip, err := strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	limit, _err := strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var query = map[string]interface{}{
		"dtime": 0,
		"name":  bson.M{"$regex": bson.RegEx{Pattern: value}},
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": u.Location.Coordinates,
				},
				"$maxDistance": 1000000,
			},
		},
		"_id": bson.M{"$ne": sess.UserID},
	}
	res, ok := user.GetByPagination(query, skip, limit)
	web.AssertNil(ok)
	var count, _ok = user.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *UserServer) HandleGetAllUserNear(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	u, err := cache_user.Get(sess.UserID)
	web.AssertNil(err)
	// ran, err := cache_rank.Get()
	var query = map[string]interface{}{
		"dtime": 0,
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": u.Location.Coordinates,
				},
				"$maxDistance": 1000000,
			},
		},
		"_id": bson.M{"$ne": sess.UserID},
	}
	res, ok := user.GetByPagination(query, skip, limit)
	web.AssertNil(ok)
	data := NewUserForm(res, sess.UserID)
	s.SendData(w, data)
}

func (s *UserServer) HandleGetAllInforUser(w http.ResponseWriter, r *http.Request) {
	var userID = r.URL.Query().Get("id")
	u, err := cache_user.Get(userID)
	web.AssertNil(err)

	var query = bson.M{
		"dtime":  0,
		"userid": userID,
	}
	cn, err := channel.GetChannel(query)
	web.AssertNil(err)

	ran, err := rank.GetRank(query)
	web.AssertNil(err)

	whfb, err := whfeedback.GetWHFeedBack(query)
	web.AssertNil(err)

	data := NewUserInfo(u, cn, ran, whfb)
	s.SendData(w, data)
}

func (s *UserServer) HandleGetBackAccount(w http.ResponseWriter, r *http.Request) {
	var id = r.URL.Query().Get("id")
	u, err := user.GetByID(id)
	web.AssertNil(err)
	newPass := math.RandString("", 8)
	println(newPass)
	u.UpdatePass(newPass)

	to := []string{}
	to = append(to, u.Username)
	ok := mail.SendSingleMail(cons.Email, to, "[Lấy lại mật khẩu]", u.GetID(), func() string {
		body := ""
		body += fmt.Sprintf("Xin chào %s\r\n\n", u.Name)
		body += fmt.Sprintf("Mật khẩu mới của bạn là : \r\n\n %s\r\n\n", newPass)
		body += fmt.Sprintf("Trân trọng!\r\n\n")
		return body
	})
	web.AssertNil(ok)
	obj.EmitUpdate(u)
	s.Success(w)
}

func (s *UserServer) HandleConfirmAccount(w http.ResponseWriter, r *http.Request) {
	var email = r.URL.Query().Get("email")
	//send email confirm
	to := []string{}
	to = append(to, email)
	u, err := user.GetByUsername(email)
	web.AssertNil(err)

	ok := mail.SendSingleMail(cons.Email, to, "[Lấy lại mật khẩu]", u.GetID(), func() string {
		link := cons.URL_RENEWPASS + u.GetID()
		body := ""
		body += fmt.Sprintf("Xin chào %s\r\n\n", u.Name)
		body += fmt.Sprintf("Chúng tôi nhận được yêu cầu lấy lại mật khẩu của bạn. Bạn vui lòng xác nhận lại vào link bên dưới.\nHệ thống sẽ gửi mật khẩu mới cho bạn\r\n\n %s\r\n\n", link)
		body += fmt.Sprintf("Trân trọng!\r\n\n")
		return body
	})
	web.AssertNil(ok)
	s.Success(w)
}

type ContactList struct {
	ContactList []string `json:"contact_list"`
}

type Phone struct {
	Phone  string `json:"phone"`
	Status bool   `json:"status"`
}

func (s *UserServer) HandleCheckListActive(w http.ResponseWriter, r *http.Request) {
	var list = &ContactList{}
	s.MustDecodeBody(r, list)

	fmt.Println(list.ContactList)
	var query = bson.M{
		"dtime": 0,
		"phone": bson.M{"$in": list.ContactList},
	}

	res, err := user.GetMany(query)
	data := []*Phone{}
	// for _, b := range list.ContactList {
	// 	f := &Phone{}
	// 	fmt.Println(b)
	// 	f.Status = false
	// 	data = append(data, f)
	// }
	for _, b := range list.ContactList {
		for _, c := range res {
			e := &Phone{}
			e.Phone = b
			if c.Phone == b {
				e.Status = true
				// e.Phone = b
			} else {
				e.Status = false
			}

			data = append(data, e)
		}
	}
	// for _, c := range res {
	// 	for _, b := range list.ContactList {
	// 		e := &Phone{}
	// 		e.Phone = b
	// 		if c.Phone == b {
	// 			e.Status = true

	// 		} else {
	// 			e.Status = false
	// 		}

	// 		data = append(data, e)
	// 	}
	// }
	web.AssertNil(err)
	s.SendData(w, data)
}
