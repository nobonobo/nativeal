package nativeal

// DelayNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/DelayNode
type DelayNode interface {
	AudioNode

	DelayTime() AudioParam
}
