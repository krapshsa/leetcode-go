package k_closest_points_to_origin

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type KClosetSuite struct {
	suite.Suite
}

func TestKClosetTestSuite(t *testing.T) {
	suite.Run(t, new(KClosetSuite))
}

func (s *KClosetSuite) SetupTest() {
}

func (s *KClosetSuite) Test_1() {
	assert.Equal(
		s.T(),
		[][]int{{-2, 2}},
		kClosest([][]int{{1, 3}, {-2, 2}}, 1),
		"1",
	)
}

func (s *KClosetSuite) Test_2() {
	assert.Equal(
		s.T(),
		[][]int{{3, 3}, {-2, 4}},
		kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2),
		"2",
	)
}

func (s *KClosetSuite) Test_3() {
	assert.Equal(
		s.T(),
		[][]int{{2, 31}, {-27, -38}, {-55, -39}, {-57, -67}, {34, -84}},
		kClosest([][]int{{68, 97}, {34, -84}, {60, 100}, {2, 31}, {-27, -38}, {-73, -74}, {-55, -39}, {62, 91}, {62, 92}, {-57, -67}}, 5),
		"3",
	)
}
