package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioCut(t *testing.T) {
	fmt.Println("**** Running Audio Cut Tests ****")
	a_8M := Audio {
		Data: []uint8{1, 2, 3, 4},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	cut_8M := Audio {
		Data: []uint8{2, 3, 4},
		Channel: 1,
		Size: 3,
		SamplingRate: 1,
		NumberOfSamples: 3,
		Length: 3,
	}

  t.Run("Audio8M cut", func(t *testing.T) {
		out, err := a_8M.Cut(int64(1), int64(3))
		if err != nil {
			t.Errorf("Audio8M cut should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8M cut was incorrect, got: <nil>, want: %d.", cut_8M)
		}
		if out != nil && !reflect.DeepEqual(cut_8M, *out) {
			t.Errorf("Audio8M cut was incorrect, got: %d, want: %d.", *out, cut_8M)
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
	
  cut_8S := Audio {
		Data: []Pair{Pair{uint8(2), uint8(2)}, Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)}},
		Channel: 2,
		Size: 6,
		SamplingRate: 1,
		NumberOfSamples: 3,
		Length: 3,
	}

  t.Run("Audio8S cut", func(t *testing.T) {
		out, err := a_8S.Cut(int64(1), int64(3))
		if err != nil {
			t.Errorf("Audio8S cut should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8S cut was incorrect, got: <nil>, want: %d.", cut_8S)
		}
		if out != nil && !reflect.DeepEqual(cut_8S, *out) {
			t.Errorf("Audio8S cut was incorrect, got: %d, want: %d.", *out, cut_8S)
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

	cut_16M := Audio {
		Data: []uint16{2, 3, 4},
		Channel: 1,
		Size: 6,
		SamplingRate: 1,
		NumberOfSamples: 3,
		Length: 3,
	}

  t.Run("Audio16M cut", func(t *testing.T) {
		out, err := a_16M.Cut(int64(1), int64(3))
		if err != nil {
			t.Errorf("Audio16M cut should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16M cut was incorrect, got: <nil>, want: %d.", cut_16M)
		}
		if out != nil && !reflect.DeepEqual(cut_16M, *out) {
			t.Errorf("Audio16M cut was incorrect, got: %d, want: %d.", *out, cut_16M)
		}
	})

	a_16S := Audio {
		Data: []Pair{ Pair{uint16(1), uint16(1)}, Pair{uint16(2), uint16(2)}, 
									Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)}},
		Channel: 2,
		Size: 16,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	cut_16S := Audio {
		Data: []Pair{Pair{uint16(2), uint16(2)}, Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)}},
		Channel: 2,
		Size: 12,
		SamplingRate: 1,
		NumberOfSamples: 3,
		Length: 3,
	}

  t.Run("Audio16S cut", func(t *testing.T) {
		out, err := a_16S.Cut(int64(1), int64(3))
		if err != nil {
			t.Errorf("Audio16S cut should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16S cut was incorrect, got: <nil>, want: %d.", cut_16S)
		}
		if out != nil && !reflect.DeepEqual(cut_16S, *out) {
			t.Errorf("Audio16S cut was incorrect, got: %d, want: %d.", *out, cut_16S)
		}
	})
}