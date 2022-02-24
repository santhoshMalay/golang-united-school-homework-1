package main

import (
	"fmt"
	"github.com/kyokomi/emoji"
)

func main() {
	message := "\"Hello" + emoji.Sprint(" :world_map:!\"")
	fmt.Println(message)
}
