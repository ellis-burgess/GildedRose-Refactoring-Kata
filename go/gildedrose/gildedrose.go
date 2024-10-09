package gildedrose

const (
	MIN_QUALITY         = 0
	MAX_QUALITY         = 50
	SULFURAS_QUALITY    = 80
	BASE_QUALITY_CHANGE = 1
)

type GildedItem interface {
	Update()
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

func CreateGildedItem(item *Item) GildedItem {
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
		gildedItem := CreateGildedItem(item)
		gildedItem.Update()
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

func (item *Item) CalculateQualityChange() int {
	if item.SellIn < 0 {
		return BASE_QUALITY_CHANGE * 2
	}
	return BASE_QUALITY_CHANGE
}

func (item *Item) UpdateQuality() {
	item.Quality = FloorSubtract(item.Quality, item.CalculateQualityChange())
}

func (item *Item) Update() {
	item.SellIn--
	item.UpdateQuality()
}

func (s *Sulfuras) UpdateQuality() {
	// Quality for Sulfuras should always be 80
	s.Quality = SULFURAS_QUALITY
}

func (s *Sulfuras) Update() {
	// SellIn does not change
	s.UpdateQuality()
}

func (p *BackstagePass) CalculateQualityChange() int {
	if p.SellIn >= 10 {
		return 1
	}

	if p.SellIn >= 5 {
		return 2
	}

	return 3
}

func (p *BackstagePass) UpdateQuality() {
	if p.SellIn >= 0 {
		p.Quality = CeilingAdd(p.Quality, p.CalculateQualityChange())
	} else {
		p.Quality = 0
	}
}

func (p *BackstagePass) Update() {
	p.SellIn--
	p.UpdateQuality()
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

func (b *AgedBrie) Update() {
	b.SellIn--
	b.UpdateQuality()
}
