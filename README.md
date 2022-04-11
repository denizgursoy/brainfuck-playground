# brainfuck-playground

## Build

Binary can be created with the help of steps in the makefile

Execute
`make build` or `go build -o brainfuck`

## Execution

Binary requires 3 paths for files: command file, input file, output file respectively. If they are not present, it will
panic

After binary is created, you can execute:

`make build-run` or `./brainfuck command.txt input.txt output.txt`

It will read the commands in the command.txt and writes "Hello World!" to output.txt