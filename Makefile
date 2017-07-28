all: build

APP=lowest-common-ancestor

PORT?=8888
CONFIG_FILE?=./cmd/default_config.json

vendor: clean
	go get -u github.com/Masterminds/glide \
	&& glide install \
	&& glide up

build: vendor
	cd cmd \
	&& go build -o ../${APP}

run:
	PORT=${PORT} CONFIG_FILE=${CONFIG_FILE} ./${APP}

clean:
	rm -f ${APP}
