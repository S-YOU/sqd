package sql_scanner

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Komei22/sqd/detector"
	"github.com/Komei22/sqd/eventor"
	"github.com/Komei22/sql-mask"
)

// SqlScanner struct
type SqlScanner struct {
	detector *detector.Detector
}

// New SqlScanner
func New(d *detector.Detector) *SqlScanner {
	s := &SqlScanner{detector: d}
	return s
}

// Scan sql_scanner
func (s *SqlScanner) Scan(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		s.detection(scanner.Text())
	}
}

func (s *SqlScanner) detection(querylog string) {
	parsedQuery, err := parser.Parse(querylog)
	if err != nil {
		fmt.Println(err)
	}
	suspiciousQuery, err := s.detector.Detect(parsedQuery)
	if err != nil {
		fmt.Println(err)
	}
	if suspiciousQuery != "" {
		eventor.Dump(os.Stdout, []string{suspiciousQuery})
	}
}
