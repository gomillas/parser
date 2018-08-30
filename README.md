# parser
A library to split any string into smaller tokens based on regular expressions.

# Install

Simply execute the following command from a terminal:
```bash
go install github.com/gomillas/parser
```

And then import the library from your project:
```go
import (
	// additional libraries...
	"github.com/gchumillas/go/parser"
)
```

# Example

The following example splits a string into smaller 'words'. **Note that `regexp` ignores initial white spaces**:
```go
const regexp = `^\s*(\w+)`
const src = `lorem
ipsum dolor`

m := parser.New(src)
for token, offset := m.Find(regexp); len(token) > 0; token, offset = m.Find(regexp) {
	fmt.Printf("%s (offset: %d)\n", token, offset)
}

// Unordered output:
// lorem (offset: 0)
// ipsum (offset: 6)
// dolor (offset: 12)
```
