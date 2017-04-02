package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"github.com/vektra/mockery/mockery"
	"os"
	"regexp"
	"runtime/pprof"
	"strings"
)

const regexMetadataChars = "\\.+*?()|[]{}^$"

type Config struct {
	fName      string
	fPrint     bool
	fOutput    string
	fOutpkg    string
	fDir       string
	fRecursive bool
	fAll       bool
	fIP        bool
	fTO        bool
	fCase      string
	fNote      string
	fProfile   string
	fVersion   bool
}

func main() {
	app := cli.NewApp()
	app.Name = mockery.Name
	app.Version = mockery.SemVer
	app.Usage = mockery.Usage

	var config Config
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "name or matching regular expression of interface to generate mock for",
			Destination: &config.fName,
		},
		cli.StringFlag{
			Name:        "cpuprofile",
			Usage:       "write cpu profile to file",
			Destination: &config.fProfile,
		},
		cli.StringFlag{
			Name:        "note",
			Usage:       "comment to insert into prologue of each generated file",
			Destination: &config.fNote,
		},
		cli.BoolFlag{
			Name:        "print",
			Usage:       "print the generated mock to stdout",
			Destination: &config.fPrint,
		},
		cli.StringFlag{
			Name:        "case",
			Usage:       "name the mocked file using casing convention",
			Value:       "camel",
			Destination: &config.fCase,
		},
		cli.StringFlag{
			Name:        "output",
			Usage:       "directory to write mocks to",
			Value:       "./mocks",
			Destination: &config.fOutput,
		},
		cli.StringFlag{
			Name:        "outpkg",
			Usage:       "name of generated package",
			Value:       "mocks",
			Destination: &config.fOutpkg,
		},
		cli.StringFlag{
			Name:        "dir",
			Usage:       "directory to search for interfaces",
			Value:       ".",
			Destination: &config.fDir,
		},
		cli.BoolFlag{
			Name:        "recursive",
			Usage:       "recurse search into sub-directories",
			Destination: &config.fRecursive,
		},
		cli.BoolFlag{
			Name:        "all",
			Usage:       "generates mocks for all found interfaces in all sub-directories",
			Destination: &config.fAll,
		},
		cli.BoolFlag{
			Name:        "inpkg",
			Usage:       "generate a mock that goes inside the original package",
			Destination: &config.fIP,
		},
		cli.BoolFlag{
			Name:        "testonly",
			Usage:       "generate a mock in a _test.go file",
			Destination: &config.fTO,
		},
	}

	app.Action = func(c *cli.Context) error {
		return run(c, &config)
	}

	app.Run(os.Args)
}

func run(c *cli.Context, config *Config) error {
	var recursive bool
	var filter *regexp.Regexp
	var err error
	var limitOne bool

	if config.fName != "" && config.fAll {
		return errors.New("Specify -name or -all, but not both")
	} else if config.fName != "" {
		recursive = config.fRecursive
		if strings.ContainsAny(config.fName, regexMetadataChars) {
			if filter, err = regexp.Compile(config.fName); err != nil {
				return errors.New("Invalid regular expression provided to -name")
			}
		} else {
			filter = regexp.MustCompile(fmt.Sprintf("^%s$", config.fName))
			limitOne = true
		}
	} else if config.fAll {
		recursive = true
		filter = regexp.MustCompile(".*")
	} else {
		return errors.New("Use -name to specify the name of the interface or -all for all interfaces found")
	}

	if config.fProfile != "" {
		f, err := os.Create(config.fProfile)
		if err != nil {
			return err
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var osp mockery.OutputStreamProvider
	if config.fPrint {
		osp = &mockery.StdoutStreamProvider{}
	} else {
		osp = &mockery.FileOutputStreamProvider{
			BaseDir:   config.fOutput,
			InPackage: config.fIP,
			TestOnly:  config.fTO,
			Case:      config.fCase,
		}
	}

	visitor := &mockery.GeneratorVisitor{
		InPackage:   config.fIP,
		Note:        config.fNote,
		Osp:         osp,
		PackageName: config.fOutpkg,
	}

	walker := mockery.Walker{
		BaseDir:   config.fDir,
		Recursive: recursive,
		Filter:    filter,
		LimitOne:  limitOne,
	}

	generated := walker.Walk(visitor)

	if config.fName != "" && !generated {
		return fmt.Errorf("Unable to find %s in any go files under this path", config.fName)
	}

	return nil
}
