package gildedrose

const MIN_QUALITY = 0
const MAX_QUALITY = 50

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateItems(items []*Item) {
	for _, item := range items {
		UpdateQuality(item)
	}
}

func FloorSubtract(val, amount, min int) int {
	if val-amount > min {
		return val - amount
	}
	return min
}

func UpdateQuality(item *Item) {
	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" && item.Name != "Sulfuras, Hand of Ragnaros" {
		item.Quality = FloorSubtract(item.Quality, 1, MIN_QUALITY)
	} else {
		if item.Quality < MAX_QUALITY {
			item.Quality = item.Quality + 1
			if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
				if item.SellIn < 11 {
					if item.Quality < MAX_QUALITY {
						item.Quality = item.Quality + 1
					}
				}
				if item.SellIn < 6 {
					if item.Quality < MAX_QUALITY {
						item.Quality = item.Quality + 1
					}
				}
			}
		}
	}

	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.SellIn = item.SellIn - 1
	}

	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				if item.Quality > MIN_QUALITY {
					if item.Name != "Sulfuras, Hand of Ragnaros" {
						item.Quality = item.Quality - 1
					}
				}
			} else {
				item.Quality = item.Quality - item.Quality
			}
		} else {
			if item.Quality < MAX_QUALITY {
				item.Quality = item.Quality + 1
			}
		}
	}
}
