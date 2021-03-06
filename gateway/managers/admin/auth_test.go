package admin

import (
	"context"
	"reflect"
	"testing"

	"github.com/spaceuptech/space-cloud/gateway/config"
)

func TestManager_createToken(t *testing.T) {
	type fields struct {
		user      *config.AdminUser
		isProd    bool
		clusterID string
	}
	type args struct {
		tokenClaims map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "valid info provided",
			fields: fields{
				user: &config.AdminUser{Secret: "some-secret"},
			},
			args:    args{tokenClaims: map[string]interface{}{"id": "admin", "role": "admin"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				user:      tt.fields.user,
				isProd:    tt.fields.isProd,
				clusterID: tt.fields.clusterID,
			}
			_, err := m.createToken(tt.args.tokenClaims)
			if (err != nil) != tt.wantErr {
				t.Errorf("createToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestManager_parseToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name:    "valid token provided",
			args:    args{token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6ImFkbWluIiwicm9sZSI6ImFkbWluIn0.N4aa9nBNQHsvnWPUfzmKjMG3YD474ChIyOM5FEUuVm4"},
			want:    map[string]interface{}{"id": "admin", "role": "admin"},
			wantErr: false,
		},
		{
			name:    "invalid token algorithm",
			args:    args{token: "eyJhbGciOiJFUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.tyh-VfuzIxCyGYDlkBA7DfyjrqmSHu6pQ2hoZuFqUSLPNY2N0mpHb3nk5K17HWP_3cYHBw7AhHale5wky6-sVA"},
			want:    nil,
			wantErr: true,
		},
	}
	m := New("", "clusterID", false, &config.AdminUser{Secret: "some-secret"})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.parseToken(context.Background(), tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
