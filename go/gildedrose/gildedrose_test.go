package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func TestQualityDecreasesBeforeSellInDate(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "test_item",
			SellIn:  10,
			Quality: 10,
		},
	}

	gildedrose.UpdateQuality(items)

	want := 9
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestSellInDateDecrements(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "test_item",
			SellIn:  10,
			Quality: 10,
		},
	}

	gildedrose.UpdateQuality(items)

	want := 9
	got := items[0].SellIn

	if want != got {
		t.Errorf("SellIn: want %d, but got %d", want, got)
	}
}

func TestQualityDecreasesTwiceAfterSellInDate(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "test_item",
			SellIn:  -1,
			Quality: 10,
		},
	}

	gildedrose.UpdateQuality(items)

	want := 8
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestQualityIsNotNegativeBeforeSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "test_item",
			SellIn:  1,
			Quality: 0,
		},
	}

	gildedrose.UpdateQuality(items)

	want := 0
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestQualityIsNotNegativeAfterSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "test_item",
			SellIn:  -1,
			Quality: 0,
		},
	}

	gildedrose.UpdateQuality(items)

	want := 0
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestAgedBrieQualityIncreases(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Aged Brie",
			SellIn:  10,
			Quality: 0,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 1
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestAgedBrieQualityIncreasesTwiceAfterSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Aged Brie",
			SellIn:  0,
			Quality: 10,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 12
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestAgedBrieQualityDoesNotExceedFiftyBeforeSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Aged Brie",
			SellIn:  10,
			Quality: 50,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 50
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestAgedBrieQualityDoesNotExceedFiftyAfterSellIn(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Aged Brie",
			SellIn:  -1,
			Quality: 50,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 50
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestSulfurasSellInDoesNotDecrementWhilePositive(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  1,
			Quality: 10,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 1
	got := items[0].SellIn

	if want != got {
		t.Errorf("SellIn: want %d, but got %d", want, got)
	}
}

func TestSulfurasSellInDoesNotDecrementWhileNegative(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  -1,
			Quality: 10,
		},
	}

	gildedrose.UpdateQuality(items)
	want := -1
	got := items[0].SellIn

	if want != got {
		t.Errorf("SellIn: want %d, but got %d", want, got)
	}
}

func TestSulfurasQualityDoesNotDecrement(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Sulfuras, Hand of Ragnaros",
			SellIn:  10,
			Quality: 80,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 80
	got := items[0].Quality

	if want != got {
		t.Errorf("SellIn: want %d, but got %d", want, got)
	}
}

func TestBackstagePassIncreasesByOneIfSellInGreaterThanTen(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  15,
			Quality: 20,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 21
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestBackstagePassQualityIncreasesByTwoIfSellInBetweenFiveAndTen(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  6,
			Quality: 20,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 22
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestBackstagePassQualityIncreasesByThreeIfSellInLessThanFive(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  3,
			Quality: 20,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 23
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}
