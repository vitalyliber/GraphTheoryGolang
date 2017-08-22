package models

type Vertex struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
	Edges []Edge `json:"edges"`
}
