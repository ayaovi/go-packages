package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioNorm(t *testing.T) {
	fmt.Println("**** Running Audio Norm Tests ****")
	a_8M := Audio {
		Data: []uint8{0, 0, 0, 1},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	norm_8M := Audio {
		Data: []uint8{0, 0, 0, 2},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio8M norm", func(t *testing.T) {
		out, err := a_8M.Norm(1.0, 1.0)
		if err != nil {
			t.Errorf("Audio8M norm should Pass.")
		}
		if !reflect.DeepEqual(*out, norm_8M) {
			t.Errorf("Audio8M norm was incorrect, got: %d want: %d.", *out, norm_8M)
		}
	})

	a_8S := Audio {
		Data: []Pair{ Pair{uint8(0), uint8(0)}, Pair{uint8(0), uint8(0)}, 
									Pair{uint8(0), uint8(0)}, Pair{uint8(1), uint8(1)}},
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	norm_8S := Audio {
		Data: []Pair{ Pair{uint8(0), uint8(0)}, Pair{uint8(0), uint8(0)}, 
									Pair{uint8(0), uint8(0)}, Pair{uint8(2), uint8(2)}},
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio8S norm", func(t *testing.T) {
		out, err := a_8S.Norm(1.0, 1.0)
		if err != nil {
			t.Errorf("Audio8S norm should Pass.")
		}
		if !reflect.DeepEqual(*out, norm_8S) {
			t.Errorf("Audio8S norm was incorrect, got: %d want: %d.", *out, norm_8S)
		}
	})

	a_16M := Audio {
		Data: []uint16{0, 0, 0, 1},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	norm_16M := Audio {
		Data: []uint16{0, 0, 0, 2},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16M norm", func(t *testing.T) {
		out, err := a_16M.Norm(1.0, 1.0)
		if err != nil {
			t.Errorf("Audio16M norm should Pass.")
		}
		if !reflect.DeepEqual(*out, norm_16M) {
			t.Errorf("Audio16M norm was incorrect, got: %d want: %d.", *out, norm_16M)
		}
	})

	a_16S := Audio {
		Data: []Pair{ Pair{uint16(0), uint16(0)}, Pair{uint16(0), uint16(0)}, 
									Pair{uint16(0), uint16(0)}, Pair{uint16(1), uint16(1)}},
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	norm_16S := Audio {
		Data: []Pair{ Pair{uint16(0), uint16(0)}, Pair{uint16(0), uint16(0)}, 
									Pair{uint16(0), uint16(0)}, Pair{uint16(2), uint16(2)}},
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16S norm", func(t *testing.T) {
		out, err := a_16S.Norm(1.0, 1.0)
		if err != nil {
			t.Errorf("Audio16S norm should Pass.")
		}
		if !reflect.DeepEqual(*out, norm_16S) {
			t.Errorf("Audio16S norm was incorrect, got: %d want: %d.", *out, norm_16S)
		}
	})
}