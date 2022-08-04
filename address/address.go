package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	pb "github.com/franciscobonand/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func promptPersonInfo() (*pb.Person, error) {
	pers := &pb.Person{}
	rd := bufio.NewReader(os.Stdin)

	fmt.Println("Enter person ID number:")
	if _, err := fmt.Fscanf(rd, "%d\n", &pers.Id); err != nil {
		return nil, err
	}

	fmt.Println("Enter person name:")
	name, err := rd.ReadString('\n')
	if err != nil {
		return nil, err
	}
	pers.Name = strings.TrimSpace(name)

	fmt.Println("Enter person email:")
	email, err := rd.ReadString('\n')
	if err != nil {
		return nil, err
	}
	pers.Email = strings.TrimSpace(email)

	for {
		fmt.Println("Enter phone number. Leave blank to proceed")
		phone, err := rd.ReadString('\n')
		if err != nil {
			return nil, err
		}
		phone = strings.TrimSpace(phone)
		if phone == "" {
			break
		}

		fmt.Println("Enter phone type (mobile|home|work):")
		ptype, err := rd.ReadString('\n')
		if err != nil {
			return nil, err
		}
		ptype = strings.TrimSpace(ptype)

		switch ptype {
		case "mobile":
			pers.Phones = append(pers.Phones,
				&pb.Person_PhoneNumber{
					Number: phone,
					Type:   pb.Person_MOBILE,
				})
		case "home":
			pers.Phones = append(pers.Phones,
				&pb.Person_PhoneNumber{
					Number: phone,
					Type:   pb.Person_HOME,
				})
		case "work":
			pers.Phones = append(pers.Phones,
				&pb.Person_PhoneNumber{
					Number: phone,
					Type:   pb.Person_WORK,
				})
		default:
			fmt.Println("Invalid phone type, should be 'mobile', 'home' or 'work'")
		}
	}

	pers.LastUpdated = timestamppb.Now()

	return pers, nil
}

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
		log.Fatalf("Unable to get current address book: %v", err)
	}

	newPers, err := promptPersonInfo()
	if err != nil {
		log.Fatalf("Failed to get new person data: %v", err)
	}

	book.People = append(book.People, newPers)

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalf("Unable to marshal updated address book: %v", err)
	}

	if err = os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalf("Unable to save changes into %s: %v", fname, err)
	}

	fmt.Println("Address Book updated successfully!")
}
