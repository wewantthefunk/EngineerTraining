package main

import (
    "io"
    "log"
    "net/http"
    "os"
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

var db *sql.DB

type userInformation struct {
    id int
    Email string
    Password string
}

func main() {
    http.HandleFunc("/", ExampleHandler)
    http.HandleFunc("/db", ShowQuery)
    http.HandleFunc("/test", ShowTest)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    var err error

    connStr := "postgres://postgres:mysecretpassword@some-postgres/mytestdb?sslmode=disable"
    db, err = sql.Open("postgres", connStr)

    if err != nil {
        log.Println("DB open error")
        panic(err)
    }

    perr := db.Ping()

    if perr != nil {
        log.Println("pinged DB error")
        panic(perr)
    } else {
      log.Println("pinged DB successful")
    }

    log.Println("** Service Started on Port " + port + " **")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}


func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"ok"}`)
}

func ShowTest(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"test"}`)
}

func ShowQuery(w http.ResponseWriter, r *http.Request) {
/*    ping := db.Ping()
    if ping != nil {
        log.Println("ping failed")
        return
    }*/
    //
    rows, err := db.Query("SELECT * FROM public.mytable")

    if err != nil {

	log.Println("status1")
        //http.Error(w, http.StatusText(500), "error with select statement")//http.StatusInternalServerError)
        w.Header().Add("Content-Type", "application/json")
        io.WriteString(w, `{"status1":err.Error()}`)
        return
    }
    defer rows.Close()

    usrs := make([]userInformation, 0)

    for rows.Next() {
        usr := userInformation{}
        err := rows.Scan(&usr.id, &usr.Email, &usr.Password)
        if err != nil {
	    log.Println("status2")
            log.Println(err)
            w.Header().Add("Content-Type", "application/json")
            io.WriteString(w, `{"status2":err.Error()}`) 
            return
        }
        usrs = append(usrs, usr)
    }

    if err = rows.Err(); err != nil {
	log.Println("status3")
        //http.Error(w, http.StatusText(500), 500)
        w.Header().Add("Content-Type", "application/json")
        io.WriteString(w, `{"status3":err.Error()}`)
        return
    }

    for _, usr := range usrs {
        fmt.Fprintf(w, "%d %s %s", usr.id, usr.Email, usr.Password)
    }

    log.Println("ping succeeded")
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status1":"ping succeeded"}`)
}
