gen: 
	@igop export -outdir ./internal/exported ./pkg/api/hanoi

clean: 
	@rm -rf internal/exported


runHanoi:
	@go run -ldflags="-checklinkname=0" ./cmd/hanoi

runBallSort:
	@go run ./cmd/ball-sort