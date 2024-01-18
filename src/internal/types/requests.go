package types

type ReqSet struct {
	Name string   `json:"name"`
	Reps []ReqRep `json:"reps"`
}

type ReqRep struct {
	Count  uint8  `json:"count,string"`
	Weight uint16 `json:"weight,string"`
}
