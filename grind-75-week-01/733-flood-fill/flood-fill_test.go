package flood_fill

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_floodFill(t *testing.T) {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	expected := [][]int{
		{2, 2, 2},
		{2, 2, 0},
		{2, 0, 1},
	}
	assert.Equal(t, expected, floodFill(image, 1, 1, 2), "From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.\nNote the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.")
}

func Test_floodFill_do_nothing(t *testing.T) {
	image := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}
	expected := [][]int{
		{0, 0, 0},
		{0, 0, 0},
	}
	assert.Equal(t, expected, floodFill(image, 0, 0, 0), "Do nothing")
}
