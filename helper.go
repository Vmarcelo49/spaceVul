package main

import "image"

func Coords(rect image.Rectangle) []*int {
	return []*int{&rect.Max.X, &rect.Max.Y}
}

func Size(rect image.Rectangle) []*int {
	return []*int{&rect.Min.X, &rect.Min.Y}
}
