package cert

import "testing"

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2018-05-31")
	if err != nil {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should not be empty. got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is invalid. Shoudl have 'GOLANG COURSE'. got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned for empty course")
	}
}

func TestCourseNameTooLong(t *testing.T) {
	_, err := New("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", "Bob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned when course's name is too long")
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Course", "", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned for empty name")
	}
}

func TestNameTooLong(t *testing.T) {
	_, err := New("Course", "Boooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooob", "2018-05-31")
	if err == nil {
		t.Errorf("Error should be returned when name is too long")
	}
}
