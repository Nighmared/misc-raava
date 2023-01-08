package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type folder struct {
	Id                int
	Name              string
	Mapped_path       string
	Description       string
	Used_quota_size   int
	Used_quota_files  int
	Last_quota_update int
	Users             []string
}

type user struct {
	Id              int
	Status          int
	Username        string
	Email           string
	Description     string
	Expiration_date int
	Password        string
	Public_keys     []string
	Home_dir        string
	Virtual_folders []folder
	//..... many more but do i even need that?

}

const (
	fifty_gigabyte = 50000000000
)

func main() {
	logged_in_user := os.Getenv("SFTPGO_LOGIND_USER")
	login_protocol := os.Getenv("SFTPGO_LOGIND_PROTOCOL")
	login_ip := os.Getenv("SFTPGO_LOGIND_IP")
	if login_protocol != "OIDC" {
		//ignore for other auth methods
		return
	}
	var parsed_user user
	json.Unmarshal([]byte(logged_in_user), &parsed_user)
	fmt.Print("{")
	fmt.Print(`"status":1,`)
	fmt.Printf(`"username":"%s",`, parsed_user.Username)
	fmt.Printf(`"email":"%s",`, parsed_user.Email)
	fmt.Printf(`"home_dir":"/srv/sftpgo/data/%s",`, parsed_user.Username)
	fmt.Printf(`"quota_size":%d,`, fifty_gigabyte)
	fmt.Print(`"description":"oidc user",`)
	fmt.Printf(`"additional_info":"last logged in from %s",`, login_ip)
	fmt.Print(`"groups": [`)
	fmt.Print(`{`)
	fmt.Print(`"name":"keycloak_users",`)
	fmt.Print(`"type":1`)
	fmt.Print(`}`)
	fmt.Print("],")
	fmt.Print(`"permissions": {`)
	fmt.Print(`"/": [`)
	fmt.Print(`"*"`)
	fmt.Print(`]`)
	fmt.Print((`}`))
	fmt.Print(`}`)

}
