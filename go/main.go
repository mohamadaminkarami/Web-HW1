package main

import (
	"context"
	"fmt"
)

var ctx = context.Background()
var rdb = initRedis()

func main() {
	setValueRedis("key", "value")

	var value, _ = getValueRedis("key")
	fmt.Println(value)
}
