# Note:
# This Makefile is only intended to work within an OS X environment. (for now)

FAUXPATH = $(GOPATH)/src/go.owls.io/fauxbox
APP = $(FAUXPATH)/fauxbox.app

.phony: all-osx

all-osx: clean-osx build-osx run-osx-dev

clean-osx:
	@echo "Deleting fauxbox.app"
	@$(shell if [ -d $(APP) ]; then rm -r $(APP); fi)

build-osx:
	@echo "Building fauxbox.app"
	@go build go.owls.io/fauxbox
	@mkdir -p $(APP)/Contents/MacOS
	@mv $(FAUXPATH)/fauxbox $(APP)/Contents/MacOS/
	@cp -r $(FAUXPATH)/resources $(APP)/Contents/MacOS/

run-osx-dev:
	@echo "Running fauxbox.app (dev)"
	$(APP)/Contents/MacOS/fauxbox

run-osx:
	@echo "Running fauxbox.app"
	open $(APP)

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
