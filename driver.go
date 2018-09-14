package main

import (
	"go-packages/audio"
	// "fmt"
	"unsafe"
	"io/ioutil"
	"os"
	"encoding/json"
	"strconv"
	"strings"
)

type AudioFiles struct {
	AudioFiles []AudioFile `json:"audioFiles"`
}

type AudioFile struct {
	DirPath string `json:"dirPath"`
	Name string `json:"name"`
	Extension string `json:"ext"`
	SamplingRate uint `json:"samplingRate"`
	Channel uint `json:"channel"`
	BitsPerSample int8 `json:"bitCount"`
	Signed bool `json:"signed"`
}

func check(e error) {
	if e != nil { panic(e) }
}

func loadAudio(file* AudioFile) (output* audio.Audio) {
	f, err := os.Open(getAudioFilePullPath(file))
	check(err)

	fi, err := f.Stat()
	check(err)
	
	output = &audio.Audio {
		Channel: file.Channel,
		SamplingRate: file.SamplingRate,
		Size: fi.Size(),
		Data: make([]byte, fi.Size()),
	}
	
	_, err = f.Read(output.Data.([]byte))
	check(err)
	defer f.Close()

	output.NumberOfSamples = output.Size / (int64(unsafe.Sizeof(output.Data.([]byte)[0])) * int64(output.Channel))
	output.Length = output.NumberOfSamples / int64(output.SamplingRate)

	return
}

func saveAudio(path string, file* audio.Audio) () {
	f, err := os.Create(path)
	check(err)
	
	_, err = f.Write(file.Data.([]byte))
	check(err)
}

func loadAudioFiles(path string) (audioFiles* AudioFiles) {
	jsonFile, err := os.Open(path)
	check(err)
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &audioFiles)
	return
}

func getAudioFilePullPath(file* AudioFile) string {
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

	saveAudio("/home/ayaovi/Downloads/input_files/countdown14sec_44100_signed_8bit_mono.raw", output)
}
