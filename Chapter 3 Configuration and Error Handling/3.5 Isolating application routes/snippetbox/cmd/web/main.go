package main

/*
Alex Edwards - Let's Go

Chapter 3 Configuration & Error Handling

Section 3.5 Isolating App Routes
*/

import (
	"flag"
	"log"
	"net/http"
	"os"
)

/*
The main() function is now limited to:

1. Parsing the runtime configuration settings for the application.
2. Establishing the dependencies for the handlers.
3. Running the HTTP server.
*/

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		// Call the new app.routes() method to get the servemux containing
		// our routes.
		Handler: app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
