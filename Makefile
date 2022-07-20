default: release

release:
	@rm -f bin/*
	go build -o bin/igo
clean:
	@echo "--> Cleaning..."
	@rm -f bin/*
	
test:
	go test ./... -covermode=atomic
