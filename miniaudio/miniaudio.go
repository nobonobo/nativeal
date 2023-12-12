package miniaudio

/*
#cgo CFLAGS: -std=gnu99 -Wno-unused-result
#cgo ma_debug CFLAGS: -DMA_DEBUG_OUTPUT=1

#cgo linux,!android LDFLAGS: -ldl -lpthread -lm
#cgo linux,arm LDFLAGS: -latomic
#cgo openbsd LDFLAGS: -ldl -lpthread -lm
#cgo netbsd LDFLAGS: -ldl -lpthread -lm
#cgo freebsd LDFLAGS: -ldl -lpthread -lm
#cgo android LDFLAGS: -lm

#cgo !noasm,!arm,!arm64 CFLAGS: -msse2
#cgo !noasm,arm,arm64 CFLAGS: -mfpu=neon -mfloat-abi=hard
#cgo noasm CFLAGS: -DMA_NO_SSE2 -DMA_NO_AVX2 -DMA_NO_AVX512 -DMA_NO_NEON

#define MINIAUDIO_IMPLEMENTATION
#include "./miniaudio/miniaudio.h"
*/
import "C"
