package audio

import (
  "fmt"
)

type Audio8M struct {
	Data []byte
	Channel uint
	Size int64
	SamplingRate uint
	NumberOfSamples int64
	Length int64
}

type Audio8S struct {
	Data []Pair8
	Channel uint
	Size int64
	SamplingRate uint
	NumberOfSamples int64
	Length int64
}

type Audio16M struct {
	Data []uint16
	Channel uint
	Size int64
	SamplingRate uint
	NumberOfSamples int64
	Length int64
}

type Audio16S struct {
	Data []Pair16
	Channel uint
	Size int64
	SamplingRate uint
	NumberOfSamples int64
	Length int64
}

type AudioError struct {
  Message string
}

type Pair8 struct {
  First uint8
  Second uint8
}

type Pair16 struct {
  First uint16
  Second uint16
}

type Volume struct {
  C1 float32
  C2 float32
}

func (ae *AudioError) Error() string {
	return ae.Message
}

func (a* Audio8M) Validate() error {
  if a.Channel != 1 { 
    return &AudioError { Message: fmt.Sprintf("incorrect channel %d for Audio8M", a.Channel) } 
  }
  return nil
}

func (a* Audio16M) Validate() error {
  if a.Channel != 1 { 
    return &AudioError { Message: fmt.Sprintf("incorrect channel %d for Audio16M", a.Channel) } 
  }
  return nil
}

func (a* Audio8S) Validate() error {
  if a.Channel != 2 { 
    return &AudioError { Message: fmt.Sprintf("incorrect channel %d for Audio8S", a.Channel) } 
  }
  return nil
}

func (a* Audio16S) Validate() error {
  if a.Channel != 2 { 
    return &AudioError { Message: fmt.Sprintf("incorrect channel %d for Audio16S", a.Channel) } 
  }
  return nil
}

func clamp8(value1 uint8, value2 uint8) uint8 {
  max := uint8(255)
  if value2 < (max - value1) { return value1 + value2 }
  return max
}

func clamp16(value1 uint16, value2 uint16) uint16 {
  max := uint16(65535)
  if value2 < (max - value1) { return value1 + value2 }
  return max
}

func Compare8M(a1* Audio8M, a2* Audio8M) error {
  if a1.Size != a2.Size {
    return &AudioError { Message: fmt.Sprintf("mismatching audio size, %d and %d", a1.Size, a2.Size) }
  }
  if a1.Length != a2.Length {
    return &AudioError { Message: fmt.Sprintf("mismatching audio length, %d and %d", a1.Length, a2.Length) }
  }
  return nil
}

func Compare8S(a1* Audio8S, a2* Audio8S) error {
  if a1.Size != a2.Size {
    return &AudioError { Message: fmt.Sprintf("mismatching audio size, %d and %d", a1.Size, a2.Size) }
  }
  if a1.Length != a2.Length {
    return &AudioError { Message: fmt.Sprintf("mismatching audio length, %d and %d", a1.Length, a2.Length) }
  }
  return nil
}

func (a1* Audio8M) Plus(a2* Audio8M) (a3 *Audio8M, err error) {
	// validate
  if err := Compare8M(a1, a2); err != nil {
    return nil, err
  }
  
  out := Audio8M {
    Data: make([]byte, a1.Size),
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }
  
  for i := int64(0); i < a1.Size; i++ {
    out.Data[i] = clamp8(a1.Data[i], a2.Data[i])
  }
  
	return &out, nil
}

func (a1* Audio8S) Plus(a2* Audio8S) (a3 *Audio8S, err error) {
	// validate
  if err := Compare8S(a1, a2); err != nil {
    return nil, err
  }
  
  out := Audio8S {
    Data: make([]Pair8, a1.Size),
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }
  
  for i := int64(0); i < a1.Size; i++ {
    out.Data[i].First = clamp8(a1.Data[i].First, a2.Data[i].First)
    out.Data[i].Second = clamp8(a1.Data[i].Second, a2.Data[i].Second)
  }
  
	return &out, nil
}


func (a1* Audio8M) Times(a2* Audio8M) (a3 *Audio8M, err error) {
  if err := Compare8M(a1, a2); err != nil {
    return nil, err
  }
  return nil, nil
}

func (a1* Audio8M) Concat(a2* Audio8M) (a3 *Audio8M, err error) {
  if err := Compare8M(a1, a2); err != nil {
    return nil, err
  }
  out := Audio8M {
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

func (a1* Audio8S) Concat(a2* Audio8S) (a3 *Audio8S, err error) {
  if err := Compare8S(a1, a2); err != nil {
    return nil, err
  }
  out := Audio8S {
    Data: make([]Pair8, a1.Size + a2.Size),
    Channel: a1.Channel,
    Size: a1.Size + a2.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length + a1.Length,
  }

  for i := int64(0); i < a1.Size; i++ {
    out.Data[i].First = a1.Data[i].First
    out.Data[i].Second = a1.Data[i].Second
  }

  for i := a1.Size; i < a1.Size + a2.Size; i++ {
    out.Data[i].First = a1.Data[i - a1.Size].First
    out.Data[i].Second = a1.Data[i - a1.Size].Second
  }

  return &out, nil
}

func (a1* Audio8M) Reverse() (a2 *Audio8M, err error) {
  out := Audio8M {
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

func (a1* Audio8S) Reverse() (a2 *Audio8S, err error) {
  out := Audio8S {
    Data: make([]Pair8, a1.Size),
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }

  for i := int64(0); i < a1.Size; i++ {
    out.Data[i].First = a1.Data[a1.Size - i - 1].First
    out.Data[i].Second = a1.Data[a1.Size - i - 1].Second
  }

  return &out, nil
}

func (a1* Audio8M) Cut(start int64, end int64) (a2 *Audio8M, err error) {
  if end >= a1.NumberOfSamples {
    return nil, &AudioError { Message: fmt.Sprintf("invaild audio range %d - %d\n", start, end) }
  }

  out := Audio8M {
    Data: make([]byte, end - start + 1),
    Channel: a1.Channel,
    Size: end - start + 1,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: end - start + 1,
    Length: (end - start + 1) / int64(a1.SamplingRate),
  }
  
  for i := start; i < end + 1; i++ {
    out.Data[i - start] = a1.Data[i]
  }

  return &out, nil
}

func (a1* Audio8S) Cut(start int64, end int64) (a2 *Audio8S, err error) {
  if end >= a1.NumberOfSamples {
    return nil, &AudioError { Message: fmt.Sprintf("invaild audio range %d - %d\n", start, end) }
  }

  out := Audio8S {
    Data: make([]Pair8, end - start + 1),
    Channel: a1.Channel,
    Size: end - start + 1,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: end - start + 1,
    Length: (end - start + 1) / int64(a1.SamplingRate),
  }
  
  for i := start; i < end + 1; i++ {
    out.Data[i - start].First = a1.Data[i].First
    out.Data[i - start].Second = a1.Data[i].Second
  }

  return &out, nil
}

func (a1* Audio8M) Amplify(vol Volume) (a2 *Audio8M, err error) {
  out := Audio8M {
    Data: make([]byte, a1.Size),
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }
  
  for i := int64(0); i < a1.Size; i++ {
    out.Data[i] = uint8(vol.C1 * float32(a1.Data[i]))
  }

  return &out, nil
}

func (a1* Audio8S) Amplify(vol Volume) (a2 *Audio8S, err error) {
  out := Audio8S {
    Data: make([]Pair8, a1.Size),
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }
  
  for i := int64(0); i < a1.Size; i++ {
    out.Data[i].First = uint8(vol.C1 * float32(a1.Data[i].First))
    out.Data[i].Second = uint8(vol.C2 * float32(a1.Data[i].Second))
  }

  return &out, nil
}