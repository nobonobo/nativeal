package nativeal

type AudioNode interface {
	Context() AudioContext
	NumberOfInputs() uint
	NumberOfOutputs() uint

	ConnectNode(destination AudioNode)
	ConnectParam(destination AudioParam)
	Disconnect()
	DisconnectNode(destination AudioNode)
	DisconnectParam(destination AudioParam)
}
