package utils

import (
	"fmt"
	"testing"
)

func TestRemoveNewline1(t *testing.T) {
	fmt.Println(RemoveNewline1("您希望这个AI编曲和写歌词的功能，具体是怎么样的呢？ 🤔 \n\n\n\n\n"))
}
