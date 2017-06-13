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
	//Db, err = sql.Open("mysql", "ayong:ayong@/test?charset=utf8")
	Db, err = sql.Open("mysql", "ayong:ayong@tcp(127.0.0.1:3306)/test?parseTime=true")

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


	return
}

/*
func dummy() (_, err error) {
	return
}
*/

/*
func login(writer http.ResponseWriter, request *http.Request) {
	tmpl_files := []string {"templates/login.layout.html", "templates/public.navbar.html", "templates/login.html"}
	templates := template.Must(template.ParseFiles(tmpl_files...))
	templates.Execute(writer, nil)
	templates.ExecuteTemplate(writer, "layout", nil)

}
*/

// parse HTML templates
// pass in a list of file names, and get a template
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}

func login(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("lll_1")
	t := parseTemplateFiles("login.layout", "public.navbar", "login")
	fmt.Println("lll_2")
	t.Execute(writer, nil)
	fmt.Println("lll_3")
}


func index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("xx0")

	var threads []Thread

	fmt.Println("xx1")
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		fmt.Println("xx_err2")
		log.Fatal(err)
	}

	fmt.Println("xx2")
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.CreatedAt); err != nil {
			log.Fatal(err)
			return
		}
		threads = append(threads, th)
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

func (user *User) CreateThread(topic string) (conv Thread, err error) {
	/*
	fmt.Println("zz1")
	statement := "insert into threads(topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
	stmt, err :=Db.Prepare(statement)
	defer stmt.Close()
	fmt.Println("zz2")
	err = stmt.QueryRow(topic, 10, time.Now()).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	fmt.Println("zz3")
	*/

	fmt.Println("zz1")
	stmt, err := Db.Prepare("INSERT threads SET topic=?,user_id=?, uuid=?, created_at=?")
	fmt.Println("zz2")
	res , err := stmt.Exec(topic, 10, "uuid123", time.Now())
	fmt.Println("zz2.1")
	if res==nil {fmt.Println("nil res !!!!")} else  { log.Println(res)}
	fmt.Println("zz2.2")
	if err==nil {fmt.Println("nil err !!!!")} else  { log.Fatal(err)}
	fmt.Println("zz2.3")


	fmt.Println("zz3")


	return

}

func createThread(writer http.ResponseWriter, request *http.Request) {
	user :=User{}
	err :=request.ParseForm()
	fmt.Println("cc1")
	if err !=nil {
		log.Println(err)
	}
	fmt.Println("cc2")
	topic := request.PostFormValue("topic")
	log.Println(topic)


	if _, err :=user.CreateThread(topic); err!=nil {
		log.Println("cannot create thread!!!")
	}
	fmt.Println("cc3")
	http.Redirect(writer, request, "/", 302)


}
func newThread(writer http.ResponseWriter, request *http.Request) {

	tmpl_files := []string{
		"templates/layout.html",
		"templates/public.navbar.html",
		"templates/new.thread.html",
	}
	templates := template.Must(template.ParseFiles(tmpl_files...))
	templates.ExecuteTemplate(writer, "layout", nil)
}


func main() {

	fmt.Println("starting server ...")
	mux := http.NewServeMux()

	//handling of files must be in main function.
	//note : found inconsistent behavior if this is inside the index function.
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/thread/new", newThread)
	mux.HandleFunc("/thread/create", createThread)

	server := &http.Server{
		Addr:    "0.0.0.0:7000",
		Handler: mux,
	}
	fmt.Println("listening...")
	server.ListenAndServe()
}
