GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get
CURRENT_DIR = $(shell pwd)

all: deps clean build

build:
	mkdir -p $(CURRENT_DIR)/tmp
	$(GOBUILD) -ldflags="-s -w" -v -o $(CURRENT_DIR)/tmp/ftp-auth-handler

clean:
	$(GOCLEAN)
	rm -vrf $(CURRENT_DIR)/tmp

deps:
	$(GOGET) -u github.com/go-sql-driver/mysql
	$(GOGET) -u golang.org/x/crypto/bcrypt
	$(GOGET) -u gopkg.in/yaml.v2
