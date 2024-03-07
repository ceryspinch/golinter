# golinter

## To use this linter in your own project:

Import package as a dependency.

`go get github.com/ceryspinch/golinter`

Call golinter.RunLinter from the main function.

```
import "github.com/ceryspinch/golinter"

func main() {
    golinter.RunLinter()
}
```

When running your code, add ` -- filenametolint.go` after `go run main.go ` to specify the file you want to lint

or add ` -- ./...` to lint all files in the project.

Example command to run linter:
` go run main.go -- example.go`

## How to interpret results:

Results output to the command line are colour coded as follows

Red => explanation of what exactly the issue is

Blue => explanation of why the issue may be problematic

Green => suggestion for improvement
