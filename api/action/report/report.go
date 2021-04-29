package report

import (
	"http/web"
	"huntsub/huntsub-map-server/api/auth/session"
	cache_report "huntsub/huntsub-map-server/cache/org/action/report"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/action/report"
	"net/http"
	"strconv"
)

var obj = event.ObjectEventSource

type ReportServer struct {
	web.JsonServer
	*http.ServeMux
}

func NewReportServer() *ReportServer {
	var s = &ReportServer{
		ServeMux: http.NewServeMux(),
	}
	s.HandleFunc("/create", s.HandleCreate)
	s.HandleFunc("/get", s.HandleGetByID)
	s.HandleFunc("/update", s.HandleUpdate)
	s.HandleFunc("/delete", s.HandleDelete)
	s.HandleFunc("/getall", s.HanldeGetAll)
	s.HandleFunc("/search", s.HanldeSearch)

	return s
}

func (s *ReportServer) HandleCreate(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var p = &report.Report{}
	p.UserID = sess.UserID
	s.MustDecodeBody(r, &p)
	web.AssertNil(p.Create())
	obj.EmitCreate(p)
	s.SendData(w, p)
}

func (s *ReportServer) mustGetReport(r *http.Request) *report.Report {
	var id = r.URL.Query().Get("id")
	var p, err = cache_report.Get(id)
	web.AssertNil(err)
	return p
}

func (s *ReportServer) HandleGetByID(w http.ResponseWriter, r *http.Request) {
	var p = s.mustGetReport(r)
	s.SendData(w, p)
}

func (s *ReportServer) HandleUpdate(w http.ResponseWriter, r *http.Request) {
	var newReport = &report.Report{}
	s.MustDecodeBody(r, &newReport)
	var p = s.mustGetReport(r)
	res, err := p.Update(newReport)
	web.AssertNil(err)
	s.SendData(w, nil)
	obj.EmitUpdate(res)
}

func (s *ReportServer) HanldeGetAll(w http.ResponseWriter, r *http.Request) {
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime": 0,
	}
	res, ok := report.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = report.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}

func (s *ReportServer) HandleDelete(w http.ResponseWriter, r *http.Request) {
	var u = s.mustGetReport(r)
	web.AssertNil(report.MarkDelete(u.ID))
	s.Success(w)
	obj.EmitMarkDelete(u)
}

func (s *ReportServer) HanldeSearch(w http.ResponseWriter, r *http.Request) {
	var sess = session.MustAuthScope(r)
	var skip, err = strconv.Atoi(r.URL.Query().Get("skip"))
	web.AssertNil(err)
	var limit, _err = strconv.Atoi(r.URL.Query().Get("limit"))
	web.AssertNil(_err)
	var _type = r.URL.Query().Get("type")
	var query = map[string]interface{}{
		"dtime":  0,
		"userid": sess.UserID,
	}
	res, ok := report.GetByPaginationAd(query, skip, limit, _type)
	web.AssertNil(ok)
	var count, _ok = report.Count(query)
	web.AssertNil(_ok)
	s.SendData(w, map[string]interface{}{
		"count": count,
		"data":  res,
	})
}
