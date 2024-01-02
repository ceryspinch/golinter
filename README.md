# golinter

## To use this linter in your own project:
Import package as a dependency.

```go get github.com/ceryspinch/golinter```

Call golinter.RunLinter from the main function.
```
import "github.com/ceryspinch/golinter"

func main() {
    golinter.RunLinter()
}
```


When running your code, add ``` -- filenametolint.go``` after ```go run main.go ``` to specify the file you want to lint


or add ``` -- ./...```  to lint all files in the project.
