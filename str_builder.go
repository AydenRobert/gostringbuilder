package gostringbuilder

import "fmt"

var SB_ARR_LEN uint = 50

type StringBuilder struct {
	posx, posy uint
	chars [][]rune
}

func newStringBuilderFunc() StringBuilder {
	chars := make([][]rune, 0, 10)
	chars = append(chars, make([]rune, SB_ARR_LEN))
	sb := StringBuilder{0, 0, chars}

	return sb
}

func (sb *StringBuilder) ar(r rune) {
	sb.chars[sb.posx][sb.posy] = r
	sb.posy++
	if sb.posy >= SB_ARR_LEN {
		addRuneArray(sb)
	}
}

func (sb *StringBuilder) as(s string) {
	runes := []rune(s)
	lenLeft := SB_ARR_LEN - sb.posy
	if uint(len(s)) > lenLeft {
		runepos := 0
		for ;; {
			for i := uint(0); i < lenLeft - 1; i++ {
				sb.chars[sb.posx][sb.posy] = runes[runepos]
				sb.posy++
				runepos++
			}
			if runepos + 1 >= len(runes) {
				return
			}
			addRuneArray(sb)
			if len(runes) - runepos + 1 < int(SB_ARR_LEN) {
				lenLeft = uint(len(runes) - runepos + 1)
			} else {
				lenLeft = SB_ARR_LEN
			}
		}
	}
	for i := range runes {
		sb.chars[sb.posx][sb.posy] = runes[i]
		sb.posy++
	}
}

func (sb *StringBuilder) writeString() string {
	outstr := ""
	for i := uint(0); i < sb.posx; i++ {
		for j := uint(0); j < SB_ARR_LEN; j++ {
			outstr += string(sb.chars[i][j])
		}
	}
	for i := uint(0); i < sb.posy; i++ {
		outstr += string(sb.chars[sb.posx][i])
	}
	return outstr
}

func addRuneArray(sb *StringBuilder) {
	sb.chars = append(sb.chars, make([]rune, SB_ARR_LEN))
	sb.posx++
	sb.posy = 0
}

func main() {
	sb := newStringBuilderFunc()
	sb.as("appending a lot of chars so i can see if it works aye, it probably won't...")
	sb.ar('\n')
	sb.as("is this enough chars to go through the complicated shit i set up?")
	fmt.Printf("%s\n", sb.writeString())
}
