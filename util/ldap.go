package util

import (
	"fmt"

	"github.com/go-ldap/ldap/v3"
)

func Authenticate(ldapUrl, username, password string) (err error) {
	upn := fmt.Sprintf("%s@joynext.com", username) // TODO: get domain from DN
	conn, err := ldap.DialURL(ldapUrl)
	if err != nil {
		return
	}
	defer conn.Close()

	err = conn.Bind(upn, password)
	if err != nil {
		return
	}
	return
}

func FetchAdAccounts(ldapUrl, username, password string) (sr *ldap.SearchResult, err error) {
	upn := fmt.Sprintf("%s@joynext.com", username) // TODO: get domain from DN
	conn, err := ldap.DialURL(ldapUrl)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.Bind(upn, password)
	if err != nil {
		return nil, err
	}
	searchRequest := ldap.NewSearchRequest(
		"OU=Internal,OU=User,OU=NB,OU=JNN,DC=Joynext,DC=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		// filter to fetch all JNN members, this is one of attributes that all members have (you can select another common one)
		fmt.Sprintf("(&(objectCategory=%s))", ldap.EscapeFilter("CN=Person,CN=Schema,CN=Configuration,DC=Joynext,DC=com")),
		// target attributes we want
		[]string{"employeeID", "cn", "title", "sAMAccountName"},
		nil,
	)
	if sr, err = conn.Search(searchRequest); err != nil {
		return
	}
	return
}
