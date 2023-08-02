package tflags

type cmd struct {
	fn func([]string)
	help string
}

type cmdMeta struct {
	Name string
	Help string
}

type flag struct {
	pbool *bool
	pint *int
	pstring *string
}

type Meta struct {
	Long string
	Short string
	Help string
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

type byName struct { cmdMetas []*cmdMeta }

func (m byName) Len() int {
	return len(m.cmdMetas)
}

func (m byName) Swap(i, j int) {
	m.cmdMetas[i], m.cmdMetas[j] = m.cmdMetas[j], m.cmdMetas[i]
}

func (m byName) Less(i, j int) bool {
	return m.cmdMetas[i].Name < m.cmdMetas[j].Name
}
