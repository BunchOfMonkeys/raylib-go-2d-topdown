package animation

type Animation struct {
	timeElapsed  float32
	currentFrame int
	frames       []Frame
}

func NewAnimation(frames []Frame) *Animation {
	return &Animation{
		timeElapsed:  0.0,
		currentFrame: 0,
		frames:       frames,
	}
}

func (a *Animation) update(delta float32) {
	a.timeElapsed += delta
	if a.timeElapsed >= a.frames[a.currentFrame].duration {
		//instead of %
		for a.timeElapsed >= a.frames[a.currentFrame].duration {
			a.timeElapsed -= a.frames[a.currentFrame].duration
		}
		//next frame
		a.currentFrame = (a.currentFrame + 1) % len(a.frames)
	}
}
