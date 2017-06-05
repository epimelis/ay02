package main

import("net/http"
	"fmt"
	"html/template"
)

func dummy() (_, err error) {
	return
}

func index(writer http.ResponseWriter, r *http.Request) {
	tmpl_file1 := "test1.html"


	/*
	tmpl_filesMany :=[]string  {
		"test1.html",
		"test2.html",
	}
	*/




	templates :=template.Must(template.ParseFiles(tmpl_file1))
	//templates2 :=template.Must(template.ParseFiles(tmpl_filesMany...))

	dummy1, _:= dummy()
	templates.ExecuteTemplate(writer, "layout", dummy1)
	//templates2.ExecuteTemplate(w, "layout", dummy1)


}
func main() {
	fmt.Println("starting server + mux")

	mux :=http.NewServeMux()

	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	server :=&http.Server{
		Addr: "0.0.0.0:7002",
		Handler : mux,

	}
	fmt.Println("listening...")
	server.ListenAndServe()
}