package vstatic

import (
	"fmt"
	"net/http"
	"strings"
)

func (v *internalVersionStatic) renderView(w http.ResponseWriter, r *http.Request) {
	var header = fmt.Sprintf("<h1> Control %v </h1>", v.dir)
	var active = fmt.Sprintf("<div> Current version <b> %s </b></div>", v.active)
	var elements = []interface{}{header, active}
	var rows = []string{}
	var versions, err = v.ListVersion()
	if err != nil {
		elements = append(elements, "list version failed "+err.Error())
	} else {
		for _, version := range versions {
			action := "Activate"
			if version.Version == v.active {
				action = "Reactivate"
			}
			click := fmt.Sprintf("<form action='./activate?version=%s' method='post'><button type='submit'>%s %s</button></form>", version.Version, action, version.Version)
			tr := fmt.Sprintf("<tr><td>%s</td><td>%v</td><td>%s</td></tr>", version.Version, version.Active, click)
			rows = append(rows, tr)
		}
		var table = "<table><thead><th>Version</th><th>Active</th></thead><tbody>" + strings.Join(rows, "") + "</tbody></table>"
		elements = append(elements, table)
	}
	check := "<form action='./check'><button type='submit'>Check Version</button></form>"
	download := "<form action='./download_latest'><button type='submit'>Download Latest</button></form>"
	elements = append(elements, check, download)
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Cache-Control", "No-Cache")
	fmt.Fprint(w, elements...)
}
