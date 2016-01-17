package main

import (
	"flag"
	"fmt"
	"net"

	"github.com/mattnenterprise/tic-tac-toe/game"
)

func main() {
	opponentTypePtr := flag.String("opponentType", "none", "The type of oppenent to play: human, ai, remote")
	connectPtr := flag.String("connect", "none", "The other players hostname to join the game")
	flag.Parse()

	if *opponentTypePtr == "none" && *connectPtr == "none" {
		fmt.Println("Please select either opponentType or connect option to play a game")
		return
	}

	if *opponentTypePtr != "none" && *connectPtr != "none" {
		fmt.Println("Both oppenentType and connect cannot be set at the same time. Please refer to the README.md on how to use this program.")
		return
	}

	if *opponentTypePtr != "none" {
		if *opponentTypePtr == "remote" {
			fmt.Println("Listening on port 8080 for another player to connect")
			ln, err := net.Listen("tcp", ":8080")
			if err != nil {
				panic(err)
			}
			conn, err := ln.Accept()
			if err != nil {
				panic(err)
			}
			player1 := game.NewLocalNetworkPlayer(game.X, conn)
			player2 := game.NewRemoteNetworkPlayer(game.O, conn)
			game.NewTicTacToeGame(player1, player2).Run()
		} else if *opponentTypePtr == "ai" {
			player1 := game.NewHumanPlayer(game.X)
			player2 := game.NewAiPlayer(game.O)
			game.NewTicTacToeGame(player1, player2).Run()
		} else if *opponentTypePtr == "human" {
			player1 := game.NewHumanPlayer(game.X)
			player2 := game.NewHumanPlayer(game.O)
			game.NewTicTacToeGame(player1, player2).Run()
		}
	}

	if *connectPtr != "none" {
		conn, err := net.Dial("tcp", *connectPtr)
		if err != nil {
			panic(err)
		}
		player1 := game.NewLocalNetworkPlayer(game.O, conn)
		player2 := game.NewRemoteNetworkPlayer(game.X, conn)
		game.NewTicTacToeGame(player1, player2).Run()
	}
}
