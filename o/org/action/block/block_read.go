package block

func getBlock(where map[string]interface{}) (*Block, error) {
	var u Block
	return &u, TableBlock.ReadOne(where, &u)
}

func GetByID(id string) (*Block, error) {
	var u Block
	return &u, TableBlock.ReadByID(id, &u)
}

func GetMany(where map[string]interface{}) ([]*Block, error) {
	var b = []*Block{}
	return b, TableBlock.UnsafeReadMany(where, &b)
}

func GetAll() ([]*Block, error) {
	var Blocks = []*Block{}
	return Blocks, TableBlock.UnsafeReadAll(&Blocks)
}

func GetSeach(key, value string) ([]*Block, error) {
	var Blocks = []*Block{}
	return Blocks, TableBlock.Search(key, value, &Blocks)
}

func Count(where map[string]interface{}) (int, error) {
	return TableBlock.UnsafeCount(where)
}

func GetByPaginationAd(where map[string]interface{}, skip, limit int, _type string) ([]*Block, error) {
	var tks = []*Block{}
	var err error
	if _type != "ascending" {
		err = TableBlock.C().Find(where).Sort("-$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	} else {
		err = TableBlock.C().Find(where).Sort("$natural").Skip((skip - 1) * limit).Limit(limit).All(&tks)
	}
	return tks, err
}
