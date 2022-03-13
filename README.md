# Fuzzing Examples

This repo is just me going through the [Go Fuzzing](https://go.dev/doc/fuzz) doc and taking notes.

## Notes

Fuzzing is a type of automated testing that continuously runs a program with multiple different inputs. Beneficial for finding security expoloits and vulnerabilites.

Fuzz targets are invoked in parallel across multiple workers and in nondeterministic order.


### In Go, a fuzz test requires the folowing:
- function named like `FuzzXxx`, which accpets only a `*testing.F`, and has no return value
- exists within `*_test.go` files
- fuzz target
  - method that accepts `*testing.T` as the first parameter, followed my all fuzzing arguments.
  - no return value
- seed corpus entries
  - These entries will represent the parameters that are passed to the fuzz target. They will need to match the order of the parameters after `*testing.T`.

### Suggestions when using fuzzing:
- Fuzz targets should be fast and deterministic so the fuzzing engine can work efficiently, and new failures and code coverage can be easily reproduced.
- State shouldn't persist past the end of each call, and behavior shouldn't depend on any global state.

### Running fuzz tests
- As unit tests: `go test`
- With fuzzing: `go test -fuzz=FuzzTestName`
- By default, all other tests will be ran before fuzzing begins.
  - This is to ensure fuzzing won't report any issues that were found in unit tests.
- By default, fuzzing will continue to run until an error is found, or the process is cancelled manually.
- While fuzzing is in process, new inputs will continue to be generated and provided to the fuzz target.

### Failures
Failures can occur from the following:
- panics
- fuzz target called `t.Fail`
- `os.Exit`
- timeouts (1 second).
  - deadlocks
  - infinite loops
  - some code behavior
  - This is why the fuzz target should be fast

When failures occurr, the fuzzing engine will attempt to minimize the input to the smallest possible and most human readable value that will still produce an error. This can be configured through custom settings

After the failure occurs, the failed input will be ran by default with `go test` to further prevent regressions once the bug has been fixed.

### Corpus files

`file2fuzz` tool located [here](https://golang.org/x/tools/cmd/file2fuzz). This can be used to convert binary files to corpus files.
