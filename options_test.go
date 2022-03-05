package factory_test

import (
	"fmt"
	"testing"

	"github.com/ipsusila/factory"
)

func TestOptions(t *testing.T) {
	op := factory.Options{
		"s":    "text",
		"b":    true,
		"i":    12341,
		"u":    10001,
		"f":    10.11,
		"ss":   []string{"item1", "item2"},
		"bs":   []bool{true, false},
		"is":   []int64{1, 3, 10},
		"fs":   []float64{10.00, 11.01, 12.02},
		"i16s": []int16{11, 12, 13, 14},
		"d":    "15m",
		"t":    "2022-01-02 15:14:00+07:00",
	}

	keys := []string{}
	for key := range op {
		keys = append(keys, key)
	}
	keys = append(keys, "-- doesn't exists --")

	for _, key := range keys {
		res := map[string]interface{}{
			"String":      op.String(key),
			"Bool":        op.Bool(key),
			"Int":         op.Int(key),
			"Uint":        op.Uint(key),
			"Float":       op.Float(key),
			"Duration":    op.Duration(key),
			"Time":        op.Time(key),
			"StringSlice": op.StringSlice(key),
			"IntSlice":    op.IntSlice(key),
			"FloatSlice":  op.FloatSlice(key),
			"BoolSlice":   op.BoolSlice(key),
		}

		fmt.Println("=== TEST KEY ", key, "===")
		for k, v := range res {
			fmt.Printf("  KEY:%s -> %+v\n", k, v)
		}
	}
}
