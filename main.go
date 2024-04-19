package main

import (
	pb "example-protobuf-go/codegen"
	"fmt"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{
				Number: "555-4321",
				Type:   pb.PhoneType_PHONE_TYPE_HOME,
			},
		},
	}

	fmt.Println(p)
}
