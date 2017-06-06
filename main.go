package main

import("net/http"
	"fmt"
	"html/template"
)

func dummy() (_, err error) {
	return
}

func index(writer http.ResponseWriter, request *http.Request) {

	//---------------------------------
	/*
	//tmpl_file1 := "test1.html"
	templates :=template.Must(template.ParseFiles("templates/test1.html"))
	dummy1, _:= dummy()
	templates.ExecuteTemplate(writer, "layout", dummy1)
	*/
	//---------------------------------


	tmpl_files :=[]string  {
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html",
	}
	templates :=template.Must(template.ParseFiles(tmpl_files...))
	dummy1, _:= dummy()
	templates.ExecuteTemplate(writer, "layout", dummy1)

	//---------------------------------
	/*
	//generateHTML(writer, threads, "layout", "public.navbar", "index")
	//generateHTML(writer, threads, "layout", "private.navbar", "index")

	func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
	files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
	*/
	//---------------------------------
}



func main() {
	fmt.Println("starting server + mux")

	mux :=http.NewServeMux()

	//handling of files must be in main function.
	//note : found inconsistent behavior if this is inside the index function.
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	server :=&http.Server{
		Addr: "0.0.0.0:8000",
		Handler : mux,

	}
	fmt.Println("listening...")
	server.ListenAndServe()
}