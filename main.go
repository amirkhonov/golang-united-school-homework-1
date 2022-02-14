package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func HelloEmoji(hello string) string {
	emojiWorld := emoji.Sprint(":world_map:")

	helloword := fmt.Sprint(hello, " ", emojiWorld, "!")
	return helloword
}

func main() {
	fmt.Printf(HelloEmoji("Hello"))
}
