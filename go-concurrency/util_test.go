package main

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"sync/atomic"
	"testing"
)

type testCase struct {
	pieces    int
	workers   int
	chunkSize int
}

func (c testCase) String() string {
	return fmt.Sprintf("pieces:%d,workers:%d,chunkSize:%d", c.pieces, c.workers, c.chunkSize)
}

var cases = []testCase{
	{
		pieces:    1000,
		workers:   10,
		chunkSize: 1,
	},
	{
		pieces:    1000,
		workers:   10,
		chunkSize: 10,
	},
	{
		pieces:    1000,
		workers:   10,
		chunkSize: 100,
	},
	{
		pieces:    999,
		workers:   10,
		chunkSize: 13,
	},
}

func TestParallelizeUntil(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.String(), func(t *testing.T) {
			seen := make([]int32, tc.pieces)
			ctx := context.Background()
			ParallelizeUntil(ctx, tc.workers, tc.pieces, func(p int) {
				atomic.AddInt32(&seen[p], 1)
			}, WithChunkSize(tc.chunkSize))

			wantSeen := make([]int32, tc.pieces)
			for i := 0; i < tc.pieces; i++ {
				wantSeen[i] = 1
			}
			if diff := cmp.Diff(wantSeen, seen); diff != "" {
				t.Errorf("bad number of visits (-want,+got):\n%s", diff)
			}
		})
	}
}

func TestHandleCrash(t *testing.T) {

	go func() {
		defer HandleCrash()
		n := 0
		fmt.Println(1/ n)
	}()
}