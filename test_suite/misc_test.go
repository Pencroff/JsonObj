package test_suite

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMisc_JsonString(t *testing.T) {
	v := make(map[string]string)
	v["a"] = "a\"a\\a/a\ba\fa\na\ra\ta"
	buf, _ := json.Marshal(v)
	for _, b := range buf {
		fmt.Printf("%c  ", b)
	}
	fmt.Println()
	for _, b := range buf {
		fmt.Printf("%x ", b)
	}
	fmt.Println()
	//{  "  a  "  :  "  a  \  "  a  \  \  a  /  a  \  u  0  0  0  8  a  \  u  0  0  0  c  a  \  n  a  \  r  a  \  t  a  "  }
	//7b 22 61 22 3a 22 61 5c 22 61 5c 5c 61 2f 61 5c 75 30 30 30 38 61 5c 75 30 30 30 63 61 5c 6e 61 5c 72 61 5c 74 61 22 7d
}
