// string to make html file:
// go run HTMLsorting.go > filename.html
package main

import (
	"Book_exercises/ch7/sorting"
	"html/template"
	"log"
	"os"
	"sort"
)

type TracksList struct {
	TotalCount int
	Items      []*sorting.Track
}

func sortTable(tl *TracksList, key string) *TracksList {
	switch key {
	case "title":
		sort.Sort(sorting.ByTitle(tl.Items))
		break
	}
	return tl
}

var tracksTable = template.Must(template.New("tracksTable").Funcs(
	template.FuncMap{
		"sortTable": sortTable,
	}).Parse(`
<h1>{{.TotalCount}} songs</h1>
<table border="1">
<tr style='text-align: left'>
	<th onclick="{{ sortTable . "title" }}">Title</th>
	<th onclick="{{ sortTable . "artist" }}">Artist</th>
	<th onclick="location.href= 'test1.html?sort=byalbum'">Album</th>
	<th onclick="location.href= 'test1.html?sort=byyear'">Year</th>
	<th onclick="location.href= 'test1.html?sort=bylength'">Length</th>
</tr>
{{range .Items}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	//http.HandleFunc("/", handler)
	//log.Fatal(http.ListenAndServe(":8000", nil))
	trackListStruct := TracksList{
		sorting.CountTracks(sorting.TracksInit),
		sorting.TracksInit,
	}

	if err := tracksTable.Execute(os.Stdout, &trackListStruct); err != nil {
		log.Fatal(err)
	}
}
