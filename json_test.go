package transfex

import (
	"testing"
)


func TestGenerateCode(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "check read",
			args: args{filename: "testdata/example.json"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateCode(tt.args.filename, "somepack"); (err != nil) != tt.wantErr {
				t.Errorf("GenerateCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}