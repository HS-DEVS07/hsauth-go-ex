package main

import (
	"fmt"
	"github.com/HS-DEVS07/hsauth-go-ex"
)

func main() {
	client := hsauth.New(
		"YOUR_APP_NAME",
		"YOUR_OWNER_ID",
		"YOUR_SECRET",
		"1.0",
	)

	fmt.Println("Initializing HS DEVS AUTH Client...")
	if !client.Init() {
		fmt.Printf("Initialization failed: %s\n", client.Response.Message)
		return
	}

	fmt.Println("Session active:", client.SessionID)

	key := "XXXX-XXXX-XXXX-XXXX"
	if client.License(key) {
		fmt.Printf("License verified! Expires: %s\n", client.UserData.ExpiresAt)
	} else {
		fmt.Printf("Verification error: %s\n", client.Response.Message)
	}
}
