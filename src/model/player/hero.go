package player

type Hero struct {
	StaticId uint16 `json:"static_id"`
	Lv       uint16 `json:"lv"`
	Exp      int    `json:"exp"`
	Rank     uint16 `json:"rank"`
}
