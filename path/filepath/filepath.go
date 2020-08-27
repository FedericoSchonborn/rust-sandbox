package filepath

import (
	"path/filepath"
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

type FilePath string

func New(path string) FilePath {
	return FilePath(path)
}

func FromSlash(path string) FilePath {
	return FilePath(filepath.FromSlash(path))
}

func Glob(pattern string) (matches []FilePath, err error) {
	vs, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	for _, v := range vs {
		matches = append(matches, FilePath(v))
	}

	return matches, nil
}

func (fp FilePath) Abs() (FilePath, error) {
	v, err := filepath.Abs(string(fp))
	return FilePath(v), err
}

func (fp FilePath) Base() string {
	return filepath.Base(string(fp))
}

func (fp FilePath) Clean() FilePath {
	return FilePath(filepath.Clean(string(fp)))
}

func (fp FilePath) Dir() FilePath {
	return FilePath(filepath.Dir(string(fp)))
}

func (fp FilePath) EvalSymlinks() (FilePath, error) {
	v, err := filepath.EvalSymlinks(string(fp))
	return FilePath(v), err
}

func (fp FilePath) Ext() string {
	return filepath.Ext(string(fp))
}

func (fp FilePath) IsAbs() bool {
	return filepath.IsAbs(string(fp))
}

func (fp FilePath) Join(elem ...string) FilePath {
	return FilePath(filepath.Join(append([]string{string(fp)}, elem...)...))
}

func (fp FilePath) Match(pattern string) (matched bool, err error) {
	return filepath.Match(pattern, string(fp))
}

func (fp FilePath) Rel(target FilePath) (FilePath, error) {
	v, err := filepath.Rel(string(fp), string(target))
	return FilePath(v), err
}

func (fp FilePath) Split() (dir FilePath, file string) {
	sdir, sfile := filepath.Split(string(fp))
	return FilePath(sdir), sfile
}

func (fp FilePath) ToSlash() string {
	return filepath.ToSlash(string(fp))
}

func (fp FilePath) VolumeName() string {
	return filepath.VolumeName(string(fp))
}

func (fp FilePath) Walk(fn WalkFunc) error {
	return filepath.Walk(string(fp), fn)
}

func (fp FilePath) String() string {
	return string(fp)
}

type List []FilePath

func NewList(list string) List {
	fpaths := filepath.SplitList(list)
	if len(fpaths) == 0 {
		return List{}
	}

	v := make(List, len(fpaths))
	for i, fpath := range fpaths {
		v[i] = FilePath(fpath)
	}

	return v
}
