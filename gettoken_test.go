// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package token

import (
	"context"
	"testing"

	"github.com/axamon/hashstring"
)

func TestCheckLocalCredentials(t *testing.T) {
	type args struct {
		ctx context.Context
		c   *Credentials
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "primo", args: args{ctx: context.TODO(),
			c: &Credentials{User: "pippo", Hashpass: hashstring.Md5Sum("pippo")},
		want: true, wantErr: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckLocalCredentials(tt.args.ctx, tt.args.c)
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
