package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	// cfg is the global configuration for the server. It's read in at startup from
	// the config.json file and enviornment variables, see config.go for more info.
	cfg *config
	// log output
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	// application database connection
	appDB *sql.DB
)

func main() {
	var err error
	cfg, err = initConfig(os.Getenv("GOLANG_ENV"))
	if err != nil {
		// panic if the server is missing a vital configuration detail
		panic(fmt.Errorf("server configuration error: %s", err.Error()))
	}

	connectToAppDb()

	s := &http.Server{}
	m := http.NewServeMux()
	m.HandleFunc("/.well-known/acme-challenge/", CertbotHandler)

	m.Handle("/", middleware(HealthCheckHandler))

	// m.Handle("/v0/users", middleware())
	// m.Handle("/v0/users/", middleware())
	m.Handle("/v0/urls", middleware(UrlHandler))
	m.Handle("/v0/urls/", middleware(UrlsHandler))

	// m.Handle("/v0/content", middleware())
	// m.Handle("/v0/content/", middleware())
	// m.Handle("/v0/metadata", middleware())
	// m.Handle("/v0/metadata/", middleware())
	// m.Handle("/v0/consensus", middleware())
	// m.Handle("/v0/consensus/", middleware())
	// m.Handle("/v0/primers", middleware())
	// m.Handle("/v0/primers/", middleware())
	// m.Handle("/v0/sources", middleware())
	// m.Handle("/v0/sources/", middleware())

	// connect mux to server
	s.Handler = m

	// print notable config settings
	printConfigInfo()

	// fire it up!
	fmt.Println("starting server on port", cfg.Port)

	// start server wrapped in a log.Fatal b/c http.ListenAndServe will not
	// return unless there's an error
	logger.Fatal(StartServer(cfg, s))
}