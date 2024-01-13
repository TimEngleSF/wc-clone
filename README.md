# wc-clone

`wc-clone` is a command-line tool built with Go, inspired by the Unix command `wc`. It provides line, word, and character counts for standard .txt files.

By default, `wc-clone` prints counts for all three statistics: lines, words, and characters. It can also process multiple files and return their stats along with calculated totals.

## How to Use

### Install

To build the executable, navigate to the root directory of the program and execute:

```bash
$ go build .
```

### Usage

While in the root directory of the program, execute the following commands:

For a single file:

```bash
$ ./wc-clone path/to/some/file.txt
```

**Output**

```bash
       1       6      22 file.txt
```

For multiple files:

```bash
$ ./wc-clone path/to/file.txt ./path/to/another-file.txt
```

**Output**

```bash
       1       6      22 file.txt
       1       5      22 another-file.txt
       2      11      44 total
```

The program can also be run from the root file without building by executing

```bash
$ go run . path/to/file.txt ./path/to/another/file.txt
```

### Options

wc-clone supports the following options:

-   `-l`: Print line count of files
-   `-w`: Print word count of files
-   `-c`: Print character count of files

**Example**

```bash
$ ./wc-clone -l -c path/to/file.txt
```

**Output**

```bash
      1      22 file.txt
```
