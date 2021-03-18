package main

import (
	// "encoding/json"
	_ "image/png"

	// pigo "github.com/esimov/pigo/core"

	// "crypto/md5"

	// "errors"
	"fmt"
	// "io"
	// "io/ioutil"
	"log"
	"net/http"
	// "os"
	// "os/signal"
	// "regexp"
	// "strconv"
	// "strings"
	// "syscall"
	// "time"

	"github.com/alecthomas/kong"
	// "github.com/olekukonko/tablewriter"
	// "github.com/davecgh/go-spew/spew"
	// execute "github.com/alexellis/go-execute/pkg/v1"
	// "github.com/go-cmd/cmd"
	// "gocv.io/x/gocv"
	// bolt "go.etcd.io/bbolt"
	// import "github.com/didi/gendry/manager"
	// import "database/sql"
	"github.com/rs/cors"
)

var CLI struct {
	Serve struct {
		Port string
	} `cmd help:"Running the server."`
}

func main() {

	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "serve":
		port := ""
		if CLI.Serve.Port == "" {
			port = "8123"
		} else {
			port = CLI.Serve.Port
		}
		serve(port)
	default:
		panic(ctx.Command())
	}
}

func serve(port string) {
	fs := http.FileServer(http.Dir("."))
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)
	mux.Handle("/demo.css", fs)
	mux.Handle("/flv.js", fs)

	// handler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(":"+port, handler))
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	htmldata := renderHtml()
	fmt.Fprintf(w, htmldata)
}
