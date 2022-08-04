BIN = proto-go-course
PROTO_DIR = proto
PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

build: 	generate
	go build -o ${BIN} .

generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. ${PROTO_DIR}/*.proto

practice: add_person list_people

add_person: address/address.go proto/addressbook.pb.go
	cd address && go build -o ../bin/add_person address.go

list_people: people/people.go proto/addressbook.pb.go
	cd people && go build -o ../bin/list_people people.go

clean:
	rm ${PROTO_DIR}/*.pb.go
	rm ${BIN}