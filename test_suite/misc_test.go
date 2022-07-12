package test_suite

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMisc_JsonString(t *testing.T) {
	v := make(map[string]string)
	v["a"] = "a\"z\\a/a\ba\fa\na\ra\ta||ṧⲟᵯȇ 𝙩𝒆𝘹𝙩"
	v["b"] = "\u0041\u0042\u0043\u0044\u0045\u0046\u0047\u0048\u0049\u004A\u004B\u004C\u004D\u004E\u004F\u0050\u0051\u0052\u0053\u0054\u0055\u0056\u0057\u0058\u0059\u005A"
	buf, _ := json.Marshal(v)
	for _, b := range buf {
		fmt.Printf("%c  ", b)
	}
	fmt.Println()
	for _, b := range buf {
		fmt.Printf("%x ", b)
	}
	fmt.Println()
	// {  "  a  "  :  "  a  \  "  z  \  \  a  /  a  \  u  0  0  0  8  a  \  u  0  0  0  c  a  \  n  a  \  r  a  \  t  a  |  |  á  ¹  §  â  ²    á  µ  ¯  È       ð      ©  ð        ð      ¹  ð      ©  "  ,  "  b  "  :  "  A  B  C  D  E  F  G  H  I  J  K  L  M  N  O  P  Q  R  S  T  U  V  W  X  Y  Z  "  }
	// 7b 22 61 22 3a 22 61 5c 22 7a 5c 5c 61 2f 61 5c 75 30 30 30 38 61 5c 75 30 30 30 63 61 5c 6e 61 5c 72 61 5c 74 61 7c 7c e1 b9 a7 e2 b2 9f   e1 b5 af c8 87   20 f0 9d   99    a9 f0 9d   92   86   f0 9d   98   b9 f0 9d   99    a9 22 2c 22 62 22 3a 22 41 42 43 44 45 46 47 48 49 4a 4b 4c 4d 4e 4f 50 51 52 53 54 55 56 57 58 59 5a 22 7d
}

func TestTimeNow(t *testing.T) {
	tm := time.Now()
	fmt.Println(tm.Format(time.RFC3339))
	// 2022-07-12T21:55:16+01:00
}
