package bignumber

import "testing"

/**
 * @Author Bestony <bestony@linux.com>
 */

func TestShort(t *testing.T) {
	type args struct {
		inputNumber int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "shortZero", args: args{inputNumber: 0}, want: "0", wantErr: false},
		{name: "shortOne", args: args{inputNumber: 1}, want: "1", wantErr: false},
		// Test for number over thousand
		{name: "shortThousand", args: args{inputNumber: 1000}, want: "1k", wantErr: false},
		{name: "shortTwoThousand", args: args{inputNumber: 2000}, want: "2k", wantErr: false},
		{name: "shortTenThousand", args: args{inputNumber: 10000}, want: "10k", wantErr: false},
		{name: "shortHundredThousand", args: args{inputNumber: 100000}, want: "100k", wantErr: false},
		// Test For Number Over Million
		{name: "shortMillion", args: args{inputNumber: 1000000}, want: "1M", wantErr: false},
		{name: "shortMillion", args: args{inputNumber: 1100000}, want: "1.1M", wantErr: false},
		{name: "shortMillion", args: args{inputNumber: 1210000}, want: "1.2M", wantErr: false},
		{name: "shortMillion", args: args{inputNumber: 1350000}, want: "1.4M", wantErr: false},
		// Test for number with 1,hundred,and others
		{name: "shortThousandWithNineHundred", args: args{inputNumber: 1900}, want: "1.9k", wantErr: false},
		{name: "shortThousandWithNineHundred", args: args{inputNumber: 1901}, want: "1.9k", wantErr: false},
		{name: "shortThousandWithNineHundred", args: args{inputNumber: 1910}, want: "1.9k", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Short(tt.args.inputNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("Short(%v) error = %v, wantErr %v", tt.args.inputNumber, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Short(%v) got = %v, want %v", tt.args.inputNumber, got, tt.want)
			}
		})
	}
}

func Test_shortNumberBiggerThanMillion(t *testing.T) {
	type args struct {
		inputNumber int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{

		{name: "shortMillion", args: args{inputNumber: 1000000}, want: "1M", wantErr: false},
		{name: "shortMillion", args: args{inputNumber: 1100000}, want: "1.1M", wantErr: false},
		{name: "shortMillion", args: args{inputNumber: 1210000}, want: "1.2M", wantErr: false},
		{name: "shortMillion", args: args{inputNumber: 1350000}, want: "1.4M", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shortNumberBiggerThanMillion(tt.args.inputNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortNumberBiggerThanMillion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("shortNumberBiggerThanMillion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortNumberBiggerThanThousandButSmallerThanMillion(t *testing.T) {
	type args struct {
		inputNumber int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{

		{name: "shortThousand", args: args{inputNumber: 1000}, want: "1k", wantErr: false},
		{name: "shortTwoThousand", args: args{inputNumber: 2000}, want: "2k", wantErr: false},
		{name: "shortTenThousand", args: args{inputNumber: 10000}, want: "10k", wantErr: false},
		{name: "shortHundredThousand", args: args{inputNumber: 100000}, want: "100k", wantErr: false},

		{name: "shortThousandWithNineHundred", args: args{inputNumber: 1900}, want: "1.9k", wantErr: false},
		{name: "shortThousandWithNineHundred", args: args{inputNumber: 1901}, want: "1.9k", wantErr: false},
		{name: "shortThousandWithNineHundred", args: args{inputNumber: 1910}, want: "1.9k", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shortNumberBiggerThanThousandButSmallerThanMillion(tt.args.inputNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortNumberBiggerThanThousandButSmallerThanMillion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("shortNumberBiggerThanThousandButSmallerThanMillion() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shortNumberMinThanThousand(t *testing.T) {
	type args struct {
		inputNumber int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "shortZero", args: args{inputNumber: 0}, want: "0", wantErr: false},
		{name: "shortOne", args: args{inputNumber: 1}, want: "1", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := shortNumberMinThanThousand(tt.args.inputNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("shortNumberMinThanThousand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("shortNumberMinThanThousand() got = %v, want %v", got, tt.want)
			}
		})
	}
}
