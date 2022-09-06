package models

type ResultByIdSeq struct {
	Record []struct {
		SequenceName string `json:"sequence_name"`
		TableName    string `json:"table_name"`
	} `json:"RECORDS"`
}
