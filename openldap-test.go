package main

import (
	"log"
	"os"
	"strconv"
        "github.com/jtblin/go-ldap-client"
)

/*  Basic test performed to validate OPEN-LDAP connection instance deployed on EC2

    $ ./envSetupRun.sh "dc=fathi,dc=com" "54.155.139.221" "ip-10-122-145-55.fathi.com" 636 "cn=Manager,dc=fathi,dc=com" "admin" "user1" "XXXXX"

    Output:

          2018/03/02 12:22:13 Connected OK to backend
          2018/03/02 12:22:13 User: map[sn:user1 mail:user1@fathi.com uid:user1 givenName:Michael]
          2018/03/02 12:22:13 Groups: []

*/
func main() {

	base := os.Args[1]
	host := os.Args[2]
	serverName := os.Args[3]
	port, _ := strconv.Atoi(os.Args[4])
	bindDN := os.Args[5]
	bindPassword := os.Args[6]
	userName := os.Args[7]
	password := os.Args[8]

	client := &ldap.LDAPClient{
		Base:               base,
		Host:               host,
		ServerName:         serverName,
		InsecureSkipVerify: true,
		Port:               port,
		UseSSL:             true,
		BindDN:             bindDN,
		BindPassword:       bindPassword,
		UserFilter:         "(uid=%s)",
		GroupFilter:        "(memberUid=%s)",
		Attributes:         []string{"givenName", "sn", "mail", "uid"},
	}

	// It is the responsibility of the caller to close the connection
	defer client.Close()

	err := client.Connect()
	if err != nil {
		log.Fatalf("Error connecting to LDAP backend")
	}

	log.Printf("Connected OK to backend")

	ok, user, err := client.Authenticate(userName, password)
	if err != nil {
		log.Fatalf("Error authenticating user %s: %+v", userName, err)
	}

	if !ok {
		log.Fatalf("Authenticating failed for user %s", userName)
	}
	log.Printf("User: %+v", user)

	groups, err := client.GetGroupsOfUser(userName)
	if err != nil {
		log.Fatalf("Error getting groups for user %s: %+v", userName, err)
	}
	log.Printf("Groups: %+v", groups)
}
