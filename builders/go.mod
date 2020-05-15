module github.com/sindbach/build-game/builders

go 1.13

require (
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/hajimehoshi/ebiten v1.11.1
	github.com/sindbach/build-game/builders/images v0.0.0
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
)

replace github.com/sindbach/build-game/builders/images => ./images
