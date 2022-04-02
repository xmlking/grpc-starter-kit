# Usage:
# make        	# compile all binary
# make clean  	# remove ALL binaries and objects
# make release  # add git TAG and push
GITHUB_REPO_OWNER 				:= xmlking
GITHUB_REPO_NAME 				:= grpc-starter-kit
GITHUB_RELEASES_UI_URL 			:= https://github.com/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)/releases
GITHUB_RELEASES_API_URL 		:= https://api.github.com/repos/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)/releases
GITHUB_RELEASE_ASSET_URL		:= https://uploads.github.com/repos/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)/releases
GITHUB_DEPLOY_API_URL			:= https://api.github.com/repos/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)/deployments
DOCKER_REGISTRY 				:= ghcr.io
# DOCKER_REGISTRY 				:= us.gcr.io
DOCKER_CONTEXT_PATH 			:= $(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME)
# DOCKER_REGISTRY 				:= docker.io
# DOCKER_CONTEXT_PATH 			:= xmlking

VERSION					:= $(shell git describe --tags || echo "HEAD")
GOPATH					:= $(shell go env GOPATH)
CODECOV_FILE 			:= build/coverage.txt
TIMEOUT  				:= 60s
# don't override
GIT_TAG					:= $(shell git describe --tags --abbrev=0 --always --match "v*")
GIT_DIRTY 				:= $(shell git status --porcelain 2> /dev/null)
GIT_BRANCH  			:= $(shell git rev-parse --abbrev-ref HEAD)
HAS_GOVVV				:= $(shell command -v govvv 2> /dev/null)
HAS_KO					:= $(shell command -v ko 2> /dev/null)
HTTPS_GIT 				:= https://github.com/$(GITHUB_REPO_OWNER)/$(GITHUB_REPO_NAME).git

# Type of service e.g api, service, web, cmd (default: "service")
TYPE = $(or $(word 2,$(subst -, ,$*)), service)
override TYPES:= service
# Target for running the action
TARGET = $(word 1,$(subst -, ,$*))

override VERSION_PACKAGE = $(shell go list ./internal/config)

# $(warning TYPES = $(TYPE), TARGET = $(TARGET))
# $(warning VERSION = $(VERSION), HAS_GOVVV = $(HAS_GOVVV), HAS_KO = $(HAS_KO))

.PHONY: all tools check_dirty clean sync
.PHONY: proto proto_lint proto_breaking proto_format proto_generate proto_shared
.PHONY: lint lint-% outdated
.PHONY: format format-%
.PHONY: build build-%
.PHONY: run run-%
.PHONY: docker_clean docker docker-% docker_push
.PHONY: kustomize build/kustomize
.PHONY: release/draft release/publish
.PHONY: deploy/e2e deploy/prod

all: build

################################################################################
# Target: tools
################################################################################

tools:
	@echo "==> Installing dev tools"
	# go install github.com/ahmetb/govvv@latest
	# go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	# go install github.com/bufbuild/buf/cmd/buf@latest
	# go install github.com/oligot/go-mod-upgrade@latest

check_dirty:
ifdef GIT_DIRTY
	$(error "Won't run on a dirty working copy. Commit or stash and try again.")
endif

clean:
	@for d in ./build/*-service; do \
		echo "Deleting $$d;"; \
		rm -f $$d; \
	done
	@for f in */*/pkged.go ; do \
		echo "Deleting $$f;"; \
		rm -f $$f; \
	done

################################################################################
# Target: go-mod                                                               #
################################################################################

sync:
	@echo Removing all go.sum files and download sync deps....
	@for d in `find * -name 'go.mod'`; do \
		pushd `dirname $$d` >/dev/null; \
	    echo delete and sync $$d; \
	    rm -f go.sum; \
	    go mod download; \
	    go mod tidy -compat=1.18; \
		popd >/dev/null; \
	done

verify:
	@echo Go mod verify and tidy....
	@for d in `find * -name 'go.mod'`; do \
		pushd `dirname $$d` >/dev/null; \
		echo verifying $$d; \
		go mod verify; \
		go mod tidy -compat=1.18; \
		popd >/dev/null; \
	done

outdated:
	@go-mod-upgrade

################################################################################
# Target: proto                                                                #
################################################################################

proto_clean:
	@echo "Deleting generated Go files....";
	@for f in ./gen/**/**/**/**/*.pb.*; do \
		echo ✓ deleting: $$f; \
		rm -f $$f; \
	done
	@for f in ./gen/**/**/**/*.pb.*; do \
		echo ✓ deleting: $$f; \
		rm -f $$f; \
	done

proto_lint:
	@echo "Linting protos";
	@${GOPATH}/bin/buf lint
	@echo "✓ Proto: Linted"

proto_breaking:
	@echo "Checking proto breaking changes";
	@${GOPATH}/bin/buf breaking --against '.git#branch=master'
#	@${GOPATH}/bin/buf breaking --against "$(HTTPS_GIT)#branch=master"
	@echo "✓ Proto: Breaking"

# I prefer VS Code's proto plugin to format my code then prototool
proto_format: proto_lint
	@echo "Formatting protos";
#	@${GOPATH}/bin/prototool format -w proto;
	@echo "✓ Proto: Formatted"

proto_check: proto_lint proto_breaking proto_format

proto_generate:
	@echo "Generating protos";
	@${GOPATH}/bin/buf generate ;
	@${GOPATH}/bin/buf generate --template buf.gen.tag.yaml ;

proto: proto_check proto_clean proto_generate

################################################################################
# Target: lints                                                                #
################################################################################

lint lint-%:
	@if [ -z $(TARGET) ]; then \
		echo "Linting all go"; \
		${GOPATH}/bin/golangci-lint run ./... --deadline=5m --config=.github/linters/.golangci.yml; \
		echo "✓ Go: Linted"; \
	else \
		echo "Linting go in ${TARGET}-${TYPE}..."; \
		${GOPATH}/bin/golangci-lint run ./${TYPE}/${TARGET}/... --config=.github/linters/.golangci.yml; \
		echo "✓ Go: Linted in ${TYPE}/${TARGET}..."; \
	fi

# @clang-format -i $(shell find . -type f -name '*.proto')

format format-%:
	@if [ -z $(TARGET) ]; then \
		echo "Formatting all go"; \
		gofmt -l -w . ; \
		echo "✓ Go: Formatted"; \
	else \
		echo "Formatting go in ${TYPE}/${TARGET}..."; \
		gofmt -l -w ./${TYPE}/${TARGET}/ ; \
		echo "✓ Go: Formatted in ${TYPE}/${TARGET}..."; \
	fi


################################################################################
# Target: generate                                                                #
################################################################################
generate:
	go generate ./...

generate-version:
	go generate ./internal/version

generate-mockery:
	go generate ./service/emailer/service
	go generate ./service/account/repository

################################################################################
# Target: build                                                                #
################################################################################

build build-%: generate-version
	@if [ -z $(TARGET) ]; then \
		for type in $(TYPES); do \
			echo "Building Type: $${type}..."; \
			for _target in $${type}/*/; do \
				temp=$${_target%%/}; target=$${temp#*/}; \
				echo "\tBuilding $${target}-$${type}"; \
				CGO_ENABLED=0 GOOS=linux go build -o build/$${target}-$${type} -a -trimpath ./$${type}/$${target}; \
			done \
		done \
	else \
		echo "Building ${TARGET}-${TYPE}"; \
		go build -o  build/${TARGET}-${TYPE} -a -trimpath ./${TYPE}/${TARGET}; \
	fi

################################################################################
# Target: tests                                                                #
################################################################################

TEST_TARGETS := test-default test-bench test-unit test-inte test-e2e test-race test-cover
.PHONY: $(TEST_TARGETS) check test tests
test-bench:   	ARGS=-run=__absolutelynothing__ -bench=. ## Run benchmarks
test-unit:   	ARGS=-short        					## Run only unit tests
test-inte:   	ARGS=-run Integration       		## Run only integration tests
test-e2e:   	ARGS=-run E2E       				## Run only E2E tests
test-race:    	ARGS=-race         					## Run tests with race detector
test-cover:   	ARGS=-cover -short -coverprofile=${CODECOV_FILE} -covermode=atomic ## Run tests in verbose mode with coverage reporting
$(TEST_TARGETS): NAME=$(MAKECMDGOALS:test-%=%)
$(TEST_TARGETS): test
check test tests: generate-version
	@if [ -z $(TARGET) ]; then \
		echo "Running $(NAME:%=% )tests for all"; \
		go test -timeout $(TIMEOUT) $(ARGS) ./... ; \
	else \
		echo "Running $(NAME:%=% )tests for ${TARGET}-${TYPE}"; \
		go test -timeout $(TIMEOUT) -v $(ARGS) ./${TYPE}/${TARGET}/... ; \
	fi

################################################################################
# Target: run                                                                  #
################################################################################

run run-%: generate-version
	@if [ -z $(TARGET) ]; then \
		echo "no  TARGET. example usage: make run TARGET=account"; \
	else \
		go run  ./${TYPE}/${TARGET} ${ARGS}; \
	fi

################################################################################
# Target: release                                                              #
################################################################################

release: verify
	@if [ -z $(TAG) ]; then \
		echo "no  TAG. Usage: make release TAG=v0.1.1"; \
	else \
		for m in `find * -name 'go.mod' -mindepth 1 -exec dirname {} \;`; do \
			hub release create -m "$$m/${TAG} release" $$m/${TAG}; \
		done \
	fi

release/draft: check_dirty
	@echo Publishing Draft: $(VERSION)
	@git tag -a $(VERSION) -m "[skip ci] Release: $(VERSION)" || true
	@git push origin $(VERSION)
	@echo "\n\nPlease inspect the release and run `make release/publish` if it looks good"
	@open "$(GITHUB_RELEASES_UI_URL)/$(VERSION)"

release/publish:
	@echo Publishing Release: $(VERSION)

deploy/e2e:
	@curl -H "Content-Type: application/json" \
		-H "Accept: application/vnd.github.ant-man-preview+json"  \
		-H "Authorization: token $(GITHUB_TOKEN)" \
    -XPOST $(GITHUB_DEPLOY_API_URL) \
    -d '{"ref": "develop", "environment": "e2e", "payload": { "what": "deployment for e2e testing"}}'

deploy/prod:
	@curl -H "Content-Type: application/json" \
		-H "Accept: application/vnd.github.ant-man-preview+json"  \
		-H "Authorization: token $(GITHUB_TOKEN)" \
    -XPOST $(GITHUB_DEPLOY_API_URL) \
    -d '{"ref": "develop", "environment": "production", "payload": { "what": "production deployment to GKE"}}'

################################################################################
# Target: docker                                                               #
################################################################################
docker docker-%:
	@if [ -z $(TARGET) ]; then \
		echo "Building images for all services..."; \
		for type in $(TYPES); do \
			echo "Building Type: $${type}..."; \
			for _target in $${type}/*/; do \
				temp=$${_target%%/}; target=$${temp#*/}; \
				echo "Building Image $${target}-$${type}..."; \
				nerdctl build \
				--build-arg VERSION=$(VERSION) \
				--build-arg TYPE=$${type} \
				--build-arg TARGET=$${target} \
				--build-arg DOCKER_REGISTRY=${DOCKER_REGISTRY} \
				--build-arg DOCKER_CONTEXT_PATH=${DOCKER_CONTEXT_PATH} \
				--build-arg VCS_REF=$(shell git rev-parse --short HEAD) \
				--build-arg BUILD_DATE=$(shell date +%FT%T%Z) \
				-t ${DOCKER_REGISTRY}/${DOCKER_CONTEXT_PATH}/$${target}-$${type}:$(VERSION) .; \
			done \
		done \
	else \
		echo "Building image for ${TARGET}-${TYPE}..."; \
		nerdctl build --progress=plain \
		--build-arg VERSION=$(VERSION) \
		--build-arg TYPE=${TYPE} \
		--build-arg TARGET=${TARGET} \
		--build-arg DOCKER_REGISTRY=${DOCKER_REGISTRY} \
		--build-arg DOCKER_CONTEXT_PATH=${DOCKER_CONTEXT_PATH} \
		--build-arg VCS_REF=$(shell git rev-parse --short HEAD) \
		--build-arg BUILD_DATE=$(shell date +%FT%T%Z) \
		-t ${DOCKER_REGISTRY}/${DOCKER_CONTEXT_PATH}/${TARGET}-${TYPE}:$(VERSION) .; \
	fi

docker_clean:
	@echo "Cleaning dangling images..."
	@nerdctl images -f "dangling=true" -q  | xargs docker rmi
	@echo "Removing microservice images..."
	@nerdctl images -f "label=org.label-schema.vendor=sumo" -q | xargs docker rmi
	@echo "Pruneing images..."
	@nerdctl image prune -f

docker_push:
	@echo "Publishing images with VCS_REF=$(shell git rev-parse --short HEAD)"
	@nerdctl images -f "label=org.opencontainers.image.ref.name=$(shell git rev-parse --short HEAD)" --format {{.Repository}}:{{.Tag}} | \
	while read -r image; do \
		echo Now pushing $$image; \
		nerdctl push $$image; \
	done;

docker_base:
	nerdctl build --build-arg BUILD_DATE=$(shell date +%FT%T%Z) --build-arg VERSION=$(VERSION) -f Dockerfile.base -t ${DOCKER_REGISTRY}/${DOCKER_CONTEXT_PATH}/base:$(VERSION) .
	nerdctl tag ${DOCKER_REGISTRY}/${DOCKER_CONTEXT_PATH}/base:$(VERSION) ${DOCKER_REGISTRY}/${DOCKER_CONTEXT_PATH}/base:latest
#	nerdctl push ${DOCKER_REGISTRY}/${DOCKER_CONTEXT_PATH}/base:$(VERSION)
#	nerdctl push ${DOCKER_REGISTRY}/${DOCKER_CONTEXT_PATH}/base:latest

################################################################################
# Target: deploy                                                               #
################################################################################

kustomize: OVERLAY 	:= local
kustomize: NS 		:= default
kustomize:
	# @kustomize build --load_restrictor none config/envs/${OVERLAY}/ | sed -e "s|\$$(NS)|${NS}|g" 		-e "s|\$$(IMAGE_VERSION)|${VERSION}|g" | kubectl apply -f -
	@kustomize build --load_restrictor none config/envs/${OVERLAY}/ | sed -e "s|\$$(NS)|${NS}|g" 		-e "s|\$$(IMAGE_VERSION)|${VERSION}|g" > build/kubernetes.yaml

build/kustomize: check_dirty
	@kustomize build --load_restrictor none config/envs/local  		| sed -e "s|\$$(NS)|default|g" 		-e "s|\$$(IMAGE_VERSION)|${VERSION}|g" > build/kubernetes.local.yaml
	@kustomize build --load_restrictor none config/envs/production/ | sed -e "s|\$$(NS)|default|g" 		-e "s|\$$(IMAGE_VERSION)|${VERSION}|g" > build/kubernetes.production.yaml
	@kustomize build --load_restrictor none config/envs/production/ | sed -e "s|\$$(NS)|mynamespace|g" 	-e "s|\$$(IMAGE_VERSION)|${VERSION}|g" > build/kubernetes.production.mynamespace.yaml
