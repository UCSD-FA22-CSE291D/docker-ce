CLI_DIR:=$(CURDIR)/components/cli
ENGINE_DIR:=$(CURDIR)/components/engine
PACKAGING_DIR:=$(CURDIR)/components/packaging
VERSION=$(shell cat VERSION)

help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test-integration-cli: $(CLI_DIR)/build/docker ## test integration of cli and engine
	$(MAKE) -C $(ENGINE_DIR) DOCKER_CLI_PATH=$< test-integration-cli

$(CLI_DIR)/build/docker:
	$(MAKE) -C $(CLI_DIR) -f docker.Makefile build

deb: ## build deb packages
	$(MAKE) VERSION=$(VERSION) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) deb

rpm: ## build rpm packages
	$(MAKE) VERSION=$(VERSION) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) rpm

static: ## build static packages
	$(MAKE) VERSION=$(VERSION) CLI_DIR=$(CLI_DIR) ENGINE_DIR=$(ENGINE_DIR) -C $(PACKAGING_DIR) static

clean: ## clean the build artifacts
	-$(MAKE) -C $(CLI_DIR) clean
	-$(MAKE) -C $(ENGINE_DIR) clean
	-$(MAKE) -C $(PACKAGING_DIR) clean

vendor: $(CLI_DIR)/build/docker
	docker run --rm -it -v $(CLI_DIR):/go/src/github.com/docker/cli -v $(ENGINE_DIR):/go/src/github.com/docker/docker docker-cli-dev sh -c 'cd /go/src/github.com/docker/docker && git init && git add . && git -c user.name=user -c user.email=email@example.com commit -m first && cd /go/src/github.com/docker/cli && vndr; rm -rf /go/src/github.com/docker/docker/.git'
