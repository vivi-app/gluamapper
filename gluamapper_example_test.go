package gluamapper

import (
	"fmt"
	"github.com/awirix/lua"
)

func ExampleMap() {
	type Role struct {
		Name string
	}

	type Person struct {
		Name      string
		Age       int
		WorkPlace string
		Role      []*Role
	}

	L := lua.NewState(nil)
	if err := L.DoString(`
    person = {
      name = "Michel",
      age  = "31", -- weakly input
      work_place = "San Jose",
      role = {
        {
          name = "Administrator"
        },
        {
          name = "Operator"
        }
      }
    }
	`); err != nil {
		panic(err)
	}
	var person Person
	if err := Map(L.GetGlobal("person").(*lua.LTable), &person); err != nil {
		panic(err)
	}
	fmt.Printf("%s %d", person.Name, person.Age)
	// Output:
	// Michel 31
}
