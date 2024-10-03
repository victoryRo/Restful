package main

import (
	jsonparse "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID string
}

type Book struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book
	absPath, _ := filepath.Abs("books.json")
	raw, readerr := os.ReadFile(absPath)
	if readerr != nil {
		log.Println("error:", readerr)
		os.Exit(1)
	}

	marshalerr := jsonparse.Unmarshal(raw, &books)
	if marshalerr != nil {
		log.Println("error:", marshalerr)
		os.Exit(1)
	}

	for _, book := range books {
		if book.ID == args.ID {
			*reply = book
			break
		}
	}
	return nil
}

func main() {
	// Create a new RPC server
	s := rpc.NewServer()
	// Register the type of data requested as JSON
	s.RegisterCodec(json.NewCodec(), "application/json")
	// Register the service by creating a new JSON server
	_ = s.RegisterService(new(JSONServer), "")

	router := mux.NewRouter()
	router.Handle("/rpc", s)
	fmt.Println("RPC server run on port :1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}

// go run .
// Ahora bien, ¿tenemos que desarrollar un cliente? No necesariamente,
// debido a que un cliente puede ser un curl programa,
// ya que el servidor RPC atiende solicitudes a través de HTTP,
// debemos publicar JSON con una ID de libro para obtener los detalles.
// Entonces, inicia otro shell y ejecuta esta curl solicitud:

/**
curl -X POST \
   http://localhost:1234/rpc \
   -H 'cache-control: no-cache' \
   -H 'content-type: application/json' \
   -d '{
     "method": "JSONServer.GiveBookDetail",
     "params": [{
       "ID": "1234"
     }],
     "id": "1"
   }'
*/
