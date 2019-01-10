package main

import (
	"fmt"
	"github.com/gen2brain/raylib-go/raylib"
	"math"
	"math/rand"
)

const N = 30000

type rect struct {
	x, y  float32
	color rl.Color
	speed float32
	w, h  int32
	phase float32
}

var rects = [N]*rect{}

func load() {
	for i := 0; i < N; i++ {
		rects[i] = &rect{
			x:     947 * rand.Float32(),
			y:     781 * rand.Float32(),
			color: rl.NewColor(uint8(rand.Int()), uint8(rand.Int()), uint8(rand.Int()), 255),
			speed: 400 * rand.Float32(),
			w:     rand.Int31n(120),
			h:     rand.Int31n(120),
			phase: 2 * math.Pi * rand.Float32(),
		}
	}
}

func update() {
	for i := 0; i < N; i++ {
		rect := rects[i]
		rect.x += rect.speed * float32(math.Sin(float64(rl.GetTime()+rect.phase))) * rl.GetFrameTime()
	}
}

func draw() {
	for i := 0; i < N; i++ {
		rect := rects[i]
		rl.DrawRectangle(int32(rect.x), int32(rect.y), rect.w, rect.h, rect.color)
	}

	rl.DrawText(fmt.Sprint("fps: ", rl.GetFPS()), 20, 20, 14, rl.White)
}

func main() {
	rl.InitWindow(947, 781, "hello, raylib!")
	rl.SetTargetFPS(60)

	load()

	for !rl.WindowShouldClose() {
		update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		draw()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
