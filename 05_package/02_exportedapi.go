package main

import (
	"fmt"
	"strings"
)

type Catalog struct {
	items []string
}

func NewCatalog() *Catalog {
	return &Catalog{}
}

func (c *Catalog) Add(name string) {
	// TODO: append the normalized item name to c.items.
	c.items = append(c.items, normalizeItemName(name))
}

func (c *Catalog) Count() int {
	// TODO: return the number of stored items.
	return len(c.items)
}

func (c *Catalog) summary() string {
	return strings.Join(c.items, ", ")
}

func normalizeItemName(name string) string {
	// TODO: return strings.ToUpper(name).
	return strings.ToUpper(name)
}

func main() {
	catalog := NewCatalog()
	// TODO: add "notebook", "pen", and "eraser" to catalog in that order.
	catalog.Add("notebook")
	catalog.Add("pen")
	catalog.Add("eraser")

	fmt.Println("Exported type: Catalog")
	fmt.Println("Items:", catalog.summary())
	fmt.Println("Count:", catalog.Count())
}
