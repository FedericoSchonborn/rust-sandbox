package filepath

import (
	"path/filepath"
	"strings"
)

const (
	Separator     = filepath.Separator
	ListSeparator = filepath.ListSeparator
)

var (
	ErrBadPattern = filepath.ErrBadPattern
	SkipDir       = filepath.SkipDir
)

type WalkFunc = filepath.WalkFunc

type FilePath struct{ inner string }

func New(path string) FilePath {
	return FilePath{inner: path}
}

func FromSlash(path string) FilePath {
	return New(filepath.FromSlash(path))
}

func Glob(pattern string) (matches []FilePath, err error) {
	vs, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, v := range vs {
		matches = append(matches, New(v))
	}

	return matches, nil
}

func (fp FilePath) Abs() (FilePath, error) {
	v, err := filepath.Abs(fp.inner)
	return New(v), err
}

func (fp FilePath) Base() string {
	return filepath.Base(fp.inner)
}

func (fp FilePath) Clean() FilePath {
	return New(filepath.Clean(fp.inner))
}

func (fp FilePath) Dir() FilePath {
	return New(filepath.Dir(fp.inner))
}

func (fp FilePath) EvalSymlinks() (FilePath, error) {
	v, err := filepath.EvalSymlinks(fp.inner)
	return New(v), err
}

func (fp FilePath) Ext() string {
	return filepath.Ext(fp.inner)
}

func (fp FilePath) IsAbs() bool {
	return filepath.IsAbs(fp.inner)
}

func (fp FilePath) Join(elem ...string) FilePath {
	return New(filepath.Join(append([]string{fp.inner}, elem...)...))
}

func (fp FilePath) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, fp.inner)
}

func (fp FilePath) Rel(target FilePath) (FilePath, error) {
	v, err := filepath.Rel(fp.inner, target.inner)
	return New(v), err
}

func (fp FilePath) Split() (dir FilePath, file string) {
	sdir, sfile := filepath.Split(fp.inner)
	return New(sdir), sfile
}

func (fp FilePath) ToSlash() string {
	return filepath.ToSlash(fp.inner)
}

func (fp FilePath) VolumeName() string {
	return filepath.VolumeName(fp.inner)
}

func (fp FilePath) Walk(fn WalkFunc) error {
	return filepath.Walk(fp.inner, fn)
}

func (fp FilePath) String() string {
	return fp.inner
}

type List []FilePath

func NewList(list string) List {
	fpaths := filepath.SplitList(list)
	if len(fpaths) == 0 {
		return List{}
	}

	v := make(List, len(fpaths))
	for i, fpath := range fpaths {
		v[i] = New(fpath)
	}

	return v
}

func (l List) String() string {
	elems := make([]string, len(l))
	for i, fpath := range l {
		elems[i] = fpath.inner
	}

	return strings.Join(elems, ":")
}
