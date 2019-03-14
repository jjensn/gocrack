// +build !auth_database,!auth_ldap

package server

import (
	_ "github.com/jjensn/gocrack/server/authentication/database"
	_ "github.com/jjensn/gocrack/server/authentication/ldap"
)
