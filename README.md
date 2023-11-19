# Echo

Echo prints its command-line arguments.

These are exercises 1.1, 1.2 and 1.3 from the book The Go Programming Language by Brian W. Kernighan and Alan Donovan.

## Requirements

Make sure you have at least Go 1.21.4 installed on your system before running the program.

## How to Run

To run the Echo program, use the following command:

```bash
go run echo.go [Arg1] [Arg2] [Arg3] ...
```

## Example Usage

```bash
go run echo.go hello world
```

## Tasks

[ x ] Modify the echo program to also print os.Args[0], the name of the command that invoked it.

[] Modify the echo program to print the index and value of each of its arguments, one per line.

[] Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses strings.Join.
