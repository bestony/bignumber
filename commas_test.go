package bignumber

import "testing"

/**
 * @Author Bestony <bestony@linux.com>
 */

func TestComma(t *testing.T) {
	type args struct {
		inputNumber int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "testZero", args: args{inputNumber: 0}, want: "0", wantErr: false},
		{name: "testOne", args: args{inputNumber: 1}, want: "1", wantErr: false},
		{name: "testTen", args: args{inputNumber: 10}, want: "10", wantErr: false},
		{name: "testHundred", args: args{inputNumber: 100}, want: "100", wantErr: false},
		{name: "testThousand", args: args{inputNumber: 1000}, want: "1,000", wantErr: false},
		{name: "testThousand", args: args{inputNumber: 1001}, want: "1,001", wantErr: false},
		{name: "testThousand", args: args{inputNumber: 2000}, want: "2,000", wantErr: false},
		{name: "testThousand", args: args{inputNumber: 5000}, want: "5,000", wantErr: false},
		{name: "testThousand", args: args{inputNumber: 9000}, want: "9,000", wantErr: false},
		{name: "testThousand", args: args{inputNumber: 1000000}, want: "1,000,000", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Comma(tt.args.inputNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Comma(%v) error = %v, wantErr %v", tt.args.inputNumber, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Comma(%v) got = %v, want %v", tt.args.inputNumber, got, tt.want)
			}
		})
	}
}
