package challenger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSliceRange(t *testing.T) {
	var frames = make([]int, 10, 20)
	testFunc := func(frames []int, startingChunkIndex int) error {
		if startingChunkIndex > len(frames) {
			return fmt.Errorf("startingChunkIndex is out of frames range, startingChunkIndex: %d, len(frames): %d", startingChunkIndex, len(frames))
		} else {
			return nil
		}
	}

	require.NoError(t, testFunc(frames, len(frames)-1))
	require.NoError(t, testFunc(frames, len(frames)))
	require.Error(t, testFunc(frames, len(frames)+1))
}
