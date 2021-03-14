package structs

type SessionWatermarkBlock struct {
Alpha float64
DataPoints []interface{}
FontSize *string //small,medium,large,huge
Name *string
Position *string //top_left,top_center,top_right,middle_left,middle_center,middle_right,bottom_left,bottom_center,bottom_right
PositionReferencePoint *string //top_left,top_center,top_right,middle_left,middle_center,middle_right,bottom_left,bottom_center,bottom_right
PositionX int
PositionY int
ScrollText *string //none,ltr,rtl
TextAlignment *string //left,right,center
TextColor *string
TextShadow bool
}
