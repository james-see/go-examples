package headliner4

import (
	"fmt"

	prose "gopkg.in/jdkato/prose.v2"
)

// GetEntities get those things from prose.
func GetEntities(instatext string) {
	fmt.Println("Example 1 entity extract is: Lebron James plays basketball in Los Angeles.")
	doc, _ := prose.NewDocument("Lebron James plays basketball in Los Angeles.")
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// Lebron James PERSON
		// Los Angeles GPE
	}
	fmt.Println("Example 2 entity extract is from the instagram text example.")
	doc, _ = prose.NewDocument(instatext)
	for _, ent := range doc.Entities() {
		fmt.Println(ent.Text, ent.Label)
		// James Campbell PERSON
	}
}
