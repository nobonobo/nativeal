package nativeal

const (
	DefaultSampleRate = 44100
)

type AudioContextState string

const (
	AudioContextStateSuspended AudioContextState = "suspended"
	AudioContextStateRunning   AudioContextState = "running"
	AudioContextStateClosed    AudioContextState = "closed"
)

type AudioContext interface {
	CurrentTime() float64
	Destination() AudioDestinationNode
	Listener() AudioListener
	SampleRate() float64
	State() AudioContextState

	CreateBuffer(numChannels, length, sampleRate uint) AudioBuffer
	CreateBufferSource() AudioBufferSourceNode
	CreateConvolver() ConvolverNode
	CreateDelay() DelayNode
	CreateDynamicsCompressor() DynamicsCompressorNode
	CreateGain() GainNode
	CreateOscillator() OscillatorNode
	CreatePanner() PannerNode
	CreateStereoPanner() StereoPannerNode
	DecodeAudioData(data []byte) AudioBuffer // promised

	BaseLatency() float64
	OutputLatency() float64

	Close()   // promised
	Resume()  // promised
	Suspend() // promised
}

func NewAudioContext() AudioContext {
	return &audioContext{}
}

type audioContext struct {
}

func (c *audioContext) CurrentTime() float64 {
	return 0
}

func (c *audioContext) Destination() AudioDestinationNode {
	return nil
}

func (c *audioContext) Listener() AudioListener {
	return nil
}

func (c *audioContext) SampleRate() float64 {
	return DefaultSampleRate
}

func (c *audioContext) State() AudioContextState {
	return AudioContextStateSuspended
}

func (c *audioContext) CreateBuffer(numChannels, length, sampleRate uint) AudioBuffer {
	return nil
}

func (c *audioContext) CreateBufferSource() AudioBufferSourceNode {
	return nil
}

func (c *audioContext) CreateConvolver() ConvolverNode {
	return nil
}

func (c *audioContext) CreateDelay() DelayNode {
	return nil
}

func (c *audioContext) CreateDynamicsCompressor() DynamicsCompressorNode {
	return nil
}

func (c *audioContext) CreateGain() GainNode {
	return nil
}

func (c *audioContext) CreateOscillator() OscillatorNode {
	return nil
}

func (c *audioContext) CreatePanner() PannerNode {
	return nil
}

func (c *audioContext) CreateStereoPanner() StereoPannerNode {
	return nil
}

func (c *audioContext) DecodeAudioData(data []byte) AudioBuffer {
	return nil
}

func (c *audioContext) BaseLatency() float64 {
	return 0
}

func (c *audioContext) OutputLatency() float64 {
	return 0
}

func (c *audioContext) Close() {
}

func (c *audioContext) Resume() {
}

func (c *audioContext) Suspend() {
}
