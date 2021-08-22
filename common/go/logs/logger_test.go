package logs

import (
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		opt *Option
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "all success",
			args: args{
				opt: &Option{
					FilePath:     "./tmp/logs/",
					FileName:     "logger.log",
					Formatter:    "JSON",
					Stdout:       false,
					ReportCaller: false,
				},
			},
			wantErr: false,
		},
		{
			name: "all success 2",
			args: args{
				opt: &Option{
					FilePath:     "./tmp/logs/",
					FileName:     "logger.log",
					Formatter:    "TEXT",
					Stdout:       true,
					ReportCaller: false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("New() got = %v", nil)
			}
		})
	}
}
