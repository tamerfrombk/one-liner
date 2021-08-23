# one-liner

A simple command line utility that reads input from `stdin` and folds it into one line on `stdout`. The goal of this program is to be simple and have a significant performance improvement over other UNIX tools that can be used to accomplish this task.

### Running

By default, `one-liner` assumes `stdin` input uses the same line ending as the operating system default. To force `one-liner` to interpret the input with a different line ending, use `-crlf` or `-lf` to have the program read CRLF and LF inputs respectively.

### Building

`go build ./cmd/one-liner`

### Installing

`go install ./cmd/one-liner`

### Testing

`go test ./...`

### Performance Testing

Navigate to the root of the source tree.

To test the performance of `one-liner`, first generate the input test files by running the `gen-test-data.py` script. This will generate 4 input test files inside of `/test-data`: `small.txt`, `medium.txt`, `large.txt`, and `huge.txt`. These files are used as input data to `one-liner`.

From there, run the `run-perf-tests.sh`. This script will build `one-liner` and execute the following against each input file in ascending size order:

1. `wc` to count the number of lines, words, and bytes in the input file
2. `tr -s <newline> ' '` to serve as a baseline for a correct implementation and performance measure
3. `one-line`

All tests assume the input files use the OS' default line ending and use the `time` program to check performance.

Here is an example snippet of `run-perf-tests.sh` for `large.txt`:

```
Input file: large.txt
----------- wc ------------
   400000  15650000 108050000 large.txt
---------------------------
----------- tr test ------------
      0 15650000 107800001

real    0m0.310s
user    0m0.270s
sys     0m0.040s
--------------------------------
----------- one-liner test ------------
      1 15650000 108050001

real    0m0.114s
user    0m0.105s
sys     0m0.010s
---------------------------------------
```

Here, we can see that `one-liner` preserved the word and byte count of `large.txt` but obviously the input was collapsed to one line. We can also see that `one-liner` is approximately 3 times as fast as `tr` in this case.

