package go2048

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type G2048 [4][4]int

func (t *G2048) mirrorV() {
	temp := new(G2048)
	for i, line := range t {
		for j, num := range line {
			temp[len(t)-i-1][j] = num
		}
	}
	*t = *temp
}

func (t *G2048) right90() {
	temp := new(G2048)
	for i, line := range t {
		for j, num := range line {
			temp[j][len(t)-i-1] = num
		}
	}
	*t = *temp
}

func (t *G2048) left90() {
	temp := new(G2048)
	for i, line := range t {
		for j, num := range line {
			temp[len(t)-j-1][i] = num
		}
	}
	*t = *temp
}

func (t *G2048) right180() {
	temp := new(G2048)
	for i, line := range t {
		for j, num := range line {
			temp[len(t)-i-1][len(t)-j-1] = num
		}
	}
	*t = *temp
}

//func (t *G2048) Print() {
//	for _, line := range t {
//		for _, num := range line {
//			fmt.Printf("%2d ", num)
//		}
//		fmt.Println()
//	}
//	tn := G2048{{1, 2, 3, 4}, {5, 8}, {9, 10, 11}, {13, 14, 16}}
//	*t = tn
//}

func (t *G2048) mergeUp() bool {
	tLength := len(t)
	changed := false
	notfull := false
	for i := 0; i < tLength; i++ {
		np := tLength
		n := 0
		for x := 0; x < np; x++ {
			if t[x][i] != 0 {
				t[n][i] = t[x][i]
				if n != x {
					changed = true
				}
				n++
			}
		}
		if n < tLength {
			notfull = true
		}

		np = n
		for x := 0; x < np-1; x++ {
			if t[x][i] == t[x+1][i] {
				t[x][i] *= 2
				t[x+1][i] = 0
				Score += t[x][i] * step
				x++
				changed = true
			}
		}
		n = 0
		for x := 0; x < np; x++ {
			if t[x][i] != 0 {
				t[n][i] = t[x][i]
				if n != x {
					changed = true
				}
				n++
			}
		}
		for x := n; x < tLength; x++ {
			t[x][i] = 0
		}
	}
	return changed || !notfull
}

func (t *G2048) mergeDown() bool {
	t.right180()
	changed := t.mergeUp()
	t.right180()
	return changed
}

func (t *G2048) mergeLeft() bool {
	t.right90()
	changed := t.mergeUp()
	t.left90()
	return changed
}

func (t *G2048) mergeRight() bool {
	t.left90()
	changed := t.mergeUp()
	t.right90()
	return changed
}

func (t *G2048) checkWinOrAdd() Status {
	for _, line := range t {
		for _, num := range line {
			if num > Max {
				return Win
			}
		}
	}
	tLength := len(t)
	i := rand.Intn(tLength)
	j := rand.Intn(tLength)
	for x := 0; x < tLength; x++ {
		for y := 0; y < tLength; y++ {
			if t[i%tLength][j%tLength] == 0 {
				t[i%tLength][j%tLength] = 2 << (rand.Uint32() % 2)
				return Add
			}
			j++
		}
		i++
	}
	return Lose
}

func (t G2048) initialize(ox, oy int) error {
	fg := termbox.ColorYellow
	bg := termbox.ColorBlack
	termbox.Clear(fg, bg)
	str := "							Score:" + fmt.Sprint(Score)
	for n, c := range str {
		termbox.SetCell(ox+n, oy-1, c, fg, bg)
	}
	str = "ESC:exit " + "Enter:replay"
	for n, c := range str {
		termbox.SetCell(ox+n, oy-2, c, fg, bg)
	}
	str = " PLAY with ARROW KEY"
	for n, c := range str {
		termbox.SetCell(ox+n, oy-3, c, fg, bg)
	}

	tLength := len(t)
	fg = termbox.ColorBlack
	bg = termbox.ColorGreen
	for i := 0; i <= tLength; i++ {
		for x := 0; x < 5*tLength; x++ {
			termbox.SetCell(ox+x, oy+i*2, '-', fg, bg)
		}
		for x := 0; x <= 2*tLength; x++ {
			if x%2 == 0 {
				termbox.SetCell(ox+i*5, oy+x, '+', fg, bg)
			} else {
				termbox.SetCell(ox+i*5, oy+x, '|', fg, bg)
			}
		}
	}
	fg = termbox.ColorYellow
	bg = termbox.ColorBlack
	for i := range t {
		for j := range t[i] {
			if t[i][j] > 0 {
				str := fmt.Sprint(t[i][j])
				for n, char := range str {
					termbox.SetCell(ox+j*5+1+n, oy+i*2+1, char, fg, bg)
				}
			}
		}
	}
	return termbox.Flush()
}

func (t *G2048) mergeAndReturKey() termbox.Key {
	var changed bool
Lable:
	changed = false
	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent() //监听键盘事件
		}
	}()
	ev := <-eventQueue
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			changed = t.mergeUp()
		case termbox.KeyArrowDown:
			changed = t.mergeDown()
		case termbox.KeyArrowLeft:
			changed = t.mergeLeft()
		case termbox.KeyArrowRight:
			changed = t.mergeRight()
		case termbox.KeyEsc, termbox.KeyEnter:
			changed = true
		default:
			changed = false
		}
		if !changed {
			goto Lable
		}
	case termbox.EventResize:
		x, y := termbox.Size()
		t.initialize(x/2-10, y/2-4)
		goto Lable
	case termbox.EventError:
		panic(ev.Err)
	}
	step++
	return ev.Key
}

func (b *G2048) clear() {
	next := new(G2048)
	Score = 0
	step = 0
	*b = *next
}

func (b *G2048) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())
A:
	b.clear()
	for {
		st := b.checkWinOrAdd()
		x, y := termbox.Size()
		b.initialize(x/2-10, y/2-4)

		switch st {
		case Win:
			str := "Win!!"
			strl := len(str)
			coverPrintStr(x/2-strl/2, y/2, str, termbox.ColorMagenta, termbox.ColorYellow)
		case Lose:
			str := "Lose!!"
			strl := len(str)
			coverPrintStr(x/2-strl/2, y/2, str, termbox.ColorBlack, termbox.ColorRed)
		case Add:
		default:
			fmt.Print("Err")
		}

		key := b.mergeAndReturKey()
		if key == termbox.KeyEsc {
			return
		}
		if key == termbox.KeyEnter {
			goto A
		}
	}
}
