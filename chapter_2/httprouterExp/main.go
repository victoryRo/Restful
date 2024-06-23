package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, arguments ...string) string {
	out, _ := exec.Command(command, arguments...).Output()
	return string(out)
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	response := getCommandOutput("/usr/local/bin/go", "version")
	_, _ = io.WriteString(w, response)
}

func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	res := getCommandOutput("/bin/cat", params.ByName("name"))
	_, _ = io.WriteString(w, res)
}

func main() {
	router := httprouter.New()

	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)

	fmt.Println("run server port :3007")
	log.Fatal(http.ListenAndServe(":3007", router))
}
