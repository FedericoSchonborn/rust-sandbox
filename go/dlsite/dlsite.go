// Package dlsite implements a parser for golang.org/dl.
package dlsite

type Version struct {
	Name  string
	Files []*File
}

func (v *Version) Source() *File {
	for _, file := range v.Files {
		if file.Kind == FileKindSource {
			return file
		}
	}

	return nil
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
	OSMacOS   OS = "macOS"
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
	ArchX86_64 Arch = "x86-64"
	ArchX86    Arch = "x86"
	ArchARMv8  Arch = "ARMv8"
	ArchARMv6  Arch = "ARMv6"
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
