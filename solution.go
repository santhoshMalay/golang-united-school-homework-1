package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	map1 := emoji.Sprint("Hello :world_map:!")
	return map1
}
