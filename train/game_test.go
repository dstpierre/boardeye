package train

import (
	"bytes"
	"slices"
	"strings"
	"testing"
)

func TestStartGame(t *testing.T) {
	g, err := Start([]string{"unit", "test"})
	if err != nil {
		t.Fatal(err)
	} else if len(g.Players) != 2 {
		t.Errorf("expected 2 players got %d", len(g.Players))
	}
}

func TestSaveAndLoadGame(t *testing.T) {
	g, err := Start([]string{"unit", "test"})
	if err != nil {
		t.Fatal(err)
	}

	g.Turns = append(g.Turns, Turn{
		Player: "unit",
		Action: ActionPickRouteCard,
	})

	var buf bytes.Buffer
	if err := g.Save(&buf); err != nil {
		t.Fatal(err)
	}

	rdr := strings.NewReader(buf.String())

	g2, err := Load(rdr)
	if err != nil {
		t.Fatal(err)
	} else if len(g2.Turns) != 1 {
		t.Errorf("expected g2 to have 1 turn got %d", len(g2.Turns))
	} else if slices.Compare(g.Players, g2.Players) != 0 {
		t.Error("expected g and g2 to have same players", g.Players, g2.Players)
	}
}
