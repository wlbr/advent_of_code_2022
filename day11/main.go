package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type operation struct {
	cmd  string
	argi int
	args string
}

func NewOperation(op string, arg string) *operation {
	newops := &operation{cmd: op}
	if arg == "old" {
		newops.args = arg
	} else {
		i, e := strconv.Atoi(arg)
		if e != nil {
			log.Fatalf("Error parsing argument '%s': %s", arg, e)
		}
		newops.argi = i
	}

	return newops
}

func (o *operation) String() string {
	if o.args == "" {
		return fmt.Sprintf("new = old %s %d", o.cmd, o.argi)
	}
	return fmt.Sprintf("new = old %s %s", o.cmd, o.args)
}

func (o *operation) Execute(old int) int {
	var arg int

	arg = o.argi
	if o.args == "old" {
		arg = old
	}
	switch o.cmd {
	case "+":
		return old + arg
	case "*":
		return old * arg
	case "-":
		return old - arg
	case "/":
		return old / arg
	default:
		log.Fatalf("Unknown operation '%s'", o.cmd)
	}
	return 0
}

type decision struct {
	cmd         string
	arg         int
	trueaction  int
	falseaction int
}

func NewDecision(dec string, arg, trueaction, falseaction int) *decision {
	newdec := &decision{cmd: dec, arg: arg, trueaction: trueaction, falseaction: falseaction}
	return newdec
}

func (d *decision) String() string {
	return fmt.Sprintf("Test: %s by %d\nIf true: throw to monkey %d\nIf false: throw to monkey %d", d.cmd, d.arg, d.trueaction, d.falseaction)
}

func (d *decision) Execute(old int) bool {
	switch d.cmd {
	case "divisible":
		return old%d.arg == 0
	case "equal":
		return old == d.arg
	case "greater":
		return old > d.arg
	case "less":
		return old < d.arg
	default:
		log.Fatalf("Unknown operation '%s'", d.cmd)
	}
	return false
}

type monkey struct {
	id          int
	items       []int
	operation   *operation
	decision    *decision
	inspections int
}

func (m monkey) String() string {
	s := fmt.Sprintf("Monkey %d: ", m.id)
	for _, i := range m.items {
		s += fmt.Sprintf("%d ", i)
	}
	return s
}

func (m *monkey) ThrowFirst(othermonkey *monkey) {
	othermonkey.items = append(othermonkey.items, m.items[0])
	m.items = m.items[1:]
}

type sortedByActivity []*monkey

func (a sortedByActivity) Len() int {
	return len(a)
}
func (a sortedByActivity) Less(i, j int) bool {
	return a[i].inspections > a[j].inspections
}
func (a sortedByActivity) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (m *monkey) Swap(othermonkey *monkey) bool {
	return m.items[0] < othermonkey.items[0]
}

func readInput(fname string) (monkeys []*monkey) {
	f, err := os.Open(fname)
	if err != nil {
		log.Fatalf("Error opening dataset '%s':  %s", fname, err)
	}
	defer f.Close()

	tmpmonkeys := make(map[int]*monkey)

	scanner := bufio.NewScanner(f)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		m := monkey{}

		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			line = scanner.Text()
		}
		_, end := fmt.Sscanf(line, "Monkey %d:\n", &m.id)
		if end != nil {
			log.Fatal("Error parsing monkey)")
		}

		scanner.Scan()
		line = scanner.Text()
		for _, item := range strings.Split(strings.Split(line, ":")[1], ",") {
			iitem, err := strconv.Atoi(strings.Trim(item, " "))
			if err != nil {
				log.Fatalf("Error parsing item '%s': %s", item, err)
			}
			m.items = append(m.items, iitem)
		}

		scanner.Scan()
		line = scanner.Text()
		var ops, arg string
		_, end = fmt.Sscanf(line, "  Operation: new = old %s %s\n", &ops, &arg)
		if end != nil {
			log.Fatalf("Error parsing monkey: %s", end)
		}
		m.operation = NewOperation(ops, arg)

		scanner.Scan()
		line = scanner.Text()
		var dec string
		var decarg, truetarget, falsetarget int
		_, end = fmt.Sscanf(line, "  Test: %s by %d\n", &dec, &decarg)
		if end != nil {
			log.Fatal("Error parsing monkey", end)
		}

		scanner.Scan()
		line = scanner.Text()
		_, end = fmt.Sscanf(line, "  If true: throw to monkey %d\n", &truetarget)
		if end != nil {
			log.Fatal("Error parsing monkey", end)
		}

		scanner.Scan()
		line = scanner.Text()
		_, end = fmt.Sscanf(line, "  If false: throw to monkey %d\n", &falsetarget)
		if end != nil {
			log.Fatal("Error parsing monkey)")
		}
		m.decision = NewDecision(dec, decarg, truetarget, falsetarget)
		tmpmonkeys[m.id] = &m
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	max := 0
	for _, m := range tmpmonkeys {
		if m.id > max {
			max = m.id
		}
	}
	monkeys = make([]*monkey, max+1)
	for _, m := range tmpmonkeys {
		monkeys[m.id] = m
	}

	return monkeys
}
func monkeybusiness(monkeys []*monkey, relief, rounds int) int {

	var worrylevel int
	bigDivisor := 1
	for _, m := range monkeys {
		bigDivisor *= m.decision.arg
	}

	for i := 0; i < rounds; i++ {
		for _, m := range monkeys {
			for {
				if len(m.items) > 0 {
					m.inspections++
					worrylevel = m.operation.Execute(m.items[0])
					if relief > 1 {
						worrylevel = worrylevel / relief
					} else {
						worrylevel = worrylevel % bigDivisor
					}
					m.items[0] = worrylevel
					item := m.items[0]
					if m.decision.Execute(item) {
						m.ThrowFirst(monkeys[m.decision.trueaction])
					} else {
						m.ThrowFirst(monkeys[m.decision.falseaction])
					}
				} else {
					break
				}
			}
		}
	}
	sort.Sort(sortedByActivity(monkeys))

	return monkeys[0].inspections * monkeys[1].inspections
}

func task1(fname string) int {
	monkeys := readInput(fname)
	return monkeybusiness(monkeys, 3, 20)
}

func task2(fname string) int {
	monkeys := readInput(fname)
	return monkeybusiness(monkeys, 1, 10000)
}

func main() {
	startOverall := time.Now()
	input := "input.txt"
	t1 := task1(input)
	afterTask1 := time.Now()
	t2 := task2(input)
	afterTask2 := time.Now()

	fmt.Println("Task 1 - positions visited by the tail of 2   =  ", t1)
	fmt.Printf("Task 2 - chars are: %d \n\n", t2)

	fmt.Println("Time task 1: ", afterTask1.Sub(startOverall))
	fmt.Println("Time task 2: ", afterTask2.Sub(afterTask1))
	fmt.Println("Total time: ", afterTask2.Sub(startOverall))
}
