package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	//"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	//CreatedAt time.Time
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}


func dummy() (_, err error) {
	return
}

func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("bb0")

	var threads []Thread

	Db, err := sql.Open("mysql", "ayong:ayong@/test?charset=utf8")
	if err != nil {
		fmt.Println("Error in ee1")
		log.Fatal(err)
	}
	rows, err := Db.Query("SELECT id, uuid, topic, user_id FROM threads ORDER BY created_at DESC")
	if err != nil {
		fmt.Println("Error in ee3")
		log.Fatal(err)
	}
	fmt.Println("bb1")

	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId); err != nil {
			log.Fatal(err)
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	fmt.Println("bb2")


	//---------------------------------
	/*
		//tmpl_file1 := "test1.html"
		templates :=template.Must(template.ParseFiles("templates/test1.html"))
		dummy1, _:= dummy()
		templates.ExecuteTemplate(writer, "layout", dummy1)
	*/
	//---------------------------------

	tmpl_files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(tmpl_files...))

	//dummy1, _ := dummy()
	//templates.ExecuteTemplate(writer, "layout", dummy1)

	templates.ExecuteTemplate(writer, "layout", threads)

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

	mux := http.NewServeMux()

	//handling of files must be in main function.
	//note : found inconsistent behavior if this is inside the index function.
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr:    "0.0.0.0:7000",
		Handler: mux,
	}
	fmt.Println("listening...")
	server.ListenAndServe()
}
