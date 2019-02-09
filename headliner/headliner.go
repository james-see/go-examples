package main

import (
    "gopkg.in/jdkato/prose.v2"
)

func main() {
    doc, _ := prose.NewDocument("Lebron James plays basketball in Los Angeles.")
    for _, ent := range doc.Entities() {
        fmt.Println(ent.Text, ent.Label)
        // Lebron James PERSON
        // Los Angeles GPE
    }
}

