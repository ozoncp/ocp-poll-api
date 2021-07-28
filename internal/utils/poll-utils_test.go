package utils

import (
	"reflect"
	"testing"
)
import "github.com/ozoncp/ocp-poll-api/internal/models"

func TestPollsToMap(t *testing.T) {
	type args struct {
		polls []models.Poll
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint64]models.Poll
		wantErr bool
	}{
		{
			name: "nilSourceMap",
			args: args{
				polls: nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct",
			args: args{
				polls: []models.Poll{{1, 1, 1, 1}, {2, 2, 2, 2}},
			},
			want: map[uint64]models.Poll{1: {1, 1, 1, 1}, 2: {2, 2, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PollsToMap(tt.args.polls)
			if (err != nil) != tt.wantErr {
				t.Errorf("PollsToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PollsToMap() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPoll(t *testing.T) {
	type args struct {
		slice []models.Poll
		size  int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]models.Poll
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
				slice: []models.Poll{{1, 1, 1, 1}, {2, 2, 2, 2}},
				size:  0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "correct1",
			args: args{
				slice: []models.Poll{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}, {5, 5, 5, 5}, {6, 6, 6, 6}},
				size:  3,
			},
			want: [][]models.Poll{{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}}, {{4, 4, 4, 4}, {5, 5, 5, 5}, {6, 6, 6, 6}}},
		},
		{
			name: "correct2",
			args: args{
				slice: []models.Poll{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}, {4, 4, 4, 4}, {5, 5, 5, 5}},
				size:  3,
			},
			want: [][]models.Poll{{{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3}}, {{4, 4, 4, 4}, {5, 5, 5, 5}}},
		},
		{
			name: "correct3",
			args: args{
				slice: []models.Poll{},
				size:  3,
			},
			want: [][]models.Poll{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitPoll(tt.args.slice, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitPoll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitPoll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
