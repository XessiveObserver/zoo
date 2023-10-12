package main

import "github.com/google/uuid"

type Animal struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Kind string    `json:"kind"`
	Diet string    `json:"diet"`
}
