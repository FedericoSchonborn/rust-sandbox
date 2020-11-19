// Package dlsite implements a parser for golang.org/dl.
package dlsite

type Index struct {
	Stable  *Version
	Archive []*Version
}

type Version struct {
	Name  string
	Files Files
}

type Files []*File

func (fs Files) Source() *File {
	for _, file := range fs {
		if file.Kind == FileKindSource {
			return file
		}
	}

	return nil
}

func (fs Files) Archives() Files {
	return fs.filter(func(f *File) bool {
		return f.Kind == FileKindArchive
	})
}

func (fs Files) Installers() Files {
	return fs.filter(func(f *File) bool {
		return f.Kind == FileKindInstaller
	})
}

func (fs Files) filter(fn func(*File) bool) Files {
	var nfs Files
	for _, file := range fs {
		if fn(file) {
			nfs = append(nfs, file)
		}
	}

	return nfs
}

type File struct {
	Name     string
	URL      string
	Kind     FileKind
	OS       OS
	Arch     Arch
	Size     int
	Checksum string
}

type OS string

// Operating systems as displayed on the site, may be extended.
const (
	OS4       OS = "4" // Used by Go 1.4
	OSMacOS   OS = "macOS"
	OSOSX106  OS = "OS X 10.6+"
	OSOSX108  OS = "OS X 10.8+"
	OSWindows OS = "Windows"
	OSLinux   OS = "Linux"
	OSFreeBSD OS = "FreeBSD"
)

func (os OS) String() string {
	return string(os)
}

// Architectures as displayed on the site, may be extended.
type Arch string

const (
	ArchBootstrap Arch = "bootstrap" // Used by Go 1.4
	ArchX86_64    Arch = "x86-64"
	ArchX86       Arch = "x86"
	ArchARMv8     Arch = "ARMv8"
	ArchARMv6     Arch = "ARMv6"
	ArchPPC64LE   Arch = "ppc64le"
	ArchS390X     Arch = "s390x"
)

func (a Arch) String() string {
	return string(a)
}

type FileKind string

// File kinds as displayed on the site, may be extended.
const (
	FileKindSource    FileKind = "Source"
	FileKindArchive   FileKind = "Archive"
	FileKindInstaller FileKind = "Installer"
)

func (fk FileKind) String() string {
	return string(fk)
}
