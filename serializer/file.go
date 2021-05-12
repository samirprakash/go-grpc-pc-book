package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

// WriteProtobufToJSONFile writes a protocol buffer message to a JSON file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message JSON : %w", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to JSON file : %w", err)
	}

	return nil
}

// WriteProtobufToBinaryFile writes a protocol buffer message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal protobuf message to binary : %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file : %w", err)
	}

	return nil
}

// ReadProtobufFromBinaryFile reads protocol message from a binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file : %w ", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary data to proto message : %w", err)
	}

	return nil
}
