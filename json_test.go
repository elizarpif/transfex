package transfex

import (
	"testing"
)

func TestReadJson(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "check read",
			args: args{name: "testadata/example.json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadJson(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
