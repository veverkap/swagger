package structs

type WatermarkImage struct {
Alpha float64
Image *Image
Position *string //top_left,top_right,bottom_left,bottom_right
}
