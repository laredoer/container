package money

import (
	"testing"

	"github.com/shopspring/decimal"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRate(t *testing.T) {
	Init(func() (map[Currency]decimal.Decimal, error) {
		return map[Currency]decimal.Decimal{
			USD: decimal.NewFromInt(1),
		}, nil
	})

	Convey("TestGetRate", t, func() {
		Convey("USD rate should be 1", func() {
			rate := GetRate(USD)
			So(rate, ShouldEqual, decimal.NewFromInt(1))
		})

		Convey("HKD rate should be 7.8", func() {
			rate := GetRate(HKD)
			So(rate, ShouldEqual, decimal.NewFromFloat(7.8))
		})

		Convey("SGD rate should be 1.35", func() {
			rate := GetRate(SGD)
			So(rate, ShouldEqual, decimal.NewFromFloat(1.35))
		})

		Convey("CNY rate should be 7.3", func() {
			rate := GetRate(CNY)
			So(rate, ShouldEqual, decimal.NewFromFloat(7.3))
		})

		Convey("EUR rate should be 0", func() {
			rate := GetRate(EUR)
			So(rate.String(), ShouldEqual, decimal.Zero.String())
		})
	})

	Convey("Test Rate not exist", t, func() {
		Convey("Test EUR exchange", func() {
			amountHKD := New(decimal.NewFromInt(20), HKD)
			eur := amountHKD.To(EUR).Value.StringFixed(2)
			So(eur, ShouldEqual, decimal.RequireFromString("0.00").StringFixed(2))
		})
	})
}
