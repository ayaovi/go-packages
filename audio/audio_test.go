package audio

import (
	"testing"
	"reflect"
)

func TestAudio(t *testing.T) {
	a0 := Audio {
		Data: []byte{1, 1, 1, 1},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Validate", func(t *testing.T) {
		err := a0.Validate()
		if err != nil {
			t.Errorf("Audio validation should Pass.")
		}
	})
  
  add0 := Audio {
		Data: []byte{2, 2, 2, 2},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  	t.Run("Addition without overboard", func(t *testing.T) {
		out, err := a0.Plus(&a0)
		if err != nil {
			t.Errorf("Audio addition should Pass.")
		}
		if !reflect.DeepEqual(add0, *out) {
			t.Errorf("Audio addition was incorrect, got: %d, want: %d.", *out, add0)
		}
	})
  
  a1 := Audio {
		Data: []byte{150, 150, 150, 150},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  add1 := Audio {
		Data: []byte{255, 255, 255, 255},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Addition with overboard", func(t *testing.T) {
		out, err := a1.Plus(&a1)
		if err != nil {
			t.Errorf("Audio addition should Pass.")
		}
		if !reflect.DeepEqual(add1, *out) {
			t.Errorf("Audio addition was incorrect, got: %d, want: %d.", *out, add1)
		}
  })
  
  concat := Audio {
		Data: []byte{150, 150, 150, 150, 150, 150, 150, 150},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 8,
	}

  t.Run("Concat", func(t *testing.T) {
		out, err := a1.Concat(&a1)
		if err != nil {
			t.Errorf("Audio concatenation should Pass.")
		}
		if !reflect.DeepEqual(concat, *out) {
			t.Errorf("Audio addition was incorrect, got: %d, want: %d.", *out, concat)
		}
	})
}