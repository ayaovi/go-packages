package audio

import (
	"testing"
	"math"
	"fmt"
)

func TestAudioRms(t *testing.T) {
	fmt.Println("**** Running Audio Rms Tests ****")
	a_8M_0 := Audio {
		Data: []uint8{1, 2, 3, 4},
		Channel: 1,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

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

	a_8S_0 := Audio {
		Data: []Pair { Pair{uint8(1), uint8(1)}, Pair{uint8(2), uint8(2)}, 
									 Pair{uint8(3), uint8(3)}, Pair{uint8(4), uint8(4)} },
		Channel: 2,
		Size: 4,
		SamplingRate: 1,
		NumberOfSamples: 4,
		Length: 4,
	}

	rms_8S_1, rms_8S_2 := math.Sqrt(float64(5)), math.Sqrt(float64(5))

  t.Run("Audio8S rms", func(t *testing.T) {
		rms1, rms2, err := a_8S_0.Rms()
		if err != nil {
			t.Errorf("Audio8S rms should Pass.")
		}
		if rms_8S_1 == rms1 && rms_8S_2 == rms2 {
			t.Errorf("Audio8S rms was incorrect, got: (%f, %f) want: (%f, %f).", rms1, rms2, rms_8S_1, rms_8S_2)
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

	rms_16M := math.Sqrt(float64(5))

  t.Run("Audio16M rms", func(t *testing.T) {
		rms, _, err := a_16M_0.Rms()
		if err != nil {
			t.Errorf("Audio8M rms should Pass.")
		}
		if rms_8M == rms {
			t.Errorf("Audio8M rms was incorrect, got: %f, want: %f.", rms, rms_16M)
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

	rms_16S_1, rms_16S_2 := math.Sqrt(float64(5)), math.Sqrt(float64(5))

  t.Run("Audio16S rms", func(t *testing.T) {
		rms1, rms2, err := a_16S_0.Rms()
		if err != nil {
			t.Errorf("Audio16S rms should Pass.")
		}
		if rms_16S_1 == rms1 && rms_16S_2 == rms2 {
			t.Errorf("Audio16S rms was incorrect, got: (%f, %f) want: (%f, %f).", rms1, rms2, rms_16S_1, rms_16S_2)
		}
	})
}