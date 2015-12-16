################################################################################
# Note:
# This Makefile is only intended to work within an OS X environment. (for now)

FAUXPATH = $(GOPATH)/src/go.owls.io/fauxbox

OSXAPP = $(FAUXPATH)/fauxbox.app

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
	@cp -r $(FAUXPATH)/images $(OSXAPP)/Contents/MacOS/
	@cp -r $(FAUXPATH)/models $(OSXAPP)/Contents/MacOS/
	@cp -r $(FAUXPATH)/resources $(OSXAPP)/Contents/MacOS/
	@cp -r $(FAUXPATH)/source $(OSXAPP)/Contents/MacOS/

run-osx-dev:
	@echo "Running fauxbox.app (dev)"
	$(OSXAPP)/Contents/MacOS/fauxbox

run-osx:
	@echo "Running fauxbox.app"
	open $(OSXAPP)

build-osx-windows:
	@echo "Building fauxbox.exe"
	@GOOS=windows GOARCH=amd64 go build go.owls.io/fauxbox

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


################################################################################
win-dev: clean-win build-win run-win-dev

win: clean-win build-win run-win

clean-win:
	@echo "Deleting fauxbox.exe"

build-win:
	@echo "Building fauxbox.exe"

run-win-dev:
	@echo "Running fauxbox.exe (dev)"

run-win:
	@echo "Running fauxbox.exe"

################################################################################
lin-dev: clean-lin build-lin run-lin-dev

lin: clean-lin build-lin run-lin

clean-lin:
	@echo "Deleting fauxbox"

build-lin:
	@echo "Building fauxbox"

run-lin-dev:
	@echo "Running fauxbox (dev)"

run-lin:
	@echo "Running fauxbox"
