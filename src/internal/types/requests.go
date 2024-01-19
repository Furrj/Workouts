package types

type ReqSet struct {
	SetID     uint8  `json:"set_id"`
	Timestamp uint64 `json:"timestamp"`
	Text      string `json:"text"`
	Name      string `json:"name"`
	Reps      string `json:"reps"`
	Weights   string `json:"weights"`
}

//type ReqSet struct {
//	Name  string   `json:"name"`
//	SetID uint8    `json:"set_id"`
//	Reps  []ReqRep `json:"reps"`
//}

type ReqRep struct {
	RepID  uint8  `json:"rep_id"`
	Count  uint8  `json:"count"`
	Weight uint16 `json:"weight"`
}
