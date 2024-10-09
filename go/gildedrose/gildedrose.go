package gildedrose

const MIN_QUALITY = 0
const MAX_QUALITY = 50
const BASE_QUALITY_CHANGE = 1

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateItems(items []*Item) {
	for _, item := range items {
		if item.Name != "Sulfuras, Hand of Ragnaros" {
			UpdateSellInDate(item)
			UpdateQuality(item)
		}
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
	qualityChange := BASE_QUALITY_CHANGE

	if item.SellIn < 0 {
		qualityChange = qualityChange * 2
	}

	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		if item.SellIn >= 0 {
			qualityChange = 3 - (item.SellIn / 5)
			item.Quality = CeilingAdd(item.Quality, qualityChange, MAX_QUALITY)
		} else {
			item.Quality = 0
		}
	} else if item.Name == "Aged Brie" {
		item.Quality = CeilingAdd(item.Quality, qualityChange, MAX_QUALITY)
	} else {
		item.Quality = FloorSubtract(item.Quality, qualityChange, MIN_QUALITY)
	}
}
