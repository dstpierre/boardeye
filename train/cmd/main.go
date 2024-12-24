package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dstpierre/boardeye/train"
)

var game train.Game

func main() {
	flag.Parse()

	if _, err := os.Stat("./game.json"); os.IsNotExist(err) {
		fmt.Println("no game created")
	} else {
		file, err := os.Open("./game.json")
		if err != nil {
			fmt.Println("error reading game file: ", err)
			os.Exit(1)
		}
		defer file.Close()

		g, err := train.Load(file)
		if err != nil {
			fmt.Println("error loading game JSON file: ", err)
			os.Exit(1)
		}

		game = g
	}

	switch flag.Arg(0) {
	case "start":
		start()
	case "route":
		route()
	case "search":
		search()
	case "turn":
		playTurn()
	default:
		fmt.Println("Available command:\n\tstart")
	}
}

func start() {
	reader := bufio.NewReader(os.Stdin)

	var players []string

	for {
		fmt.Print("Next player name (empty to stop): ")
		s, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("error reading stdin: ", err)
			os.Exit(1)
		}

		name := strings.Replace(s, "\n", "", -1)
		if len(name) == 0 {
			break
		}

		players = append(players, name)
	}

	g, err := train.Start(players)
	if err != nil {
		fmt.Println("error creating game: ", err)
		os.Exit(1)
	}

	save(g)
}

func save(g train.Game) {
	f, err := os.Create("./game.json")
	if err != nil {
		fmt.Println("error writing data file: ", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := g.Save(f); err != nil {
		fmt.Println("error writing file: ", err)
		os.Exit(1)
	}
}

func route() {
	start := train.FindCity(game.Graph, flag.Arg(1))
	end := train.FindCity(game.Graph, flag.Arg(2))

	results, err := train.FindShortesPath(game.Graph, start, end)
	if err != nil {
		fmt.Println("error finding shortest route: ", err)
		os.Exit(1)
	}

	for _, station := range results {
		fmt.Println(station)
	}
}

func search() {
	n, err := strconv.Atoi(flag.Arg(2))
	if err != nil {
		fmt.Println("the 2nd argument must be the number of card of that color")
		os.Exit(1)
	}

	color := strings.ToUpper(flag.Arg(1))

	stations := train.RoutesByNumberOfColor(game.Graph, color, n)
	for _, station := range stations {
		fmt.Println(station)
	}
}

func playTurn() {
	action := flag.Arg(1)
	extra := flag.Arg(2)

	player := game.Players[game.NextPlayer]

	fmt.Println(player, " turn.")

	var station train.Station

	at := train.ActionPickRouteCard
	if action == "draw" {
		at = train.ActionDraw
	} else if action == "connect" {
		at = train.ActionConnect
	}

	if at == train.ActionConnect {
		target, ok := train.FindRoute(game.Graph, extra)
		if !ok {
			fmt.Println("cannot find this route: ", extra)
			os.Exit(1)
		}

		station = target

		tagAsTaken(target)
	}

	game.Turns = append(game.Turns, train.Turn{
		Player:  player,
		Action:  train.ActionType(at),
		Station: station,
		Point:   calculatePoint(station),
	})

	game.NextPlayer += 1
	if game.NextPlayer >= len(game.Players) {
		game.NextPlayer = 0
	}

	save(game)

	fmt.Println("OK, done")
}

func tagAsTaken(target train.Station) {
	reverse := train.Station{
		Start:  target.End,
		End:    target.Start,
		Color:  target.Color,
		Length: target.Length,
	}

	for _, stations := range game.Graph {
		for idx, station := range stations {
			if station.Equal(target) || station.Equal(reverse) {
				stations[idx].Taken = true
			}
		}
	}
}

func calculatePoint(station train.Station) int {
	if station.Length <= 2 {
		return 0
	} else if station.Length == 3 {
		return 5
	} else if station.Length == 4 {
		return 7
	} else if station.Length == 5 {
		return 10
	}

	return 15
}
