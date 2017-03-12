package prototype

import (
	"testing"
)

func TestPrototype(t *testing.T) {
	resume := NewResume()

	resume.setPersonalInfo("hclA", "男", "22")
	resume.setWorkExperience("3", "Apple")
	resume.display()

	cloneresume := resume.clone()
	cloneresume.setPersonalInfo("hclB", "女", "22")
	cloneresume.setWorkExperience("3", "HW")
	cloneresume.display()
}
