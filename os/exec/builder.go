package exec

import (
	"context"
	"io"
	"os/exec"
)

type Builder struct {
	ctx    context.Context
	name   string
	args   []string
	dir    string
	envs   []string
	stderr io.Writer
	stdout io.Writer
	stdin  io.Reader
}

func NewBuilder(name string) *Builder {
	return NewBuilderContext(context.TODO(), name)
}

func NewBuilderContext(ctx context.Context, name string) *Builder {
	return &Builder{
		ctx:  ctx,
		name: name,
		args: []string{},
		envs: []string{},
	}
}

func (b *Builder) Context(ctx context.Context) *Builder {
	b.ctx = ctx
	return b
}

func (b *Builder) Arg(arg string) *Builder {
	b.args = append(b.args, arg)
	return b
}

func (b *Builder) Args(args []string) *Builder {
	b.args = append(b.args, args...)
	return b
}

func (b *Builder) Dir(dir string) *Builder {
	b.dir = dir
	return b
}

func (b *Builder) Env(key, value string) *Builder {
	b.envs = append(b.envs, key+"="+value)
	return b
}

func (b *Builder) Envs(envs map[string]string) *Builder {
	for key, value := range envs {
		b.envs = append(b.envs, key+"="+value)
	}
	return b
}

func (b *Builder) Stderr(w io.Writer) *Builder {
	b.stderr = w
	return b
}

func (b *Builder) Stdout(w io.Writer) *Builder {
	b.stdout = w
	return b
}

func (b *Builder) Stdin(r io.Reader) *Builder {
	b.stdin = r
	return b
}

func (b *Builder) Build() *exec.Cmd {
	cmd := exec.CommandContext(b.ctx, b.name, b.args...)
	cmd.Dir = b.dir
	cmd.Env = b.envs
	cmd.Stderr = b.stderr
	cmd.Stdout = b.stdout
	cmd.Stdin = b.stdin

	return cmd
}
