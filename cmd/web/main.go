package main

import (
	"flag"
	"os"
	"log"
	"net/http"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parses the command-line flag, giving a value to addr
	flag.Parse()

	// Loggers for making information and error messages
	infoLog := log.New(
		os.Stdout, 
		"INFO\t", 
		log.Ldate|log.Ltime)

	errorLog := log.New(
		os.Stderr, 
		"ERROR\t", 
		log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
