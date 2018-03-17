package audio

import (
	"testing"
	"reflect"
)

func TestAudio(t *testing.T) {
	a0 := Audio {
		Data: []byte{1, 2, 3, 4},
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
		Data: []byte{2, 4, 6, 8},
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
			t.Errorf("Audio concatenation was incorrect, got: %d, want: %d.", *out, concat)
		}
  })

  reverse := Audio {
		Data: []byte{4, 3, 2, 1},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  t.Run("Reverse", func(t *testing.T) {
		out, err := a0.Reverse()
		if err != nil {
			t.Errorf("Audio reverse should Pass.")
		}
		if !reflect.DeepEqual(reverse, *out) {
			t.Errorf("Audio reverse was incorrect, got: %d, want: %d.", *out, reverse)
		}
  })
  
  cut := Audio {
		Data: []byte{2, 3, 4},
		Channel: 1,
		Size: 3,
		SamplingRate: 1,
		NumberOfSamples: 3,
		Length: 3,
	}

  t.Run("Cut", func(t *testing.T) {
		out, err := a0.Cut(int64(1), int64(3))
		if err != nil {
			t.Errorf("Audio cut should Pass.")
		}
		if !reflect.DeepEqual(cut, *out) {
			t.Errorf("Audio cut was incorrect, got: %d, want: %d.", *out, cut)
		}
  })
  
  amplify := Audio {
		Data: []byte{0, 1, 1, 2},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

  t.Run("Amplify", func(t *testing.T) {
		out, err := a0.Amplify(Volume{ C1: 0.5, C2: 0.5 })
		if err != nil {
			t.Errorf("Audio amplify should Pass.")
		}
		if !reflect.DeepEqual(amplify, *out) {
			t.Errorf("Audio amplify was incorrect, got: %d, want: %d.", *out, amplify)
		}
	})
}