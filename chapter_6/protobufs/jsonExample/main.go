package main

import (
	"encoding/json"
	"fmt"

	pb "github.com/victoryRo/Restful/chapter_6/protobufs/profiles"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	body, _ := json.Marshal(p)
	fmt.Println(string(body))
}
