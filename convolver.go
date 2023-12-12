package nativeal

// ConvolverNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/ConvolverNode
type ConvolverNode interface {
	AudioNode

	Buffer() AudioBuffer
	SetBuffer(buffer AudioBuffer)
	Normalize() bool
	SetNormalize(normalize bool)
}
