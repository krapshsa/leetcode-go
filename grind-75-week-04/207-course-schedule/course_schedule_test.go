package course_schedule

import "testing"
import "github.com/stretchr/testify/suite"
import "github.com/stretchr/testify/assert"

type CourseScheduleSuite struct {
	suite.Suite
}

func TestCourseScheduleTestSuite(t *testing.T) {
	suite.Run(t, new(CourseScheduleSuite))
}

func (s *CourseScheduleSuite) SetupTest() {
}

func (s *CourseScheduleSuite) Test_1() {
	assert.Equal(s.T(), true, canFinish(2, [][]int{[]int{1, 0}}), "")
}

func (s *CourseScheduleSuite) Test_2() {
	assert.Equal(s.T(), false, canFinish(2, [][]int{[]int{1, 0}, []int{0, 1}}), "")
}
