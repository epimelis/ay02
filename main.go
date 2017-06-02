package main

import("net/http"
	"fmt"
	"html/template"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
}

func Threads() (threads []Thread, err error) {

	return
}


func TTT() (_, err error) {
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aa1")
	tmpl_files :=[]string  {
		"layout.html",
	}
	fmt.Println("aa2")
	templates :=template.Must(template.ParseFiles(tmpl_files...))
	fmt.Println("aa3")
	ttt, _:= TTT()

	templates.ExecuteTemplate(w, "layout", ttt)

	/*
		if err == nil {
			fmt.Println("aaa4_noErr")
			templates.ExecuteTemplate(w, "layout", ttt)
		} else {
			fmt.Println("aaa5_ERR")
		}
	*/

}
func main() {
	fmt.Println("starting server")
	mux :=http.NewServeMux()
	mux.HandleFunc("/", index)
	server :=&http.Server{
		Addr: "0.0.0.0:8080",
		Handler : mux,

	}
	fmt.Println("listening on port 8080...")
	server.ListenAndServe()
}