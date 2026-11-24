# HS DEVS AUTH — Go Client SDK

Official Go client library for HS DEVS AUTH authentication platform.

## Install

```bash
go get github.com/HS-DEVS07/hsauth-go-ex
```

## Usage

```go
client := hsauth.New("AppName", "OwnerID", "Secret", "1.0")
if client.Init() {
    if client.Login("username", "password") {
        fmt.Println("Welcome", client.UserData.Username)
    }
}
```
