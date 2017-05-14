package ticker

import (
	"fmt"
	"time"
)

type SimpleWriter struct{}

func (s SimpleWriter) Write(storage Storage) {
	format := "%s: %6.2f    Active sources: (%d of %d);\t"
	result := ""
	for title, pair := range storage {
		value, active, all := getPairData(pair)

		result += fmt.Sprintf(format, title, value, active, all)
	}

	fmt.Printf("%s\n", result)
}

func getPairData(rates map[string]Rate) (value float64, active int, all int) {
	now := time.Now()
	for _, r := range rates {
		all++
		if r.ExpiredAt.After(now) {
			active++
			value += r.Value
		}
	}

	value = value / float64(active)

	return
}
