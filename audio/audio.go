package audio

import (
)

type Audio struct {
	Data []byte
	Channel uint
	Size int64
	SamplingRate uint
	NumberOfSamples int64
	Length int64
}

type AudioError struct {
  Message string
}

func (ae *AudioError) Error() string {
	return ae.Message
}

func (a* Audio) Validate() (yn bool, err error) {
	return true , &AudioError { Message: "no errors."}
}

func (a1* Audio) Plus(a2* Audio) (a3 *Audio, err error) {
	// validate
	return nil,
	&AudioError { Message: "no error." }
}