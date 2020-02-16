package tcpros

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/aler9/goroslib/msgs"
)

type Parent struct {
	A msgs.String
}

var casesMessage = []struct {
	name string
	msg  interface{}
	byts []byte
}{
	{
		"empty",
		&struct{}{},
		[]byte{0x00, 0x00, 0x00, 0x00},
	},
	{
		"base types",
		&struct {
			A msgs.Bool
			B msgs.Byte
			C msgs.Char
			D msgs.Int8
			E msgs.Uint8
			F msgs.Int16
			G msgs.Uint16
			H msgs.Int32
			I msgs.Uint32
			J msgs.Int64
			K msgs.Uint64
			L msgs.Float32
			M msgs.Float64
			N msgs.String
			O msgs.Time
			P msgs.Duration
		}{
			true, 15, 'a', -1, 2, -3, 4, -5, 6, -7, 8, 9, 10, "abc",
			time.Date(2010, 11, 12, 13, 14, 15, 16, time.UTC),
			time.Duration(5 * time.Second),
		},
		[]byte{
			0x44, 0x00, 0x00, 0x00, 0x01, 0x0F, 0x61, 0xFF,
			0x02, 0xFD, 0xFF, 0x04, 0x00, 0xFB, 0xFF, 0xFF,
			0xFF, 0x06, 0x00, 0x00, 0x00, 0xF9, 0xFF, 0xFF,
			0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x08, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10,
			0x41, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x24,
			0x40, 0x03, 0x00, 0x00, 0x00, 0x61, 0x62, 0x63,
			0xa7, 0x3d, 0xdd, 0x4c, 0x10, 0x00, 0x00, 0x00,
			0x05, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	},
	{
		"empty string",
		&struct {
			A msgs.String
		}{
			"",
		},
		[]byte{
			0x04, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		},
	},
	{
		"empty time",
		&struct {
			A msgs.Time
		}{
			msgs.Time{},
		},
		[]byte{
			0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
		},
	},
	{
		"empty duration",
		&struct {
			A msgs.Duration
		}{
			msgs.Duration(0),
		},
		[]byte{
			0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00,
		},
	},
	{
		"variable array",
		&struct {
			A msgs.Uint8
			B []msgs.Uint32
		}{
			1, []msgs.Uint32{2, 3},
		},
		[]byte{
			0x0d, 0x00, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00,
			0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00,
			0x00,
		},
	},
	{
		"fixed array",
		&struct {
			A msgs.Uint8
			B [2]msgs.Uint32
		}{
			1, [2]msgs.Uint32{2, 3},
		},
		[]byte{
			0x09, 0x00, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00,
			0x00, 0x03, 0x00, 0x00, 0x00,
		},
	},
	{
		"variable array of parent",
		&struct {
			A msgs.Uint8
			B []Parent
		}{
			1, []Parent{{"abc"}, {"def"}},
		},
		[]byte{
			0x13, 0x00, 0x00, 0x00, 0x01, 0x02, 0x00, 0x00,
			0x00, 0x03, 0x00, 0x00, 0x00, 0x61, 0x62, 0x63,
			0x03, 0x00, 0x00, 0x00, 0x64, 0x65, 0x66,
		},
	},
	{
		"fixed array of parent",
		&struct {
			A msgs.Uint8
			B [2]Parent
		}{
			1, [2]Parent{{"abc"}, {"def"}},
		},
		[]byte{
			0x0f, 0x00, 0x00, 0x00, 0x01, 0x03, 0x00, 0x00,
			0x00, 0x61, 0x62, 0x63, 0x03, 0x00, 0x00, 0x00,
			0x64, 0x65, 0x66,
		},
	},
}

func TestMessageDecode(t *testing.T) {
	for _, c := range casesMessage {
		t.Run(c.name, func(t *testing.T) {
			msg := reflect.New(reflect.TypeOf(c.msg).Elem()).Interface()
			err := messageDecode(bytes.NewBuffer(c.byts), msg)
			require.NoError(t, err)
			require.Equal(t, c.msg, msg)
		})
	}
}
