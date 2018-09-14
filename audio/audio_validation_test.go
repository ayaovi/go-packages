package audio

import (
	"testing"
	"reflect"
	"fmt"
)

func TestAudioValidation(t *testing.T) {
	fmt.Println("**** Running Audio Validation Tests ****")
	a_8M_0 := Audio {
		Data: []uint8{1, 2, 3, 4},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Validate Audio8M ok", func(t *testing.T) {
		err := a_8M_0.Validate()
		if err != nil {
			t.Errorf("Audio validation should Pass.")
		}
	})

	t.Run("Validate Audio8M NOT ok", func(t *testing.T) {
		a := Audio {
			Data: []uint8{1, 2, 3, 4},
			Channel: 2,
		}
		expected := &AudioError { Message: "*** audio validation ***\nincorrect channel 2 for Audio8M" }
		err := a.Validate()
		if !reflect.DeepEqual(err, expected) {
			t.Errorf("Audio8M validation error was incorrect, got: %s, want: %s.", err, expected)
		}
	})

	a_16M_0 := Audio {
		Data: []uint16{1, 2, 3, 4},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	t.Run("Validate Audio16M ok", func(t *testing.T) {
		err := a_16M_0.Validate()
		if err != nil {
			t.Errorf("Audio validation should Pass.")
		}
	})

	t.Run("Validate Audio16M NOT ok", func(t *testing.T) {
		a := Audio {
			Data: []uint16{1, 2, 3, 4},
			Channel: 2,
		}
		expected := &AudioError { Message: "*** audio validation ***\nincorrect channel 2 for Audio16M" }
		err := a.Validate()
		if !reflect.DeepEqual(err, expected) {
			t.Errorf("Audio16M validation error was incorrect, got: %s, want: %s.", err, expected)
		}
	})

	a_8S_0 := Audio {
		Data: []Pair { Pair{uint8(1), uint8(1)}, Pair{uint8(2), uint8(2)}, 
									 Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)} },
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Validate Audio8S ok", func(t *testing.T) {
		err := a_8S_0.Validate()
		if err != nil {
			t.Errorf("Audio8S validation should Pass.")
		}
	})

	t.Run("Validate Audio8S NOT ok", func(t *testing.T) {
		a := Audio {
			Data: []Pair { Pair{uint8(1), uint8(1)}, Pair{uint8(2), uint8(2)}, 
										 Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)} },
			Channel: 1,
		}
		expected := &AudioError { Message: "*** audio validation ***\nincorrect channel 1 for Audio8S" }
		err := a.Validate()
		if !reflect.DeepEqual(err, expected) {
			t.Errorf("Audio8S validation error was incorrect, got: %s, want: %s.", err, expected)
		}
	})

	a_16S_0 := Audio {
		Data: []Pair { Pair{uint16(1), uint16(1)}, Pair{uint16(2), uint16(2)}, 
									 Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)} },
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Validate Audio16S ok", func(t *testing.T) {
		err := a_16S_0.Validate()
		if err != nil {
			t.Errorf("Audio16S validation should Pass.")
		}
	})

	t.Run("Validate Audio16S NOT ok", func(t *testing.T) {
		a := Audio {
			Data: []Pair { Pair{uint16(1), uint16(1)}, Pair{uint16(2), uint16(2)}, 
			Pair{uint16(3), uint16(3)}, Pair{uint16(4), uint16(4)} },
			Channel: 1,
		}
		expected := &AudioError { Message: "*** audio validation ***\nincorrect channel 1 for Audio16S" }
		err := a.Validate()
		if !reflect.DeepEqual(err, expected) {
			t.Errorf("Audio8S validation error was incorrect, got: %s, want: %s.", err, expected)
		}
	})
}