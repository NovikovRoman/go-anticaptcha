# Anticaptcha

[Русская версия документации](https://anticaptcha.atlassian.net/wiki/spaces/API/pages/196633)

[Documentation in English](https://anticaptcha.atlassian.net/wiki/spaces/API/pages/196635/Documentation+in+English)

```go
package main

import (
	"encoding/json"
    "fmt"
    "github.com/NovikovRoman/go-anticaptcha"
	"log"
)

func main() {
	client := &Client{
		ClientKey: "[API_KEY]",
	}
	imageTextTask := &ImageToTextTask{
		Body:    "base64 image",
		Numeric: 1,
	}
	res, err := client.SendImageToTextTask(imageTextTask)
	if err != nil {
		log.Fatalf("Create task %s", err)
	}
	resJson, err := json.Marshal(&res)
	if err != nil {
		log.Fatalf("Error marshal %s", err)
	}
	fmt.Printf("%s\n", string(resJson))
}
```