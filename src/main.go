package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bughomenoise/bsgender/seed"
	"github.com/bughomenoise/bsgender/stdout"
	tea "github.com/charmbracelet/bubbletea"
)

type modelType string

type model struct {
	cursor   *int
	choices  []string
	msg      string
	selected map[int]struct{}
	exit     *bool
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("bsgender")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			*m.exit = true
			return m, tea.Quit
		case "up", "k":
			if *m.cursor > 0 {
				*m.cursor--
			}
		case "down", "j":
			if *m.cursor < (len(m.choices) - 1) {
				*m.cursor++
			}
		case "enter", " ":
			m.selected[*m.cursor] = struct{}{}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("%s\n", m.msg)
	for i, choices := range m.choices {
		cursor := "[ ]"
		if *m.cursor == i {
			cursor = "[x]"
		}
		s += fmt.Sprintf("%s %s\n", cursor, choices)
	}
	return s
}

func initModel(msg string, choices []string, result *int, exit *bool) model {
	return model{
		cursor:   result,
		choices:  choices,
		msg:      msg,
		selected: make(map[int]struct{}),
		exit:     exit,
	}
}

func loopStrInput(n uint8) (seed.Seed, error) {
	var result seed.Seed
	reader := bufio.NewReader(os.Stdin)
	var s []string
	for i := range n {
		fmt.Printf("Input Text Entropy (%d/%d)\n", i+1, n)
		text, err := reader.ReadString('\n')
		if err != nil {
			return result, err
		}
		s = append(s, text)
	}
	if n == 16 {
		arr := [16]string{}
		copy(arr[:], s[:16])
		return seed.StringsTo12W(arr)
	} else if n == 32 {
		arr := [32]string{}
		copy(arr[:], s[:32])
		return seed.StringsTo24W(arr)
	}
	return result, errors.New("ErrorFunc: loopStrInput")
}

func loopByteInput(n uint8) (seed.Seed, error) {
	var result seed.Seed
	reader := bufio.NewReader(os.Stdin)
	var s []byte
	for i := range n {
		fmt.Printf("Input Byte(Number 0-255: Default = 0) Entropy (%d/%d)\n", i+1, n)
		var b byte
		for {
			str, err := reader.ReadString('\n')
			str = strings.Trim(str, "\n")
			if str == "" {
				str = "0"
			}
			if err != nil {
				return result, err
			}
			u, err := strconv.ParseUint(str, 0, 64)
			if err != nil || u > 255 {
				fmt.Printf("wrong number try again!\n")
				continue
			}
			b = byte(u)
			break
		}
		s = append(s, b)
	}
	if n == 16 {
		arr := [16]byte{}
		copy(arr[:], s[:16])
		return seed.BytesTo12W(arr)
	} else if n == 32 {
		arr := [32]byte{}
		copy(arr[:], s[:32])
		return seed.ByteTo24W(arr)
	}
	return result, errors.New("ErrorFunc: loopByteInput")

}

func getOptions(str string, choices []string) (int, error) {
	var iResult int
	var exit bool
	p1 := tea.NewProgram(initModel(str, choices, &iResult, &exit))
	_, err := p1.Run()
	if err != nil {
		return iResult, err
	}
	if exit {
		exitError(nil)
	}
	return iResult, nil
}

func getSeed(isStr bool, is12Word bool) (seed.Seed, error) {
	var n uint8
	if is12Word {
		n = 16
	} else {
		n = 32
	}
	if isStr {
		return loopStrInput(n)
	} else {
		return loopByteInput(n)
	}
}

func exitError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

func main() {
	r1, err := getOptions("What is your entropy type?", []string{"String", "Byte"})
	if err != nil {
		exitError(err)
	}
	var isStr bool
	if r1 == 0 {
		isStr = true
	}
	r2, err := getOptions("What is your seed length?", []string{"12 word", "24 word"})
	if err != nil {
		exitError(err)
	}
	var is12Word bool
	if r2 == 0 {
		is12Word = true
	}

	rSeed, err := getSeed(isStr, is12Word)
	if err != nil {
		exitError(err)
	}

	stdout.PrintSeedSignerQRCode(rSeed)
	fmt.Println("######## SeedWord ########")
	for i, v := range rSeed.GetWords() {
		fmt.Printf("(%d) %s\n", i+1, v)
	}
	fmt.Println("######## SeedWord ########")
}
