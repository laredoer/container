package money

import (
	"fmt"
	"sync"
	"time"

	"git.5th.im/lb-public/gear/log"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// 基准货币为美元

var r *exchangeRates

type exchangeRates struct {
	Rates map[Currency]decimal.Decimal
	mu    sync.RWMutex
}

func (e *exchangeRates) UpdateRates(db *gorm.DB) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	type rate struct {
		CurrencyCode    Currency        `gorm:"column:currency_code"`
		UsdExchangeRate decimal.Decimal `gorm:"column:usd_exchange_rate"`
	}

	var rates []rate
	if err := db.Raw(
		"SELECT * FROM ads_exchange_rate_d WHERE dt = (SELECT MAX(dt) FROM ads_exchange_rate_d LIMIT 1)").
		Scan(&rates).Error; err != nil {
		return err
	}

	for _, v := range rates {
		e.Rates[v.CurrencyCode] = decimal.NewFromInt(1).Div(v.UsdExchangeRate)
	}

	log.Infof("[Success] Update exchange rates: %v", e.Rates)

	return nil
}

func NewExchangeRates(db *gorm.DB) {
	r = &exchangeRates{
		Rates: make(map[Currency]decimal.Decimal),
	}
	r.UpdateRates(db)

	go func(db *gorm.DB) {
		t := time.NewTicker(time.Minute * 10)
		for {
			<-t.C
			if err := r.UpdateRates(db); err != nil {
				continue
			}
		}
	}(db)
}

// only for test
func NewTestExchangeRates(testRates map[Currency]decimal.Decimal) {
	r = &exchangeRates{
		Rates: map[Currency]decimal.Decimal{
			AUD: decimal.NewFromFloat(1.50),
			CAD: decimal.NewFromFloat(1.37),
			CNY: decimal.NewFromFloat(7.31),
			EUR: decimal.NewFromFloat(0.91),
			GBP: decimal.NewFromFloat(0.77),
			HKD: decimal.NewFromFloat(7.82),
			IDR: decimal.NewFromFloat(16492.56),
			INR: decimal.NewFromFloat(84.03),
			JPY: decimal.NewFromFloat(157.29),
			KRW: decimal.NewFromFloat(1400.83),
			MYR: decimal.NewFromFloat(4.70),
			PHP: decimal.NewFromFloat(59.64),
			SGD: decimal.NewFromFloat(1.34),
			THB: decimal.NewFromFloat(36.35),
			USD: decimal.NewFromFloat(1),
			VND: decimal.NewFromFloat(25479.39),
		},
	}

	for k, v := range testRates {
		r.Rates[k] = v
	}

	log.Warnf("Test [Success] Update exchange rates: %v", r.Rates)
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

// GreaterThanOrEqual 判断是否大于
func (m Money) GreaterThanOrEqual(om Money) bool {
	if m.SameCurrency(om) {
		return m.Value.GreaterThanOrEqual(om.Value)
	}
	mUSD := toUSD(m)
	omUSD := toUSD(om)
	return mUSD.Value.GreaterThanOrEqual(omUSD.Value)
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
	r.mu.RLock()
	defer r.mu.RUnlock()

	if rate, ok := r.Rates[currency]; ok {
		return rate
	}
	return r.Rates[USD]
}

// 检查汇率是否存在
func CheckRate(currency Currency) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	_, ok := r.Rates[currency]
	return ok
}
