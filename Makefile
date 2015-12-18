################################################################################
# Note:
# This Makefile is not designed to work within a Windows environment (for now).

FAUXPATH = $(GOPATH)/src/go.owls.io/fauxbox

OSXAPP = $(FAUXPATH)/fauxbox.app
LINUXAPP = $(FAUXPATH)/fauxbox_linux
WINDOWSAPP = $(FAUXPATH)/fauxbox.exe

phony:
	@echo -ne "\033[0;33mAvailable commands:\033[0m\n\n\
make linux\n\
make osx[-dev]\n\
make windows[-dev] (not configured)\n"

clean: clean-osx clean-linux clean-windows

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

build-osx-windows:
	@echo "Building fauxbox.exe"
	@GOOS=windows GOARCH=amd64 go build go.owls.io/fauxbox

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

################################################################################
windows-dev: clean-windows build-windows run-windows-dev

windows: clean-windows build-windows run-windows

clean-windows:
	@echo "Deleting fauxbox.exe"

build-windows:
	@echo "Building fauxbox.exe"

run-windows-dev:
	@echo "Running fauxbox.exe (dev)"

run-windows:
	@echo "Running fauxbox.exe"
