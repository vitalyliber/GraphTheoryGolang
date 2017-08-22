package handlers

import (
	"../models"
	"encoding/json"
	"fmt"
	"net/http"
)

var Vertexes = []models.Vertex{
	models.Vertex{
		Name:  "First",
		Value: 1,
		Edges: []models.Edge{
			models.Edge{
				Value: 2,
				Name:  "Edge",
			},
		},
	},
	models.Vertex{
		Name:  "Last",
		Value: 2,
	},
}

func CreateVertex(w http.ResponseWriter, r *http.Request) {
	Vertexes = append(Vertexes, models.Vertex{
		Name:  "Last",
		Value: 2,
	})

	out, err := json.Marshal(Vertexes)
	if err != nil {
		panic(err)
	}

	fmt.Println(r.FormValue("value"))
	fmt.Println(r.FormValue("name"))
	fmt.Fprintf(w, "%s", out)
}
