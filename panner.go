package nativeal

// PannerNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/PannerNode
type PannerNode interface {
	AudioNode

	ConeInnerAngle() float64
	SetConeInnerAngle(angle float64)
	ConeOuterAngle() float64
	SetConeOuterAngle(angle float64)
	DistanceModel() PannerDistanceModel
	SetDistanceModel(model PannerDistanceModel)
	MaxDistance() float64
	SetMaxDistance(distance float64)
	OrientationX() AudioParam
	OrientationY() AudioParam
	OrientationZ() AudioParam
	PanningModel() PannerPanningModel
	SetPanningModel(model PannerPanningModel)
	PositionX() AudioParam
	PositionY() AudioParam
	PositionZ() AudioParam
	RefDistance() float64
	SetRefDistance(distance float64)
	RolloffFactor() float64
	SetRolloffFactor(factor float64)
}

type PannerDistanceModel string

const (
	PannerDistanceModelLinear      PannerDistanceModel = "linear"
	PannerDistanceModelInverse     PannerDistanceModel = "inverse"
	PannerDistanceModelExponential PannerDistanceModel = "exponential"
)

type PannerPanningModel string

const (
	PannerPanningModelEqualPower PannerPanningModel = "equalpower"
	PannerPanningModelHRTF       PannerPanningModel = "HRTF"
)

// StereoPannerNode as described here:
// https://developer.mozilla.org/en-US/docs/Web/API/StereoPannerNode
type StereoPannerNode interface {
	AudioNode
	Pan() AudioParam
}
