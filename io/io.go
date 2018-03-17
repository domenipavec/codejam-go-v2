package io

import (
	"bufio"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type TestCaseFunc func(*Input, *Output)

type Parser struct {
	f             TestCaseFunc
	input         *Input
	output        *Output
	compareOutput *CompareOutput

	debugNumber *int
	profile     *bool

	baseFn    string
	inputFn   string
	outputFn  string
	correctFn string
	profileFn string
}

func TestCases(f TestCaseFunc) {
	log.SetFlags(0)

	parser := Parser{
		f:           f,
		debugNumber: flag.Int("d", -1, "Debug output only for this case"),
		profile:     flag.Bool("p", false, "Run profiler"),
	}

	inputFileFlag := flag.String("i", "", "Input file")

	flag.Parse()

	inputFile := ""
	if *inputFileFlag != "" {
		inputFile = *inputFileFlag
	} else if len(flag.Args()) == 1 {
		inputFile = flag.Args()[0]
	} else {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		downloadFiles, err := filepath.Glob(path.Join(usr.HomeDir, "Downloads/*.in"))
		if err != nil {
			log.Fatal(err)
		}
		if len(downloadFiles) > 1 {
			log.Fatal("Multiple input files in Downloads")
		} else if len(downloadFiles) == 1 {
			inputFile = filepath.Base(downloadFiles[0])
			os.Rename(downloadFiles[0], inputFile)
		} else {
			localFiles, err := filepath.Glob("*.in")
			if err != nil {
				log.Fatal(err)
			}
			if len(localFiles) > 1 {
				log.Fatal("Multiple input files found, please specify one!")
			} else if len(localFiles) == 1 {
				inputFile = localFiles[0]
			} else {
				log.Fatal("No input files found")
			}
		}
	}

	if !strings.HasSuffix(inputFile, ".in") {
		if inputFile[len(inputFile)-1] != '.' {
			inputFile += ".in"
		} else {
			inputFile += "in"
		}
	}

	log.Println("Using input file:", inputFile)

	parser.SetFn(inputFile)
	parser.ParseFile()
}

func (parser *Parser) SetFn(inputFn string) {
	parser.inputFn = inputFn
	parser.baseFn = strings.TrimSuffix(inputFn, ".in")
	parser.outputFn = parser.baseFn + ".out"
	parser.correctFn = parser.baseFn + ".correct"
	parser.profileFn = parser.baseFn + ".prof"
}

func (parser *Parser) ParseFile() {
	inputF, err := os.Open(parser.inputFn)
	if err != nil {
		log.Fatalln("Error opening input file:", err)
	}
	defer inputF.Close()

	outputF, err := os.Create(parser.outputFn)
	if err != nil {
		log.Fatalln("Error creating output file:", err)
	}
	defer outputF.Close()

	var outputWriter io.Writer = outputF
	if parser.inputFn == "example.in" {
		outputWriter = io.MultiWriter(outputF, os.Stdout)
	}

	scanner := bufio.NewScanner(inputF)
	scanner.Split(bufio.ScanWords)

	parser.output = newOutput(outputWriter)
	parser.input = newInput(scanner)

	parser.compareOutput = nil
	if _, err := os.Stat(parser.correctFn); err == nil {
		correctF, err := os.Open(parser.correctFn)
		if err != nil {
			log.Fatalln("Error opening correct file:", err)
		}
		defer correctF.Close()

		parser.compareOutput = NewCompareOutput(correctF)
	}

	T := parser.input.Int()

	if *parser.debugNumber != -1 {
		log.SetOutput(ioutil.Discard)
	}

	if *parser.profile {
		f, err := os.Create(parser.profileFn)
		if err != nil {
			log.Fatalln("Error creating profile file:", err)
		}
		pprof.StartCPUProfile(f)
	}

	startTime := time.Now()
	for i := 1; i <= T; i++ {
		if *parser.debugNumber == i {
			log.SetOutput(os.Stderr)
		}
		parser.output.startTime = time.Now()
		parser.runTestCase(i)
		if *parser.debugNumber == i {
			log.SetOutput(ioutil.Discard)
		}
	}
	totalTime := time.Now().Sub(startTime)

	if *parser.profile {
		parser.printProfile()
	}

	for key, timer := range parser.output.timers {
		log.Printf("Timer %s: %v", key, timer.Total)
	}
	log.Println("Total time:", totalTime)
}

func (parser *Parser) runTestCase(i int) {
	warningTimer := time.NewTimer(1 * time.Second)
	stopProfileTimer := time.NewTimer(10 * time.Second)
	periodicPrintTicker := time.NewTicker(1 * time.Second)

	doneChan := make(chan bool)

	go func() {
		parser.output.init(parser.input, i)
		parser.input.init()

		parser.f(parser.input, parser.output)

		if parser.compareOutput != nil && parser.compareOutput.HasOutput(i) {
			parser.output.AssertEqual(string(parser.compareOutput.GetOutput(i)))
		}

		parser.output.flush()
		parser.writeChart(i)
		doneChan <- true
	}()

loop:
	for {
		select {
		case <-warningTimer.C:
			parser.output.Debug("Long calculation")
		case <-periodicPrintTicker.C:
			parser.output.triggerPeriodic()
		case <-stopProfileTimer.C:
			if *parser.profile {
				parser.printProfile()
				log.Fatalln("Profile finished")
			}
		case <-doneChan:
			break loop
		}
	}

	periodicPrintTicker.Stop()
	parser.output.resetPeriodic()
}

func (parser *Parser) printProfile() {
	pprof.StopCPUProfile()
	out, err := exec.Command("go", "tool", "pprof", "-top", os.Args[0], parser.profileFn).CombinedOutput()
	if err != nil {
		log.Fatalln("Error running profile tool:", err)
	}
	parser.output.Debug("CPUProfile:", string(out))

	err = exec.Command("go", "tool", "pprof", "-web", os.Args[0], parser.profileFn).Run()
	if err != nil {
		log.Fatalln("Error running profile tool:", err)
	}
}

func (parser *Parser) writeChart(i int) {
	if len(parser.output.points) == 0 {
		return
	}

	p, err := plot.New()
	if err != nil {
		log.Fatalln("Error creating plot:", err)
	}

	err = plotutil.AddLinePoints(p, "", parser.output.points)
	if err != nil {
		log.Fatalln("Error adding linepoints:", err)
	}

	err = p.Save(4*vg.Inch, 4*vg.Inch, parser.baseFn+strconv.Itoa(i)+".png")
	if err != nil {
		log.Fatalln("Error saving img:", err)
	}

	parser.output.points = parser.output.points[:0]
}
