package helper

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
	var n uint64
	var sign bool
	l := len(s)
	if l > 0 && s[0] == '-' {
		sign = true
		i++
	}
	if i == l {
		return 0, false
	}
	if s[i] == '.' {
		return 0, false
	}
	for ; i < l; i++ {
		smb := s[i] - '0'
		if smb > 9 {
			if smb+'0' == '.' {
				break
			} else {
				return 0, false
			}
		}
		n1 := n*10 + uint64(smb)
		if n1 < n {
			if sign {
				return MinInt, false
			} else {
				return MaxInt, false
			}
		}
		n = n1
	}
	if !sign && n > MaxIntUint {
		return MaxInt, false
	}
	if n > MinIntUint {
		return MinInt, false
	}
	if sign {
		return -int64(n), true
	}
	return int64(n), true
}

// StringToUint - gjson inspiration with optimizations for behavior similar to JavaScript JSON.parse
// https://github.com/tidwall/gjson/blob/e20a0bfa61962f4b430a2187ee5fbffe17f06856/gjson.go#L2568
func StringToUint(s string) (n uint64, ok bool) {
	var i int
	l := len(s)
	if i == l {
		return 0, false
	}
	if s[i] == '.' {
		return 0, false
	}
	for ; i < l; i++ {
		smb := s[i] - '0'
		if smb > 9 {
			if smb+'0' == '.' {
				return n, true
			} else {
				return 0, false
			}
		}
		n1 := n*10 + uint64(smb)
		if n1 < n {
			return MaxUint, false
		}
		n = n1
	}
	return n, true
}
