package path

import (
	"path"
)

var (
	ErrBadPattern = path.ErrBadPattern
)

type Path struct{ inner string }

func New(path string) Path {
	return Path{inner: path}
}

func (p Path) Base() string {
	return path.Base(p.inner)
}

func (p Path) Clean() Path {
	return New(path.Clean(p.inner))
}

func (p Path) Dir() Path {
	return New(path.Dir(p.inner))
}

func (p Path) Ext() string {
	return path.Ext(p.inner)
}

func (p Path) IsAbs() bool {
	return path.IsAbs(p.inner)
}

func (p Path) Join(elem ...string) Path {
	return New(path.Join(append([]string{p.inner}, elem...)...))
}

func (p Path) Match(pattern string) (matched bool, err error) {
	return path.Match(pattern, p.inner)
}

func (p Path) Split() (dir Path, file string) {
	sdir, sfile := path.Split(p.inner)
	return New(sdir), sfile
}

func (p Path) String() string {
	return p.inner
}
