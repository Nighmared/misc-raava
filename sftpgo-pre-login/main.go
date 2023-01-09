package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Customfields struct {
	Storage_access bool
}

type user struct {
	Id                 int
	Status             int
	Username           string
	Email              string
	Oidc_custom_fields Customfields
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
	if parsed_user.Id != 0 {
		fmt.Printf(`{"additional_info":"last logged in from %s"}`, login_ip)
		//no need for all of the info if user already exists
		return
	}
	if !parsed_user.Oidc_custom_fields.Storage_access {
		//user isnt supposed to have storage access...
		fmt.Print(`{"status":0}`)
		return
	}

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
