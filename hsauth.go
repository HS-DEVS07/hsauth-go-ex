// Package hsauth provides the official Go Client SDK for HS DEVS AUTH.
package hsauth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type UserData struct {
	Username  string `json:"username"`
	IP        string `json:"ip"`
	HWID      string `json:"hwid"`
	ExpiresAt string `json:"expiry"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type Client struct {
	Name        string
	OwnerID     string
	Secret      string
	Version     string
	APIUrl      string
	SessionID   string
	Initialized bool
	UserData    UserData
	Response    Response
	httpClient  *http.Client
}

func New(name, ownerID, secret, version string) *Client {
	return &Client{
		Name:       name,
		OwnerID:    ownerID,
		Secret:     secret,
		Version:    version,
		APIUrl:     "https://auth.hsdevs.pro/api/1.3/",
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
}

func (c *Client) Init() bool {
	data := url.Values{
		"type":    {"init"},
		"name":    {c.Name},
		"ownerid": {c.OwnerID},
		"ver":     {c.Version},
	}

	resp, err := c.httpClient.PostForm(c.APIUrl, data)
	if err != nil {
		c.Response = Response{Success: false, Message: err.Error()}
		return false
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		c.Response = Response{Success: false, Message: "Invalid JSON response"}
		return false
	}

	success, _ := res["success"].(bool)
	message, _ := res["message"].(string)
	c.Response = Response{Success: success, Message: message}

	if success {
		c.SessionID, _ = res["sessionid"].(string)
		c.Initialized = true
		return true
	}
	return false
}

func (c *Client) Login(username, password string) bool {
	if !c.Initialized && !c.Init() {
		return false
	}

	data := url.Values{
		"type":      {"login"},
		"username":  {username},
		"pass":      {password},
		"sessionid": {c.SessionID},
		"name":      {c.Name},
		"ownerid":   {c.OwnerID},
	}

	resp, err := c.httpClient.PostForm(c.APIUrl, data)
	if err != nil {
		c.Response = Response{Success: false, Message: err.Error()}
		return false
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		c.Response = Response{Success: false, Message: "Invalid JSON response"}
		return false
	}

	success, _ := res["success"].(bool)
	message, _ := res["message"].(string)
	c.Response = Response{Success: success, Message: message}

	if success {
		if info, ok := res["info"].(map[string]interface{}); ok {
			c.UserData.Username, _ = info["username"].(string)
			c.UserData.IP, _ = info["ip"].(string)
			c.UserData.HWID, _ = info["hwid"].(string)
			c.UserData.ExpiresAt, _ = info["expiry"].(string)
		}
		return true
	}
	return false
}

func (c *Client) License(key string) bool {
	if !c.Initialized && !c.Init() {
		return false
	}

	data := url.Values{
		"type":      {"license"},
		"key":       {key},
		"sessionid": {c.SessionID},
		"name":      {c.Name},
		"ownerid":   {c.OwnerID},
	}

	resp, err := c.httpClient.PostForm(c.APIUrl, data)
	if err != nil {
		c.Response = Response{Success: false, Message: err.Error()}
		return false
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var res map[string]interface{}
	if err := json.Unmarshal(body, &res); err != nil {
		c.Response = Response{Success: false, Message: "Invalid JSON response"}
		return false
	}

	success, _ := res["success"].(bool)
	message, _ := res["message"].(string)
	c.Response = Response{Success: success, Message: message}

	if success {
		if info, ok := res["info"].(map[string]interface{}); ok {
			c.UserData.Username, _ = info["username"].(string)
			c.UserData.ExpiresAt, _ = info["expiry"].(string)
		}
		return true
	}
	return false
}
