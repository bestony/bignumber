# bignumber

![Go](https://github.com/bestony/bignumber/workflows/Go/badge.svg)

Format numbers for human consumption


## Install

```bash
go get github.com/bestony/bignumber
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/bestony/bignumber"
)

func main() {
	num, _ := bignumber.Short(123456)
	fmt.Println(fmt.Sprintf("fmt %d is %s", 123456, num))
	num2, _ := bignumber.Comma(1234567890)
	fmt.Println(fmt.Sprintf("fmt %d is %s", 1234567890, num2))
}

```

## LICENSE

[MIT LICENSE](LICENSE)