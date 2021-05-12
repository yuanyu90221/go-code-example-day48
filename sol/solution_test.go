package sol

import (
	"reflect"
	"testing"
)

func TestRightShiftLinkedListByKSteps(t *testing.T) {
	type args struct {
		l *ListNode
		k int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example1",
			args: args{
				l: ConvertList(&[]int{1, 2, 3, 4, 5}),
				k: 2,
			},
			want: ConvertList(&[]int{4, 5, 1, 2, 3}),
		},
		{
			name: "Example2",
			args: args{
				l: ConvertList(&[]int{1, 2, 3}),
				k: 3,
			},
			want: ConvertList(&[]int{1, 2, 3}),
		},
		{
			name: "Example3",
			args: args{
				l: ConvertList(&[]int{1, 2, 3}),
				k: 4,
			},
			want: ConvertList(&[]int{3, 1, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RightShiftLinkedListByKSteps(tt.args.l, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RightShiftLinkedListByKSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
