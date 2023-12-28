package db

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccounts_AddressPageQuery(t *testing.T) {
	type args struct {
		limit  int
		page int
	}
	tests := []struct {
		name      string
		args      args
		wantErr   bool
	}{
		{
			name: "test",
			args: args{
				limit:  300,
				page: 62,
			},
			wantErr: false,
		},
	}
	
	_,err := Connect("root:rootPassword@tcp(192.168.150.40:23306)/sig_turbo_tianzhou")
	require.NoError(t, err, "connect db error")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			offset := (tt.args.page - 1) *  tt.args.limit
			gotAddrs, err := Accounts{}.AddressPageQuery(tt.args.limit, offset)
			if (err != nil) != tt.wantErr {
				t.Errorf("Accounts.AddressPageQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(gotAddrs.Address())
		})
	}
}
