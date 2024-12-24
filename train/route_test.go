package train

import (
	"testing"
)

func TestFindShortesRoute(t *testing.T) {
	graph, err := buildGraph()
	if err != nil {
		t.Fatal(err)
	}

	//route, err := findShortestRoute(graph, "Miami", "Winnipeg")
	route, err := FindShortesPath(graph, "MIAMI", "WINNIPEG")
	if err != nil {
		t.Fatal(err)
	}

	cities := []string{"MIAMI", "ATLANTA", "NASHVILLE", "SAINT LOUIS", "CHICAGO", "DULUTH", "WINNIPEG"}

	if len(cities) != len(route) {
		if len(route) < 10 {
			t.Log(cities)
			for _, r := range route {
				t.Log(r.Start)
			}
		}
		t.Fatalf("shortest route does not have 7 station, got %d", len(route))
	}

	for idx, station := range route {
		if station.Start != cities[idx] {
			t.Errorf("expected %d station to be %s got %s", idx, cities[idx], station.Start)
		}
	}
}

func TestAltShortestRoute(t *testing.T) {
	graph, err := buildGraph()
	if err != nil {
		t.Fatal(err)
	}

	// take the Duluth -> Winnipeg track
	duluth, ok := graph["DULUTH"]
	if !ok {
		t.Fatal("cannot find Duluth")
	}

	for idx, station := range duluth {
		if station.End == "WINNIPEG" {
			duluth[idx].Taken = true
			break
		}
	}

	route, err := FindShortesPath(graph, "MIAMI", "WINNIPEG")
	if err != nil {
		t.Fatal(err)
	}

	cities := []string{"MIAMI", "CHARLESTON", "RALEIGH", "PITTSBURGH", "TORONTO", "SAULT STE. MARIE", "WINNIPEG"}

	if len(cities) != len(route) {
		if len(route) < 10 {
			t.Log(cities)
			for _, r := range route {
				t.Log(r.Start)
			}
		}
		t.Fatalf("shortest route does not have 7 station, got %d", len(route))
	}

	for idx, station := range route {
		if station.Start != cities[idx] {
			t.Errorf("expected %d station to be %s got %s", idx, cities[idx], station.Start)
		}
	}
}
