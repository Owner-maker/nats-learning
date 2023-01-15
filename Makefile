docker:
	docker-compose build && docker-compose up

swag:
	swag init -g cmd/subscriber/main.go

sub:
	go run github.com/Owner-maker/nats-learning/cmd/subscriber

pub:
	go run github.com/Owner-maker/nats-learning/cmd/publisher

