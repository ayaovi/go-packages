package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioFadeOut(t *testing.T) {
	fmt.Println("**** Running Audio Fade-Out Tests ****")

	audio_8M := Audio {
		Data: []uint8{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100},
		Channel: 1,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeOut_8M := Audio {
		Data: []uint8{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 100},
		Channel: 1,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio8M fade-out", func(t *testing.T) {
		output, err := audio_8M.FadeOut(10)
		if err != nil {
			t.Errorf("Audio8M fade-out should Pass.")
		}
		if !reflect.DeepEqual(fadeOut_8M, *output) {
			t.Errorf("Audio8M fade-out was incorrect, got: %d, want: %d.", *output, fadeOut_8M)
		}
	})

	audio_16M := Audio {
		Data: []uint16{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100},
		Channel: 1,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeOut_16M := Audio {
		Data: []uint16{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 100},
		Channel: 1,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio16M fade-out", func(t *testing.T) {
		output, err := audio_16M.FadeOut(10)
		if err != nil {
			t.Errorf("Audio16M fade-out should Pass.")
		}
		if !reflect.DeepEqual(fadeOut_16M, *output) {
			t.Errorf("Audio16M fade-out was incorrect, got: %d, want: %d.", *output, fadeOut_16M)
		}
	})
}