package zero_one_matrix

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type matrixSuite struct {
	suite.Suite
}

func TestMatrixTestSuite(t *testing.T) {
	suite.Run(t, new(matrixSuite))
}

func (s *matrixSuite) SetupTest() {
}

func (s *matrixSuite) Test_1() {
	assert.Equal(
		s.T(),
		[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}),
		"test 1",
	)
}

func (s *matrixSuite) Test_2() {
	assert.Equal(
		s.T(),
		[][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}),
		"test 2",
	)
}

func (s *matrixSuite) Test_3() {
	assert.Equal(
		s.T(),
		[][]int{{2, 1, 0, 0, 1, 0, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 1, 1, 2, 2, 1}, {1, 1, 1, 0, 0, 1, 2, 2, 1, 0}, {0, 1, 2, 1, 0, 1, 2, 3, 2, 1}, {0, 0, 1, 2, 1, 2, 1, 2, 1, 0}, {1, 1, 2, 3, 2, 1, 0, 1, 1, 1}, {0, 1, 2, 3, 2, 1, 1, 0, 0, 1}, {1, 2, 1, 2, 1, 0, 0, 1, 1, 2}, {0, 1, 0, 1, 1, 0, 1, 2, 2, 3}, {1, 2, 1, 0, 1, 0, 1, 2, 3, 4}},
		updateMatrix([][]int{{1, 1, 0, 0, 1, 0, 0, 1, 1, 0}, {1, 0, 0, 1, 0, 1, 1, 1, 1, 1}, {1, 1, 1, 0, 0, 1, 1, 1, 1, 0}, {0, 1, 1, 1, 0, 1, 1, 1, 1, 1}, {0, 0, 1, 1, 1, 1, 1, 1, 1, 0}, {1, 1, 1, 1, 1, 1, 0, 1, 1, 1}, {0, 1, 1, 1, 1, 1, 1, 0, 0, 1}, {1, 1, 1, 1, 1, 0, 0, 1, 1, 1}, {0, 1, 0, 1, 1, 0, 1, 1, 1, 1}, {1, 1, 1, 0, 1, 0, 1, 1, 1, 1}}),
		"test 3",
	)
}
