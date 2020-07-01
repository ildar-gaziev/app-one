package count

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestCount(t *testing.T) {
	s := "qwerrtyeege"

	require.Equal(t, Count(s, 'e'), 4, "counting 'e' in " + s)
	require.Equal(t, Count(s, 'r'), 2, "counting 'r' in " + s)
	require.Equal(t, Count(s, 'q'), 1, "counting 'q' in " + s)

}
