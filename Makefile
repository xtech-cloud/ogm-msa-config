APP_NAME := ogm-config
BUILD_VERSION   := $(shell git tag --contains)
BUILD_TIME      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD )

.PHONY: build
build:
	go build -ldflags \
		"\
		-X 'main.BuildVersion=${BUILD_VERSION}' \
		-X 'main.BuildTime=${BUILD_TIME}' \
		-X 'main.CommitID=${COMMIT_SHA1}' \
		"\
		-o ./bin/${APP_NAME}

.PHONY: run
run:
	./bin/${APP_NAME}

.PHONY: install
install:
	go install

.PHONY: clean
clean:
	rm -rf /tmp/ogm-config.db

.PHONY: call
call:
	gomu --registry=etcd --client=grpc call xtc.ogm.config Healthy.Echo '{"msg":"hello"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Write '{"path":"/test/1.json", "content":"{\"msg\":\"test1\"}"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Write '{"path":"/test/2.json", "content":"{\"msg\":\"test2\"}"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Read '{"path":"/test/1.json"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Read '{"path":"/test/2.json"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.List '{"offset":1, "count":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.List '{"offset":0, "count":1}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Search '{"offset":0, "count":2, "path": "test"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Search '{"offset":1, "count":1, "path": "test"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Get '{"uuid":"df4d9fca6c229eb68c6b0500e622c0fc"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Delete '{"uuid":"df4d9fca6c229eb68c6b0500e622c0fc"}'
	gomu --registry=etcd --client=grpc call xtc.ogm.config Text.Delete '{"uuid":"163146e36dab318699ee9dd1f39f358b"}'

.PHONY: post
post:
	curl -X POST -d '{"msg":"hello"}' localhost/ogm/config/Healthy/Echo                                                                                     1

.PHONY: dist
dist:
	mkdir dist
	tar -zcf dist/${APP_NAME}-${BUILD_VERSION}.tar.gz ./bin/${APP_NAME}

.PHONY: docker
docker:
	docker build -t xtechcloud/${APP_NAME}:${BUILD_VERSION} .
	docker rm -f ${APP_NAME}
	docker run --restart=always --name=${APP_NAME} --net=host -v /data/${APP_NAME}:/ogm -e MSA_REGISTRY_ADDRESS='localhost:2379' -e MSA_CONFIG_DEFINE='{"source":"file","prefix":"/ogm/config","key":"${APP_NAME}.yaml"}' -d xtechcloud/${APP_NAME}:${BUILD_VERSION}
	docker logs -f ${APP_NAME}
