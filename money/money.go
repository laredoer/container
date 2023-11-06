package money

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// 基准货币为美元

type rates struct {
	defaultRates map[Currency]decimal.Decimal // 默认汇率
	base         map[Currency]decimal.Decimal // 系统汇率
	fn           func() (map[Currency]decimal.Decimal, error)
}

var r = &rates{
	defaultRates: map[Currency]decimal.Decimal{
		USD: decimal.NewFromFloat(1),
		HKD: decimal.NewFromFloat(7.8),
		SGD: decimal.NewFromFloat(1.35),
		CNY: decimal.NewFromFloat(7.3),
	},
	base: map[Currency]decimal.Decimal{},
}

// Money 货币
type Money struct {
	Value    decimal.Decimal `json:"value"`
	Currency Currency        `json:"currency"`
}

// New 创建 Money
func New(value decimal.Decimal, currency Currency) Money {
	return Money{
		Value:    value.Truncate(4),
		Currency: currency,
	}
}

func (m Money) String() string {
	return fmt.Sprintf("%s %s", m.Value, m.Currency)
}

// Add 加法
func (m Money) Add(value Money) Money {

	var addVal decimal.Decimal
	// 判断货币是否相同，不相同将 value 转换为美元的货币，再转换为 m 货币
	if !m.SameCurrency(value) {
		addVal = toCurrency(toUSD(value), m.Currency).Value.Truncate(4)
	} else {
		addVal = value.Value
	}

	return Money{
		Value:    m.Value.Add(addVal).Truncate(4),
		Currency: m.Currency,
	}
}

// Div 除法
func (m Money) Div(value Money) Money {

	var divVal decimal.Decimal
	// 判断货币是否相同，不相同将 value 转换为美元的货币，再转换为 m 货币
	if !m.SameCurrency(value) {
		divVal = toCurrency(toUSD(value), m.Currency).Value.Truncate(4)
	} else {
		divVal = value.Value
	}

	m.Value = m.Value.Div(divVal).Truncate(4)
	return m
}

// GreaterThan 判断是否大于
func (m Money) GreaterThan(om Money) bool {

	if m.SameCurrency(om) {
		return m.Value.GreaterThan(om.Value)
	}

	mUSD := toUSD(m)
	omUSD := toUSD(om)
	return mUSD.Value.GreaterThan(omUSD.Value)
}

// LessThan 判断是否小于
func (m Money) LessThan(om Money) bool {

	if m.SameCurrency(om) {
		return m.Value.LessThan(om.Value)
	}

	mUSD := toUSD(m)
	omUSD := toUSD(om)
	return mUSD.Value.LessThan(omUSD.Value)
}

// SameCurrency 判断货币是否相同
func (m Money) SameCurrency(om Money) bool {
	return m.Currency == om.Currency
}

// To 转换货币
func (m Money) To(currency Currency) Money {
	if m.Currency == currency {
		return m
	}

	if GetRate(currency) == decimal.Zero {
		return m
	}

	return toCurrency(toUSD(m), currency)
}

func toUSD(in Money) Money {
	if in.Currency == USD {
		return in
	}
	return Money{
		Value:    in.Value.Div(GetRate(in.Currency)).Truncate(4),
		Currency: USD,
	}
}

func toCurrency(in Money, to Currency) Money {
	if in.Currency == to {
		return in
	}
	return Money{
		Value:    in.Value.Mul(GetRate(to)).Truncate(4),
		Currency: to,
	}
}

// GetRate 获取汇率，不存在获取默认汇率
func GetRate(currency Currency) decimal.Decimal {
	// 获取系统汇率
	rate, ok := r.base[currency]
	if !ok {
		rateMap, err := r.fn()
		if err != nil {
			// 获取默认汇率
			return r.defaultRates[currency]
		}

		for k, v := range rateMap {
			r.base[k] = v
		}

		if baseRate, ok := r.base[currency]; !ok || baseRate == decimal.Zero {
			return r.defaultRates[currency]
		}

		return r.base[currency]
	}
	return rate
}

func Init(f func() (map[Currency]decimal.Decimal, error)) {
	r.fn = f
}
