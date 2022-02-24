package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	message := "\"Hello" + emoji.Sprint(" :world_map:!\"")
	fmt.Println(message)
}
