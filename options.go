package factory

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"time"
)

// Supported time layout
var tmLayouts = []string{
	time.RFC3339,
	"2006-01-02 15:04:05Z07:00",
	"02/01/2006 15:04:05Z07:00",
	time.RFC3339Nano,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.ANSIC,
	time.Layout,
	time.RubyDate,
	time.UnixDate,
}

// Options for object construction
type Options map[string]interface{}

// Has return true if specified key exists,
func (o Options) Has(key string) bool {
	_, ok := o[key]
	return ok
}

// toString convert interface{} to string
func (o Options) toString(val interface{}, defV string) string {
	switch v := val.(type) {
	case string:
		return v
	case bool:
		return strconv.FormatBool(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case int:
		return strconv.Itoa(v)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case int32:
		return strconv.Itoa(int(v))
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case int16:
		return strconv.Itoa(int(v))
	case uint16:
		return strconv.Itoa(int(v))
	case int8:
		return strconv.Itoa(int(v))
	case uint8:
		return strconv.Itoa(int(v))
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case fmt.Stringer:
		return v.String()
	}
	return defV
}

// toBool convert value to boolean or default value
func (o Options) toBool(val interface{}, defV bool) bool {
	switch v := val.(type) {
	case bool:
		return v
	case int64:
		return v != 0
	case uint64:
		return v != 0
	case int:
		return v != 0
	case uint:
		return v != 0
	case int32:
		return v != 0
	case uint32:
		return v != 0
	case int16:
		return v != 0
	case uint16:
		return v != 0
	case int8:
		return v != 0
	case uint8:
		return v != 0
	case float64:
		return v != 0
	case float32:
		return v != 0
	case string:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return defV
		}
		return b
	case fmt.Stringer:
		b, err := strconv.ParseBool(v.String())
		if err != nil {
			return defV
		}
		return b
	}

	return defV
}

// convert interface val to 64-integer
func (o Options) toInt(val interface{}, defV int64) int64 {
	switch v := val.(type) {
	case int64:
		return v
	case uint64:
		if v > math.MaxInt64 {
			return defV
		}
		return int64(v)
	case int:
		return int64(v)
	case uint:
		if v > math.MaxInt64 {
			return defV
		}
		return int64(v)
	case int32:
		return int64(v)
	case uint32:
		return int64(v)
	case int16:
		return int64(v)
	case uint16:
		return int64(v)
	case int8:
		return int64(v)
	case uint8:
		return int64(v)
	case float32:
		iv := int64(v)
		if float32(iv) == v {
			return iv
		}
		return defV
	case float64:
		iv := int64(v)
		if float64(iv) == v {
			return iv
		}
		return defV
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		res, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return defV
		}
		return res
	case fmt.Stringer:
		res, err := strconv.ParseInt(v.String(), 10, 64)
		if err != nil {
			return defV
		}
		return res
	}

	return defV
}

// convert to unsigned integer
func (o Options) toUint(val interface{}, defV uint64) uint64 {
	switch v := val.(type) {
	case uint64:
		return v
	case int64:
		if v < 0 {
			return defV
		}
		return uint64(v)
	case int:
		if v < 0 {
			return defV
		}
		return uint64(v)
	case uint:
		return uint64(v)
	case int32:
		if v < 0 {
			return defV
		}
		return uint64(v)
	case uint32:
		return uint64(v)
	case int16:
		if v < 0 {
			return defV
		}
		return uint64(v)
	case uint16:
		return uint64(v)
	case int8:
		if v < 0 {
			return defV
		}
		return uint64(v)
	case uint8:
		return uint64(v)
	case float32:
		if v < 0 {
			return defV
		}

		// convertible if no fraction
		iv := uint64(v)
		if float32(iv) == v {
			return iv
		}
		return defV
	case float64:
		if v < 0 {
			return defV
		}

		iv := uint64(v)
		if float64(iv) == v {
			return iv
		}
		return defV
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		res, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return defV
		}
		return res
	case fmt.Stringer:
		res, err := strconv.ParseUint(v.String(), 10, 64)
		if err != nil {
			return defV
		}
		return res
	}

	return defV
}

func (o Options) toFloat(val interface{}, defV float64) float64 {
	// maximum integer that exactly
	// can be represented as float
	const maxI = int64(1) << 53
	const minI = -maxI

	switch v := val.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int64:
		if v > maxI || v < minI {
			return defV
		}
		return float64(v)
	case uint64:
		if v > uint64(maxI) {
			return defV
		}
		return float64(v)
	case int:
		iv := int64(v)
		if iv > maxI || iv < minI {
			return defV
		}
		return float64(v)
	case uint:
		uv := uint64(v)
		if uv > uint64(maxI) {
			return defV
		}
		return float64(v)
	case int32:
		return float64(v)
	case uint32:
		return float64(v)
	case int16:
		return float64(v)
	case uint16:
		return float64(v)
	case int8:
		return float64(v)
	case uint8:
		return float64(v)
	case bool:
		if v {
			return 1
		}
		return 0
	case string:
		fv, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return defV
		}
		return fv
	case fmt.Stringer:
		fv, err := strconv.ParseFloat(v.String(), 64)
		if err != nil {
			return defV
		}
		return fv
	}

	return defV
}

// convert to duration
func (o Options) toDuration(val interface{}, defV time.Duration) time.Duration {
	switch v := val.(type) {
	case time.Duration:
		return v
	case string:
		d, err := time.ParseDuration(v)
		if err != nil {
			return defV
		}
		return d
	case fmt.Stringer:
		d, err := time.ParseDuration(v.String())
		if err != nil {
			return defV
		}
		return d
	default:
		iDef := int64(defV)
		iv := o.toInt(val, iDef)
		return time.Duration(iv)
	}
}

func (o Options) parseTime(str string, defV time.Time) time.Time {
	for _, layout := range tmLayouts {
		if tm, err := time.Parse(layout, str); err == nil {
			return tm
		}
	}
	return defV
}

// toTime convert interface value to time.Time
// or default value if invalid/not specified.
func (o Options) toTime(val interface{}, defV time.Time) time.Time {
	switch v := val.(type) {
	case time.Time:
		return v
	case string:
		return o.parseTime(v, defV)
	case fmt.Stringer:
		return o.parseTime(v.String(), defV)
	default:
		// get from timestamp
		iDef := defV.Unix()
		iv := o.toInt(val, iDef)
		return time.Unix(iv, 0)
	}
}

// String return string value from options.
func (o Options) String(key string, def ...string) string {
	var defV string
	if len(def) > 0 {
		defV = def[0]
	}
	val, ok := o[key]
	if !ok || val == nil {
		return defV
	}

	return o.toString(val, defV)
}

// Bool return boolean value or default if specified
func (o Options) Bool(key string, def ...bool) bool {
	var defV bool
	if len(def) > 0 {
		defV = def[0]
	}

	val, ok := o[key]
	if !ok || val == nil {
		return defV
	}

	return o.toBool(val, defV)
}

// Int return integer value or default value if specified.
// If default value is not specified, 0 will be used as default value
func (o Options) Int(key string, def ...int64) int64 {
	var defV int64
	if len(def) > 0 {
		defV = def[0]
	}

	val, ok := o[key]
	if !ok || val == nil {
		return defV
	}

	return o.toInt(val, defV)
}

// Uint return integer value or default value if specified.
// If default value is not specified, 0 will be used as default value
func (o Options) Uint(key string, def ...uint64) uint64 {
	var defV uint64
	if len(def) > 0 {
		defV = def[0]
	}

	val, ok := o[key]
	if !ok || val == nil {
		return defV
	}

	return o.toUint(val, defV)
}

// Float return float64 value of default if specified
func (o Options) Float(key string, def ...float64) float64 {
	var defV float64
	if len(def) > 0 {
		defV = def[0]
	}

	val, ok := o[key]
	if !ok || val == nil {
		return defV
	}

	return o.toFloat(val, defV)
}

// Duration return time.Duration
func (o Options) Duration(key string, def ...time.Duration) time.Duration {
	var defV time.Duration
	if len(def) > 0 {
		defV = def[0]
	}

	val, ok := o[key]
	if !ok || val == nil {
		return defV
	}

	return o.toDuration(val, defV)
}

// Time return time.Time
func (o Options) Time(key string, def ...time.Time) time.Time {
	var defV time.Time
	if len(def) > 0 {
		defV = def[0]
	}

	val, ok := o[key]
	if !ok || val == nil {
		return defV
	}

	return o.toTime(val, defV)
}

// StringSlice returns slice of string or default value
func (o Options) StringSlice(key string, def ...string) []string {
	val, ok := o[key]
	if !ok || val == nil {
		return def
	}

	t := reflect.TypeOf(val)
	if t.Kind() != reflect.Slice {
		return def
	}

	switch t.Elem().Kind() {
	case reflect.String:
		return val.([]string)
	default:
		rv := reflect.ValueOf(val)
		n := rv.Len()
		if n == 0 || !rv.Index(0).CanInterface() {
			return nil
		}

		items := make([]string, 0, n)
		for i := 0; i < n; i++ {
			ev := rv.Index(i)
			items = append(items, o.toString(ev.Interface(), ""))
		}
		return items
	}
}

// FloatSlice return the value as given slice
func (o Options) FloatSlice(key string, def ...float64) []float64 {
	val, ok := o[key]
	if !ok || val == nil {
		return def
	}

	items := []float64{}
	switch v := val.(type) {
	case []float64:
		return v
	case []float32:
		for _, sv := range v {
			items = append(items, float64(sv))
		}
	case []int:
		for _, sv := range v {
			items = append(items, o.toFloat(sv, 0))
		}
	case []uint:
		for _, sv := range v {
			items = append(items, o.toFloat(sv, 0))
		}
	case []int64:
		for _, sv := range v {
			items = append(items, o.toFloat(sv, 0))
		}
	case []uint64:
		for _, sv := range v {
			items = append(items, o.toFloat(sv, 0))
		}
	case []int32:
		for _, sv := range v {
			items = append(items, float64(sv))
		}
	case []uint32:
		for _, sv := range v {
			items = append(items, float64(sv))
		}
	case []int16:
		for _, sv := range v {
			items = append(items, float64(sv))
		}
	case []uint16:
		for _, sv := range v {
			items = append(items, float64(sv))
		}
	case []int8:
		for _, sv := range v {
			items = append(items, float64(sv))
		}
	case []uint8:
		for _, sv := range v {
			items = append(items, float64(sv))
		}
	case []interface{}:
		for _, sv := range v {
			items = append(items, o.toFloat(sv, 0))
		}
	default:
		return def
	}

	return items
}

// IntSlice return the value as given slice
func (o Options) IntSlice(key string, def ...int64) []int64 {
	val, ok := o[key]
	if !ok || val == nil {
		return def
	}

	items := []int64{}
	switch v := val.(type) {
	case []int64:
		return v
	case []float64:
		for _, sv := range v {
			items = append(items, o.toInt(sv, 0))
		}
	case []float32:
		for _, sv := range v {
			items = append(items, o.toInt(sv, 0))
		}
	case []int:
		for _, sv := range v {
			items = append(items, int64(sv))
		}
	case []uint:
		for _, sv := range v {
			items = append(items, o.toInt(sv, 0))
		}
	case []uint64:
		for _, sv := range v {
			items = append(items, o.toInt(sv, 0))
		}
	case []int32:
		for _, sv := range v {
			items = append(items, int64(sv))
		}
	case []uint32:
		for _, sv := range v {
			items = append(items, int64(sv))
		}
	case []int16:
		for _, sv := range v {
			items = append(items, int64(sv))
		}
	case []uint16:
		for _, sv := range v {
			items = append(items, int64(sv))
		}
	case []int8:
		for _, sv := range v {
			items = append(items, int64(sv))
		}
	case []uint8:
		for _, sv := range v {
			items = append(items, int64(sv))
		}
	case []interface{}:
		for _, sv := range v {
			items = append(items, o.toInt(sv, 0))
		}
	default:
		return def
	}

	return items
}

// BoolSlice convert items into slice of boolean value.
func (o Options) BoolSlice(key string, def ...bool) []bool {
	val, ok := o[key]
	if !ok || val == nil {
		return def
	}

	items := []bool{}
	switch v := val.(type) {
	case []bool:
		return v
	case []string:
		for _, sv := range v {
			items = append(items, o.toBool(sv, false))
		}
	case []interface{}:
		for _, sv := range v {
			items = append(items, o.toBool(sv, false))
		}
	default:
		return def
	}

	return items
}
