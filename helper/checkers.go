package helper

func IsTimeFormat(v []byte) bool {
	l := len(v)
	if l < minTimeFormatLen || l > maxTimeFormatLen {
		return false
	}
	if v[0] != '"' || v[l-1] != '"' {
		return false
	}
	// head
	if !NumCh[v[1]] || !NumCh[v[2]] || !NumCh[v[3]] || !NumCh[v[4]] ||
		v[5] != '-' ||
		!NumCh[v[6]] || !NumCh[v[7]] ||
		v[8] != '-' ||
		!NumCh[v[9]] || !NumCh[v[10]] ||
		v[11] != 'T' ||
		!NumCh[v[12]] || !NumCh[v[13]] ||
		v[14] != ':' ||
		!NumCh[v[15]] || !NumCh[v[16]] ||
		v[17] != ':' ||
		!NumCh[v[18]] || !NumCh[v[19]] {
		return false
	}

	// tail
	tailIdx := l - 2
	if v[tailIdx] != 'Z' && !(NumCh[v[tailIdx-0]] && NumCh[v[tailIdx-1]] &&
		v[tailIdx-2] == ':' &&
		NumCh[v[tailIdx-3]] && NumCh[v[tailIdx-4]] &&
		(v[tailIdx-5] == '+' || v[tailIdx-5] == '-')) {
		return false
	}
	if tailIdx == 20 || tailIdx-5 == 20 {
		return true
	}
	// time-secfrac
	if v[20] == '.' {
		if !NumCh[v[21]] {
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
			if !NumCh[v[tailIdx]] {
				return false
			}
			tailIdx -= 1
		}
	}
	return true
}
