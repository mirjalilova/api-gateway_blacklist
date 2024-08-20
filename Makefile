CURRENT_DIR=$(shell pwd)

proto-gen:
	./internal/scripts/gen-proto.sh ${CURRENT_DIR}
exp:
	export DBURL='postgres://postgres:1234@localhost:5432/blacklist?sslmode=disable'

mig-run:
	migrate create -ext sql -dir migrations -seq create_table

mig-up:
	migrate -database 'postgres://postgres:1234@localhost:5432/blacklist?sslmode=disable' -path migrations up

mig-down:
	migrate -database 'postgres://postgres:1234@localhost:5432/blacklist?sslmode=disable' -path migrations down


gen-protoAll:
	protoc --go_out=./ \
	--go-grpc_out=./ \
	blacklist_submodule/black_list_service/*.proto

swag-gen:
	~/go/bin/swag init -g ./api/router.go -o docs force 1	
run:
	go run main.go