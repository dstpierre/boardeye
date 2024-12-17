package train

import (
	"fmt"
	"testing"
)

func mustBuildGraph(t *testing.T) Graph {
	graph, err := buildGraph()
	if err != nil {
		t.Fatal(err)
	}

	return graph
}

func TestFindRoute(t *testing.T) {
	graph := mustBuildGraph(t)

	station, ok := findRoute(graph, "mon-bos-g")
	if !ok {
		t.Error("cannot find route via 'mbg'")
	} else if station.Start != "MONTREAL" {
		t.Errorf("start should be BOSTON got %s", station.Start)
	} else if station.End == "MONTREAL" {
		t.Errorf("end should be MONTREAL, got %s", station.End)
	}
}

func TestEnsureCityNameDoesNotCollide(t *testing.T) {
	t.Skip()

	graph := mustBuildGraph(t)

	tla := make(map[string]Station)
	for _, stations := range graph {
		for _, station := range stations {
			s := fmt.Sprintf("%s%s%s", station.Start[:1], station.End[:1], station.Color[:1])

			if other, ok := tla[s]; ok {
				t.Errorf("tla for station %s taken by %s", station, other)
			}

			tla[s] = station
		}
	}
}

func TestTmp(t *testing.T) {
	graph := mustBuildGraph(t)

	colors := make(map[string]int)

	for _, stations := range graph {
		for _, station := range stations {
			n := colors[station.Color]
			n += station.Length
			colors[station.Color] = n
		}
	}

	for k, v := range colors {
		fmt.Println(k, v)
	}
}