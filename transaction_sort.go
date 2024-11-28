package main

import "time"

type Transaction struct {
	Payer     string    `json:"payer"`
	Points    int       `json:"points"`
	Timestamp time.Time `json:"timestamp"`
}

// TransactionsByTimestamp implements sort.Interface to sort transactions by Timestamp
type TransactionsByTimestamp []Transaction

func (t TransactionsByTimestamp) Len() int { return len(t) }

func (t TransactionsByTimestamp) Less(i, j int) bool { return t[i].Timestamp.Before(t[j].Timestamp) }

func (t TransactionsByTimestamp) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
