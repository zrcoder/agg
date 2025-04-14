gen: 
	@igop export -outdir ./internal/exported ./pkg/export/hanoi

clean: 
	@rm -rf internal/exported


run:
	@go run -ldflags="-checklinkname=0" ./cmd
