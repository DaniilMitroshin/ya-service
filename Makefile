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
migrate-force:
	migrate -path migrations -database "$(CONN_STRING)" force $(version)

migrate-version:
	migrate -path migrations -database "$(CONN_STRING)" version

.PHONY: run migrate-up migrate-down migrate-new migrate-force migrate-version