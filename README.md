# GoBrainfuck

![Go Version](https://img.shields.io/badge/Version-1.22-242B36?style=for-the-badge&logo=go&logoColor=white&labelColor=1A222E)

A [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck) language interpreter with a few slight additions, written in Go.

## Usage

### Running a compiled file

* Windows

```bash
./brainfuck-windows.exe [FILE PATH]
./brainfuck-windows-32.exe [FILE PATH]
./brainfuck-windows-arm64.exe [FILE PATH]
```

* macOS

```bash
./brainfuck-macos-intel [FILE PATH] # For Macs with Intel processors
./brainfuck-macos-arm [FILE PATH] # For Macs with Apple M processors 
```

* Linux

```bash
./brainfuck-linux [FILE PATH]
./brainfuck-linux-32 [FILE PATH]
./brainfuck-linux-arm [FILE PATH]
./brainfuck-linux-arm64 [FILE PATH]
```

> [!INFO]
> If you see `permission denied: ./interpreter` error when running the file, make the file executable
> by typing `chmod +x brainfuck`

### Running from source

To do this, you must have the [Go](https://go.dev/) 1.22+ compiler installed.

```bash
go run brainfuck [FILE PATH]
```

## Commands

| Command | Go Equivalent                       | Meaning                                            |
|---------|-------------------------------------|----------------------------------------------------|
| `>`     | `cursor++`                          | Move to the next cell                              |
| `<`     | `cursor--`                          | Move to the previous cell                          |
| `+`     | `memory[cursor]++`                  | Increase the value in the current cell by 1        |
| `-`     | `memory[cursor]--`                  | Decrease the value in the current cell by 1        |
| `.`     | `fmt.Print(string(memory[cursor]))` | Print the contents of the current cell             |
| `,`     | `fmt.Scanln(&byte)`                 | Store user input in a cell                         |
| `[`     | `for memory[cursor] != 0 {`         | Start of loop                                      |
| `]`     | `}`                                 | End of loop                                        |
| `*`     | `fmt.Print(memory[cursor])`         | Print the contents of the current cell in raw form |

> The `*` command is not present in the official language specification, nevertheless, it was introduced in this
> implementation for simplicity of debugging.
> Whenever possible, use code that fully conforms to the official specification, for example
> `n>>+++++++++[<++++++>-]----[<+>-]<.`.

## I/O

### Input

Input using `,` command works with `fmt.Scanln()`. The first byte of input is recorded.
Usually, ASCII characters are accepted, but you can specify a number directly if you start the line with `//`.
For example, `//69` will write the number 69 to the cell.

### Output

Output is implemented via `fmt.Print()`. Usually, the character under the number written in the cell is displayed,
however,
you can print the number itself using the `*` command.

## Building the project

> [!INFO]
> Please note that compiled builds are available in Actions.

1. Install the latest version of the Go compiler from the [official site](https://go.dev/dl/)
2. Initiate the compilation using `go build brainfuck`

Example for Linux on the amd64 platform with root permissions:

```bash
wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
go build brainfuck
```