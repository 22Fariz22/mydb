package entity

type Command string

type Pair struct{
	Key string `json:"key"`
	Value string `json:"value"`
}

type PairsList map[string]string


