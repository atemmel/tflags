package tflags

import "os"

type flag struct {
	pbool *bool
	pint *int
	pstring *string
}

var(
	flags = map[string]flag{}
	unmatched = []string{}
)

type Meta struct {
	Long string
	Short string
	//TODO: present help
	Help string
}

func String(s *string, meta Meta) {
	if meta.Short != "" {
		flags["-" + meta.Short] = sflag(s)
	}
	if meta.Long != "" {
		flags["--" + meta.Long] = sflag(s)
	}
}

func Int(i *int, meta Meta) {
	if meta.Short != "" {
		flags["-" + meta.Short] = iflag(i)
	}
	if meta.Long != "" {
		flags["--" + meta.Long] = iflag(i)
	}
}

func Bool(b *bool, meta Meta) {
	if meta.Short != "" {
		flags["-" + meta.Short] = bflag(b)
	}
	if meta.Long != "" {
		flags["--" + meta.Long] = bflag(b)
	}
}

func Parse() {
	ParseThem(os.Args)
}

func ParseThem(args []string) {
	n := len(args)
	for i := 0; i < n; i++ {
		arg := args[i]
		flag, ok := flags[arg]
		if !ok {
			unmatched = append(unmatched, arg)
			continue
		} 

		if flag.pbool != nil {
			*flag.pbool = true
			continue
		}

		if flag.pstring != nil {
			//TODO: range check
		}

		if flag.pint != nil {
			//TODO: range check
		}
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
