package audio

import (
  "fmt"
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

type Volume struct {
  C1 float32
  C2 float32
}

func (ae *AudioError) Error() string {
	return ae.Message
}

func (a* Audio) Validate() error {
	return nil
}

func clamp(value1 uint8, value2 uint8) uint8 {
  max := uint8(255)
  if value2 < (max  - value1) { return value1 + value2 }
  return max
}

func Compare(a1* Audio, a2* Audio) error {
  if a1.Size != a2.Size {
    return &AudioError { Message: fmt.Sprintf("mismatching audio size, %d and %d", a1.Size, a2.Size) }
  }
  if a1.Length != a2.Length {
    return &AudioError { Message: fmt.Sprintf("mismatching audio length, %d and %d", a1.Length, a2.Length) }
  }
  return nil
}

func (a1* Audio) Plus(a2* Audio) (a3 *Audio, err error) {
	// validate
  if err := Compare(a1, a2); err != nil {
    return nil, err
  }
  
  out := Audio {
    Data: make([]byte, a1.Size),
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }
  
  for i := int64(0); i < a1.Size; i++ {
    out.Data[i] = clamp(a1.Data[i], a2.Data[i])
  }
  
	return &out, nil
}

func (a1* Audio) Times(a2* Audio) (a3 *Audio, err error) {
  if err := Compare(a1, a2); err != nil {
    return nil, err
  }
  return nil, nil
}

func (a1* Audio) Concat(a2* Audio) (a3 *Audio, err error) {
  if err := Compare(a1, a2); err != nil {
    return nil, err
  }
  out := Audio {
    Data: make([]byte, a1.Size + a2.Size),
    Channel: a1.Channel,
    Size: a1.Size + a2.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length + a1.Length,
  }

  for i := int64(0); i < a1.Size; i++ {
    out.Data[i] = a1.Data[i]
  }

  for i := a1.Size; i < a1.Size + a2.Size; i++ {
    out.Data[i] = a1.Data[i - a1.Size]
  }

  return &out, nil
}

func (a1* Audio) Reverse() (a2 *Audio, err error) {
  out := Audio {
    Data: make([]byte, a1.Size),
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }

  for i := int64(0); i < a1.Size; i++ {
    out.Data[i] = a1.Data[a1.Size - i - 1]
  }

  return &out, nil
}