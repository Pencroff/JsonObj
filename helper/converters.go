package helper

import "math/bits"

const (
	// MaxUint is the maximum value of uint
	MaxUint    = ^uint64(0)
	MaxIntUint = MaxUint >> 1
	MinIntUint = MaxIntUint + 1
	// MaxInt is the maximum value of int
	MaxInt = int64(MaxIntUint)
	// MinInt is the minimum value of int
	MinInt = -MaxInt - 1
	// MinSafeInt is the minimum value of int that is safe to use in Float64
	// https://tc39.es/ecma262/#sec-number.min_safe_integer
	MinSafeInt = int64(-(1<<53 - 1))
	// MaxSafeInt is the maximum value of int that is safe to use in Float64
	// https://tc39.es/ecma262/#sec-number.max_safe_integer
	MaxSafeInt = int64(1<<53 - 1)
)

// FloatToInt - transforms float to the nearest integer
// decided not to change default behavior
func FloatToInt(val float64) int64 {
	if val < 0 {
		return int64(val - 0.5)
	}
	return int64(val + 0.5)
}

// StringToInt - gjson inspiration with optimizations for behavior similar to JavaScript JSON.parse
// https://github.com/tidwall/gjson/blob/e20a0bfa61962f4b430a2187ee5fbffe17f06856/gjson.go#L2583
func StringToInt(s string) (int64, bool) {
	var i int
	sign := int64(1)
	l := len(s)
	if l > 0 && s[0] == '-' {
		sign = -1
		i++
	}
	if i == l {
		return 0, false
	}
	carryOut := uint64(0)
	n := uint64(0)
	for ; i < l; i++ {
		character := s[i]
		if character == '.' {
			break
		}
		if character < '0' || character > '9' {
			return 0, false
		}
		carryOut, n = bits.Mul64(n, 10)
		if carryOut > 0 {
			break
		}
		n, carryOut = bits.Add64(n, uint64(character-'0'), carryOut)
		if carryOut > 0 {
			break
		}
	}
	if carryOut > 0 {
		if sign == -1 {
			return MinInt, false
		} else {
			return MaxInt, false
		}
	}
	if sign == 1 && n > MaxIntUint {
		return MaxInt, false
	}
	if n > MinIntUint {
		return MinInt, false
	}
	return int64(n) * sign, true
}

// StringToUint - gjson inspiration with optimizations for behavior similar to JavaScript JSON.parse
// https://github.com/tidwall/gjson/blob/e20a0bfa61962f4b430a2187ee5fbffe17f06856/gjson.go#L2568
func StringToUint(s string) (n uint64, ok bool) {
	var i int
	l := len(s)
	if i == l {
		return 0, false
	}
	carryOut := uint64(0)
	for ; i < l; i++ {
		character := s[i]
		if character == '.' {
			break
		}
		if character < '0' || character > '9' {
			return 0, false
		}
		carryOut, n = bits.Mul64(n, 10)
		if carryOut > 0 {
			return MaxUint, false
		}
		n, carryOut = bits.Add64(n, uint64(character-'0'), carryOut)
		if carryOut > 0 {
			return MaxUint, false
		}
	}

	return n, true
}
