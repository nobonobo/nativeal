package nativeal

type AudioDestinationNode interface {
	AudioNode

	MaxChannelCount() uint32
}
