# GoKafka
1. cd $GOPATH/src/GoKafka
docker-compose -f docker-compose.yml up
2. To Run Producer
cd $GOPATH/src/GoKafka/Producer
go run main.go
3. To Run Consumer
cd $GOPATH/src/GoKafka/Consumer
go run main.go

## Update producer.go to change Topic name or message.
