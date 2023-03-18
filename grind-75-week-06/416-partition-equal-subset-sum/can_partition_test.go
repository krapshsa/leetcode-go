package partition_equal_subset_sum

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type PartitionSuite struct {
	suite.Suite
}

func TestPartitionTestSuite(t *testing.T) {
	suite.Run(t, new(PartitionSuite))
}

func (s *PartitionSuite) SetupTest() {
}

func (s *PartitionSuite) Test_1() {
	assert.Equal(s.T(), true, canPartition([]int{1, 5, 11, 5}), "")
}

func (s *PartitionSuite) Test_2() {
	assert.Equal(s.T(), false, canPartition([]int{1, 2, 3, 5}), "")
}

func (s *PartitionSuite) Test_3() {
	assert.Equal(s.T(), true, canPartition([]int{3, 3, 3, 4, 5}), "")
}

func (s *PartitionSuite) Test_4() {
	assert.Equal(s.T(), false, canPartition([]int{1, 2, 5}), "")
}
