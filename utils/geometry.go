package utils

import "image"

func GetCenter(x, y int) image.Point {
	return image.Point{
		X: x / 2,
		Y: y / 2,
	}
}
