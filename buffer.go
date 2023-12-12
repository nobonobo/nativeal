package nativeal

import "github.com/mokiat/gog/opt"

type AudioBuffer interface {
	SampleRate() float64
	Length() uint
	Duration() float64
	NumberOfChannels() uint

	// TODO:
	// GetChannelData(channel uint) Float32Array
}

type AudioBufferSourceNode interface {
	AudioScheduledSourceNode

	Buffer() AudioBuffer
	SetBuffer(buffer AudioBuffer)
	Detune() AudioParam
	Loop() bool
	SetLoop(loop bool)
	LoopStart() float64
	SetLoopStart(start float64)
	LoopEnd() float64
	SetLoopEnd(end float64)
	PlaybackRate() AudioParam

	StartDetailed(when, offset float64, duration opt.T[float64])
}
