# K8-Label-Encoder

A Go module for encoding special characters to be compliant with the character limitations of Kubernetes Labels and Selectors relying on a similar en

## Installation

```bash
go get github.com/CactiNotch/K8-Label-Encoder
```

## Usage

```go
import "github.com/CactiNotch/K8-Label-Encoder"

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

## List of supported mappings

| Character | Encoded Form | Description       |
| --------- | ------------ | ----------------- |
| Space     | `_20`        | Space character   |
| `!`       | `_21`        | Exclamation mark  |
| `"`       | `_22`        | Double quote      |
| `#`       | `_23`        | Hash/Pound        |
| `$`       | `_24`        | Dollar sign       |
| `%`       | `_25`        | Percent           |
| `&`       | `_26`        | Ampersand         |
| `'`       | `_27`        | Single quote      |
| `(`       | `_28`        | Left parenthesis  |
| `)`       | `_29`        | Right parenthesis |
| `*`       | `_2A`        | Asterisk          |
| `+`       | `_2B`        | Plus sign         |
| `,`       | `_2C`        | Comma             |
| `/`       | `_2F`        | Forward slash     |
| `:`       | `_3A`        | Colon             |
| `;`       | `_3B`        | Semicolon         |
| `<`       | `_3C`        | Less than         |
| `=`       | `_3D`        | Equals            |
| `>`       | `_3E`        | Greater than      |
| `?`       | `_3F`        | Question mark     |
| `@`       | `_40`        | At sign           |
| `[`       | `_5B`        | Left bracket      |
| `\`       | `_5C`        | Backslash         |
| `]`       | `_5D`        | Right bracket     |
| `^`       | `_5E`        | Caret             |
| `` ` ``   | `_60`        | Backtick          |
| `{`       | `_7B`        | Left brace        |
| `\|`      | `_7C`        | Vertical bar      |
| `}`       | `_7D`        | Right brace       |
| `~`       | `_7E`        | Tilde             |

**Note:** Alphanumeric characters, hyphens (`-`), underscores (`_`), and periods (`.`) are not encoded as they are valid in Kubernetes labels.
