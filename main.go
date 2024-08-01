package main

import (
	"html/template"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// film := map[string][]Film{
	// 	"Films": {
	// 		{Title: "Koe", Director: "Phongphat"},
	// 		{Title: "Minkwan", Director: "Rinlada"},
	// 	},
	// }

	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Fprintf(w, "%s", "hello")

		tmpl.Execute(w, nil)
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("monthly.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Fprintf(w, "%s", "hello")

		tmpl.Execute(w, nil)
	}

	h3 := func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("yearly.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//fmt.Fprintf(w, "%s", "hello")

		tmpl.Execute(w, nil)
	}
	// add := func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("Hello")
	// 	log.Println("f", film)
	// 	film["Films"] = append(film["Films"], Film{Title: "Love", Director: "Forever"})
	// 	log.Println("f", film)
	// 	htmlStr := fmt.Sprintf("<p>%s %s</p>", film["Films"][len(film["Films"])-1].Title, film["Films"][len(film["Films"])-1].Director)

	// 	tmpl, _ := template.New("t").Parse(htmlStr)
	// 	tmpl.Execute(w, nil)
	// }

	http.HandleFunc("/", h1)
	http.HandleFunc("/monthly", h2)
	http.HandleFunc("/yearly", h3)
	http.ListenAndServe(":8000", nil)

}
