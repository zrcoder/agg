gen: 
	@go generate

clean: 
	@rm -rf internal/exported


run:
	@go run ./cmd
