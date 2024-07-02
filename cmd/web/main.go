package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	host := flag.String("host", "127.0.0.1", "HTTP network Host")
	flag.Parse()

	// Open files for writing
	infoWriteLog, _err := os.OpenFile("info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666) 
	if _err != nil {
		log.Fatal(_err)
	}
	defer infoWriteLog.Close()

	errorWriteLog, _err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if _err != nil {
		log.Fatal(_err)
	}
	defer errorWriteLog.Close()



	// lOGGING
	// For generic logging
	infoLog := log.New(infoWriteLog, "INFO\t", log.Ldate|log.Ltime)
	// For error logging
	errorLog := log.New(errorWriteLog, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)






	mux := http.NewServeMux()

	//	routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
	

	// Command line flags

	// Static files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := &http.Server {
		Addr: *host + *addr,
		Handler: mux,
		ErrorLog: errorLog,
	}


	infoLog.Println("Starting server on http://" + *host + *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
