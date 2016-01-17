tic-tac-toe
===========

This is a simple command line tic-tac-toe game written in go.

The board has the followoing layout

  1 | 2 | 3 |

 \-\-\-\-\-\-\-\-\-\-\-\-\-

  4 | 5 | 6 |

 \-\-\-\-\-\-\-\-\-\-\-\-\-

  7 | 8 | 9 |

## Build

To build simple do:

```shell
$ go build -o tic-tac-toe ./
```

### Usage

This game has modes to play with another player either remotely or locally.
You can also play with ai.

To play with the ai do:

```shell
$ ./tic-tac-toe --opponentType ai
```

To player with another player locally do:

```shell
$ ./tic-tac-toe --opponentType human
```

To player a multiplayer game one person will start a server and the other will connect.

To start the server do:

```shell
$ ./tic-tac-toe --opponentType remote
```

To connect to the server do:

```shell
$ ./tic-tac-toe --connect <host-name>:8080
```
