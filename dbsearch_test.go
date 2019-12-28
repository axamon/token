// Copyright 2019 Alberto Bregliano. All rights reserved.

// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.

package token_test

import (
	"context"
	"testing"
	"time"

	"github.com/axamon/token"

	// mysql import
	_ "github.com/go-sql-driver/mysql"
)

func TestCheckCredentialsDBCtx(t *testing.T) {
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
		{name: "Should go in error",
			args:    args{ctx: context.TODO(), c: &token.Credentials{User: "pippo", Hashpass: "fsd"}},
			want:    false,
			wantErr: true},
		{name: "Should timeout",
			args:    args{ctx: ctxshort, c: &token.Credentials{User: "pippo", Hashpass: "fsd"}},
			want:    false,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := token.CheckCredentialsDBCtx(tt.args.ctx, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckCredentialsDBCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckCredentialsDBCtx() = %v, want %v", got, tt.want)
			}
		})
	}
}
