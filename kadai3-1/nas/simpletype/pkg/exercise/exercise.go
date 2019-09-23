package exercise

// Exercise has questions and now question number
type Exercise struct {
	Questions      []string
	NowQuestionNum int
}

// Next return Question stil has questions or not
func (e *Exercise) Next() bool {
	if e.NowQuestionNum >= len(e.Questions) {
		e.NowQuestionNum++
		return false
	}
	e.NowQuestionNum++
	return true
}

// Get return question
func (e *Exercise) Get() string {
	if e.NowQuestionNum == 0 {
		return ""
	}
	if e.NowQuestionNum > len(e.Questions) {
		return ""
	}
	return e.Questions[e.NowQuestionNum-1]
}
