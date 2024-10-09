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

func CeilingAdd(val, amount, max int) int {
	if val+amount < max {
		return val + amount
	}
	return max
}

func UpdateSellInDate(item *Item) {
	item.SellIn -= 1
}

func UpdateQuality(item *Item) {
	if item.Name == "Sulfuras, Hand of Ragnaros" {
		return
	}

	UpdateSellInDate(item)

	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" {
		item.Quality = FloorSubtract(item.Quality, 1, MIN_QUALITY)
	} else {
		item.Quality = CeilingAdd(item.Quality, 1, MAX_QUALITY)
		if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
			if item.SellIn < 10 {
				item.Quality = CeilingAdd(item.Quality, 1, MAX_QUALITY)
			}
			if item.SellIn < 5 {
				item.Quality = CeilingAdd(item.Quality, 1, MAX_QUALITY)
			}
		}
	}

	if item.SellIn < 0 {
		if item.Name != "Aged Brie" {
			if item.Name != "Backstage passes to a TAFKAL80ETC concert" {
				item.Quality = FloorSubtract(item.Quality, 1, MIN_QUALITY)
			} else {
				item.Quality = MIN_QUALITY
			}
		} else {
			item.Quality = CeilingAdd(item.Quality, 1, MAX_QUALITY)
		}
	}
}
