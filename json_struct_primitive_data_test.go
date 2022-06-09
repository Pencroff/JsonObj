package JsonStruct

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNullOps(t *testing.T) {
	js := JsonStruct{}
	assert.Equal(t, true, js.IsNull())
	assert.Equal(t, "null", js.ToJson())
	js.SetInt(1)
	assert.Equal(t, false, js.IsNull())
	js.SetNull()
	assert.Equal(t, true, js.IsNull())

}

func TestIntOps(t *testing.T) {
	js := JsonStruct{}
	js.SetInt(1)
	assert.Equal(t, 1, js.Int())
	assert.Equal(t, true, js.IsNumber())
	assert.Equal(t, true, js.IsInt())
	assert.Equal(t, false, js.IsFloat())
	assert.Equal(t, "1", js.ToJson())
}

func TestFloatOps(t *testing.T) {
	js := JsonStruct{}
	js.SetFloat(3.1415926535897932385)
	assert.Equal(t, 3.141592653589793, js.Float())
	assert.Equal(t, true, js.IsNumber())
	assert.Equal(t, false, js.IsInt())
	assert.Equal(t, true, js.IsFloat())
	assert.Equal(t, "3.141592653589793", js.ToJson())
	js.SetFloat(3.1415)
	assert.Equal(t, "3.1415", js.ToJson())
}

func TestBoolOps(t *testing.T) {
	js := JsonStruct{}
	js.SetBool(true)
	assert.Equal(t, true, js.Bool())
	assert.Equal(t, true, js.IsBool())
	assert.Equal(t, "true", js.ToJson())
	js.SetBool(false)
	assert.Equal(t, false, js.Bool())
	assert.Equal(t, "false", js.ToJson())
}

func TestStringOps(t *testing.T) {
	js := JsonStruct{}
	js.SetString("hello")
	assert.Equal(t, "hello", js.String())
	assert.Equal(t, true, js.IsString())
	assert.Equal(t, `"hello"`, js.ToJson())
}

func TestTimeOps(t *testing.T) {
	js := JsonStruct{}
	tm, err := time.Parse(time.RFC3339, "2015-01-01T12:34:56Z")
	assert.NoError(t, err)
	js.SetTime(tm)
	assert.Equal(t, tm, js.Time())
	assert.Equal(t, true, js.IsTime())
	assert.Equal(t, `"2015-01-01T12:34:56Z"`, js.ToJson())
}
