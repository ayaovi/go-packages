package main

import (
	// "fmt"
	"go-packages/audio"
	"unsafe"
	"io/ioutil"
	"os"
	"encoding/json"
	"encoding/binary"
	"strconv"
	"strings"
)

type AudioFiles struct {
	AudioFiles []AudioFile `json:"audioFiles"`
}

type AudioFile struct {
	DirPath string `json:"dirPath"` // e.g. /home/some-dir/
	Name string `json:"name"`
	Extension string `json:"ext"`	// e.g. .raw
	SamplingRate uint `json:"samplingRate"`
	Channel uint `json:"channel"`
	BitCount uint `json:"bitCount"`
	Signed bool `json:"signed"`
}

func check(e error) {
	if e != nil { panic(e) }
}

func loadAudio(file* AudioFile) (output* audio.Audio) {
	audioDataMakers := make(map[uint]func(*os.File, int64) interface{})
	
	// 8-bit mono.
	audioDataMakers[9] = func(file *os.File, size int64) interface{} {
		data := make([]byte, size)
		_, err := file.Read(data)
		check(err)
		return data
	}
	// 8-bit stereo.
	audioDataMakers[10] = func(file *os.File, size int64) interface{} {
		data := make([]audio.Pair, size / 2)
		bytes := make([]byte, size)
		_, err := file.Read(bytes)
		check(err)

		for i := int64(0); i < size; i += 2 {
			data[i / 2].First = bytes[i]
			data[i / 2].Second = bytes[i + 1]
		}
		return data
	}
	// 16-bit mono.
	audioDataMakers[17] = func(file *os.File, size int64) interface{} {
		var data []uint16
		bytes := make([]byte, size)
		_, err := file.Read(bytes)
		check(err)

		for i := 0; i < len(bytes); i += 2 {
			data = append(data, binary.LittleEndian.Uint16(bytes[i:i + 2]))
		}
		return data
	}
	// 16-bit stereo.
	audioDataMakers[18] = func(file *os.File, size int64) interface{} {
		data := make([]audio.Pair, size / 4)
		bytes := make([]byte, size)
		var twoBytes []uint16
		_, err := file.Read(bytes)
		check(err)

		// 1st transform the []byte to []uint16.
		for i := int64(0); i < size; i += 2 {
			twoBytes = append(twoBytes, binary.LittleEndian.Uint16(bytes[i:(i + 2)]))
		}
		
		// then populate the data.
		for i := 0; i < len(twoBytes); i += 2 {
			data[i / 2].First = twoBytes[i]
			data[i / 2].Second = twoBytes[i + 1]
		}
		return data
	}
	
	f, err := os.Open(getFullAudioFilePath(file))
	check(err)

	fi, err := f.Stat()
	check(err)

	output = &audio.Audio {
		Channel: file.Channel,
		SamplingRate: file.SamplingRate,
		Size: fi.Size(),
		Data: audioDataMakers[file.Channel + file.BitCount](f, fi.Size()),
	}

	defer f.Close()

	switch output.Data.(type) {
	case []uint8:
		output.NumberOfSamples = output.Size
		break
	case []uint16:
		output.NumberOfSamples = output.Size / 2
		break
	case []audio.Pair:
		switch output.Data.([]audio.Pair)[0].First.(type) {
		case uint8:
			output.NumberOfSamples = output.Size / 2
			break
		case uint16:
			output.NumberOfSamples = output.Size / 4
			break
		}
		break
	}
	
	output.Length = output.NumberOfSamples / int64(output.SamplingRate)

	return
}

func saveAudio(path string, file* audio.Audio) () {
	audioWriters := make(map[uint]func(*os.File, *audio.Audio))

	// 8-bit mono
	audioWriters[9] = func (writer* os.File, a* audio.Audio) {
		_, err := writer.Write(a.Data.([]byte))
		check(err)
	}
	// 8-bit stereo
	audioWriters[10] = func (writer* os.File, a* audio.Audio) {
		bytes := make([]byte, a.Size)

		for i := 0; i < len(bytes); i += 2 {
			bytes[i] = a.Data.([]audio.Pair)[i / 2].First.(uint8)
			bytes[i + 1] = a.Data.([]audio.Pair)[i / 2].Second.(uint8)
		}
		_, err := writer.Write(bytes)
		check(err)
	}
	// 16-bit mono
	audioWriters[17] = func (writer* os.File, a* audio.Audio) {
		bytes := make([]byte, a.Size)
		for i := 0; i < len(bytes); i += 2 {
			x := (*[2]byte)(unsafe.Pointer(&(a.Data.([]uint16)[i / 2])))[:]
			bytes[i] = x[0]
			bytes[i + 1] = x[1]
		}

		_, err := writer.Write(bytes)
		check(err)
	}
	// 16-bit stereo
	audioWriters[18] = func (writer* os.File, a* audio.Audio) {
		bytes := make([]byte, a.Size)

		for i := 0; i < len(bytes); i += 4 {
			x := (*[2]byte)(unsafe.Pointer(&a.Data.([]audio.Pair)[i / 4].First))[:]
			y := (*[2]byte)(unsafe.Pointer(&a.Data.([]audio.Pair)[i / 4].Second))[:]
			bytes[i] = x[0]
			bytes[i + 1] = x[1]
			bytes[i + 2] = y[0]
			bytes[i + 3] = y[1]
		}

		_, err := writer.Write(bytes)
		check(err)
	}

	f, err := os.Create(path)
	check(err)

	switch file.Data.(type) {
	case []uint8:
		audioWriters[file.Channel + 8](f, file)
		break
	case []uint16:
		audioWriters[file.Channel + 16](f, file)
		break
	case []audio.Pair:
		switch file.Data.([]audio.Pair)[0].First.(type) {
		case uint8:
			audioWriters[file.Channel + 8](f, file)
			break
		case uint16:
			audioWriters[file.Channel + 16](f, file)
			break
		}
	}
}

func loadAudioFiles(path string) (audioFiles* AudioFiles) {
	jsonFile, err := os.Open(path)
	check(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &audioFiles)
	return
}

func getFullAudioFilePath(file* AudioFile) string {
	var fullPath strings.Builder
	fullPath.WriteString(file.DirPath)
	fullPath.WriteString(file.Name)
	fullPath.WriteString(file.Extension)
	return fullPath.String()
}

func audioAddition(args []string, audioFiles* AudioFiles) (output* audio.Audio) {
	index1, err := strconv.ParseInt(args[0], 10, 8)
	check(err)

	index2, err := strconv.ParseInt(args[1], 10, 8)
	check(err)

	input1 := loadAudio(&audioFiles.AudioFiles[index1])
	input2 := loadAudio(&audioFiles.AudioFiles[index2])

	output, err = input1.Plus(input2)
	check(err)

	return
}

func audioCut(args []string, audioFiles* AudioFiles) (output* audio.Audio) {
	startRange, err := strconv.ParseInt(args[0], 10, 64)
	check(err)

	endRange, err := strconv.ParseInt(args[1], 10, 64)
	check(err)

	audioFileIndex, err := strconv.ParseInt(args[2], 10, 8)
	check(err)

	input := loadAudio(&audioFiles.AudioFiles[audioFileIndex])

	output, err = input.Cut(startRange, endRange)
	check(err)
	return
}

func main() {
	// driver [<ops>] soundFileIndex1 [soundFileIndex2]
	// <ops> can be one of the following
	// -add: add soundFile1 and soundFile2
	// -cut r1 r2: remove samples over range [r1,r2]
	// -radd r1 r2 s1 s2: add soundFile1 and soundFile2 over sub-ranges indicated (in seconds). The ranges must be equal in length. 
	// -cat: concatenate soundFile1 and soundFile2
	// -v r1 r2: volume factor for left/right audio (def=1.0/1.0) (assumes one sound file)
	// -rev: reverse sound file (assumes one sound file only)
	// -rms: Prints out the RMS of the soundfile (assumes one sound file only).

	audioFiles := loadAudioFiles("../audio-files.json")

	args := os.Args[1:]

	operations := make(map[string]func([]string, *AudioFiles) *audio.Audio)

	operations["-add"] = audioAddition
	operations["-cut"] = audioCut
	operations["-radd"] = audioAddition
	operations["-cat"] = audioAddition

	output := operations[args[0]](args[1:], audioFiles)

	// output, err := input.Cut(0, 617400)
	// check(err)

	saveAudio("/home/ayaovi/Downloads/input_files/countdown14sec_44100_signed_16bit_mono.raw", output)
}