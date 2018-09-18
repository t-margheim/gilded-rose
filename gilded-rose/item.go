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
		if item.name != "Aged Brie" && item.name != "Backstage passes to a TAFKAL80ETC concert" {
			if item.quality > 0 {
				if item.name != "Sulfuras, Hand of Ragnaros" {
					item.quality = item.quality - 1
				}
			}
		} else {
			if item.quality < 50 {
				item.quality = item.quality + 1
				if item.name == "Backstage passes to a TAFKAL80ETC concert" {
					if item.days < 11 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
					if item.days < 6 {
						if item.quality < 50 {
							item.quality = item.quality + 1
						}
					}
				}
			}
		}

		if item.name != "Sulfuras, Hand of Ragnaros" {
			item.days = item.days - 1
		}

		if item.days < 0 {
			if item.name != "Aged Brie" {
				if item.name != "Backstage passes to a TAFKAL80ETC concert" {
					if item.quality > 0 {
						if item.name != "Sulfuras, Hand of Ragnaros" {
							item.quality = item.quality - 1
						}
					}
				} else {
					item.quality = item.quality - item.quality
				}
			} else {
				if item.quality < 50 {
					item.quality = item.quality + 1
				}
			}
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
