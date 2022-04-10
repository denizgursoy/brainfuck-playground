.PHONY: build
build:
	go build -o brainfuck

.PHONY: build-run
build-run:
	rm -f brainfuck
	make build
	./brainfuck command.txt input.txt output.txt