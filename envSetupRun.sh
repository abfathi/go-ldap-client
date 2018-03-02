#!/bin/bash

export GOPATH=$PWD
export GOBIN=$GOPATH/bin
go get "github.com/jtblin/go-ldap-client"
go run openldap-test.go $1 $2 $3 $4 $5 $6 $7 $8
