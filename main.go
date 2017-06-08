package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
	"time"
)

var (
	Db *sql.DB
)
type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}
type User struct {
	Id int
	Uuid string
	Name string
	Email string
	Password string
	CreatedAt time.Time
}
type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}
func init() {
	fmt.Println("init0")
	var err error
	//Db, err = sql.Open("postgres", "postgres://user:pass@localhost/bookstore")
	Db, err = sql.Open("mysql", "ayong:ayong@/test?charset=utf8")

	if err != nil {
		fmt.Println("init_err1")
		log.Fatal(err)
	}

	if err = Db.Ping(); err != nil {
		fmt.Println("init_err2")
		log.Fatal(err)
	}
	fmt.Println("init1")
}

//Go method - attach a function to thread
func (thread *Thread) NumReplies() (count int) {
	fmt.Println("tt0")
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = ?", thread.Id)
	fmt.Println("tt1")
	if err != nil {
		fmt.Println("tt_err1")
		log.Fatal(err)
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			fmt.Println("tt_err2")
			log.Fatal(err)
			return
		}
	}
	rows.Close()
	fmt.Println("tt2")
	return
}

func (thread *Thread) User() (user User) {
	user=User{}
	fmt.Println("uu0")
	Db.QueryRow("select id, uuid, name, email from users where id=?", thread.UserId).Scan(&user.Id, &user.Uuid, &user.Name, &user.Email)
	fmt.Println("uu1")
	fmt.Println("-----start---------")
	fmt.Println(thread.UserId)
	fmt.Println(user.Name)
	fmt.Println("-----end--------")

	return
}
/*
func dummy() (_, err error) {
	return
}
*/


func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("xx0")

	if Db==nil {
		fmt.Println("Error : Db is nil!!!!!")
	} else {
		fmt.Println("Db initialized ...")
	}

	var threads []Thread

	fmt.Println("xx1")
	rows, err := Db.Query("SELECT id, uuid, topic, user_id FROM threads ORDER BY created_at DESC")
	if err != nil {
		fmt.Println("xx_err2")
		log.Fatal(err)
	}

	fmt.Println("xx2")
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId); err != nil {
			log.Fatal(err)
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	fmt.Println("xx3")


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

/*
func init() {

	fmt.Println("init : Open database...")
	Db, err := sql.Open("mysql", "ayong:ayong@/test?charset=utf8")
	if err != nil {
		fmt.Println("init_err1")
		log.Fatal(err)
	}
	if Db==nil {
		fmt.Println("arrrrrgh!")
	} else {
		fmt.Println("Db opened successfully....OK!")
	}
	return;

}
*/

func main() {


	fmt.Println("starting db ...")
	/*Db, err := sql.Open("mysql", "ayong:ayong@/test?charset=utf8")
	if Db==nil {
		fmt.Println("Error : Db is nil!!!!!")
		log.Fatal(err)
	} else {
		fmt.Println("Db initialized ...")
	}
	*/

	fmt.Println("starting server ...")
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
