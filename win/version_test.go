package win

import (
	"log"
	"testing"
)

func TestGetExeVersion(t *testing.T) {
	log.Println(GetVersion("D:\\software\\WeChat\\[3.6.0.18]\\WeChatWin.dll"))
}
