################################################################################
# Note:
# This Makefile is not designed to work within a Windows environment (for now).

FAUXPATH = $(GOPATH)/src/go.owls.io/fauxbox

OSXAPP = $(FAUXPATH)/fauxbox.app
LINUXAPP = $(FAUXPATH)/fauxbox_linux
WINDOWSAPP = $(FAUXPATH)/fauxbox.exe

phony:
	@echo -ne "\033[0;33mAvailable commands:\033[0m\n\n\
make osx[-dev]\n\
make linux[-dev]\n\
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
linux-dev: clean-linux build-linux run-linux-dev

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

# build-osx-to-windows:
# 	CGO_ENABLED=1 \
# 	CGO_CFLAGS="-I/opt/mingw-w64-1.0-bin_i686-darwin_20120227/include" \
# 	# CGO_LDFLAGS="-L/opt/mingw-w64-1.0-bin_i686-darwin_20120227/lib -lallegro_acodec -lallegro_audio -lallegro_color -lallegro_dialog -lallegro_image -lallegro_main -lallegro_memfile -lallegro_physfs -lallegro_primitives -lallegro_ttf -lallegro_font -lallegro" \
# 	CC=/opt/mingw-w64-1.0-bin_i686-darwin_20120227/bin/x86_64-w64-mingw32-gcc \
# 	CXX=/opt/mingw-w64-1.0-bin_i686-darwin_20120227/bin/x86_64-w64-mingw32-g++ \
# 	GOOS=windows GOARCH=amd64 go build go.owls.io/fauxbox

# build-osx-to-linux:

# build-allegro-windows:
# 	rm -rf ./allegro
# 	git clone git://github.com/liballeg/allegro5.git allegro
# 	cd allegro && git checkout --track origin/5.0
# 	mkdir allegro/buildw64
# 	cd allegro/buildw64 && cmake \
# 		-DCMAKE_SYSTEM_NAME=Windows \
# 		-DCMAKE_FIND_ROOT_PATH=/usr/x86_64-w64-mingw32 \
# 		-DCMAKE_C_COMPILER=x86_64-w64-mingw32-gcc \
# 		-DCMAKE_CXX_COMPILER=x86_64-w64-mingw32-g++ \
# 		.. && \
# 	make

# build-game-windows:
# 	CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -ldflags="-extld=$CC"
