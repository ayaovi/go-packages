package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioAmplify(t *testing.T) {
	fmt.Println("**** Running Audio Amplify Tests ****")
	a_8M := Audio {
		Data: []uint8{1, 2, 3, 4},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	amp_8M := Audio {
		Data: []uint8{0, 1, 1, 2},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

  t.Run("Audio8M amplify", func(t *testing.T) {
		out, err := a_8M.Amplify(Volume{ C1: 0.5, C2: 0.5 })
		if err != nil {
			t.Errorf("Audio8M amplify should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8M amplify was incorrect, got: <nil>, want: %d.", amp_8M)
		}
		if out != nil && !reflect.DeepEqual(amp_8M, *out) {
			t.Errorf("Audio8M amplify was incorrect, got: %d, want: %d.", *out, amp_8M)
		}
	})

	a_8S := Audio {
		Data: []Pair { Pair{uint8(1), uint8(1)}, Pair{uint8(2), uint8(2)}, 
									 Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)} },
		Channel: 2,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	amp_8S := Audio {
		Data: []Pair{ Pair{uint8(0), uint8(1)}, Pair{uint8(1), uint8(2)}, 
									Pair{uint8(1), uint8(3)}, Pair{uint8(2), uint8(4)}},
		Channel: 2,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

  t.Run("Audio8S amplify", func(t *testing.T) {
		out, err := a_8S.Amplify(Volume{ C1: 0.5, C2: 1.0 })
		if err != nil {
			t.Errorf("Audio8S amplify should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8S amplify was incorrect, got: <nil>, want: %d.", amp_8S)
		}
		if out != nil && !reflect.DeepEqual(amp_8S, *out) {
			t.Errorf("Audio8S amplify was incorrect, got: %d, want: %d.", *out, amp_8S)
		}
	})

	a_16M := Audio {
		Data: []uint16{1, 2, 3, 4},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	amp_16M := Audio {
		Data: []uint16{0, 1, 1, 2},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16M amplify", func(t *testing.T) {
		out, err := a_16M.Amplify(Volume{ C1: 0.5, C2: 1.0 })
		if err != nil {
			t.Errorf("Audio16M amplify should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16M amplify was incorrect, got: <nil>, want: %d.", amp_16M)
		}
		if out != nil && !reflect.DeepEqual(amp_16M, *out) {
			t.Errorf("Audio16M amplify was incorrect, got: %d, want: %d.", *out, amp_16M)
		}
	})

	a_16S := Audio {
		Data: []Pair { Pair{uint16(1), uint16(1)}, Pair{uint16(2), uint16(2)}, 
									 Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)} },
		Channel: 2,
		Size: 16,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	amp_16S := Audio {
		Data: []Pair{ Pair{uint16(0), uint16(1)}, Pair{uint16(1), uint16(2)}, 
									Pair{uint16(1), uint16(3)}, Pair{uint16(2), uint16(4)}},
		Channel: 2,
		Size: 16,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16S amplify", func(t *testing.T) {
		out, err := a_16S.Amplify(Volume{ C1: 0.5, C2: 1.0 })
		if err != nil {
			t.Errorf("Audio16S amplify should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16S amplify was incorrect, got: <nil>, want: %d.", amp_16S)
		}
		if out != nil && !reflect.DeepEqual(amp_16S, *out) {
			t.Errorf("Audio16S amplify was incorrect, got: %d, want: %d.", *out, amp_16S)
		}
	})
}