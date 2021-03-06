// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	// mysql import.
	_ "github.com/go-sql-driver/mysql"
)

var userdb, passdb, addr, d string
var db *sql.DB

var isdbok bool

func init() {
	userdb = "pippo"
	passdb = "pippo"
	addr = "127.0.0.1:3306"
	d = "app"

	isdbok = true
	// Apre la conessione al DB.
	dbconn, err := sql.Open("mysql", userdb+":"+passdb+"@tcp("+addr+")/"+d)

	// Se c'è un errore esce.
	if err != nil {
		log.Printf("DB access not possible: %v", err)
	}
	err = dbconn.PingContext(context.TODO())
	if err != nil {
		log.Printf("DB not rechable: %v", err)
	}
	if err != nil {
		isdbok = false
	}

	db = dbconn
}

// QueryCredentials is a query that executed on DB returns true if both user
// and password are found in a record.
const QueryCredentials = "SELECT IF(COUNT(*),'true','false') FROM app.credentials WHERE username = ? AND password = ?"

// CheckCredentialsDBCtx looks for credentials in the DB.
func CheckCredentialsDBCtx(ctx context.Context, c *Credentials) (bool, error) {

	if isdbok == false {
		return false, fmt.Errorf("DB was not rechable during INIT")
	}

	// Crea il contesto base.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Chiude la connesione al DB alla fine.
	defer db.Close()

	err := db.PingContext(ctx)
	if err != nil {
		return false, fmt.Errorf("DB not rechable: %v", err)
	}

	var isAuthenticated bool
	err = db.QueryRowContext(ctx,
		QueryCredentials, c.User, c.Hashpass).Scan(&isAuthenticated)
	if err != nil {
		return false, fmt.Errorf("db access not possible: %v", err)
	}

	select {
	case <-ctx.Done():
		return false, fmt.Errorf("Timeout: %v", ctx.Err())
	default:
		return isAuthenticated, err
	}
}
