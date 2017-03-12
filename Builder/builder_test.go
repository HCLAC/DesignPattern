package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	thin := Thin{}
	fat := Fat{}

	director := Director{&thin}
	director.CreatePerson()

	directorF := Director{&fat}
	directorF.CreatePerson()
}
