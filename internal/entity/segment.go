package entity

//type segment string

//const (
//	VoiceMessage   segment = "AVITO_VOICE_MESSAGES"
//	PerformanceVas segment = "AVITO_PERFORMANCE_VAS"
//	Discount30     segment = "AVITO_DISCOUNT_30"
//	Discount50     segment = "AVITO_DISCOUNT_50"
//)

type Segment struct {
	ID      int    `json:"id"`
	Segment string `json:"segment"`
}
