package animation

import rl "github.com/gen2brain/raylib-go/raylib"

type Frame struct {
	texture  *rl.Texture2D
	src      rl.Rectangle
	duration float32
}

func NewFrame(texture *rl.Texture2D, src rl.Rectangle, duration float32) Frame {
	return Frame{
		texture:  texture,
		src:      src,
		duration: duration,
	}
}
