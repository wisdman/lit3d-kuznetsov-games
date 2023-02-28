.PHONY: lib firmware clean run

lib:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" go build -mod=vendor -buildmode=c-shared -ldflags="-s -w -buildid= -H windowsgui" -o build/winspool.drv ./lib/winspool

firmware:
	cd lit3d-kuznetsov-games; make

clean:
	rm -rf build/*

run:
	go run -mod=vendor ./cmd/hid-test
