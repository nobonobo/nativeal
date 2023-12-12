package nativeal

// OscillatorNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode
type OscillatorNode interface {
	AudioScheduledSourceNode

	Frequency() AudioParam
	Detune() AudioParam
	Type() OscillatorType
	SetType(oType OscillatorType)
}

// OscillatorType as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/OscillatorNode/type
type OscillatorType string

const (
	OscillatorTypeSine     OscillatorType = "sine"
	OscillatorTypeSquare   OscillatorType = "square"
	OscillatorTypeSawTooth OscillatorType = "sawtooth"
	OscillatorTypeTriangle OscillatorType = "triangle"
	OscillatorTypeCustom   OscillatorType = "custom"
)
