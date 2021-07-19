# albacore
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Albacore is a simple Golang client for [Stanford CoreNLP server](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html).

# Install
```shell
go get github.com/shota3506/albacore
```

# Usage
## Run Stanford CoreNLP server
Please run the Stanford CoreNLP server following [the official documentation](https://stanfordnlp.github.io/CoreNLP/corenlp-server.html).
Or you can run the server under [docker](https://stanfordnlp.github.io/CoreNLP/other-languages.html#docker).

Make sure you use version 4.0.0 or above.

## Tokenize
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

	fmt.Println(doc.Sentences[0].Tokens[1].Word) // quick
	fmt.Println(doc.Sentences[0].Tokens[1].Pos)  // JJ

	fmt.Println(doc.Sentences[0].Tokens[3].Word) // fox
	fmt.Println(doc.Sentences[0].Tokens[3].Pos)  // NN
}
```
