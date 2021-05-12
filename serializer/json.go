package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshal := protojson.MarshalOptions{
		Indent:          " ",
		EmitUnpopulated: true,
		UseEnumNumbers:  false,
		UseProtoNames:   false,
	}
	jsonBytes, err := marshal.Marshal(message)
	return string(jsonBytes), err
}
