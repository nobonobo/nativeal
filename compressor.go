package nativeal

// DynamicsCompressorNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/DynamicsCompressorNode
type DynamicsCompressorNode interface {
	AudioNode

	Threshold() AudioParam
	Knee() AudioParam
	Ratio() AudioParam
	Reduction() float64
	Attack() AudioParam
	Release() AudioParam
}
