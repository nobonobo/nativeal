package nativeal

// AudioScheduledSourceNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/AudioScheduledSourceNode
type AudioScheduledSourceNode interface {
	AudioNode

	Start(when float64)
	Stop(when float64)
}
