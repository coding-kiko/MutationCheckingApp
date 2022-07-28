package mutant

type DnaRequest struct {
	Dna []string `json:"dna,omitempty"`
}

type Stats struct {
	CountMutantDna int64    `json:"count_mutant_dna"`
	CountHumanDna  int64    `json:"count_human_dna"`
	Ratio          *float64 `json:"ratio"` // pointer to float beacuse it can be undefined => nil
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

type MutantDnaResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
