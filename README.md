# Iran IDValidator




Iran National Id, Bank Card Number, Mobile Number Validator for golang


### Installation

```
go get -u github.com/mohammadv184/idvalidator
```

### Usage

```go
package main

import (
	"fmt"
	
	"github.com/mohammadv184/idvalidator/validate/nationalid"
	
	"github.com/mohammadv184/idvalidator/validate/card"
)

func main() {
	ok, err := card.IsValid("1234567890123456")
	if !ok {
		fmt.Println(err)
	}

	ok, err = nationalid.IsValid("1234567890123456")
	if !ok {
		fmt.Println(err)
	}

}
```
## License

The MIT License (MIT). Please see [License File](LICENSE) for more information.