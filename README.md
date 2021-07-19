# albacore
Albacore is a simple Golang client for [Stanford CoreNLP server](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html).

# Install
```shell
go get github.com/shota3506/albacore
```

# Usage
```go
package main

import (
	"context"
	"fmt"
	"log"

	corenlp "github.com/shota3506/albacore/stanfordcorenlp"
)

func main() {
	ctx := context.Background()

	// create client for Stanford CoreNLP
	client := corenlp.NewClient(ctx, "http://localhost:9000")

	// sample text
	text := "The quick brown fox jumped over the lazy dog."
	doc, err := client.Tokenize(ctx, text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(doc) // The quick brown fox jumped over the lazy dog .
}
```
