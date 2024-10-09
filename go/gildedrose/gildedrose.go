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

func FloorSubtract(val, amount int) int {
	if val-amount > MIN_QUALITY {
		return val - amount
	}
	return MIN_QUALITY
}

func CeilingAdd(val, amount int) int {
	if val+amount < MAX_QUALITY {
		return val + amount
	}
	return MAX_QUALITY
}

func UpdateSellInDate(item *Item) {
	item.SellIn -= 1
}

func CalculateQualityChange(item *Item) int {
	qualityChange := BASE_QUALITY_CHANGE

	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		return 3 - (item.SellIn / 5)
	}

	if item.SellIn < 0 {
		qualityChange = qualityChange * 2
	}

	return qualityChange
}

func UpdateQuality(item *Item) {
	qualityChange := CalculateQualityChange(item)

	switch name := item.Name; name {
	case "Backstage passes to a TAFKAL80ETC concert":
		if item.SellIn < 0 {
			item.Quality = MIN_QUALITY
		} else {
			item.Quality = CeilingAdd(item.Quality, qualityChange)
		}
	case "Aged Brie":
		item.Quality = CeilingAdd(item.Quality, qualityChange)
	default:
		item.Quality = FloorSubtract(item.Quality, qualityChange)
	}
}
