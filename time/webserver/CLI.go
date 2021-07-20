package poker

import (
	"bufio"
	"io"
	"strings"
	"time"
)

type CLI struct {
	playerStore  PlayerStore
	scanner      *bufio.Scanner
	blindAlerter BlindAlerter
}

func NewCLI(store PlayerStore, in io.Reader, alerter BlindAlerter) *CLI {
	return &CLI{
		playerStore:  store,
		scanner:      bufio.NewScanner(in),
		blindAlerter: alerter,
	}
}

func (cli *CLI) PlayPoker() {
	cli.scheduleBlindAlerts()
	cli.playerStore.RecordWin(extractWinner(cli.readLine()))
}

func (cli *CLI) scheduleBlindAlerts() {
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		cli.blindAlerter.ScheduleAlertAt(blindTime, blind)
		blindTime += 10 * time.Minute
	}
}

func extractWinner(input string) string {
	return strings.Replace(input, " wins", "", 1)
}
func (cli *CLI) readLine() string {
	cli.scanner.Scan()
	return cli.scanner.Text()
}
