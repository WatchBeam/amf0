package amf0

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getEncoded() []byte {
	s := NewString()
	s.SetBody("こんにちは")
	n := NewNumber()
	n.SetNumber(42)
	bytes := append([]byte{MARKER_STRING}, s.EncodeBytes()...)
	bytes = append(bytes, MARKER_NUMBER)
	bytes = append(bytes, n.EncodeBytes()...)
	return bytes
}

func TestDecoderComplete(t *testing.T) {
	bytes := getEncoded()

	r := &reluctantReader{src: bytes}
	kind, err := Decode(r)
	assert.Nil(t, err)
	assert.Equal(t, "こんにちは", kind.(*String).GetBody())
	kind, err = Decode(r)
	assert.Nil(t, err)
	assert.Equal(t, float64(42), kind.(*Number).GetNumber())
	_, err = Decode(r)
	assert.NotNil(t, err)
}

func TestDecodeFrom(t *testing.T) {
	bytes := getEncoded()

	kind, n, err := DecodeFrom(bytes, 0)
	assert.Nil(t, err)
	assert.Equal(t, "こんにちは", kind.(*String).GetBody())
	kind, n, err = DecodeFrom(bytes, n)
	assert.Nil(t, err)
	assert.Equal(t, float64(42), kind.(*Number).GetNumber())
}

func BenchmarkDecoder(b *testing.B) {
	s := NewBoolean()
	s.Set(true)
	bytes := append([]byte{MARKER_BOOLEAN}, s.EncodeBytes()...)

	for i := 0; i < b.N; i++ {
		DecodeFrom(bytes, 0)
	}
}