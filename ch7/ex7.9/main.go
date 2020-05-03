package main

import (
	tr "github.com/Mengman/the_go_programming_language_exercises/ch7/ex7.9/track"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func main() {
	const tpl = `
	<!DOCTYPE html>
	<html>
	<head>
	<meta charset="UTF-8"><title>Song tracks</title>
	<style>
	table,
td {
    border: 1px solid #333;
}

table a {
	text-decoration: none;
	color: #fff;
}

tbody tr {
    background: antiquewhite;

}

thead,
tfoot {
    background-color: #333;
    color: #fff;
}
	</style>
	</head>

	<body>
	<div>
	<table>
	<thead>
	<tr>
		<th> <a href="/?sortBy=1"> Title  </a> </th>  
		<th> <a href="/?sortBy=2"> Artist </a> </th>  
		<th> <a href="/?sortBy=3"> Album  </a> </th>  
		<th> <a href="/?sortBy=4"> Year   </a> </th>  
		<th> <a href="/?sortBy=5"> Length </a> </th>  
  	</tr>
	</thead>
	<tbody>
	{{ range . }}
	<tr>
	<th>{{.Title}}</th>
	<th>{{.Artist}}</th>
	<th>{{.Album}}</th>
	<th>{{.Year}}</th>
	<th>{{.Length}}</th>
  </tr>
	{{end}}
	
	</tbody>
	</table>
	</div>
	</body>
	</html>
	`

	t, err := template.New("webpage").Parse(tpl)
	if err != nil {
		log.Fatal(err)
	}

	var tracks = []*tr.Track{
		{"Go", "Delilah", "From the Roots Up", 2012, tr.Length("3m38s")},
		{"Go", "Moby", "Moby", 1992, tr.Length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, tr.Length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, tr.Length("4m24s")},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s := r.FormValue("sortBy")
		idx, err := strconv.Atoi(s)
		bmrc := tr.NewByMostRecentlyColumns(tracks, 5)

		if err == nil && idx < 6 && idx > 0 {
			switch idx {
			case 1:
				bmrc.AddCmp(bmrc.LessTitle)
			case 2:
				bmrc.AddCmp(bmrc.LessArtist)
			case 3:
				bmrc.AddCmp(bmrc.LessAlbum)
			case 4:
				bmrc.AddCmp(bmrc.LessYear)
			case 5:
				bmrc.AddCmp(bmrc.LessLength)
			}

			sort.Sort(bmrc)
		}
		t.Execute(w, tracks)
	})

	log.Fatal(http.ListenAndServe(":8888", nil))
}
