package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// CubeState represents the state of the entire cube
type CubeState struct {
	Pieces []Piece `json:"pieces"`
}

type ScrambleRequest struct {
	MovesCount int       `json:"movesCount"`
	CubeState  CubeState `json:"cubeState"`
}

// RotationRequest represents the JSON structure for rotation requests

type RotationRequest struct {
	Move      string    `json:"move"`
	CubeState CubeState `json:"cubeState"`
}

func main() {
	router := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"}, // Specify your frontend URL
		AllowedMethods:   []string{"POST", "OPTIONS", "GET", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	router.HandleFunc("/scramble", func(w http.ResponseWriter, r *http.Request) {
		var req ScrambleRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		rotationSequence, newState := scrambleCube(req.CubeState.Pieces, req.MovesCount)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newState)
	})

	router.HandleFunc("/rotate", func(w http.ResponseWriter, r *http.Request) {
		var req RotationRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		if req.Move == "" {
			http.Error(w, "Missing 'move' field", http.StatusBadRequest)
			return
		}
		var newState CubeState

		switch req.Move {
		case "U":
			newState = CubeState{Pieces: TopClockwise(req.CubeState.Pieces)}
		case "U'":
			newState = CubeState{Pieces: TopCounterClockwise(req.CubeState.Pieces)}
		case "B":
			newState = CubeState{Pieces: BottomClockwise(req.CubeState.Pieces)}
		case "B'":
			newState = CubeState{Pieces: BottomCounterClockwise(req.CubeState.Pieces)}
		case "R":
			newState = CubeState{Pieces: RightClockwise(req.CubeState.Pieces)}
		case "R'":
			newState = CubeState{Pieces: RightCounterClockwise(req.CubeState.Pieces)}
		case "L":
			newState = CubeState{Pieces: LeftClockwise(req.CubeState.Pieces)}
		case "L'":
			newState = CubeState{Pieces: LeftCounterClockwise(req.CubeState.Pieces)}
		case "F":
			newState = CubeState{Pieces: FrontClockwise(req.CubeState.Pieces)}
		case "F'":
			newState = CubeState{Pieces: FrontCounterClockwise(req.CubeState.Pieces)}
		case "B":
			newState = CubeState{Pieces: BackClockwise(req.CubeState.Pieces)}
		case "B'":
			newState = CubeState{Pieces: BackCounterClockwise(req.CubeState.Pieces)}
		default:
			http.Error(w, "Unknown move: "+req.Move, http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newState)
	}).Methods("POST")

	log.Printf("Server running on port :8000")

	handler := c.Handler(router)

	http.ListenAndServe(":8000", handler)
}
