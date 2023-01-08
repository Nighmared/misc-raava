package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type folder struct {
	id                int
	name              string
	mapped_path       string
	description       string
	used_quota_size   int
	used_quota_files  int
	last_quota_update int
	users             []string
}

type user struct {
	id              int
	status          int
	username        string
	email           string
	description     string
	expiration_date int
	password        string
	public_keys     []string
	home_dir        string
	virtual_folders []folder
	//..... many more but do i even need that?

}

func main() {
	logged_in_user := os.Getenv("SFTPGO_LOGIND_USER")
	login_protocol := os.Getenv("SFTPGO_LOGIND_PROTOCOL")
	if login_protocol != "OIDC" {
		return
	}
	var parsed_user user
	json.Unmarshal([]byte(logged_in_user), &parsed_user)
	fmt.Println(parsed_user)

}
