# K8-Label-Encoder

A Go module for encoding special characters to be compliant with the character limitations of Kubernetes Labels and Selectors relying on a similar en

## Installation

```bash
go get github.com/YourUsername/K8-Label-Encoder
```

## Usage

```go
import "github.com/YourUsername/K8-Label-Encoder"

func main() {
    label := "hello world!"
    encoded := encoder.EncodeSpecialChars(label)
    decoded := encoder.DecodeSpecialChars(encoded)
}
```

## Features

- Encodes special characters to Kubernetes-safe format
- Decodes encoded labels back to original form
- Validates labels against Kubernetes restrictions
- Maintains bidirectional character mapping
