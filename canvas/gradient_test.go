package canvas

import (
	"image/color"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHorizontalGradient(t *testing.T) {
	horizontal := NewHorizontalGradient(color.Black, color.Transparent)

	img := horizontal.Generate(50, 5)
	assert.Equal(t, color.RGBA{0, 0, 0, 0xff}, img.At(0, 0))
	for i := 0; i < 5; i++ {
		assert.Equal(t, color.RGBA{0, 0, 0, 0x7f}, img.At(25, i))
	}
	assert.Equal(t, color.RGBA{0, 0, 0, 0x00}, img.At(50, 0))
}

func TestNewVerticalGradient(t *testing.T) {
	vertical := NewVerticalGradient(color.Black, color.Transparent)
	imgVert := vertical.Generate(5, 50)
	assert.Equal(t, color.RGBA{0, 0, 0, 0xff}, imgVert.At(0, 0))
	for i := 0; i < 5; i++ {
		assert.Equal(t, color.RGBA{0, 0, 0, 0x7f}, imgVert.At(i, 25))
	}
	assert.Equal(t, color.RGBA{0, 0, 0, 0x00}, imgVert.At(50, 0))
}

func TestNewRadialGradient(t *testing.T) {
	circle := NewRadialGradient(color.Black, color.Transparent)
	imgCircle := circle.Generate(10, 10)
	assert.Equal(t, color.RGBA{0, 0, 0, 0xff}, imgCircle.At(5, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0xcc}, imgCircle.At(4, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0x99}, imgCircle.At(3, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0x66}, imgCircle.At(2, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0x33}, imgCircle.At(1, 5))

	circle.CenterOffsetX = 0.1
	circle.CenterOffsetY = 0.1
	imgCircleOffset := circle.Generate(10, 10)
	assert.Equal(t, color.RGBA{0, 0, 0, 0xc3}, imgCircleOffset.At(5, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0xa0}, imgCircleOffset.At(4, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0x79}, imgCircleOffset.At(3, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0x50}, imgCircleOffset.At(2, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0x26}, imgCircleOffset.At(1, 5))

	circle.CenterOffsetX = -0.1
	circle.CenterOffsetY = -0.1
	imgCircleOffset = circle.Generate(10, 10)
	assert.Equal(t, color.RGBA{0, 0, 0, 0xc3}, imgCircleOffset.At(5, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0xd5}, imgCircleOffset.At(4, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0xc3}, imgCircleOffset.At(3, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0xa0}, imgCircleOffset.At(2, 5))
	assert.Equal(t, color.RGBA{0, 0, 0, 0x79}, imgCircleOffset.At(1, 5))
}
