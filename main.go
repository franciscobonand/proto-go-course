package main

import (
	"fmt"
	"reflect"

	pb "github.com/franciscobonand/proto-go-course/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          42,
		IsSimple:    true,
		Name:        "A name",
		SampleLists: []int32{1, 2, 3, 4, 5, 6},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   42,
			Name: "My name",
		},
		MultDummies: []*pb.Dummy{
			{Id: 43, Name: "Another name"},
			{Id: 44, Name: "And another name"},
			{Id: 45, Name: "And another one"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: 3, //pb.EyeColor_EYE_COLOR_BROWN,
	}
}

func doOneOf(msg interface{}) {
	switch x := msg.(type) {
	case *pb.Result_Id:
		fmt.Println(msg.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(msg.(*pb.Result_Message).Message)
	default:
		fmt.Println(fmt.Errorf("message has unexpected type: %v", x))
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":    {Id: 42},
			"anotId":  {Id: 43},
			"anotId1": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"

	writeToFile(path, p)
	msg := &pb.Simple{}
	readFromFile(path, msg)

	fmt.Println(msg)
}

func doToJSON(p proto.Message) string {
	return toJSON(p)
}

func doFromJSON(jsonStr string, t reflect.Type) proto.Message {
	msg := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonStr, msg)

	return msg
}

func main() {
	// fmt.Println(doSimple())

	// fmt.Println(doComplex())

	// fmt.Println(doEnum())

	// doOneOf(&pb.Result_Id{Id: 42})
	// doOneOf(&pb.Result_Message{Message: "A new message has arrived"})
	// doOneOf(2.0)

	// fmt.Println(doMap())

	// doFile(doSimple())

	// jsonStr := doToJSON(doSimple())
	// msg := doFromJSON(jsonStr, reflect.TypeOf(pb.Simple{}))
	// fmt.Println(jsonStr)
	// fmt.Println(msg)
	// fmt.Println(doFromJSON(`{"id":42, "unkw": "test"}`, reflect.TypeOf(pb.Simple{})))
}
