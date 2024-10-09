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

func SortItem(item *Item) ItemInterface {
	switch name := item.Name; name {
	case "Sulfuras, Hand of Ragnaros":
		return &Sulfuras{item}
	case "Backstage passes to a TAFKAL80ETC concert":
		return &BackstagePass{item}
	case "Aged Brie":
		return &AgedBrie{item}
	default:
		return item
	}
}

func UpdateItems(items []*Item) {
	for _, item := range items {
		sortedItem := SortItem(item)
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

func (item *Item) CalculateQualityChange() int {
	if item.SellIn < 0 {
		return BASE_QUALITY_CHANGE * 2
	}
	return BASE_QUALITY_CHANGE
}

func (item *Item) UpdateQuality() {
	item.Quality = FloorSubtract(item.Quality, item.CalculateQualityChange())
}

func (s *Sulfuras) CalculateQualityChange() int {
	return 80
}

func (s *Sulfuras) UpdateQuality() {
	s.Quality = s.CalculateQualityChange()
}

func (s *Sulfuras) UpdateSellInDate() {
	// sellin does not change for sulfuras
}

func (p *BackstagePass) CalculateQualityChange() int {
	if p.SellIn >= 10 {
		return 1
	} else if p.SellIn >= 5 {
		return 2
	} else {
		return 3
	}
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
