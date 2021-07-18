package geometry

type Rectangle struct {
	height int
	widht int
}



func  areaOfRectangle(r *Rectangle) int {
	return (*r).height*(*r).widht
}
