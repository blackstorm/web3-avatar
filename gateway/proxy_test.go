package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func TestGetAvatar(t *testing.T) {
	web3, err := NewProxy(os.Getenv("INFURA_TOKEN"))
	if err != nil {
		t.Error(err)
	}

	avatar, err := web3.GetAvatar("0x5Bf717B477D4360a9EfcCa0Db3F512F732d21398", "default")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(avatar)
}
