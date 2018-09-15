package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioReverse(t *testing.T) {
	fmt.Println("**** Running Audio Reverse Tests ****")
	a_8M := Audio {
		Data: []uint8{1, 2, 3, 4},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	rev_8M := Audio {
		Data: []uint8{4, 3, 2, 1},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  t.Run("Audio8M reverse", func(t *testing.T) {
		out, err := a_8M.Reverse()
		if err != nil {
			t.Errorf("Audio8M reverse should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8M reverse was incorrect, got: <nil>, want: %d.", rev_8M)
		}
		if out != nil && !reflect.DeepEqual(rev_8M, *out) {
			t.Errorf("Audio8M reverse was incorrect, got: %d, want: %d.", *out, rev_8M)
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

	rev_8S := Audio {
		Data: []Pair{ Pair{uint8(4), uint8(4)}, Pair{uint8(3), uint8(3)}, 
									Pair{uint8(2), uint8(2)}, Pair{uint8(1), uint8(1)}},
		Channel: 2,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  t.Run("Audio8S reverse", func(t *testing.T) {
		out, err := a_8S.Reverse()
		if err != nil {
			t.Errorf("Audio8S reverse should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8S reverse was incorrect, got: <nil>, want: %d.", rev_8S)
		}
		if out != nil && !reflect.DeepEqual(rev_8S, *out) {
			t.Errorf("Audio8S reverse was incorrect, got: %d, want: %d.", *out, rev_8S)
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

	rev_16M := Audio {
		Data: []uint16{4, 3, 2, 1},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16M reverse", func(t *testing.T) {
		out, err := a_16M.Reverse()
		if err != nil {
			t.Errorf("Audio16M reverse should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16M reverse was incorrect, got: <nil>, want: %d.", rev_16M)
		}
		if out != nil && !reflect.DeepEqual(rev_16M, *out) {
			t.Errorf("Audio16M reverse was incorrect, got: %d, want: %d.", *out, rev_16M)
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

	rev_16S := Audio {
		Data: []Pair{ Pair{uint16(4), uint16(4)}, Pair{uint16(3), uint16(3)}, 
									Pair{uint16(2), uint16(2)}, Pair{uint16(1), uint16(1)}},
		Channel: 2,
		Size: 16,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  t.Run("Audio16S reverse", func(t *testing.T) {
		out, err := a_16S.Reverse()
		if err != nil {
			t.Errorf("Audio16S reverse should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16S reverse was incorrect, got: <nil>, want: %d.", rev_16S)
		}
		if out != nil && !reflect.DeepEqual(rev_16S, *out) {
			t.Errorf("Audio16S reverse was incorrect, got: %d, want: %d.", *out, rev_16S)
		}
	})
}