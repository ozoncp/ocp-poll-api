package utils

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		slice []int
		size  int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]int
		wantErr bool
	}{
		{
			name: "nilSlice",
			args: args{
				slice: nil,
				size:  3,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "sizeIs0",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				size:  0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct1",
			args: args{
				slice: []int{1, 2, 3, 4, 5, 6},
				size:  3,
			},
			want: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name: "correct2",
			args: args{
				slice: []int{1, 2, 3, 4, 5},
				size:  3,
			},
			want: [][]int{{1, 2, 3}, {4, 5}},
		},
		{
			name: "correct3",
			args: args{
				slice: []int{},
				size:  3,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Split(tt.args.slice, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("Split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRevertMap(t *testing.T) {
	type args struct {
		sourceMap map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "nilSourceMap",
			args: args{
				sourceMap: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct",
			args: args{
				sourceMap: map[string]string{"key1": "value1", "key2": "value2"},
			},
			want: map[string]string{"value1": "key1", "value2": "key2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := RevertMap(tt.args.sourceMap)
			if (err != nil) != tt.wantErr {
				t.Errorf("RevertMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RevertMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args struct {
		slice []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "nilSlice",
			args: args{
				slice: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct",
			args: args{
				slice: []string{"Hello", "World", "Value1", "value2"},
			},
			want: []string{"value2"},
		},
		{
			name: "correct2",
			args: args{
				slice: []string{"Hello", "World", "value1", "value2"},
			},
			want: []string{"value1", "value2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Filter(tt.args.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("Filter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
