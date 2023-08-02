package tflags

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

const BadExitCode = 1
const HelpExitCode = 0

type flag struct {
	pbool *bool
	pint *int
	pstring *string
}

type byShort struct { Meta []*Meta }

func (m byShort) Len() int {
	return len(m.Meta)
}

func (m byShort) Swap(i, j int) {
	m.Meta[i], m.Meta[j] = m.Meta[j], m.Meta[i]
}

func (m byShort) Less(i, j int) bool {
	return m.Meta[i].Short < m.Meta[j].Short
}

type cmdfn func([]string)

type cmd struct {
	fn cmdfn
	help string
}

var(
	cmds = map[string]cmd{}
	flags = map[string]flag{}
	cmdMetas = []*cmdMeta{}
	flagMetas = []*Meta{}
	unmatched = []string{}

	About = ""
)

type Meta struct {
	Long string
	Short string
	Help string
}

type cmdMeta struct {
	Name string
	Help string
}


func String(s *string, meta *Meta) {
	flagMetas = append(flagMetas, meta)
	if meta.Short != "" {
		flags["-" + meta.Short] = sflag(s)
	}
	if meta.Long != "" {
		flags["--" + meta.Long] = sflag(s)
	}
}

func Int(i *int, meta *Meta) {
	flagMetas = append(flagMetas, meta)
	if meta.Short != "" {
		flags["-" + meta.Short] = iflag(i)
	}
	if meta.Long != "" {
		flags["--" + meta.Long] = iflag(i)
	}
}

func Bool(b *bool, meta *Meta) {
	flagMetas = append(flagMetas, meta)
	if meta.Short != "" {
		flags["-" + meta.Short] = bflag(b)
	}
	if meta.Long != "" {
		flags["--" + meta.Long] = bflag(b)
	}
}

func Cmd(fn cmdfn, name, help string) {
	cmds[name] = cmd{
		fn: fn,
		help: help,
	}
	cmdMetas = append(cmdMetas, &cmdMeta{
		Name: name,
		Help: help,
	})
}

func Parse() *cmdfn {
	return ParseThem(os.Args, true)
}

func ParseThem(args []string, exitOnHelp bool) *cmdfn {
	defer func(){
		flagMetas = flagMetas[:0]
		flags = make(map[string]flag)
	}()
	var err error
	n := len(args)

	for i := 0; i < n; i++ {
		arg := args[i]
		checkHelp(arg, exitOnHelp)
		flag, ok := flags[arg]
		if !ok {
			cmd, ok := cmds[arg]
			if ok {
				unmatched = args[i:]
				return &cmd.fn
			}
			unmatched = append(unmatched, arg)
			continue
		} 

		if flag.pbool != nil {
			*flag.pbool = true
			continue
		}

		next := nextArg(i, args)
		i++

		if flag.pstring != nil {
			*flag.pstring = next
		}

		if flag.pint != nil {
			*flag.pint, err = strconv.Atoi(next)
			if err != nil {
				fmt.Fprintf(os.Stderr, 
					"Argument '%s' expects an integer, provided: '%s'\n", 
					arg,
					next,
				);
				os.Exit(BadExitCode)
			}
		}
	}
	return nil
}

func nextArg(current int, all []string) string {
	next := current + 1 
	if len(all) <= next {
		fmt.Fprintf(os.Stderr, 
			"Argument '%s' expects an argument, none provided\n", 
			all[current],
		);
		os.Exit(BadExitCode)
	}
	return all[next]
}

func checkHelp(arg string, exit bool) {
	helpFlags := []string{
		"-h",
		"-help",
		"--help",
	}
	for _, h := range helpFlags {
		if arg == h {
			Help()
			if exit {
				os.Exit(HelpExitCode)
			}
			return
		}
	}
}

func Help() {
	HelpForeword(About)
}

func HelpForeword(foreword string) {
	sort.Sort(byShort{flagMetas})
	max := 0
	for _, m := range flagMetas {
		if len(m.Long) > max {
			max = len(m.Long)
		}
	}
	for _, m := range flagMetas {
		fmt.Fprintf(os.Stderr, "  -%s", m.Short)
		offset := len(m.Long)
		if m.Long != "" {
			fmt.Fprintf(os.Stderr, ", --%s", m.Long)
			offset += 6
		} else {
			offset += 2
		}
		pad(8 + max - offset, os.Stderr)
		fmt.Fprintf(os.Stderr, "%s\n", m.Help)
	}
}

func pad(n int, w io.Writer) {
	for i := 0; i < n; i++ {
		fmt.Fprintf(w, " ")
	}
}

func Unmatched() []string {
	return unmatched
}

func bflag(b *bool) flag {
	return flag{
		pbool: b,
		pint: nil,
		pstring: nil,
	}
}

func iflag(i *int) flag {
	return flag{
		pbool: nil,
		pint: i,
		pstring: nil,
	}
}

func sflag(s *string) flag {
	return flag{
		pbool: nil,
		pint: nil,
		pstring: s,
	}
}
