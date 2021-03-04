package benchs

import (
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestDoWork(t *testing.T) {
	runtime.GOMAXPROCS(12)
	type args struct {
		c        int
		endPrime uint64
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{name: "i7",
			args: args{
				c:        6,
				endPrime: 200000,
			},
			want: 1 * time.Second,
			/*
				1	5.6s
				5	5.7s
				6	5.8s	+0.1
						6	6.0s	GOMAXPROCS=6
				7	6.4s	+0.6
				8	6.8s	+0.4
				9	7.1s	+0.3
				11 	7.5s	+0.4
				12	7.9s	+0.4
						12	7.9s	GOMAXPROCS=12
						12 	11.8s	GOMAXPROCS=6
				13	8.6s	+0.7
				14	9.3s	+0.7
				15	10.0s	+0.7
			*/
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//if got := DoWork(tt.args.c, tt.args.endPrime); got != tt.want {
			if got := DoWork(tt.args.c, tt.args.endPrime); got < tt.want {
				t.Errorf("DoWork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindAllPrime(t *testing.T) {
	type args struct {
		start uint64
		end   uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAllPrime(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAllPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.value); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_getCPUSample(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		wantUser  uint64
// 		wantIdle  uint64
// 		wantTotal uint64
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotUser, gotIdle, gotTotal := getCPUSample()
// 			if gotUser != tt.wantUser {
// 				t.Errorf("getCPUSample() gotUser = %v, want %v", gotUser, tt.wantUser)
// 			}
// 			if gotIdle != tt.wantIdle {
// 				t.Errorf("getCPUSample() gotIdle = %v, want %v", gotIdle, tt.wantIdle)
// 			}
// 			if gotTotal != tt.wantTotal {
// 				t.Errorf("getCPUSample() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
// 			}
// 		})
// 	}
// }
//
// func Test_handleV(t *testing.T) {
// 	type args struct {
// 		c *gin.Context
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 		})
// 	}
// }
