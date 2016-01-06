package main

import (
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

//随机设置字符单元的属性
func draw() {
	w, h := termbox.Size()
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault) //清除内部缓存区
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			termbox.SetCell(x, y, ' ', termbox.ColorDefault, termbox.Attribute(rand.Int()%8)+1)
		}
	}
	termbox.Flush() //刷新后台缓存到界面里
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent() //监听键盘事件
		}
	}()

	draw()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey && ev.Key == termbox.KeyEsc {
				return
			}
		default:
			draw()
			time.Sleep(10 * time.Millisecond)
		}
	}
}
