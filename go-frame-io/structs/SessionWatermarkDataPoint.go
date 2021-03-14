package structs

type SessionWatermarkDataPoint struct {
Order int
Type *string //email,username,ip_address,timestamp,custom_text,user_input
Value *string
}
