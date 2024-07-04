package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {	
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	host := flag.String("host", "127.0.0.1", "HTTP network Host")

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	
	
	
	// lOGGING
	// For generic logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// For error logging
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	
	
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
		
	
	// Initialize a new instance of application containing the dependencies
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	
	server := &http.Server {
		Addr: *host + *addr,
		Handler: app.routes(),
		ErrorLog: errorLog,
	}


	infoLog.Println("Starting server on http://" + *host + *addr)
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}



// openDB opens a connection to the MySQL database
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	
	// Ping the database to test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}