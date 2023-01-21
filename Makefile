docker:
	docker-compose build && docker-compose up

swag:
	swag init -g cmd/subscriber/main.go

pub:
	go run github.com/Owner-maker/nats-learning/cmd/publisher

test:
	go test -v ./...