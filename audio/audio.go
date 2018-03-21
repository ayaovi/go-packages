package audio

import (
  "fmt"
  "math"
)

type Audio struct {
	Data interface{}
	Channel uint
	Size int64
	SamplingRate uint
	NumberOfSamples int64
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
    for i := int64(0); i < a1.Size; i++ {
      a3.Data.([]uint8)[i] = clamp8(a1.Data.([]uint8)[i], a2.Data.([]uint8)[i])
    }
    break
  case []Pair:
    a3.Data = make([]Pair, a1.Size)
    switch a1.Data.([]Pair)[0].First.(type){
      case uint8:
        for i := int64(0); i < a1.Size; i++ {
          a3.Data.([]Pair)[i].First = clamp8(a1.Data.([]Pair)[i].First.(uint8), 
          a2.Data.([]Pair)[i].First.(uint8))
          a3.Data.([]Pair)[i].Second = clamp8(a1.Data.([]Pair)[i].Second.(uint8), 
          a2.Data.([]Pair)[i].Second.(uint8))
        }
        break
      case uint16:
        break
    }
    break
  }
  
	return
}

func (a1* Audio) Times(a2* Audio) (a3 *Audio, err error) {
  if err := Compare(a1, a2); err != nil {
    return nil, err
  }
  return
}


func (a1* Audio) Concat(a2* Audio) (a3 *Audio, err error) {
  if err := Compare(a1, a2); err != nil {
    return nil, err
  }

  a3 = &Audio {
    Channel: a1.Channel,
    Size: a1.Size + a2.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples + a2.NumberOfSamples,
    Length: a1.Length + a2.Length,
  }
  switch a1.Data.(type) {
  case []uint8:
    a3.Data = make([]uint8, a1.Size + a2.Size)
    // copy content of a1 into a3.
    for i := int64(0); i < a1.Size; i++ {
      a3.Data.([]uint8)[i] = a1.Data.([]uint8)[i]
    }
    // then append content od a2 to a3.
    for i := a1.Size; i < a1.Size + a2.Size; i++ {
      a3.Data.([]uint8)[i] = a1.Data.([]uint8)[i - a1.Size]
    }
    break
  case []Pair:
    a3.Data = make([]Pair, a1.Size + a2.Size)
    // copy content of a1 into a3.
    for i := int64(0); i < a1.Size; i++ {
      a3.Data.([]Pair)[i].First = a1.Data.([]Pair)[i].First
      a3.Data.([]Pair)[i].Second = a1.Data.([]Pair)[i].Second
    }
    // then append content od a2 to a3.
    for i := a1.Size; i < a1.Size + a2.Size; i++ {
      a3.Data.([]Pair)[i].First = a1.Data.([]Pair)[i - a1.Size].First
      a3.Data.([]Pair)[i].Second = a1.Data.([]Pair)[i - a1.Size].Second
    }
    break
  }

  return
}

func (a1* Audio) Reverse() (a2 *Audio, err error) {
  a2 = &Audio {
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }
  switch a1.Data.(type) {
  case []uint8:
    a2.Data = make([]uint8, a1.Size)
    for i := int64(0); i < a1.Size; i++ {
      a2.Data.([]uint8)[i] = a1.Data.([]uint8)[a1.Size - i - 1]
    }
    break
  case []Pair:
    a2.Data = make([]Pair, a1.Size)
    for i := int64(0); i < a1.Size; i++ {
      a2.Data.([]Pair)[i].First = a1.Data.([]Pair)[a1.Size - i - 1].First
      a2.Data.([]Pair)[i].Second = a1.Data.([]Pair)[a1.Size - i - 1].Second
    }
    break
  }

  return
}

func (a1* Audio) Cut(start int64, end int64) (a2 *Audio, err error) {
  if end >= a1.NumberOfSamples {
    return nil, &AudioError { Message: fmt.Sprintf("invaild audio range %d - %d\n", start, end) }
  }

  a2 = &Audio {
    Channel: a1.Channel,
    Size: end - start + 1,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: end - start + 1,
    Length: (end - start + 1) / int64(a1.SamplingRate),
  }

  switch a1.Data.(type) {
  case []uint8:
    a2.Data = make([]byte, end - start + 1)
    for i := start; i < end + 1; i++ {
      a2.Data.([]uint8)[i - start] = a1.Data.([]uint8)[i]
    }
    break
  case []Pair:
    a2.Data = make([]Pair, end - start + 1)
    for i := start; i < end + 1; i++ {
      a2.Data.([]Pair)[i - start].First = a1.Data.([]Pair)[i].First
      a2.Data.([]Pair)[i - start].Second = a1.Data.([]Pair)[i].Second
    }
    break
  }
  return
}

func (a1* Audio) Amplify(vol Volume) (a2 *Audio, err error) {
  a2 = &Audio {
    Channel: a1.Channel,
    Size: a1.Size,
    SamplingRate: a1.SamplingRate,
    NumberOfSamples: a1.NumberOfSamples,
    Length: a1.Length,
  }

  switch a1.Data.(type) {
  case []uint8:
    a2.Data = make([]uint8, a1.Size)
    for i := int64(0); i < a1.Size; i++ {
      a2.Data.([]uint8)[i] = uint8(vol.C1 * float32(a1.Data.([]uint8)[i]))
    }
    break
  case []Pair:
    a2.Data = make([]Pair, a1.Size)
    switch a1.Data.([]Pair)[0].First.(type) {
    case uint8:
      for i := int64(0); i < a1.Size; i++ {
        a2.Data.([]Pair)[i].First = uint8(vol.C1 * float32(a1.Data.([]Pair)[i].First.(uint8)))
        a2.Data.([]Pair)[i].Second = uint8(vol.C2 * float32(a1.Data.([]Pair)[i].Second.(uint8)))
      }
      break
    case uint16:
      for i := int64(0); i < a1.Size; i++ {
        a2.Data.([]Pair)[i].First = uint16(vol.C1 * float32(a1.Data.([]Pair)[i].First.(uint16)))
        a2.Data.([]Pair)[i].Second = uint16(vol.C2 * float32(a1.Data.([]Pair)[i].Second.(uint16)))
      }
    }
    break
  }

  return
}

func (a* Audio) Rms() (value1 float64, value2 float64 , err error) {
  sum1 := float64(0)
  sum2 := float64(0)
  switch a.Data.(type) {
  case []uint8:
    for _, v := range(a.Data.([]uint8)) {
      sum1 += float64(v * v)
    }
    value1 = math.Sqrt(sum1 / float64(a.Size))
    break
  case []uint16:
    for _, v := range(a.Data.([]uint16)) {
      sum1 += float64(v * v)
    }
    value1 = math.Sqrt(sum1 / float64(a.Size))
    break
  case []Pair:
    switch a.Data.([]Pair)[0].First.(type) {
    case uint8:
      for _, v := range(a.Data.([]Pair)) {
        sum1 += float64(v.First.(uint8) * v.First.(uint8))
        sum2 += float64(v.Second.(uint8) * v.Second.(uint8))
      }
      break
    case uint16:
      for _, v := range(a.Data.([]Pair)) {
        sum1 += float64(v.First.(uint16) * v.First.(uint16))
        sum2 += float64(v.Second.(uint16) * v.Second.(uint16))
      }
      break
    }
    value1 = math.Sqrt(sum1 / float64(a.Size))
    value2 = math.Sqrt(sum2 / float64(a.Size))
    break
  }
  
  return
}

// func (a1* Audio) Norm(rms_d float64) (a2* Audio, err error) {
//   //validate
//   rms_c, err := a1.Rms()
//   if err != nil {
//     return
//   }
//   a2 = &Audio8M {
//     Data: make([]byte, a1.Size),
//     Channel: a1.Channel,
//     Size: a1.Size,
//     SamplingRate: a1.SamplingRate,
//     NumberOfSamples: a1.NumberOfSamples,
//     Length: a1.Length,
//   }
//   for i := int64(0); i < a1.Size; i++ {
//     a2.Data[i] = uint8((rms_d * float64(a1.Data[1])) / rms_c);
//   }
//   return
// }

// func (a1* Audio8S) Norm(rms_d_1 float64, rms_d_2 float64) (a2* Audio8S, err error) {
//   //validate
//   rms_c_1, rms_c_2, err := a1.Rms()
//   if err != nil {
//     return
//   }
//   a2 = &Audio8S {
//     Data: make([]Pair8, a1.Size),
//     Channel: a1.Channel,
//     Size: a1.Size,
//     SamplingRate: a1.SamplingRate,
//     NumberOfSamples: a1.NumberOfSamples,
//     Length: a1.Length,
//   }
//   for i := int64(0); i < a1.Size; i++ {
//     a2.Data[i].First = uint8((rms_d_1 * float64(a1.Data[1].First)) / rms_c_1);
//     a2.Data[i].Second = uint8((rms_d_2 * float64(a1.Data[1].Second)) / rms_c_2);
//   }
//   return
// }