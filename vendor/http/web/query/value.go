package query

import (
	"http/web"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Query struct {
	url.Values
}

func GetQuery(r *http.Request) Query {
	return Query{Values: r.URL.Query()}
}

func (q Query) MustAuthScope(key string) string {
	value := q.Get(key)
	if value == "" {
		panic(web.BadRequest("missing " + key))
	}
	return value
}

func (q Query) MustGetString(key string) string {
	return q.MustAuthScope(key)
}

func (q Query) MustGetInt64(key string) int64 {
	str := q.MustAuthScope(key)
	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(web.BadRequest(key + " must be int"))
	}
	return v
}

func (q Query) MustGetFloat(key string) float64 {
	str := q.MustAuthScope(key)
	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(web.BadRequest(key + " must be number"))
	}
	return v
}

func (q Query) MustGetArrString(key string, sep string) []string {
	var value = q.MustAuthScope(key)
	if len(value) < 1 {
		return []string{}
	}
	return strings.Split(value, sep)
}

func (q Query) GetArrString(key string, sep string) []string {
	var value = q.Get(key)
	if len(value) < 1 {
		return []string{}
	}
	return strings.Split(value, sep)
}
