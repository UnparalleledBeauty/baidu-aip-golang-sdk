package censor

import (
	"encoding/base64"
	"fmt"
	"os"
	"testing"
)

func TestContentCensorClient_ImageCensorCommonAccurateBasic(t *testing.T) {
	client := NewClient("xxxx", "xxxxx")
	s, err := os.ReadFile("/Users/xxxxx/go/src/beauty/golang-sdk/data/11.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	str := base64.StdEncoding.EncodeToString(s)
	res := client.ImageCensorCommonAccurateBasic(str, nil)
	fmt.Println(res)
}
