package headliner

import (
	"fmt"

	prose "gopkg.in/jdkato/prose.v2"
)

// GetEntities get those things from prose.
func GetEntities() {
	doc, _ := prose.NewDocument("Lebron James plays basketball in Los Angeles.")
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// Lebron James PERSON
		// Los Angeles GPE
	}
}
