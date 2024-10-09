package gildedrose_test

import (
	"reflect"
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

func TestBackStagePassQualityBecomesZeroAfterSellInDate(t *testing.T) {
	var items = []*gildedrose.Item{
		{
			Name:    "Backstage passes to a TAFKAL80ETC concert",
			SellIn:  0,
			Quality: 50,
		},
	}

	gildedrose.UpdateQuality(items)
	want := 0
	got := items[0].Quality

	if want != got {
		t.Errorf("Quality: want %d, but got %d", want, got)
	}
}

func TestOutputMatchesExpectedOutput(t *testing.T) {
	var items = []*gildedrose.Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	var itemsExpectedDay1 = []*gildedrose.Item{
		{"+5 Dexterity Vest", 9, 19},
		{"Aged Brie", 1, 1},
		{"Elixir of the Mongoose", 4, 6},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Conjured Mana Cake", 2, 5},
	}

	var itemsExpectedDay5 = []*gildedrose.Item{
		{"+5 Dexterity Vest", 5, 15},
		{"Aged Brie", -3, 8},
		{"Elixir of the Mongoose", 0, 2},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Conjured Mana Cake", -2, 0},
	}

	var itemsExpectedDay10 = []*gildedrose.Item{
		{"+5 Dexterity Vest", 0, 10},
		{"Aged Brie", -8, 18},
		{"Elixir of the Mongoose", -5, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 35},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -5, 0},
		{"Conjured Mana Cake", -7, 0},
	}

	i := 1
	gildedrose.UpdateQuality(items)
	if !reflect.DeepEqual(items, itemsExpectedDay1) {
		for idx := 0; idx < len(items); idx++ {
			if !reflect.DeepEqual(items[idx], itemsExpectedDay1[idx]) {
				t.Logf("Unexpected item at Day 1, expected %v, got %v", items[idx], itemsExpectedDay1[idx])
			}
		}
		t.Errorf("Unexpected outcome at Day 1")
	}

	for i < 5 {
		gildedrose.UpdateQuality(items)
		i++
	}

	if !reflect.DeepEqual(items, itemsExpectedDay5) {
		for idx := 0; idx < len(items); idx++ {
			if !reflect.DeepEqual(items[idx], itemsExpectedDay5[idx]) {
				t.Logf("Unexpected item at Day 5, expected %v, got %v", items[idx], itemsExpectedDay5[idx])
			}
		}
		t.Errorf("Unexpected outcome at Day 5")
	}

	for i < 10 {
		gildedrose.UpdateQuality(items)
		i++
	}

	if !reflect.DeepEqual(items, itemsExpectedDay10) {
		for idx := 0; idx < len(items); idx++ {
			if !reflect.DeepEqual(items[idx], itemsExpectedDay10[idx]) {
				t.Logf("Unexpected item at Day 10, expected %v, got %v", items[idx], itemsExpectedDay10[idx])
			}
		}
		t.Errorf("Unexpected outcome at Day 10")
	}
}
