package gildedrose

const (
	MIN_QUALITY         = 0
	MAX_QUALITY         = 50
	BASE_QUALITY_CHANGE = 1
)

type ItemInterface interface {
	CalculateQualityChange() int
	UpdateQuality()
	UpdateSellInDate()
}

type Item struct {
	Name            string
	SellIn, Quality int
}

type Sulfuras struct {
	*Item
}

type BackstagePass struct {
	*Item
}

type AgedBrie struct {
	*Item
}

func UpdateItems(items []*Item) {
	for _, item := range items {
		var sortedItem ItemInterface
		switch name := item.Name; name {
		case "Sulfuras, Hand of Ragnaros":
			sortedItem = &Sulfuras{item}
		case "Backstage passes to a TAFKAL80ETC concert":
			sortedItem = &BackstagePass{item}
		case "Aged Brie":
			sortedItem = &AgedBrie{item}
		default:
			sortedItem = item
		}
		sortedItem.UpdateSellInDate()
		sortedItem.UpdateQuality()
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

func (item *Item) UpdateSellInDate() {
	item.SellIn -= 1
}

func (s *Sulfuras) CalculateQualityChange() int {
	return 80
}

func (s *Sulfuras) UpdateQuality() {
	s.Quality = s.CalculateQualityChange()
}

func (s *Sulfuras) UpdateSellInDate() {
	s.SellIn = s.SellIn * 1
}

func (p *BackstagePass) CalculateQualityChange() int {
	return 3 - (p.SellIn / 5)
}

func (p *BackstagePass) UpdateQuality() {
	if p.SellIn >= 0 {
		p.Quality = CeilingAdd(p.Quality, p.CalculateQualityChange())
	} else {
		p.Quality = 0
	}
}

func (b *AgedBrie) CalculateQualityChange() int {
	if b.SellIn >= 0 {
		return BASE_QUALITY_CHANGE
	} else {
		return 2 * BASE_QUALITY_CHANGE
	}
}

func (b *AgedBrie) UpdateQuality() {
	b.Quality = CeilingAdd(b.Quality, b.CalculateQualityChange())
}

func (item *Item) CalculateQualityChange() int {
	if item.SellIn < 0 {
		return BASE_QUALITY_CHANGE * 2
	}
	return BASE_QUALITY_CHANGE
}

func (item *Item) UpdateQuality() {
	item.Quality = FloorSubtract(item.Quality, item.CalculateQualityChange())
}
