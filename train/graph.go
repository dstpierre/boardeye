package train

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Station struct {
	Start  string
	End    string
	Color  string
	Length int
	Taken  bool
}

func (s Station) String() string {
	return fmt.Sprintf("%s -> %s (%d %s)", s.Start, s.End, s.Length, s.Color)
}

func (s Station) UniqueKey() string {
	x := []string{s.Start, s.End}
	slices.Sort(x)
	return strings.Join(x, ",")
}

func (s Station) Equal(x Station) bool {
	return s.Start == x.Start && s.End == x.End && s.Color == x.Color
}

type Graph map[string][]Station

func buildGraph() (Graph, error) {
	graph := make(map[string][]Station)

	stations, err := parseCities()
	if err != nil {
		return nil, err
	}

	for _, station := range stations {
		g := graph[station.Start]
		g = append(g, station)
		graph[station.Start] = g
	}

	return graph, nil
}

func parseCities() ([]Station, error) {
	var stations []Station

	f, err := os.Open("./cities.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rdr := csv.NewReader(f)

	for {
		rec, err := rdr.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return nil, err
		} else if rec == nil {
			break
		}

		i, err := strconv.Atoi(rec[3])
		if err != nil {
			return nil, err
		}

		stations = append(stations, Station{
			Start:  strings.ToUpper(rec[0]),
			End:    strings.ToUpper(rec[1]),
			Color:  strings.ToUpper(rec[2]),
			Length: i,
		})

		stations = append(stations, Station{
			Start:  strings.ToUpper(rec[1]),
			End:    strings.ToUpper(rec[0]),
			Color:  strings.ToUpper(rec[2]),
			Length: i,
		})
	}

	return stations, nil
}
