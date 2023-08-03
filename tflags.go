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

var (
	cmds = map[string]cmd{}
	cmdMetas = []*cmdMeta{}
	flags = map[string]flag{}
	flagMetas = []*Meta{}
	unmatched = []string{}
	About = ""
)

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

func Cmd(fn func([]string), name, help string) {
	cmds[name] = cmd{
		fn: fn,
		help: help,
	}
	cmdMetas = append(cmdMetas, &cmdMeta{
		Name: name,
		Help: help,
	})
}

func Parse() *func([]string) {
	return ParseThem(os.Args, true)
}

func ParseThem(args []string, exitOnHelp bool) *func([]string) {
	defer func(){
		cmds = map[string]cmd{}
		cmdMetas = cmdMetas[:0]
		flags = map[string]flag{}
		flagMetas = flagMetas[:0]
	}()

	unmatched = unmatched[:0]
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
	if About != "" {
		fmt.Fprintln(os.Stderr, About)
	}
	helpCmds()
	if len(cmdMetas) > 0 {
		fmt.Fprintln(os.Stderr)
	}
	helpFlags()
}

func helpCmds() {
	sort.Sort(byName{cmdMetas})
	max := 0
	for _, m := range cmdMetas {
		if len(m.Name) > max {
			max = len(m.Name)
		}
	}
	for _, m := range cmdMetas {
		fmt.Fprintf(os.Stderr, "  %s", m.Name)
		offset := len(m.Name)
		pad(2 + max - offset, os.Stderr)
		fmt.Fprintf(os.Stderr, "%s\n", m.Help)
	}
}

func helpFlags() {
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
