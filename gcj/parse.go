package gcj

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tamaxyo/go-utils/log"
	"sync"
)

type Solver func(lines []string) string

func Parse(solver Solver) {
	var err error

	q := flag.String("f", "q.txt", "question file")
	a := flag.String("o", "", "answer file")
	s := flag.Int("s", 1, "number of header lines to be skipped")
	l := flag.Int("l", 1, "number of lines for each problem")
	flag.Parse()

	f, err := os.Open(*q)
	log.CheckFatal(err, "could not open file", *q)
	defer f.Close()
	reader := bufio.NewReaderSize(f, 16)

	var o *os.File
	if *a == "" {
		o = os.Stdout
	} else {
		o, err = os.Create(*a)
		log.CheckFatal(err, "could not create file", *a)
		defer o.Close()
	}

	for i := 0; i < *s; i++ {
		reader.ReadLine()
	}

	ans := make(map[int]string)
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	var cnt int

	for cnt = 1; ; cnt++ {
		lines := make([]string, *l)
		for j := 0; j < *l; j++ {
			lines[j] = ""
			for {
				b, p, err := reader.ReadLine()
				if err != nil {
					goto PRINT // EOF
				}
				lines[j] += string(b)
				if !p {
					break
				}
			}
		}

		wg.Add(1)
		go func(i int, lines []string) {
			mu.Lock()
			ans[i] = solver(lines)
			mu.Unlock()

			wg.Done()
		}(cnt, lines)
	}

PRINT:
	wg.Wait()
	for i := 1; i < cnt; i++ {
		o.WriteString(fmt.Sprintf("Case #%d: %s\n", i, ans[i]))
	}
}

func SplitIntoInts(line, delim string) []int {
	splitted := strings.Split(line, delim)
	values := make([]int, len(splitted))

	for i := 0; i < len(splitted); i++ {
		v, err := strconv.Atoi(splitted[i])
		log.CheckFatal(err, "could not split line into integer values - line:", line, ", delim:", delim)
		values[i] = v
	}
	return values
}
