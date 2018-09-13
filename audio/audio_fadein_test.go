package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioFadeIn(t *testing.T) {
	fmt.Println("**** Running Audio Fade-In Tests ****")

	audio_8M := Audio {
		Data: []uint8{1, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100},
		Channel: 1,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeIn_8M := Audio {
		Data: []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 100},
		Channel: 1,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio8M fade-in", func(t *testing.T) {
		output, err := audio_8M.FadeIn(10)
		if err != nil {
			t.Errorf("Audio8M fade-in should Pass.")
		}
		if !reflect.DeepEqual(fadeIn_8M, *output) {
			t.Errorf("Audio8M fade-in was incorrect, got: %d, want: %d.", *output, fadeIn_8M)
		}
	})

	audio_16M := Audio {
		Data: []uint16{1, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100},
		Channel: 1,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeIn_16M := Audio {
		Data: []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 100},
		Channel: 1,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio16M fade-in", func(t *testing.T) {
		output, err := audio_16M.FadeIn(10)
		if err != nil {
			t.Errorf("Audio16M fade-in should Pass.")
		}
		if !reflect.DeepEqual(fadeIn_16M, *output) {
			t.Errorf("Audio16M fade-in was incorrect, got: %d, want: %d.", *output, fadeIn_16M)
		}
	})
}