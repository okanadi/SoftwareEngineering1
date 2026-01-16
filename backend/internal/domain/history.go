package domain

import (
	"time"
)

type HistoryDB struct {
	ID        string    `db:"id"`
	StepId    string    `db:"step_id"`
	Status    string    `db:"status"`
	Note      string    `db:"note"`
	UserName  string    `db:"user_name"`
	Timestamp time.Time `db:"timestamp"`
	Photos    []string  `db:"photos"`
}

type ProjectStepHistoryDTO struct {
	ProjectStepDB
	History []HistoryDetailDTO `json:"history"`
}

type HistoryDetailDTO struct {
	ID        string     `json:"id"`
	Status    string     `json:"status"`
	Note      string     `json:"note"`
	UserName  string     `json:"user_name"`
	Timestamp time.Time  `json:"timestamp"`
	Photos    []PhotoDTO `json:"photos"`
}

type PhotoDTO struct {
	ID       string `json:"id"`
	S3Key    string `json:"s3_key"`
	FileType string `json:"file_type"`
	Url      string `json:"url"`
}
