################################################################################
# Note:
# This Makefile is not designed to work within a Windows environment.

FAUXPATH = $(GOPATH)/src/go.owls.io/fauxbox

OSXAPP = $(FAUXPATH)/fauxbox.app
LINUXAPP = $(FAUXPATH)/fauxbox_linux

phony:
	@echo -ne "\033[0;33mAvailable commands:\033[0m\n\n\
make linux\n\
make osx[-dev]\n"

clean: clean-osx clean-linux

################################################################################
osx-dev: clean-osx build-osx run-osx-dev

osx: clean-osx build-osx run-osx

clean-osx:
	@echo "Deleting fauxbox.app"
	@$(shell if [ -d $(OSXAPP) ]; then rm -r $(OSXAPP); fi)

build-osx:
	@echo "Building fauxbox.app"
	@go build go.owls.io/fauxbox
	@mkdir -p $(OSXAPP)/Contents/MacOS
	@mv $(FAUXPATH)/fauxbox $(OSXAPP)/Contents/MacOS/
	@cp -r $(FAUXPATH)/resources $(OSXAPP)/Contents/MacOS/
	@mkdir -p $(OSXAPP)/Contents/Resources
	@mv $(OSXAPP)/Contents/MacOS/resources/icons $(OSXAPP)/Contents/Resources
	@mv $(OSXAPP)/Contents/MacOS/resources/Info.plist $(OSXAPP)/Contents

run-osx-dev:
	@echo "Running fauxbox.app (dev)"
	$(OSXAPP)/Contents/MacOS/fauxbox

run-osx:
	@echo "Running fauxbox.app"
	open $(OSXAPP)

################################################################################
linux: clean-linux build-linux run-linux

clean-linux:
	@echo "Deleting fauxbox_linux"
	@$(shell if [ -f $(LINUXAPP) ]; then rm $(LINUXAPP); fi)

build-linux:
	@echo "Building fauxbox_linux"
	@go build go.owls.io/fauxbox
	@mv $(FAUXPATH)/fauxbox $(LINUXAPP)
	@chmod +x $(LINUXAPP)

run-linux:
	@echo "Running fauxbox_linux"
	@$(LINUXAPP)
