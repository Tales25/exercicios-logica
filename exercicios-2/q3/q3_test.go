package q3

import (
	"fmt"
	"testing"
)

func TestChooseVideo(t *testing.T) {
	tests := []struct {
		videos  []Video
		time    int
		want    int
		wantErr bool
	}{
		{
			time: 9,
			videos: []Video{
				createVideo(1, 3),
				createVideo(5, 4),
				createVideo(7, 7),
				createVideo(6, 1),
				createVideo(6, 9),
			},
			want: 3,
		},
		{
			time: 4,
			videos: []Video{
				createVideo(4, 1),
				createVideo(3, 2),
				createVideo(3, 3),
				createVideo(2, 4),
			},
			want: 2,
		},
		{
			time: 7,
			videos: []Video{
				createVideo(5, 2),
				createVideo(5, 1),
				createVideo(5, 3),
				createVideo(5, 9),
				createVideo(5, 7),
			},
			want: 3,
		},
		{
			time: 33,
			videos: []Video{
				createVideo(54, 42),
				createVideo(71, 24),
				createVideo(69, 99),
				createVideo(96, 1),
			},
			wantErr: true,
		},
		{
			time: 179,
			videos: []Video{
				createVideo(55, 66),
				createVideo(77, 88),
			},
			want: 2,
		},
		{
			time: 0,
			videos: []Video{
				createVideo(1, 1),
				createVideo(2, 2),
				createVideo(3, 3),
				createVideo(4, 4),
				createVideo(5, 5),
			},
			wantErr: true,
		},
		{
			time:    10,
			videos:  []Video{},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("videos=%v,time=%d", test.videos, test.time), func(t *testing.T) {
			for i := range test.videos {
				test.videos[i].ID = i + 1
			}

			got, err := ChooseVideo(test.videos, test.time)
			if test.wantErr {
				if err == nil {
					t.Errorf("ChooseVideo() error = %v, wantErr %v", err, test.wantErr)
				}
				return
			}
			if got != test.videos[test.want-1] {
				t.Errorf("ChooseVideo() got = %v, want %v", got, test.videos[test.want-1])
			}
		})
	}
}

func createVideo(duration int, entertainment int) Video {
	return Video{
		Duration:      duration,
		Entertainment: entertainment,
	}
}
