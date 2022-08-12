package helper

const (
	minTimeFormatLen = 22
	maxTimeFormatLen = 35
)

var SpaceCh = [256]bool{
	0x09: true, // tab
	0x0A: true, // line feed
	0x0D: true, // carriage return
	0x20: true, // space
}

var HexCh = [256]bool{
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
	'a': true,
	'b': true,
	'c': true,
	'd': true,
	'e': true,
	'f': true,
	'A': true,
	'B': true,
	'C': true,
	'D': true,
	'E': true,
	'F': true,
}

var NumCh = [256]bool{
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

var NullTokenData = []byte(`null`)
var FalseTokenData = []byte(`false`)
var TrueTokenData = []byte(`true`)

const (
	BackSlashCh      = 0x5c // '\'
	QuoteCh          = 0x22 // '"'
	TabCh            = 0x09 // '\t'
	NewLineCh        = 0x0a // '\n'
	CarriageReturnCh = 0x0d // '\r'
	OpenBraceCh      = 0x7b // '{'
	CloseBraceCh     = 0x7d // '}'
	OpenBracketCh    = 0x5b // '['
	CloseBracketCh   = 0x5d // ']'
	CommaCh          = 0x2c // ','
	PointCh          = 0x2e // '.'
	ExpCh            = 0x45 // 'E'
	ExpSmCh          = 0x65 // 'e'
	MinusCh          = 0x2d // '-'
)

const (
	// MaxUint is the maximum value of uint
	MaxUint    = ^uint64(0) //= 18446744073709551615
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
