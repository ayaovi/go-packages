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

	audio_8S := Audio {
		Data: []Pair { Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)}, 
									 Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)},
									 Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)},
									 Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)},
									 Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)},
									 Pair{uint8(100), uint8(100)} },
		Channel: 2,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeOut_8S := Audio {
		Data: []Pair { Pair{uint8(10), uint8(10)}, Pair{uint8(9), uint8(9)}, 
									 Pair{uint8(8), uint8(8)}, Pair{uint8(7), uint8(7)},
									 Pair{uint8(6), uint8(6)}, Pair{uint8(5), uint8(5)},
									 Pair{uint8(4), uint8(4)}, Pair{uint8(3), uint8(3)},
									 Pair{uint8(2), uint8(2)}, Pair{uint8(1), uint8(1)},
									 Pair{uint8(100), uint8(100)} },
		Channel: 2,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio8S fade-out", func(t *testing.T) {
		output, err := audio_8S.FadeOut(10)
		if err != nil {
			t.Errorf("Audio8S fade-out should Pass.")
		}
		if !reflect.DeepEqual(fadeOut_8S, *output) {
			t.Errorf("Audio8S fade-out was incorrect, got: %d, want: %d.", *output, fadeOut_8S)
		}
	})

	audio_16S := Audio {
		Data: []Pair { Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)}, 
									 Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)},
									 Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)},
									 Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)},
									 Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)},
									 Pair{uint16(100), uint16(100)} },
		Channel: 2,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeOut_16S := Audio {
		Data: []Pair { Pair{uint16(10), uint16(10)}, Pair{uint16(9), uint16(9)}, 
									 Pair{uint16(8), uint16(8)}, Pair{uint16(7), uint16(7)},
									 Pair{uint16(6), uint16(6)}, Pair{uint16(5), uint16(5)},
									 Pair{uint16(4), uint16(4)}, Pair{uint16(3), uint16(3)},
									 Pair{uint16(2), uint16(2)}, Pair{uint16(1), uint16(1)},
									 Pair{uint16(100), uint16(100)} },
		Channel: 2,
		Size: 11,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio16S fade-out", func(t *testing.T) {
		output, err := audio_16S.FadeOut(10)
		if err != nil {
			t.Errorf("Audio16S fade-out should Pass.")
		}
		if !reflect.DeepEqual(fadeOut_16S, *output) {
			t.Errorf("Audio16S fade-out was incorrect, got: %d, want: %d.", *output, fadeOut_16S)
		}
	})
}