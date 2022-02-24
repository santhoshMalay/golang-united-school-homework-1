package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {
	map := emoji.SPrint("Hello :world_map:!")
	return map
}
