package experiment

import (
	"regexp"
	"time"
)

var IsTimeStrRe = regexp.MustCompile(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)$`)
var IsTimeStr6Re = regexp.MustCompile(`^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d{1,6})?))(Z|[\+-]\d{2}:\d{2})?)$`)

const (
	minLen = 22
	maxLen = 34
)

// Validated quoted string
// "2015-05-14T12:34:56.379+02:00"
// "2015-05-14T12:34:56.379Z"
func IsTimeStrReFn(v []byte) bool {
	l := len(v)
	if l < minLen || l > maxLen {
		return false
	}
	if v[0] != '"' || v[l-1] != '"' {
		return false
	}
	return IsTimeStrRe.MatchString(string(v[1 : l-1]))
}

func IsTimeStrRe6Fn(v []byte) bool {
	l := len(v)
	if l < minLen || l > maxLen {
		return false
	}
	if v[0] != '"' || v[l-1] != '"' {
		return false
	}
	return IsTimeStr6Re.MatchString(string(v[1 : l-1]))
}

func IsTimeStrTime(v []byte) bool {
	l := len(v)
	if l < minLen || l > maxLen {
		return false
	}
	if v[0] != '"' || v[l-1] != '"' {
		return false
	}
	_, e := time.Parse(time.RFC3339, string(v[1:l-1]))
	//if e != nil {
	//	fmt.Println(string(v[1:l-1])+":", e)
	//}
	return e == nil
}

var numCh = [256]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

// Assuming time-secfrac is not more then 6 digits - μs
func IsTimeStr6Fn(v []byte) bool {
	l := len(v)
	timeZonePos := 20
	if l < minLen || l > maxLen {
		return false
	}
	if v[0] != '"' || v[l-1] != '"' {
		return false
	}
	for idx := 1; idx < l-1; idx++ {
		ch := v[idx]
		if (idx == 5 || idx == 8) && ch == '-' {
			continue
		} else if idx == 11 && ch == 'T' {
			continue
		} else if (idx == 14 || idx == 17) && ch == ':' {
			continue
		} else if idx == timeZonePos && ch == 'Z' {
			idx += 1
			ch = v[idx]
			if ch != '"' || idx != l-1 {
				return false
			}
			break
		} else if idx == 20 && ch == '.' {
			idx += 1
			ch = v[idx]
			if !numCh[ch] {
				return false
			}
			for i := 1; i < 7; i++ {
				idx += 1
				ch = v[idx]
				if ch == 'Z' || ch == '+' || ch == '-' {
					timeZonePos = idx
					idx -= 1
					break
				}
				if !numCh[ch] {
					return false
				}
			}
			timeZonePos = idx + 1
		} else if idx == timeZonePos && (ch == '+' || ch == '-') {
			idx += 1
			ch = v[idx]
			if !numCh[ch] {
				return false
			}
			idx += 1
			ch = v[idx]
			if !numCh[ch] {
				return false
			}
			idx += 1
			ch = v[idx]
			if ch != ':' {
				return false
			}
			idx += 1
			ch = v[idx]
			if !numCh[ch] {
				return false
			}
			idx += 1
			ch = v[idx]
			if !numCh[ch] {
				return false
			}
			idx += 1
			ch = v[idx]
			if ch != '"' || idx != l-1 {
				return false
			}
		} else if numCh[ch] {
			continue
		} else {
			return false
		}
	}
	return true
}

//func IsTimeStrFn(v []byte) bool {
//	l := len(v)
//
//	if l < minLen || l > maxLen {
//		return false
//	}
//	if v[0] != '"' || v[l-1] != '"' {
//		return false
//	}
//	if !numCh[v[1]] || !numCh[v[2]] || !numCh[v[3]] || !numCh[v[4]] ||
//		v[5] != '-' ||
//		!numCh[v[6]] || !numCh[v[7]] ||
//		v[8] != '-' ||
//		!numCh[v[9]] || !numCh[v[10]] ||
//		v[11] != 'T' ||
//		!numCh[v[12]] || !numCh[v[13]] ||
//		v[14] != ':' ||
//		!numCh[v[15]] || !numCh[v[16]] ||
//		v[17] != ':' ||
//		!numCh[v[18]] || !numCh[v[19]] {
//		return false
//	}
//	timeZonePos := 20
//	for idx := 20; idx < l-1; idx++ {
//		ch := v[idx]
//		if idx == timeZonePos && ch == 'Z' {
//			idx += 1
//			ch = v[idx]
//			if ch != '"' || idx != l-1 {
//				return false
//			}
//			break
//		} else if idx == 20 && ch == '.' {
//			idx += 1
//			ch = v[idx]
//			if !numCh[ch] {
//				return false
//			}
//			for {
//				idx += 1
//				ch = v[idx]
//				if ch == 'Z' || ch == '+' || ch == '-' {
//					timeZonePos = idx
//					idx -= 1
//					break
//				}
//				if !numCh[ch] {
//					return false
//				}
//			}
//			timeZonePos = idx + 1
//		} else if idx == timeZonePos && (ch == '+' || ch == '-') {
//			if !numCh[v[idx+1]] || !numCh[v[idx+2]] ||
//				v[idx+3] != ':' ||
//				!numCh[v[idx+4]] || !numCh[v[idx+5]] ||
//				v[idx+6] != '"' || idx+6 != l-1 {
//				return false
//			}
//		} else {
//			return false
//		}
//	}
//	return true
//}

func IsTimeStrHeadTailFn(v []byte) bool {
	l := len(v)
	if l < minLen || l > maxLen {
		return false
	}
	if v[0] != '"' || v[l-1] != '"' {
		return false
	}
	// head
	if !numCh[v[1]] || !numCh[v[2]] || !numCh[v[3]] || !numCh[v[4]] ||
		v[5] != '-' ||
		!numCh[v[6]] || !numCh[v[7]] ||
		v[8] != '-' ||
		!numCh[v[9]] || !numCh[v[10]] ||
		v[11] != 'T' ||
		!numCh[v[12]] || !numCh[v[13]] ||
		v[14] != ':' ||
		!numCh[v[15]] || !numCh[v[16]] ||
		v[17] != ':' ||
		!numCh[v[18]] || !numCh[v[19]] {
		return false
	}

	// tail
	tailIdx := l - 2
	if v[tailIdx] != 'Z' && !(numCh[v[tailIdx-0]] && numCh[v[tailIdx-1]] &&
		v[tailIdx-2] == ':' &&
		numCh[v[tailIdx-3]] && numCh[v[tailIdx-4]] &&
		(v[tailIdx-5] == '+' || v[tailIdx-5] == '-')) {
		return false
	}
	if tailIdx == 20 || tailIdx-5 == 20 {
		return true
	}
	// time-secfrac
	if v[20] == '.' {
		if !numCh[v[21]] {
			return false
		}
		if v[tailIdx] == 'Z' {
			tailIdx -= 1
		} else {
			tailIdx -= 6
		}
		for {
			if tailIdx <= 22 {
				break
			}
			if !numCh[v[tailIdx]] {
				return false
			}
			tailIdx -= 1
		}
	}
	return true
}
