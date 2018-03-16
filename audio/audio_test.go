package audio

import (
	"testing"
	"reflect"
)

func TestAuto(t *testing.T) {
	a := Audio {
		data: []uint8{1, 1, 1, 1},
		channel: 2,
		size: 2,
		samplingRate: 44,
		numberOfSamples: 10,
		length: 2,
	}
	// test validate resulting in true.
	t.Run("Validate", func(t *testing.T) {
		yn, err := a.Validate()
		expectedErr := &AudioError{ Message: "no errors." }
		if yn {
			t.Errorf("Audio validation should Pass.")
		}
		if !reflect.DeepEqual(err, expectedErr) {
			t.Errorf("Time error message was incorrect, got: %s, want: %s.", err, expectedErr)
		}
	})
}