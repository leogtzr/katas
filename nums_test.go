package main

import "testing"

func Test_numToWord(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "67 -> sixty seven",
			args: args{
				num: 67,
			},
			want: "sixty seven",
		},
		{
			name: "321 -> three hundred and twenty one",
			args: args{
				num: 321,
			},
			want: "three hundred and twenty one",
		},
		{
			name: "0 - zero",
			args: args{
				num: 0,
			},
			want: "zero",
		},
		{
			name: "11 - eleven",
			args: args{
				num: 11,
			},
			want: "eleven",
		},
		{
			name: "123720",
			args: args{
				num: 123720,
			},
			want: "one hundred and twenty three thousand seven hundred and twenty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numToWord(tt.args.num); got != tt.want {
				t.Errorf("numToWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
