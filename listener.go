package nativeal

type AudioListener interface {
	PositionX() AudioParam
	PositionY() AudioParam
	PositionZ() AudioParam
	ForwardX() AudioParam
	ForwardY() AudioParam
	ForwardZ() AudioParam
	UpX() AudioParam
	UpY() AudioParam
	UpZ() AudioParam
}
