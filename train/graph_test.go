package train

import (
	"testing"
)

func TestCreateCitiesGraph(t *testing.T) {
	graph, err := buildGraph()
	if err != nil {
		t.Fatal(err)
	} else if len(graph) != 37 {
		t.Errorf("expected 37 cities got %d", len(graph))
	}
}

func TestGraphIntegrity(t *testing.T) {
	graph, err := buildGraph()
	if err != nil {
		t.Fatal(err)
	}

	bos := Station{Start: "MONTREAL", End: "BOSTON", Color: "GRAY", Length: 2}
	ny := Station{Start: "MONTREAL", End: "NEW YORK", Color: "BLUE", Length: 3}
	tor := Station{Start: "MONTREAL", End: "TORONTO", Color: "GRAY", Length: 3}
	ssm := Station{Start: "MONTREAL", End: "SAULT STE. MARIE", Color: "BLACK", Length: 5}

	montreal, ok := graph["MONTREAL"]
	if !ok {
		t.Error("cannot find Montreal station")
	} else if len(montreal) != 5 {
		t.Errorf("expected 5 destination for Montreal got %d", len(montreal))
		t.Log(montreal)
	} else if montreal[0] != bos {
		t.Error("first montreal dest isn't Boston", montreal[0])
	} else if montreal[1] != ny {
		t.Error("2nd dest for Montreal isn't New York", montreal[1])
	} else if montreal[2] != tor {
		t.Error("3rd dest for Montreal isn't Toronto", montreal[2])
	} else if montreal[3] != ssm {
		t.Error("4th dest for Montreal isn't SSM", montreal[3])
	}
}

func TestEnsureGraphHasNoExtraCities(t *testing.T) {
	graph := mustBuildGraph(t)

	checks := make(map[string]int)

	for _, stations := range graph {
		for _, station := range stations {
			n := checks[station.String()]
			n += 1
			checks[station.String()] = n
		}
	}

	for k, v := range checks {
		if v > 2 {
			t.Errorf("station %s has more than 2 got %d (start <-> end) tracks", k, v)
		}
	}
}
