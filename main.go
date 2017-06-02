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


func dummy() (_, err error) {
	return
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("aa1")
	tmpl_file1 := "test.html"
	/*tmpl_files :=[]string  {
		"test.html",
	}*/

	fmt.Println("aa2")
	//templates :=template.Must(template.ParseFiles(tmpl_files...))
	templates :=template.Must(template.ParseFiles(tmpl_file1))
	fmt.Println("aa3")
	dummy1, _:= dummy()

	templates.ExecuteTemplate(w, "layout", dummy1)


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