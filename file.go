package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func writeToFile(fname string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Unable to serialize to bytes: ", err)
	}

	if err = ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Unable to write to file: ", err)
	}

	fmt.Println("Data has been written!")
}

func readFromFile(fname string, pb proto.Message) proto.Message {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Unable to read file: ", err)
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Unable to unmarshal file: ", err)
	}

	return pb
}
