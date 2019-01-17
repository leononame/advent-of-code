package aoc

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type Config struct {
	File       string
	Input      []string
	Logger     *logrus.Logger
	SubCommand string
	Year, Day  int
	All        bool
	Part       int
}

func (c *Config) SetToDefaultFilePath() {
	c.File = fmt.Sprintf("./input/%d/day%02d", c.Year, c.Day)
}

func (c *Config) ReadFile() error {
	if c.File == "" {
		c.SetToDefaultFilePath()
	}
	s, err := os.Stat(c.File)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("file %s can't be read: is directory", c.File)
	}
	f, err := os.Open(c.File)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	c.Input = c.Input[:0]
	for sc.Scan() {
		c.Input = append(c.Input, sc.Text())
	}
	return nil
}
