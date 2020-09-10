package exec

import (
	"context"
	"io"
	"os/exec"
)

type Builder struct {
	inner *exec.Cmd
}

func NewBuilder(name string, args ...string) *Builder {
	return &Builder{inner: exec.Command(name, args...)}
}

func NewBuilderContext(ctx context.Context, name string, args ...string) *Builder {
	return &Builder{inner: exec.CommandContext(ctx, name, args...)}
}

func (b *Builder) Arg(arg string) *Builder {
	b.inner.Args = append(b.inner.Args, arg)
	return b
}

func (b *Builder) Args(args ...string) *Builder {
	b.inner.Args = append(b.inner.Args, args...)
	return b
}

func (b *Builder) Dir(dir string) *Builder {
	b.inner.Dir = dir
	return b
}

func (b *Builder) Env(key, value string) *Builder {
	b.inner.Env = append(b.inner.Env, key+"="+value)
	return b
}

func (b *Builder) Envs(envs map[string]string) *Builder {
	for key, value := range envs {
		b.inner.Env = append(b.inner.Env, key+"="+value)
	}
	return b
}

func (b *Builder) Stderr(w io.Writer) *Builder {
	b.inner.Stderr = w
	return b
}

func (b *Builder) Stdout(w io.Writer) *Builder {
	b.inner.Stdout = w
	return b
}

func (b *Builder) Stdin(r io.Reader) *Builder {
	b.inner.Stdin = r
	return b
}

func (b *Builder) Finish() *exec.Cmd {
	return b.inner
}
