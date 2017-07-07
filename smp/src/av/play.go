package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player

	switch mtype {
	case "mp3":
		fallthrough
	case "MP3":
		p = &MP3Player{}
	case "wav":
		fallthrough
	case "WAV":
		p = &WAVPlayer{}

	default:
		fmt.Println("Unsupported media type ", mtype)
		return

	}

	p.Play(source)

}
