// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token_test

import (
	"context"
	rand "math/rand"
	"testing"
	"time"

	"github.com/axamon/hashstring"
	"github.com/axamon/token"
)

func TestCheckLocalCredentials(t *testing.T) {
	ctxshort, cancel := context.WithTimeout(context.Background(), time.Microsecond)
	defer cancel()
	type args struct {
		ctx context.Context
		c   *token.Credentials
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "first", args: args{ctx: context.TODO(),
			c: &token.Credentials{
				User:     "pippo",
				Hashpass: hashstring.Md5Sum("pippo")}},
			want: true, wantErr: false},
		{name: "second", args: args{ctx: context.TODO(),
			c: &token.Credentials{
				User:     "pluto",
				Hashpass: hashstring.Md5Sum("pippo")}},
			want: false, wantErr: true},
		{name: "third", args: args{ctx: context.TODO(),
			c: &token.Credentials{
				User:     "pippo",
				Hashpass: hashstring.Md5Sum("pipp")}},
			want: false, wantErr: false},
		{name: "forth", args: args{ctx: ctxshort,
			c: &token.Credentials{
				User:     "pipp",
				Hashpass: hashstring.Md5Sum("pippo")}},
			want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Sleep(time.Millisecond) // modified after bloomfilter introduction.
			got, err := token.CheckLocalCredentials(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLocalCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckLocalCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckLocalCredentials2(t *testing.T) {
	token.CredentialsJSONFile = "file.json"
	type args struct {
		ctx context.Context
		c   *token.Credentials
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "first", args: args{ctx: context.TODO(),
			c: &token.Credentials{
				User:     "pippo",
				Hashpass: hashstring.Md5Sum("pippo")}},
			want: false, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := token.CheckLocalCredentials(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckLocalCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckLocalCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateToken(t *testing.T) {
	rand.Seed(int64(99))
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "first", args: args{ctx: context.TODO()}, want: "75ed1842-49e9-bc19-675e-4d1f766213da", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := token.GenerateCtx(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
