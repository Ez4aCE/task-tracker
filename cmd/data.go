package cmd

import "time"

type Status string

const (
	StatusOpen Status="open"
	StatusInProgress Status="in-progress"
	StatusDone Status = "done"
)

type Priority string

const(
	PriorityLow Priority="low"
	PriorityMedium Priority="medium"
	PriorityHigh Priority="high"
)

type Task struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status Status `json:"status"`
	Priority Priority `json:"priority"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type Store struct{
	Next_ID int `json:"next_id"`
	Tasks []Task `json:"tasks"`
}