package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioAddition(t *testing.T) {
	fmt.Println("**** Running Audio Addition Tests ****")
	a_8M_0 := Audio {
		Data: []uint8{1, 2, 3, 4},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	add_8M_0 := Audio {
		Data: []uint8{2, 4, 6, 8},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Audio8M addition without overboard", func(t *testing.T) {
		out, err := a_8M_0.Plus(&a_8M_0)
		if err != nil {
			t.Errorf("Audio8M addition should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8M addition was incorrect, got: <nil>, want: %d.", add_8M_0)
		}
		if out != nil && !reflect.DeepEqual(add_8M_0, *out) {
			t.Errorf("Audio8M addition was incorrect, got: %d, want: %d.", *out, add_8M_0)
		}
	})

	a_8S_0 := Audio {
		Data: []Pair { Pair{uint8(1), uint8(1)}, Pair{uint8(2), uint8(2)}, 
									 Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)} },
		Channel: 2,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
	add_8S_0 := Audio {
		Data: []Pair{ Pair{uint8(2), uint8(2)}, Pair{uint8(4), uint8(4)}, 
									Pair{uint8(6), uint8(6)}, Pair{uint8(8), uint8(8)}},
		Channel: 2,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Audio8S addition without overboard", func(t *testing.T) {
		out, err := a_8S_0.Plus(&a_8S_0)
		if err != nil {
			t.Errorf("Audio8S addition should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8S addition was incorrect, got: <nil>, want: %d.", add_8S_0)
		}
		if out != nil && !reflect.DeepEqual(add_8S_0, *out) {
			t.Errorf("Audio8S addition was incorrect, got: %d, want: %d.", *out, add_8S_0)
		}
	})
  
  a_8M_1 := Audio {
		Data: []uint8{150, 150, 150, 150},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  add_8M_1 := Audio {
		Data: []uint8{255, 255, 255, 255},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Audio8M addition with overboard", func(t *testing.T) {
		out, err := a_8M_1.Plus(&a_8M_1)
		if err != nil {
			t.Errorf("Audio8M addition should Pass.")
		}
		if out == nil {
			t.Errorf("Audio8M addition was incorrect, got: <nil>, want: %d.", add_8M_1)
		}
		if out != nil && !reflect.DeepEqual(add_8M_1, *out) {
			t.Errorf("Audio8M addition was incorrect, got: %d, want: %d.", *out, add_8M_1)
		}
	})
	
	a_16M_0 := Audio {
		Data: []uint16{1, 2, 3, 4},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	add_16M_0 := Audio {
		Data: []uint16{2, 4, 6, 8},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16M addition without overboard", func(t *testing.T) {
		out, err := a_16M_0.Plus(&a_16M_0)
		if err != nil {
			t.Errorf("Audio16M addition should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16M addition was incorrect, got: <nil>, want: %d.", add_16M_0)
		}
		if out != nil && !reflect.DeepEqual(add_16M_0, *out) {
			t.Errorf("Audio16M addition was incorrect, got: %d, want: %d.", *out, add_16M_0)
		}
	})

	a_16M_1 := Audio {
		Data: []uint16{32800, 32800, 32800, 32800},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
	
	add_16M_1 := Audio {
		Data: []uint16{65535, 65535, 65535, 65535},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16M addition with overboard", func(t *testing.T) {
		out, err := a_16M_1.Plus(&a_16M_1)
		if err != nil {
			t.Errorf("Audio16M addition should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16M addition was incorrect, got: <nil>, want: %d.", add_16M_1)
		}
		if out != nil && !reflect.DeepEqual(add_16M_1, *out) {
			t.Errorf("Audio16M addition was incorrect, got: %d, want: %d.", *out, add_16M_1)
		}
	})

	a_16S_0 := Audio {
		Data: []Pair { Pair{uint16(1), uint16(1)}, Pair{uint16(2), uint16(2)}, 
									 Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)} },
		Channel: 2,
		Size: 16,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	add_16S_0 := Audio {
		Data: []Pair{ Pair{uint16(2), uint16(2)}, Pair{uint16(4), uint16(4)}, 
									Pair{uint16(6), uint16(6)}, Pair{uint16(8), uint16(8)}},
		Channel: 2,
		Size: 16,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16S addition without overboard", func(t *testing.T) {
		out, err := a_16S_0.Plus(&a_16S_0)
		if err != nil {
			t.Errorf("Audio16S addition should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16S addition was incorrect, got: <nil>, want: %d.", add_16S_0)
		}
		if out != nil && !reflect.DeepEqual(add_16S_0, *out) {
			t.Errorf("Audio16S addition was incorrect, got: %d, want: %d.", *out, add_16S_0)
		}
	})

	a_16S_1 := Audio {
		Data: []Pair { Pair{uint16(32800), uint16(32800)}, Pair{uint16(32800), uint16(32800)}, 
									 Pair{uint16(32800), uint16(32800)}, Pair{uint16(32800), uint16(32800)} },
		Channel: 2,
		Size: 16,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	add_16S_1 := Audio {
		Data: []Pair{ Pair{uint16(65535), uint16(65535)}, Pair{uint16(65535), uint16(65535)}, 
									Pair{uint16(65535), uint16(65535)}, Pair{uint16(65535), uint16(65535)}},
		Channel: 2,
		Size: 16,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Audio16S addition with overboard", func(t *testing.T) {
		out, err := a_16S_1.Plus(&a_16S_1)
		if err != nil {
			t.Errorf("Audio16S addition should Pass.")
		}
		if out == nil {
			t.Errorf("Audio16S addition was incorrect, got: <nil>, want: %d.", add_16S_1)
		}
		if out != nil && !reflect.DeepEqual(add_16S_1, *out) {
			t.Errorf("Audio16S addition was incorrect, got: %d, want: %d.", *out, add_16S_1)
		}
	})
}