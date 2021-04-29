package user

import (
	"db/mgo"
	"huntsub/huntsub-map-server/o/model"
	"reflect"
)

var TableUser = model.NewTable("huntsub-server", "users", "usr")

func NewUserID() string {
	return TableUser.Next()
}

func (s *User) MakeID() string {
	return TableUser.IdMaker.Next()
}

func (s *User) Insert() error {
	s.BeforeCreate()

	if s.GetID() == "" {
		s.SetID(TableUser.IdMaker.Next())
	}

	if err := s.ensureUniqueUsername(); err != nil {
		return err
	}
	var p = password(s.Password)
	// replace
	if err := p.HashTo(&s.Password); err != nil {
		return err
	}
	var _ = TableUser.EnsureAddressIndex()
	return TableUser.UnsafeInsert(s)
}

func (b *User) Create() error {
	if err := b.ensureUniqueUsername(); err != nil {
		return err
	}
	var p = password(b.Password)
	// replace
	if err := p.HashTo(&b.Password); err != nil {
		return err
	}
	var _ = TableUser.EnsureAddressIndex()
	return TableUser.Create(b)
}

func MarkDelete(id string) error {
	return TableUser.MarkDelete(id)
}

func (v *User) Update(newValue *User) (*User, error) {
	var values = map[string]interface{}{}

	if newValue.GetPhone() != v.GetPhone() && newValue.GetPhone() != "" {
		values["phone"] = newValue.GetPhone()
	}

	if newValue.GetName() != v.GetName() && newValue.GetName() != "" {
		values["name"] = newValue.GetName()
	}

	if newValue.GetSex() != v.GetSex() && newValue.GetSex() != "" {
		values["sex"] = newValue.GetSex()
	}

	if newValue.GetEmail() != v.GetEmail() && newValue.GetEmail() != "" {
		values["email"] = newValue.GetEmail()
	}

	if newValue.GetBirthday() != v.GetBirthday() && newValue.GetBirthday() != "" {
		values["birthday"] = newValue.GetBirthday()
	}

	if newValue.GetNationality() != v.GetNationality() && newValue.GetNationality() != "" {
		values["nationality"] = newValue.GetNationality()
	}

	if newValue.GetIDCard() != v.GetIDCard() && newValue.GetIDCard() != "" {
		values["idcard"] = newValue.GetIDCard()
	}

	if newValue.GetAvatar() != v.GetAvatar() && newValue.GetAvatar() != "" {
		values["avatar"] = newValue.GetAvatar()
	}

	if newValue.GetBackGround() != v.GetBackGround() && newValue.GetBackGround() != "" {
		values["background"] = newValue.GetBackGround()
	}

	if newValue.GetReadedLaw() != v.GetReadedLaw() {
		values["readed_law"] = newValue.GetReadedLaw()
	}

	if newValue.GetVerify() != v.GetVerify() && !v.GetVerify() {
		values["verify"] = newValue.GetVerify()
	}

	if newValue.GetActive() != v.GetActive() {
		values["active"] = newValue.GetActive()
	}

	if newValue.GetExpoToken() != v.GetExpoToken() && newValue.GetExpoToken() != "" {
		values["expotoken"] = newValue.GetExpoToken()
	}

	if !reflect.DeepEqual(newValue.GetStatusActive(), v.GetStatusActive()) {
		values["statusactive"] = newValue.GetStatusActive()
	}

	if newValue.Getlanguage() != v.Getlanguage() && newValue.Getlanguage() != "" {
		values["language"] = newValue.Getlanguage()
	}

	if !reflect.DeepEqual(newValue.GetLocation(), v.GetLocation()) {
		if newValue.Location.Type != "" {
			values["location"] = newValue.GetLocation()
		}
	}

	if !reflect.DeepEqual(newValue.GetHome(), v.GetHome()) {
		if newValue.Location.Type != "" {
			values["home"] = newValue.GetHome()
		}
	}

	var _ = TableUser.EnsureAddressIndex()
	err := TableUser.UnsafeUpdateByID(v.ID, values)
	res, err := GetByID(v.ID)
	return res, err
}

func (v *User) UpdatePass(newValue string) error {
	var update = map[string]interface{}{
		"password": newValue,
	}

	if len(newValue) > 0 {
		var p = password(newValue)
		if err := p.HashTo(&newValue); err != nil {
			return err
		}
		update["password"] = newValue
	}
	return TableUser.UnsafeUpdateByID(v.ID, update)
}

var _ = TableUser.EnsureIndex(mgo.Index{
	Key:        []string{"name"},
	Background: true,
})

var _ = TableUser.EnsureAddressIndex()
