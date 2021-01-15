package haberdasher

import (
	"context"
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/sony/sonyflake"
	"github.com/tullo/microservice/rpc/haberdasher"
)

func TestServer_MakeHat(t *testing.T) {
	type fields struct {
		db        *sqlx.DB
		sonyflake *sonyflake.Sonyflake
	}
	type args struct {
		ctx  context.Context
		size *haberdasher.Size
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *haberdasher.Hat
		wantErr bool
	}{
		{
			name:    "Hat too small",
			fields:  fields{nil, nil},
			args:    args{context.TODO(), &haberdasher.Size{Centimeters: 0}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &Server{
				db:        tt.fields.db,
				sonyflake: tt.fields.sonyflake,
			}
			got, err := svc.MakeHat(tt.args.ctx, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.MakeHat() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.MakeHat() = %v, want %v", got, tt.want)
			}
		})
	}
}
