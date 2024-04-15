build:
	go build -o ./bin/juaz ./../../

biometry: 
	./bin/juaz -file=./protos/biometry/biometry.juaz

user: 
	./bin/juaz -file=./protos/user/user.juaz

juaz: build biometry user
