package main

import (
	"fmt"

	gildedrose "github.com/t-margheim/gilded-rose/gilded-rose"
)

var things = []struct {
	name    string
	sellIn  int
	quality int
}{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	// {"Conjured Mana Cake", 3, 6},
}

func main() {
	var items []*gildedrose.Item

	for _, thing := range things {
		item := gildedrose.New(thing.name, thing.sellIn, thing.quality)
		items = append(items, item)
		fmt.Println(item)
	}

	gildedrose.UpdateQuality(items...)

	for _, item := range items {
		fmt.Println(item)
	}
}
