BIN_DIR = bin
BROKER_DIR = broker/cmd/api
BROKER_BIN = broker
CLIENT_DIR = client/cmd/api
CLIENT_BIN = client
FRONT_DIR = front/cmd
FRONT_BIN = front
SECPOSRET_DIR = secposretriever/cmd
SECPOSRET_BIN = secposret

test:
	cd bin && pwd
	pwd

build_broker:
	@echo "Building binary..."
	cd ${BROKER_DIR} && env GOOS=linux CGO_ENABLED=0 go build -o ../../../${BIN_DIR}/${BROKER_BIN} 
	@echo "Done!"

build_client:
	@echo "Building binary..."
	cd ${CLIENT_DIR} && env GOOS=linux CGO_ENABLED=0 go build -o ../../../${BIN_DIR}/${CLIENT_BIN} 
	@echo "Done!"

build_front:
	@echo "Building ${FRONT_BIN} binary..."
	cd ${FRONT_DIR} && env GOOS=linux CGO_ENABLED=0 go build -o ../../${BIN_DIR}/${FRONT_BIN} 
	@echo "Done!"

build_secposret:
	@echo "Building ${SECPOSRET_BIN} binary..."
	cd ${SECPOSRET_DIR} && env GOOS=linux CGO_ENABLED=0 go build -o ../../${BIN_DIR}/${SECPOSRET_BIN} 
	@echo "Done!"
