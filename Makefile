FAUXPATH = $(GOPATH)/src/go.owls.io/fauxbox
APP = $(FAUXPATH)/fauxbox.app

.phony: all

all: clean build run

clean:
	@echo "Deleting fauxbox.app"
	@$(shell if [ -d $(APP) ]; then rm -r $(APP); fi)

build:
	@echo "Building fauxbox.app"
	@go build go.owls.io/fauxbox
	@mkdir -p $(APP)/Contents/MacOS
	@mv $(FAUXPATH)/fauxbox $(APP)/Contents/MacOS/
	@cp -r $(FAUXPATH)/resources $(APP)/Contents/MacOS/

run:
	@echo "Running fauxbox.app"
	@open $(APP)
