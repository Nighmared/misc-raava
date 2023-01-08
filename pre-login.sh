if [ "$SFTPGO_LOGIND_PROTOCOL" == "OIDC" ]
then
    echo "{
        \"status\": 1,
        \"username\": \"${SFTPGO_LOGIND_USER}\",
        \"email\": \"a@b.de\",
        \"home_dir\": \"/srv/sftpgo/data/${SFTPGO_LOGIND_USER}\",
        \"quota_size\": 50000000000,  
        \"description\": \"created via api\",
        \"additional_info\": \"last logged in from ${SFTPGO_LOGIND_IP}\",
        \"groups\": [
            {
                \"name\": \"keycloak users\",
                \"type\":1
            }
        ],
        \"permissions\": {
            \"/\": [
                \"*\"
            ]
        }
    }"
fi