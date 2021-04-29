package conversation

func getConversation(where map[string]interface{}) (*Conversation, error) {
	var u Conversation
	return &u, TableConversation.ReadOne(where, &u)
}

func GetByID(id string) (*Conversation, error) {
	var u Conversation
	return &u, TableConversation.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Conversation, error) {
	var b = []*Conversation{}
	return b, TableConversation.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Conversation, error) {
	var Conversations = []*Conversation{}
	return Conversations, TableConversation.UnsafeReadAll(&Conversations)
}

func GetSeach(key, value string) ([]*Conversation, error) {
	var Conversations = []*Conversation{}
	return Conversations, TableConversation.Search(key, value, &Conversations)
}

func Count(where map[string]interface{}) (int, error) {
	return TableConversation.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Conversation, error) {
	var tks = []*Conversation{}
	var err error
	if _type != "ascending" {
		err = TableConversation.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableConversation.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}
