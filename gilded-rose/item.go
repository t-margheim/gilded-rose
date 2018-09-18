package gildedrose

import (
	"fmt"
)

// Item describes an item sold by the Gilded Rose Inn
type Item struct {
	name    string
	days    int
	quality int
}

// String outputs a string representation of an Item
func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.name, i.days, i.quality)
}

// New creates a new Item
func New(name string, days, quality int) *Item {
	return &Item{
		name:    name,
		days:    days,
		quality: quality,
	}
}

// UpdateQuality ages the item by a day, and updates the quality of the item
func UpdateQuality(items ...*Item) {
	for _, item := range items {
		if item.name == "normal" {
			normalTick(item)
			return
		}
		if item.name == "Aged Brie" {
			agedBrieTick(item)
			return
		}
		if item.name == "Sulfuras, Hand of Ragnaros" {
			sulfurasTick(item)
			return
		}
		if item.name == "Backstage passes to a TAFKAL80ETC concert" {
			passesTick(item)
			return
		}
	}
}

func normalTick(item *Item) {
	item.quality--
	item.days--
	if item.days < 0 {
		item.quality--
	}
	if item.quality < 0 {
		item.quality = 0
	}
}

func agedBrieTick(item *Item) {
	item.quality++
	item.days--

	if item.days < 0 {
		item.quality++
	}

	if item.quality > 50 {
		item.quality = 50
	}
}

func sulfurasTick(item *Item) {

}

func passesTick(item *Item) {
	item.days--
	item.quality++

	if item.days < 10 {
		item.quality++
	}

	if item.days < 5 {
		item.quality++
	}

	if item.days < 0 {
		item.quality = 0
	}

	if item.quality > 50 {
		item.quality = 50
	}
}
