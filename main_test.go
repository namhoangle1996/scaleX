package main

import (
	"reflect"
	"testing"
)

func Test_sol1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		{
			name: "",
			args: args{
				n: 2,
			},
			wantRes: 5,
		},
		{
			name: "",
			args: args{
				n: 3,
			},
			wantRes: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := sol1(tt.args.n); gotRes != tt.wantRes {
				t.Errorf("sol1() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func Test_getFibonacci(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		args    args
		wantRes []int64
	}{
		{
			name: "",
			args: args{
				n: 10,
			},
			wantRes: []int64{5, 8, 13, 21, 34, 55},
		}, {
			name: "",
			args: args{
				n: 4,
			},
			wantRes: []int64{1, 1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := getFibonacci(tt.args.n); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("getFibonacci() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
