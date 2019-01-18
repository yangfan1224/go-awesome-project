package template

import (
	"html/template"
	"net/http"
)

var testTemplate *template.Template
type Widget struct {
	Name  string
	Price int
}


type ViewRangeData struct {
	Name    string
	Widgets []Widget
}

type ViewData struct {
	Name string
}

func HtmlTemplateNested(){
	var err error
	testTemplate, err = template.ParseFiles("hello.gohtml","hellothere.gohtml","range.gohtml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/there", handlerThere)
	http.HandleFunc("/range", handlerRange)
	http.ListenAndServe(":3000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	vd := ViewData{"John Smith"}
	err := testTemplate.ExecuteTemplate(w,"hello.gohtml", vd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerThere(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := testTemplate.ExecuteTemplate(w,"hellothere.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handlerRange(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := testTemplate.ExecuteTemplate(w,"range.gohtml", ViewRangeData{
		Name: "Range",
		Widgets: []Widget{
			Widget{"Blue Widget", 12},
			Widget{"Red Widget", 12},
			Widget{"Green Widget", 12},
		},
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}