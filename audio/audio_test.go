package audio

import (
	"testing"
	"reflect"
	"math"
)

func TestAudio(t *testing.T) {
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
		expected := &AudioError { Message: "incorrect channel 2 for Audio8M" }
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
		expected := &AudioError { Message: "incorrect channel 2 for Audio16M" }
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
		expected := &AudioError { Message: "incorrect channel 1 for Audio8S" }
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
		expected := &AudioError { Message: "incorrect channel 1 for Audio16S" }
		err := a.Validate()
		if !reflect.DeepEqual(err, expected) {
			t.Errorf("Audio8S validation error was incorrect, got: %s, want: %s.", err, expected)
		}
	})
  
  add_8M_0 := Audio {
		Data: []byte{2, 4, 6, 8},
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
		if !reflect.DeepEqual(add_8M_0, *out) {
			t.Errorf("Audio8M addition was incorrect, got: %d, want: %d.", *out, add_8M_0)
		}
	})

	add_8S_0 := Audio {
		Data: []Pair{ Pair{uint8(2), uint8(2)}, Pair{uint8(4), uint8(4)}, 
									Pair{uint8(6), uint8(6)}, Pair{uint8(8), uint8(8)}},
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
	t.Run("Audio8S addition without overboard", func(t *testing.T) {
		out, err := a_8S_0.Plus(&a_8S_0)
		if err != nil {
			t.Errorf("Audio8S addition should Pass.")
		}
		if !reflect.DeepEqual(add_8S_0, *out) {
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
		if !reflect.DeepEqual(add_8M_1, *out) {
			t.Errorf("Audio8M addition was incorrect, got: %d, want: %d.", *out, add_8M_1)
		}
  })
  
  concat_8M := Audio {
		Data: []uint8{150, 150, 150, 150, 150, 150, 150, 150},
		Channel: 1,
		Size: 8,
		SamplingRate: 1,
		NumberOfSamples: 8,
		Length: 8,
	}

  t.Run("Audio8M concat", func(t *testing.T) {
		out, err := a_8M_1.Concat(&a_8M_1)
		if err != nil {
			t.Errorf("Audio8M concatenation should Pass.")
		}
		if !reflect.DeepEqual(concat_8M, *out) {
			t.Errorf("Audio8M concatenation was incorrect, got: %d, want: %d.", *out, concat_8M)
		}
	})
	
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
		out, err := a_8S_0.Concat(&a_8S_0)
		if err != nil {
			t.Errorf("Audio8S concatenation should Pass.")
		}
		if !reflect.DeepEqual(concat_8S, *out) {
			t.Errorf("Audio8S concatenation was incorrect, got: %d, want: %d.", *out, concat_8S)
		}
  })

  rev_8M := Audio {
		Data: []uint8{4, 3, 2, 1},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  t.Run("Audio8M reverse", func(t *testing.T) {
		out, err := a_8M_0.Reverse()
		if err != nil {
			t.Errorf("Audio8M reverse should Pass.")
		}
		if !reflect.DeepEqual(rev_8M, *out) {
			t.Errorf("Audio8M reverse was incorrect, got: %d, want: %d.", *out, rev_8M)
		}
	})
	
	rev_8S := Audio {
		Data: []Pair{ Pair{uint8(4), uint8(4)}, Pair{uint8(3), uint8(3)}, 
									Pair{uint8(2), uint8(2)}, Pair{uint8(1), uint8(1)}},
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}
  
  t.Run("Audio8S reverse", func(t *testing.T) {
		out, err := a_8S_0.Reverse()
		if err != nil {
			t.Errorf("Audio8S reverse should Pass.")
		}
		if !reflect.DeepEqual(rev_8S, *out) {
			t.Errorf("Audio8S reverse was incorrect, got: %d, want: %d.", *out, rev_8S)
		}
  })
  
  cut_8M := Audio {
		Data: []uint8{2, 3, 4},
		Channel: 1,
		Size: 3,
		SamplingRate: 1,
		NumberOfSamples: 3,
		Length: 3,
	}

  t.Run("Audio8M cut", func(t *testing.T) {
		out, err := a_8M_0.Cut(int64(1), int64(3))
		if err != nil {
			t.Errorf("Audio8M cut should Pass.")
		}
		if !reflect.DeepEqual(cut_8M, *out) {
			t.Errorf("Audio8M cut was incorrect, got: %d, want: %d.", *out, cut_8M)
		}
	})
	
  cut_8S := Audio {
		Data: []Pair{Pair{uint8(2), uint8(2)}, Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)}},
		Channel: 2,
		Size: 3,
		SamplingRate: 1,
		NumberOfSamples: 3,
		Length: 3,
	}

  t.Run("Audio8S cut", func(t *testing.T) {
		out, err := a_8S_0.Cut(int64(1), int64(3))
		if err != nil {
			t.Errorf("Audio8S cut should Pass.")
		}
		if !reflect.DeepEqual(cut_8S, *out) {
			t.Errorf("Audio8S cut was incorrect, got: %d, want: %d.", *out, cut_8S)
		}
  })
  
  amp_8M := Audio {
		Data: []uint8{0, 1, 1, 2},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

  t.Run("Audio8M amplify", func(t *testing.T) {
		out, err := a_8M_0.Amplify(Volume{ C1: 0.5, C2: 0.5 })
		if err != nil {
			t.Errorf("Audio8M amplify should Pass.")
		}
		if !reflect.DeepEqual(amp_8M, *out) {
			t.Errorf("Audio8M amplify was incorrect, got: %d, want: %d.", *out, amp_8M)
		}
	})

	amp_8S := Audio {
		Data: []Pair{ Pair{uint8(0), uint8(1)}, Pair{uint8(1), uint8(2)}, 
									Pair{uint8(1), uint8(3)}, Pair{uint8(2), uint8(4)}},
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

  t.Run("Audio8S amplify", func(t *testing.T) {
		out, err := a_8S_0.Amplify(Volume{ C1: 0.5, C2: 1.0 })
		if err != nil {
			t.Errorf("Audio8S amplify should Pass.")
		}
		if !reflect.DeepEqual(amp_8S, *out) {
			t.Errorf("Audio8S amplify was incorrect, got: %d, want: %d.", *out, amp_8S)
		}
	})

	rms_8M := math.Sqrt(float64(5))

  t.Run("Audio8M rms", func(t *testing.T) {
		rms, _, err := a_8M_0.Rms()
		if err != nil {
			t.Errorf("Audio8M rms should Pass.")
		}
		if rms_8M == rms {
			t.Errorf("Audio8M rms was incorrect, got: %f, want: %f.", rms, rms_8M)
		}
	})

	rms_8S_1, rms_8S_2 := math.Sqrt(float64(5)), math.Sqrt(float64(5))

  t.Run("Audio8S rms", func(t *testing.T) {
		rms1, rms2, err := a_8S_0.Rms()
		if err != nil {
			t.Errorf("Audio8S amplify should Pass.")
		}
		if rms_8S_1 == rms1 && rms_8S_2 == rms2 {
			t.Errorf("Audio8S amplify was incorrect, got: (%f, %f) want: (%f, %f).", rms1, rms2, rms_8S_1, rms_8S_2)
		}
	})

	a_8M_2 := Audio {
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
		out, err := a_8M_2.Norm(1.0, 1.0)
		if err != nil {
			t.Errorf("Audio8M norm should Pass.")
		}
		if reflect.DeepEqual(*out, norm_8M) {
			t.Errorf("Audio8M norm was incorrect, got: %d want: %d.", *out, norm_8M)
		}
	})
}