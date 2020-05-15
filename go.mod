module github.com/sindbach/build-game

go 1.13

require (
	github.com/hajimehoshi/ebiten v1.11.1
	github.com/sindbach/build-game/builders v0.0.0
	github.com/sindbach/build-game/builders/images v0.0.0

)

replace github.com/sindbach/build-game/builders => ./builders

replace github.com/sindbach/build-game/builders/images => ./builders/images
