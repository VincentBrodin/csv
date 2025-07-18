# CSV

A fast and easy-to-use Go library for streaming CSV encoding/decoding with struct tags, inspired by encoding/json.

## Installation

Use the package manager to install.

```bash
go get github.com/VincentBrodin/csv
```

## Usage
### Decoder
```go
package main 

import (
	"fmt"
	"io"
	"os"
	"github.com/VincentBrodin/csv"
)

type MyStruct struct {
	Name  string `csv:"name"`
	Email string `csv:"email"`
	Age   int    `csv:"age"`
}

func main() {
	file, err := os.Open("myfile.csv") // RO
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := csv.NewDecoder(file)

	for {
		s := MyStruct{}
		if err := decoder.Decode(&s); err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Println(s)
	}
}
```

### Encoder
```go
package main

import (
    "os"
	"github.com/VincentBrodin/csv"
)

type MyStruct struct {
	Name  string `csv:"name"`
	Email string `csv:"email"`
	Age   int    `csv:"age"`
}

func main() {
	file, err := os.Create("myfile.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := csv.NewEncoder(file)

	myStruct := MyStruct{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		Age:   21,
	}
    if err := encoder.Encode(myStruct); err != nil {
        panic(err)
    }

    if err := encoder.Flush(); err != nil {
        panic(err)
    }
}

```

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://github.com/VincentBrodin/csv/blob/main/LICENSE)
