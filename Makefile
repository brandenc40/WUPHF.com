REACT_PATH = wuphf-frontend/

.PHONY: run mod mod-run test drop-db build-js build-go dev-react

# Build final react bundle then run the go server
run: build-js
	@echo Running..
	@go run main.go
	
# run react app for development
dev-react:
	@npm run-script start -prefix $(REACT_PATH)

# build go binary file
build-go: 
	@echo Building..
	@go build

# build final react app bundle
build-js: 
	@echo Building React App..
	@npm run-script build -prefix $(REACT_PATH)

# clean up the go mod
mod:
	@echo Cleaning up the Go module..
	@go mod tidy 
	@go mod vendor

# run tests and display coverage
test:
	@go test ./... -coverprofile _test.out
	@rm _test.out

test-explore:
	@go test ./... -coverprofile _test.out
	@go tool cover -html=_test.out
	@rm _test.out