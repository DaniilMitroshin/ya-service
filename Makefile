include .env
export

run:
	go run main.go

migrate-up:
	migrate -path migrations -database ${CONN_STRING} up

migrate-down:
	migrate -path migrations -database ${CONN_STRING} down 1

migrate-new: 
	migrate create -ext sql -dir migrations -seq ${name}
