/* Test project based on Upload.go from Christopher
 */

package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

const version string = "1.0"

var uploadTemplate, _ = template.ParseFiles("/Users/Christopher/Documents/" +
	"Programmering/go/libs/src/github.com/christopherL91/Upload/Upload.html")

func fileserve(rw http.ResponseWriter, req *http.Request) {
	uploadTemplate.Execute(rw, nil)
}

func upload(rw http.ResponseWriter, req *http.Request) {

	fmt.Println("Incoming message...")
	if req.Method != "POST" {
		fmt.Println("ERROR not POST")
		uploadTemplate.Execute(rw, nil)
		return
	}
	file, handler, err := req.FormFile("file")
	defer file.Close()
	if err != nil {
		fmt.Println("Something happended")
	}

	fmt.Println("Name of file incoming ", handler.Filename)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintln(rw, "ERROR")
		fmt.Println("Error reading data...", err)
	}
	err = ioutil.WriteFile(handler.Filename, data, 0700)
	fmt.Println("Writing file to disc")
	if err != nil {
		fmt.Fprintln(rw, "ERROR")
		fmt.Println("Error writing file")
	}
	timeNow := time.Now()
	fmt.Fprintf(rw, "Successfull uploading file: %s\n", handler.Filename)
	fmt.Fprintf(rw, "Time : %s\n", timeNow.Format(time.Kitchen))
	fmt.Println("Successfull upload ", handler.Filename)
}

func sayDate(rw http.ResponseWriter, req *http.Request) {
	timeNow := time.Now()
	fmt.Fprintf(rw, "Time now %s", timeNow.Format(time.Kitchen))

}

func sayName(rw http.ResponseWriter, req *http.Request) {
	remPartOfURL := req.URL.Path[len("/name/"):]
	fmt.Fprintf(rw, "Hello %s", remPartOfURL)
}

func sayVersion(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "Version: %s", version)
}

func main() {
	cores := flag.Int("cores", 1, "The number of cores used")
	port := flag.Int("port", 4000, "The port number that the server will use")
	flag.Parse()
	runtime.GOMAXPROCS(*cores)

	fmt.Println("Server started on port:", *port)
	http.HandleFunc("/name/", sayName)
	http.HandleFunc("/date/", sayDate)
	http.HandleFunc("/upload/", upload)
	http.HandleFunc("/version/", sayVersion)
	http.HandleFunc("/html/", fileserve)
	//http.Handle("/look/", http.StripPrefix("/file", http.FileServer(http.Dir("."))))
	http.Handle("/look/", http.StripPrefix("/look/", http.FileServer(http.Dir("/app"))))

	err := http.ListenAndServe(":"+strconv.Itoa(*port), nil)

	if err != nil {
		log.Fatal("Error ListenAndServe", err)
		fmt.Println("ERROR")
	}
}
