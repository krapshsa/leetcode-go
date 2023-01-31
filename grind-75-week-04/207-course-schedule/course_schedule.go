package course_schedule

type Course struct {
	prerequisites []int
	walked        bool
	finish        bool
}

func fill(course int, m *map[int]*Course) bool {
	c := (*m)[course]
	if true == (*c).walked {
		return (*c).finish
	}

	(*c).walked = true
	for i := 0; i < len((*c).prerequisites); i++ {
		result := fill((*c).prerequisites[i], m)
		if false == result {
			return false
		}
	}

	(*c).finish = true
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make(map[int]*Course)
	for i := 0; i < numCourses; i++ {
		m[i] = &Course{make([]int, 0), false, false}
	}

	for i := 0; i < len(prerequisites); i++ {
		course, preRequest := prerequisites[i][0], prerequisites[i][1]
		c := m[course]
		(*c).prerequisites = append((*c).prerequisites, preRequest)
	}

	for i := 0; i < numCourses; i++ {
		result := fill(i, &m)
		if false == result {
			return false
		}
	}

	return true
}
