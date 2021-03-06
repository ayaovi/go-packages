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
		if output == nil {
			t.Errorf("Audio8M fade-in was incorrect, got: <nil>, want: %d.", fadeIn_8M)
		}
		if output != nil && !reflect.DeepEqual(fadeIn_8M, *output) {
			t.Errorf("Audio8M fade-in was incorrect, got: %d, want: %d.", *output, fadeIn_8M)
		}
	})

	audio_8S := Audio {
		Data: []Pair { Pair{uint8(1), uint8(1)}, Pair{uint8(10), uint8(10)}, 
									 Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)},
									 Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)},
									 Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)},
									 Pair{uint8(10), uint8(10)}, Pair{uint8(10), uint8(10)},
									 Pair{uint8(100), uint8(100)} },
		Channel: 2,
		Size: 22,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeIn_8S := Audio {
		Data: []Pair { Pair{uint8(0), uint8(0)}, Pair{uint8(1), uint8(1)}, 
									 Pair{uint8(2), uint8(2)}, Pair{uint8(3), uint8(3)},
									 Pair{uint8(4), uint8(4)}, Pair{uint8(5), uint8(5)},
									 Pair{uint8(6), uint8(6)}, Pair{uint8(7), uint8(7)},
									 Pair{uint8(8), uint8(8)}, Pair{uint8(9), uint8(9)},
									 Pair{uint8(100), uint8(100)} },
		Channel: 2,
		Size: 22,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio8S fade-in", func(t *testing.T) {
		output, err := audio_8S.FadeIn(10)
		if err != nil {
			t.Errorf("Audio8S fade-in should Pass.")
		}
		if output == nil {
			t.Errorf("Audio8S fade-in was incorrect, got: <nil>, want: %d.", fadeIn_8S)
		}
		if output != nil && !reflect.DeepEqual(fadeIn_8S, *output) {
			t.Errorf("Audio8S fade-in was incorrect, got: %d, want: %d.", *output, fadeIn_8S)
		}
	})

	audio_16M := Audio {
		Data: []uint16{1, 10, 10, 10, 10, 10, 10, 10, 10, 10, 100},
		Channel: 1,
		Size: 22,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeIn_16M := Audio {
		Data: []uint16{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 100},
		Channel: 1,
		Size: 22,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio16M fade-in", func(t *testing.T) {
		output, err := audio_16M.FadeIn(10)
		if err != nil {
			t.Errorf("Audio16M fade-in should Pass.")
		}
		if output == nil {
			t.Errorf("Audio16M fade-in was incorrect, got: <nil>, want: %d.", fadeIn_16M)
		}
		if output != nil && !reflect.DeepEqual(fadeIn_16M, *output) {
			t.Errorf("Audio16M fade-in was incorrect, got: %d, want: %d.", *output, fadeIn_16M)
		}
	})

	audio_16S := Audio {
		Data: []Pair { Pair{uint16(1), uint16(1)}, Pair{uint16(10), uint16(10)}, 
									 Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)},
									 Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)},
									 Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)},
									 Pair{uint16(10), uint16(10)}, Pair{uint16(10), uint16(10)},
									 Pair{uint16(100), uint16(100)} },
		Channel: 2,
		Size: 44,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	fadeIn_16S := Audio {
		Data: []Pair { Pair{uint16(0), uint16(0)}, Pair{uint16(1), uint16(1)}, 
									 Pair{uint16(2), uint16(2)}, Pair{uint16(3), uint16(3)},
									 Pair{uint16(4), uint16(4)}, Pair{uint16(5), uint16(5)},
									 Pair{uint16(6), uint16(6)}, Pair{uint16(7), uint16(7)},
									 Pair{uint16(8), uint16(8)}, Pair{uint16(9), uint16(9)},
									 Pair{uint16(100), uint16(100)} },
		Channel: 2,
		Size: 44,
		SamplingRate: 1,
		NumberOfSamples: 11,
		Length: 11,
	}

	t.Run("Audio16S fade-in", func(t *testing.T) {
		output, err := audio_16S.FadeIn(10)
		if err != nil {
			t.Errorf("Audio16S fade-in should Pass.")
		}
		if output == nil {
			t.Errorf("Audio16S fade-in was incorrect, got: <nil>, want: %d.", fadeIn_16S)
		}
		if output != nil && !reflect.DeepEqual(fadeIn_16S, *output) {
			t.Errorf("Audio16S fade-in was incorrect, got: %d, want: %d.", *output, fadeIn_16S)
		}
	})
}