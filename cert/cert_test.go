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

	var expectedName string = "GOLANG COURSE"
	if c.Course != expectedName {
		t.Errorf("Course name is invalid. Shoudl have "+expectedName+". got=%v", c.Course)
	}
}
