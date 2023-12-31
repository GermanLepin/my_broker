FRONT_END_BINARY=binary_file/frontEndServiceApp
AUTHENTICATION_BINARY=binary_file/authenticationServiceApp
BROKER_BINARY=binary_file/brokerServiceApp
LOGGER_BINARY=binary_file/loggerServiceApp
MAIL_BINARY=binary_file/mailServiceApp
LISTENER_BINARY=binary_file/listenerServiceApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_authentication build_broker build_logger build_mail build_listener
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_authentication: builds the authentication binary as a linux executable
build_authentication:
	@echo "Building authentication binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTHENTICATION_BINARY} ./cmd/service
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/service
	@echo "Done!"

## build_logger: builds the logger binary as a linux executable
build_logger:
	@echo "Building logger binary..."
	cd ../logger-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LOGGER_BINARY} ./cmd/service
	@echo "Done!"

## build_mail: builds the mail binary as a linux executable
build_mail:
	@echo "Building mail binary..."
	cd ../mail-service && env GOOS=linux CGO_ENABLED=0 go build -o ${MAIL_BINARY} ./cmd/service
	@echo "Done!"

## build_listener: builds the listener binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ../listener-service && env GOOS=linux CGO_ENABLED=0 go build -o ${LISTENER_BINARY} ./cmd/service
	@echo "Done!"

## build_front_end: builds the frone end binary
build_front_end:
	@echo "Building front-end binary..."
	cd ../front-end-service && env CGO_ENABLED=0 go build -o ${FRONT_END_BINARY} ./cmd/web
	@echo "Done!"

## front_end_start: starts the front-end
front_end_start: build_front_end
	@echo "Starting front-end"
	cd ../front-end-service && ./${FRONT_END_BINARY} &

## front_end_stop: stops the front-end
front_end_stop:
	@echo "Stopping front-end..."
	@-pkill -SIGTERM -f "./${FRONT_END_BINARY}"
	@echo "Stopped front-end!"
