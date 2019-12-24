// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/axamon/token"
)

func TestCredentials_Autenticato(t *testing.T) {
	CredentialsJSONFile = "credentialsdb.json"
	ctxshort, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		c    Credentials
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "first", c: Credentials{User: "pippo", Hashpass: "f583ca884c1d93458fb61ed137ff44f6"}, args: args{ctx: context.TODO()}, want: true},
		{name: "second", c: Credentials{User: "pippo", Hashpass: "f583ca884c1d93458fb61ed137ff44f"}, args: args{ctx: context.TODO()}, want: false},
		{name: "third", c: Credentials{User: "pippo", Hashpass: "f583ca884c1d93458fb61ed137ff44f6"}, args: args{ctx: ctxshort}, want: false},
	}
	for _, tt := range tests {
		time.Sleep(time.Millisecond) // to trigger ctx timeout in CI.
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Autenticato(tt.args.ctx); got != tt.want {
				t.Errorf("Credentials.Autenticato() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCredentials_GetToken(t *testing.T) {
	rand.Seed(int64(99))
	ctxshort, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		c    Credentials
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "first", c: Credentials{User: "pippo", Hashpass: "f583ca884c1d93458fb61ed137ff44f6"}, args: args{ctx: context.TODO()}, want: "75ed1842-49e9-bc19-675e-4d1f766213da"},
		{name: "second", c: Credentials{User: "pippo", Hashpass: "f583ca884c1d93458fb61ed137ff44f"}, args: args{ctx: context.TODO()}, want: ""},
		{name: "third", c: Credentials{User: "pippo", Hashpass: "f583ca884c1d93458fb61ed137ff44f6"}, args: args{ctx: ctxshort}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.GetToken(tt.args.ctx); got != tt.want {
				t.Errorf("Credentials.GetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCredentials_Autenticato() {

	var c = Credentials{User: "pippo", Hashpass: "f583ca884c1d93458fb61ed137ff44f6"}

	r := c.Autenticato(context.TODO())

	fmt.Println(r)
	// Output:
	// true
}
