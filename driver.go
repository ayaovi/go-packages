package main

import (
	"go-packages/audio"
	"fmt"
	"unsafe"
	"os"
)

func check(e error) {
	if e != nil { panic(e) }
}

func loadAudio(path string) (output* audio.Audio) {
	f, err := os.Open(path)
	check(err)

	fi, err := f.Stat()
	check(err)
	
	output = &audio.Audio {
		Channel: 1,
		SamplingRate: 44100,
		Size: fi.Size(),
		Data: make([]byte, fi.Size()),
	}
	
	_, err = f.Read(output.Data.([]byte))
	check(err)
	defer f.Close()
	return
}

func saveAudio(path string, file* audio.Audio) () {
	f, err := os.Create(path)
	check(err)
	
	_, err = f.Write(file.Data.([]byte))
	check(err)
}

func main() {
	input := loadAudio("/home/ayaovi/Downloads/input_files/countdown40sec_44100_signed_8bit_mono.raw")
	
	input.NumberOfSamples = input.Size / (int64(unsafe.Sizeof(input.Data.([]byte)[0])) * int64(input.Channel))
	input.Length = input.NumberOfSamples / int64(input.SamplingRate)

	fmt.Printf("size of data is %d bytes.\n", input.Size)
	fmt.Printf("samplingRate is %d bytes long.\n", input.SamplingRate)
	fmt.Printf("numberOfSamples is %d bytes long.\n", input.NumberOfSamples)
	fmt.Printf("length is %d second(s).\n", input.Length)

	output, err := input.Cut(0, 617400)
	check(err)

	saveAudio("/home/ayaovi/Downloads/input_files/countdown14sec_44100_signed_8bit_mono.raw", output)
}
