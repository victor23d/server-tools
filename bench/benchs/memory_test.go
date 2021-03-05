package benchs

import (
	"reflect"
	"testing"
)

func TestCleanUpMem(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func TestGetMemUsage(t *testing.T) {
	tests := []struct {
		name string
		want Memory
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMemUsage(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemUsage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeBlock(t *testing.T) {
	type args struct {
		mb int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t13",
			args: args{
				mb: 13,
			},
			want: 13,
		},
		{
			name: "t130",
			args: args{
				mb: 130,
			},
			want: 130,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeBlock(tt.args.mb)
			if tt.want != len(got) {
				t.Errorf("MakeBlock() = %v, want %v", len(got), tt.want)
			}
			//t.Log(got)
		})
	}
}
