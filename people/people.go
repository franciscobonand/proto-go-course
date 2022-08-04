package main

import (
	"fmt"
	"log"
	"os"

	pb "github.com/franciscobonand/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s ADDRESS_BOOK_FILE\n", os.Args[0])
	}

	fname := os.Args[1]

	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalf("Unable to read file %s: %v", fname, err)
	}

	book := &pb.AddressBook{}
	if err = proto.Unmarshal(in, book); err != nil {
		log.Fatalf("Failed to unmarshal address book: %v", err)
	}

	for _, person := range book.People {
		fmt.Printf("ID: %d\nName: %s\nEmail: %s\n", person.Id, person.Name, person.Email)
		for _, phone := range person.Phones {
			switch phone.Type {
			case pb.Person_HOME:
				fmt.Printf("Home number: %s\n", phone.Number)
			case pb.Person_WORK:
				fmt.Printf("Work number: %s\n", phone.Number)
			case pb.Person_MOBILE:
				fmt.Printf("Mobile number: %s\n", phone.Number)
			}
		}
		fmt.Println()
	}
}
