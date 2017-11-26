BUILD_PATH		=		$(CURDIR)
GO				=		$(shell which go)
GO_INSTALL		=		$(GO) install
GO_CLEAN		=		$(GO) clean
GO_GET			=		$(GO) get

BIN_NAME		=		main

export GOPATH	=		$(CURDIR)

makedir:
	@echo "starting building tree..."
	@if [ ! -d $(BUILD_PATH)/bin ]; then mkdir -p $(BUILD_PATH)/bin; fi
	@if [ ! -d $(BUILD_PATH)/pkg ]; then mkdir -p $(BUILD_PATH)/pkg; fi
	@echo "tree built"

get:
	@$(GO_GET) github.com/Sirupsen/logrus

build:
	@echo "start building binary..."
	$(GO_INSTALL) $(BIN_NAME)
	@echo "all done"

clean:
	@echo "cleaning..."
	@rm -rf $(BUILD_PATH)/bin/$(BIN_NAME)
	@rm -rf $(BUILD_PATH)/pkg
	@rm -rf $(BUILD_PATH)/src/github.com
	@echo "cleaned"

all: makedir get build