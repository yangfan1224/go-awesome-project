package proto

import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func WriteMessage()(error){
	book := &AddressBook{}
	book.People = []*Person{
		&Person{   Id:    1234,
			Name:  "John Doe",
			Email: "jdoe@example.com",
			Phones: []*Person_PhoneNumber{
				{Number: "555-4321", Type: Person_HOME},
			},},
	}

	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
		return err
	}
	if err := ioutil.WriteFile("adressbook", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
		return err
	}
	return nil
}

func ReadMessage()(error){
	in, err := ioutil.ReadFile("adressbook")
	if err != nil {
		log.Fatalln("Error reading file:", err)
		return err
	}
	book := &AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
		return err
	}
	log.Printf("book is %v", book)
	return nil
}
