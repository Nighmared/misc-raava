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

func main() {
	logged_in_user := os.Getenv("SFTPGO_LOGIND_USER")
	login_protocol := os.Getenv("SFTPGO_LOGIND_PROTOCOL")
	if login_protocol != "OIDC" {
		fmt.Println(login_protocol)
		return
	}
	var parsed_user user
	json.Unmarshal([]byte(logged_in_user), &parsed_user)
	fmt.Fprint(os.Stderr, parsed_user)

}
