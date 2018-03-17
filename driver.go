package main

import (
	"go-packages/audio"
	//"fmt"
	"unsafe"
	"os"
)

//func load()

func main() {
	f, err := os.Open("C:\\Users\\djaye\\Downloads\\sample_input\\beez18sec_44100_signed_8bit_mono.raw")
	if err != nil {
		panic(err)
	}
	fi, err := f.Stat()
	if err != nil {
		// Could not obtain stat, handle error
		panic(err)
	}
	
	a := audio.Audio {
		Channel: 1,
		SamplingRate: 44100,
		Size: fi.Size(),
		Data: make([]byte, fi.Size()),
	}
	_, err1 := f.Read(a.Data)
	
	if err1 != nil {
		// could not read file content.
		panic(err1)
	}
	defer f.Close()
	
	a.NumberOfSamples = a.Size / (int64(unsafe.Sizeof(a.Data[0])) * int64(a.Channel))
	a.Length = a.NumberOfSamples / int64(a.SamplingRate)

	//fmt.Printf("size of data is %d bytes.\n", a.Size)
	//fmt.Printf("samplingRate is %d bytes long.\n", a.SamplingRate)
	//fmt.Printf("numberOfSamples is %d bytes long.\n", a.NumberOfSamples)
	//fmt.Printf("length is %d bytes long.\n", a.Length)
  
  //for i := 0; i < 10; i++ {
    //fmt.Printf("index %d: %d\n", i ,a.Data[i])
  //}
  
  a.Plus(&a)
  
  //for i := 0; i < 10; i++ {
    //fmt.Printf("index %d: %d\n", i ,out.Data[i])
  //}
}
