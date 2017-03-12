package abstractfactory

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	uData := User{1, "u"}
	dData := Department{1, "d"}
	data := DataAccess{}

	iU := data.createUser("access")
	iU.insert(&uData)
	gU := iU.getUser(1)
	fmt.Println(gU)

	fmt.Println("================")

	iD := data.createDepartment("sqlserver")
	iD.insert(&dData)
	gD := iD.getDepartment(1)
	fmt.Println(gD)

	if iS := data.createDepartment("a"); iS != nil {
		t.Error()
	}
}
