gen:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative proto/*.proto

clean:
	rm pb/*.go

local:
	go run main.go