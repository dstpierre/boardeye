package train

import "strings"

func findRoute(graph Graph, info string) (Station, bool) {
	info = strings.Replace(info, "-", "", -1)
	desired, ok := parseInfo(graph, info)
	if !ok {
		return Station{}, ok
	}

	for _, stations := range graph {
		for _, station := range stations {
			if desired.Equal(station) {
				return station, true
			}
		}
	}

	return Station{}, false
}

func parseInfo(graph Graph, s string) (Station, bool) {
	if len(s) < 7 {
		return Station{}, false
	}

	start := s[:3]
	end := s[3:6]
	color := strings.ToUpper(s[6:])

	station := Station{
		Start: findCity(graph, start),
		End:   findCity(graph, end),
		Color: color,
	}

	return station, true
}

func findCity(graph Graph, s string) string {
	for city := range graph {
		tmp := strings.Replace(city, " ", "", -1)

		if strings.EqualFold(s, tmp[:3]) {
			return city
		}
	}

	return ""
}

func routesByNumberOfColor(graph Graph, color string, n int) []Station {
	uniq := make(map[string]bool)

	var matches []Station

	for _, stations := range graph {
		for _, station := range stations {
			if station.Taken || station.Color != color || station.Length != n {
				continue
			}

			if _, ok := uniq[station.UniqueKey()]; !ok {
				matches = append(matches, station)
				uniq[station.UniqueKey()] = true
			}
		}
	}

	return matches
}
