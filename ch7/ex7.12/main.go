package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

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
		<th>Item</th>  
		<th>Price</th>
  	</tr>
	</thead>
	<tbody>
	{{ range $key, $value := . }}
	<tr>
	<th>{{$key}}</th>
	<th>{{$value}}</th>
  </tr>
	{{end}}
	
	</tbody>
	</table>
	</div>
	</body>
	</html>
	`

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

type dollars float32

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	t, err := template.New("page").Parse(tpl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	t.Execute(w, db)
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%0.3f\n", price)
}
