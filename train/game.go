package train

import (
	"encoding/json"
	"io"
)

type Game struct {
	Players    []string
	Graph      Graph
	Turns      []Turn
	NextPlayer int
}

type ActionType string

const (
	ActionDraw          = "draw"
	ActionConnect       = "connect"
	ActionPickRouteCard = "pick"
)

type Turn struct {
	Player  string
	Action  ActionType
	Station Station
	Point   int
}

func (g Game) Save(w io.Writer) error {
	return json.NewEncoder(w).Encode(g)
}

func Start(players []string) (Game, error) {
	graph, err := buildGraph()
	if err != nil {
		return Game{}, err
	}

	return Game{
		Players: players,
		Graph:   graph,
	}, nil
}

func Load(r io.Reader) (Game, error) {
	var g Game
	if err := json.NewDecoder(r).Decode(&g); err != nil {
		return Game{}, err
	}

	return g, nil
}
