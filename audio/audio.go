package audio

import (
  "fmt"
  "math"
)

type Audio struct {
	Data interface{}
	Channel uint
	Size int64  /* file size in bytes. */
	SamplingRate uint /* in Hz. */
	NumberOfSamples int64 /* size / (sample_size * channel) */
	Length int64
}

type AudioError struct {
  Message string
}

type Pair struct {
  First interface{}
  Second interface{}
}

type Volume struct {
  C1 float32
  C2 float32
}

func (ae *AudioError) Error() string {
	return ae.Message
}

func (a* Audio) Validate() error {
  switch a.Data.(type) {
  case []uint8:
    if a.Channel != 1 { 
      return &AudioError { Message: fmt.Sprintf("incorrect channel %d for Audio8M", a.Channel) } 
    }
    break
  case []uint16:
    if a.Channel != 1 { 
      return &AudioError { Message: fmt.Sprintf("incorrect channel %d for Audio16M", a.Channel) } 
    }
    break
  case []Pair:
    switch a.Data.([]Pair)[0].First.(type){
      case uint8:
        if a.Channel != 2 { 
          return &AudioError { Message: fmt.Sprintf("incorrect channel %d for Audio8S", a.Channel) } 
        }
        break
      case uint16:
        if a.Channel != 2 { 
          return &AudioError { Message: fmt.Sprintf("incorrect channel %d for Audio16S", a.Channel) } 
        }
        break
    }
    break
  default:
    return nil
  }
  return nil
}

func clamp(value interface{}, size uint8) interface{} {
  switch size {
  case uint8(8):
    if value.(uint16) > uint16(255) {
      return uint8(255)
    }
    return uint8(value.(uint16))
  case uint8(16):
    if value.(uint32) > uint32(65535) {
      return uint16(65535)
    }
    return uint16(value.(uint32))
  }
  return nil
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
  a3 = &Audio {
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }
  switch a1.Data.(type) {
  case []uint8:
    a3.Data = make([]uint8, a1.Size)
    for i := int64(0); i < a1.NumberOfSamples; i++ {
      a3.Data.([]uint8)[i] = clamp(uint16(a1.Data.([]uint8)[i]) + uint16(a2.Data.([]uint8)[i]), 
      uint8(8)).(uint8)
    }
    break
  case []uint16:
    a3.Data = make([]uint16, a1.Size)
    for i := int64(0); i < a1.NumberOfSamples; i++ {
      a3.Data.([]uint16)[i] = clamp(uint32(a1.Data.([]uint16)[i]) + uint32(a2.Data.([]uint16)[i]), 
      uint8(16)).(uint16)
    }
    break
  case []Pair:
    a3.Data = make([]Pair, a1.Size)
    switch a1.Data.([]Pair)[0].First.(type){
      case uint8:
        for i := int64(0); i < a1.NumberOfSamples; i++ {
          a3.Data.([]Pair)[i].First = clamp(uint16(a1.Data.([]Pair)[i].First.(uint8)) + 
          uint16(a2.Data.([]Pair)[i].First.(uint8)), uint8(8)).(uint8)
          a3.Data.([]Pair)[i].Second = clamp(uint16(a1.Data.([]Pair)[i].Second.(uint8)) + 
          uint16(a2.Data.([]Pair)[i].Second.(uint8)), uint8(8)).(uint8)
        }
        break
      case uint16:
        for i := int64(0); i < a1.NumberOfSamples; i++ {
          a3.Data.([]Pair)[i].First = clamp(uint32(a1.Data.([]Pair)[i].First.(uint16)) + 
          uint32(a2.Data.([]Pair)[i].First.(uint16)), uint8(16)).(uint16)
          a3.Data.([]Pair)[i].Second = clamp(uint32(a1.Data.([]Pair)[i].Second.(uint16)) + 
          uint32(a2.Data.([]Pair)[i].Second.(uint16)), uint8(16)).(uint16)
        }
        break
    }
    break
  }
  
	return
}

func (input1* Audio) Concat(input2* Audio) (output *Audio, err error) {
  if err := Compare(input1, input2); err != nil {
    return nil, err
  }

  output = &Audio {
    Channel: input1.Channel,
    Size: input1.Size + input2.Size,
    SamplingRate: input1.SamplingRate,
    NumberOfSamples: input1.NumberOfSamples + input2.NumberOfSamples,
    Length: input1.Length + input2.Length,
  }
  switch input1.Data.(type) {
  case []uint8:
    output.Data = make([]uint8, input1.Size + input2.Size)
    // copy content of input1 into output.
    for i := int64(0); i < input1.NumberOfSamples; i++ {
      output.Data.([]uint8)[i] = input1.Data.([]uint8)[i]
    }
    // then append content of input2 to output.
    for i := input1.NumberOfSamples; i < output.NumberOfSamples; i++ {
      output.Data.([]uint8)[i] = input2.Data.([]uint8)[i - input1.NumberOfSamples]
    }
    break
  case []uint16:
    output.Data = make([]uint16, input1.Size + input2.Size)
    // copy content of input1 into output.
    for i := int64(0); i < input1.NumberOfSamples; i++ {
      output.Data.([]uint16)[i] = input1.Data.([]uint16)[i]
    }
    // then append content of input2 to output.
    for i := input1.NumberOfSamples; i < output.NumberOfSamples; i++ {
      output.Data.([]uint16)[i] = input2.Data.([]uint16)[i - input1.NumberOfSamples]
    }
    break
  case []Pair:
    output.Data = make([]Pair, input1.Size + input2.Size)
    // copy content of input1 into output.
    for i := int64(0); i < input1.Size; i++ {
      output.Data.([]Pair)[i].First = input1.Data.([]Pair)[i].First
      output.Data.([]Pair)[i].Second = input1.Data.([]Pair)[i].Second
    }
    // then append content of input2 to output.
    for i := input1.NumberOfSamples; i < output.NumberOfSamples; i++ {
      output.Data.([]Pair)[i].First = input2.Data.([]Pair)[i - input1.NumberOfSamples].First
      output.Data.([]Pair)[i].Second = input2.Data.([]Pair)[i - input1.NumberOfSamples].Second
    }
    break
  }

  return
}

func (input* Audio) Reverse() (output *Audio, err error) {
  output = &Audio {
    Channel: input.Channel,
    Size: input.Size,
    SamplingRate: input.SamplingRate,
    NumberOfSamples: input.NumberOfSamples,
    Length: input.Length,
  }
  switch input.Data.(type) {
  case []uint8:
    output.Data = make([]uint8, input.Size)
    for i := int64(0); i < input.NumberOfSamples; i++ {
      output.Data.([]uint8)[i] = input.Data.([]uint8)[input.NumberOfSamples - i - 1]
    }
    break
  case []uint16:
    output.Data = make([]uint16, input.Size)
    for i := int64(0); i < input.NumberOfSamples; i++ {
      output.Data.([]uint16)[i] = input.Data.([]uint16)[input.NumberOfSamples - i - 1]
    }
    break
  case []Pair:
    output.Data = make([]Pair, input.Size)
    for i := int64(0); i < input.NumberOfSamples; i++ {
      output.Data.([]Pair)[i].First = input.Data.([]Pair)[input.NumberOfSamples - i - 1].First
      output.Data.([]Pair)[i].Second = input.Data.([]Pair)[input.NumberOfSamples - i - 1].Second
    }
    break
  }

  return
}

func (input* Audio) Cut(start int64, end int64) (output* Audio, err error) {
  if end >= input.NumberOfSamples {
    return nil, &AudioError { Message: fmt.Sprintf("invaild audio range %d - %d\n", start, end) }
  }

  output = &Audio {
    Channel: input.Channel,
    Size: end - start + 1,
    SamplingRate: input.SamplingRate,
    NumberOfSamples: end - start + 1,
    Length: (end - start + 1) / int64(input.SamplingRate),
  }

  switch input.Data.(type) {
  case []uint8:
    output.Data = make([]uint8, end - start + 1)
    for i := start; i < end + 1; i++ {
      output.Data.([]uint8)[i - start] = input.Data.([]uint8)[i]
    }
    break
  case []uint16:
    output.Data = make([]uint16, end - start + 1)
    for i := start; i < end + 1; i++ {
      output.Data.([]uint16)[i - start] = input.Data.([]uint16)[i]
    }
    break
  case []Pair:
    output.Data = make([]Pair, end - start + 1)
    for i := start; i < end + 1; i++ {
      output.Data.([]Pair)[i - start].First = input.Data.([]Pair)[i].First
      output.Data.([]Pair)[i - start].Second = input.Data.([]Pair)[i].Second
    }
    break
  }
  return
}

func (input* Audio) Amplify(vol Volume) (output* Audio, err error) {
  output = &Audio {
    Channel: input.Channel,
    Size: input.Size,
    SamplingRate: input.SamplingRate,
    NumberOfSamples: input.NumberOfSamples,
    Length: input.Length,
  }

  switch input.Data.(type) {
  case []uint8:
    output.Data = make([]uint8, input.Size)
    for i := int64(0); i < input.NumberOfSamples; i++ {
      output.Data.([]uint8)[i] = uint8(vol.C1 * float32(input.Data.([]uint8)[i]))
    }
    break
  case []uint16:
    output.Data = make([]uint16, input.Size)
    for i := int64(0); i < input.NumberOfSamples; i++ {
      output.Data.([]uint16)[i] = uint16(vol.C1 * float32(input.Data.([]uint16)[i]))
    }
    break
  case []Pair:
    output.Data = make([]Pair, input.Size)
    switch input.Data.([]Pair)[0].First.(type) {
    case uint8:
      for i := int64(0); i < input.NumberOfSamples; i++ {
        output.Data.([]Pair)[i].First = uint8(vol.C1 * float32(input.Data.([]Pair)[i].First.(uint8)))
        output.Data.([]Pair)[i].Second = uint8(vol.C2 * float32(input.Data.([]Pair)[i].Second.(uint8)))
      }
      break
    case uint16:
      for i := int64(0); i < input.NumberOfSamples; i++ {
        output.Data.([]Pair)[i].First = uint16(vol.C1 * float32(input.Data.([]Pair)[i].First.(uint16)))
        output.Data.([]Pair)[i].Second = uint16(vol.C2 * float32(input.Data.([]Pair)[i].Second.(uint16)))
      }
    }
    break
  }

  return
}

func (input* Audio) Rms() (value1 float64, value2 float64 , err error) {
  sum1 := float64(0)
  sum2 := float64(0)
  switch input.Data.(type) {
  case []uint8:
    for _, v := range(input.Data.([]uint8)) {
      sum1 += float64(v * v)
    }
    value1 = math.Sqrt(sum1 / float64(input.NumberOfSamples))
    break
  case []uint16:
    for _, v := range(input.Data.([]uint16)) {
      sum1 += float64(v * v)
    }
    value1 = math.Sqrt(sum1 / float64(input.NumberOfSamples))
    break
  case []Pair:
    switch input.Data.([]Pair)[0].First.(type) {
    case uint8:
      for _, v := range(input.Data.([]Pair)) {
        sum1 += float64(v.First.(uint8) * v.First.(uint8))
        sum2 += float64(v.Second.(uint8) * v.Second.(uint8))
      }
      break
    case uint16:
      for _, v := range(input.Data.([]Pair)) {
        sum1 += float64(v.First.(uint16) * v.First.(uint16))
        sum2 += float64(v.Second.(uint16) * v.Second.(uint16))
      }
      break
    }
    value1 = math.Sqrt(sum1 / float64(input.NumberOfSamples))
    value2 = math.Sqrt(sum2 / float64(input.NumberOfSamples))
    break
  }
  
  return
}

func (input* Audio) Norm(rms_d1 float64, rms_d2 float64) (output* Audio, err error) {
  //validate
  rms_c1, rms_c2, err := input.Rms()
  
  if err != nil {
    return
  }
  
  output = &Audio {
    Channel: input.Channel,
    Size: input.Size,
    SamplingRate: input.SamplingRate,
    NumberOfSamples: input.NumberOfSamples,
    Length: input.Length,
  }
  switch input.Data.(type) {
  case []uint8:
    output.Data = make([]uint8, input.Size)
    for i := int64(0); i < input.NumberOfSamples; i++ {
      output.Data.([]uint8)[i] = clamp(uint16((rms_d1 * float64(input.Data.([]uint8)[i])) / rms_c1), uint8(8)).(uint8)
    }
    break
  case []uint16:
    output.Data = make([]uint16, input.Size)
    for i := int64(0); i < input.NumberOfSamples; i++ {
      output.Data.([]uint16)[i] = clamp(uint32((rms_d1 * float64(input.Data.([]uint16)[i])) / rms_c1), uint8(16)).(uint16)
    }
    break
  case []Pair:
    output.Data = make([]Pair, input.Size)
    switch input.Data.([]Pair)[0].First.(type) {
    case uint8:
      for i := int64(0); i < input.NumberOfSamples; i++ {
        output.Data.([]Pair)[i].First = clamp(uint16((rms_d1 * float64(input.Data.([]Pair)[i].First.(uint8))) / rms_c1), 
        uint8(8)).(uint8)
        output.Data.([]Pair)[i].Second = clamp(uint16((rms_d2 * float64(input.Data.([]Pair)[i].Second.(uint8))) / rms_c2), 
        uint8(8)).(uint8)
      }
      break
    case uint16:
      for i := int64(0); i < input.NumberOfSamples; i++ {
        output.Data.([]Pair)[i].First = clamp(uint32((rms_d1 * float64(input.Data.([]Pair)[i].First.(uint16))) / rms_c1), 
        uint8(16)).(uint16)
        output.Data.([]Pair)[i].Second = clamp(uint32((rms_d2 * float64(input.Data.([]Pair)[i].Second.(uint16))) / rms_c2), 
        uint8(16)).(uint16)
      }
      break
    }
  }
  
  return
}

func (input* Audio) FadeIn(second float64) (output* Audio, err error) {
  output = &Audio {
    Channel: input.Channel,
    Size: input.Size,
    SamplingRate: input.SamplingRate,
    NumberOfSamples: input.NumberOfSamples,
    Length: input.Length,
  }
  rampLength := int64(second * float64(input.SamplingRate))

  //check that the input audio is at leat as long as the fade-in second.
  if input.NumberOfSamples < rampLength {
    err = &AudioError { Message: fmt.Sprintf("input audio is too short") }
    return
  }

  switch input.Data.(type) {
  case []uint8:
    output.Data = make([]uint8, output.Size)
    for i := int64(0); i < rampLength; i++ {
      output.Data.([]uint8)[i] = uint8(float64(i) / float64(rampLength) * float64(input.Data.([]uint8)[i]))
    }
    for i := rampLength; i < output.NumberOfSamples; i++ {
      output.Data.([]uint8)[i] = input.Data.([]uint8)[i]
    }
    break
  case []uint16:
    output.Data = make([]uint16, output.Size)
    for i := int64(0); i < rampLength; i++ {
      output.Data.([]uint16)[i] = uint16(float64(i) / float64(rampLength) * float64(input.Data.([]uint16)[i]))
    }
    for i := rampLength; i < output.NumberOfSamples; i++ {
      output.Data.([]uint16)[i] = input.Data.([]uint16)[i]
    }
    break
  case []Pair:
    output.Data = make([]Pair, output.Size)
    switch input.Data.([]Pair)[0].First.(type) {
    case uint8:
      for i := int64(0); i < rampLength; i++ {
        output.Data.([]Pair)[i].First = uint8(float64(i) / float64(rampLength) * float64(input.Data.([]Pair)[i].First.(uint8)))
        output.Data.([]Pair)[i].Second = uint8(float64(i) / float64(rampLength) * float64(input.Data.([]Pair)[i].Second.(uint8)))
      }
      for i := rampLength; i < output.NumberOfSamples; i++ {
        output.Data.([]Pair)[i].First = input.Data.([]Pair)[i].First
        output.Data.([]Pair)[i].Second = input.Data.([]Pair)[i].Second
      }
      break
    case uint16:
      for i := int64(0); i < rampLength; i++ {
        output.Data.([]Pair)[i].First = uint16(float64(i) / float64(rampLength) * float64(input.Data.([]Pair)[i].First.(uint16)))
        output.Data.([]Pair)[i].Second = uint16(float64(i) / float64(rampLength) * float64(input.Data.([]Pair)[i].Second.(uint16)))
      }
      for i := rampLength; i < output.NumberOfSamples; i++ {
        output.Data.([]Pair)[i].First = input.Data.([]Pair)[i].First
        output.Data.([]Pair)[i].Second = input.Data.([]Pair)[i].Second
      }
      break
    }
  }
  
  return
}

func (input* Audio) FadeOut(second float64) (output* Audio, err error) {
  output = &Audio {
    Channel: input.Channel,
    Size: input.Size,
    SamplingRate: input.SamplingRate,
    NumberOfSamples: input.NumberOfSamples,
    Length: input.Length,
  }
  rampLength := int64(second * float64(input.SamplingRate))

  //check that the input audio is at leat as long as the fade-in second.
  if input.NumberOfSamples < rampLength {
    err = &AudioError { Message: fmt.Sprintf("input audio is too short") }
    return
  }

  switch input.Data.(type) {
  case []uint8:
    output.Data = make([]uint8, output.Size)
    for i := int64(0); i < rampLength; i++ {
      output.Data.([]uint8)[i] = uint8(math.Round((float64(1) - float64(i) / float64(rampLength)) * float64(input.Data.([]uint8)[i])))
    }
    for i := rampLength; i < output.NumberOfSamples; i++ {
      output.Data.([]uint8)[i] = input.Data.([]uint8)[i]
    }
    break
  case []uint16:
    output.Data = make([]uint16, output.Size)
    for i := int64(0); i < rampLength; i++ {
      output.Data.([]uint16)[i] = uint16(math.Round((float64(1) - float64(i) / float64(rampLength)) * float64(input.Data.([]uint16)[i])))
    }
    for i := rampLength; i < output.NumberOfSamples; i++ {
      output.Data.([]uint16)[i] = input.Data.([]uint16)[i]
    }
    break

  case []Pair:
    output.Data = make([]Pair, output.Size)
    switch input.Data.([]Pair)[0].First.(type) {
    case uint8:
      for i := int64(0); i < rampLength; i++ {
        output.Data.([]Pair)[i].First = uint8(math.Round((float64(1) - float64(i) / float64(rampLength)) * float64(input.Data.([]Pair)[i].First.(uint8))))
        output.Data.([]Pair)[i].Second = uint8(math.Round((float64(1) - float64(i) / float64(rampLength)) * float64(input.Data.([]Pair)[i].Second.(uint8))))
      }
      for i := rampLength; i < output.NumberOfSamples; i++ {
        output.Data.([]Pair)[i].First = input.Data.([]Pair)[i].First
        output.Data.([]Pair)[i].Second = input.Data.([]Pair)[i].Second
      }
      break
    case uint16:
      for i := int64(0); i < rampLength; i++ {
        output.Data.([]Pair)[i].First = uint16(math.Round((float64(1) - float64(i) / float64(rampLength)) * float64(input.Data.([]Pair)[i].First.(uint16))))
        output.Data.([]Pair)[i].Second = uint16(math.Round((float64(1) - float64(i) / float64(rampLength)) * float64(input.Data.([]Pair)[i].Second.(uint16))))
      }
      for i := rampLength; i < output.NumberOfSamples; i++ {
        output.Data.([]Pair)[i].First = input.Data.([]Pair)[i].First
        output.Data.([]Pair)[i].Second = input.Data.([]Pair)[i].Second
      }
      break
    }
  }

  return
}