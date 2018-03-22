package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioConcat(t *testing.T) {
	fmt.Println("**** Running Audio Concat Tests ****")
	a_8M := Audio {
		Data: []uint8{150, 150, 150, 150},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
	
	concat_8M := Audio {
		Data: []uint8{150, 150, 150, 150, 150, 150, 150, 150},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 8,
		Length: 8,
	}

  t.Run("Audio8M concat", func(t *testing.T) {
		out, err := a_8M.Concat(&a_8M)
		if err != nil {
			t.Errorf("Audio8M concatenation should Pass.")
		}
		if !reflect.DeepEqual(concat_8M, *out) {
			t.Errorf("Audio8M concatenation was incorrect, got: %d, want: %d.", *out, concat_8M)
		}
	})
	
	a_8S := Audio {
		Data: []Pair { Pair{uint8(1), uint8(1)}, Pair{uint8(2), uint8(2)}, 
									 Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)} },
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

  concat_8S := Audio {
		Data: []Pair{ Pair{uint8(1), uint8(1)}, Pair{uint8(2), uint8(2)}, 
									Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)}, 
									Pair{uint8(1), uint8(1)}, Pair{uint8(2), uint8(2)}, 
									Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)}},
		Channel: 2,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 8,
		Length: 8,
	}

  t.Run("Audio8S concat", func(t *testing.T) {
		out, err := a_8S.Concat(&a_8S)
		if err != nil {
			t.Errorf("Audio8S concatenation should Pass.")
		}
		if !reflect.DeepEqual(concat_8S, *out) {
			t.Errorf("Audio8S concatenation was incorrect, got: %d, want: %d.", *out, concat_8S)
		}
	})
	
	a_16M := Audio {
		Data: []uint16{150, 150, 150, 150},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
	
	concat_16M := Audio {
		Data: []uint16{150, 150, 150, 150, 150, 150, 150, 150},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 8,
		Length: 8,
	}

  t.Run("Audio16M concat", func(t *testing.T) {
		out, err := a_16M.Concat(&a_16M)
		if err != nil {
			t.Errorf("Audio16M concatenation should Pass.")
		}
		if !reflect.DeepEqual(concat_16M, *out) {
			t.Errorf("Audio16M concatenation was incorrect, got: %d, want: %d.", *out, concat_16M)
		}
	})

	a_16S := Audio {
		Data: []Pair { Pair{uint16(1), uint16(1)}, Pair{uint16(2), uint16(2)}, 
									 Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)} },
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

  concat_16S := Audio {
		Data: []Pair{ Pair{uint16(1), uint16(1)}, Pair{uint16(2), uint16(2)}, 
									Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)}, 
									Pair{uint16(1), uint16(1)}, Pair{uint16(2), uint16(2)}, 
									Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)}},
		Channel: 2,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 8,
		Length: 8,
	}

  t.Run("Audio16S concat", func(t *testing.T) {
		out, err := a_16S.Concat(&a_16S)
		if err != nil {
			t.Errorf("Audio16S concatenation should Pass.")
		}
		if !reflect.DeepEqual(concat_16S, *out) {
			t.Errorf("Audio16S concatenation was incorrect, got: %d, want: %d.", *out, concat_16S)
		}
  })
}