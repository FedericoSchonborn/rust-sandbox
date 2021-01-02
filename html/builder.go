package html

type TagKind int

const (
	TagKindMeta TagKind = iota
	TagKindSelfClosing
	TagKindInline
	TagKindBlock
)

type Builder struct {
	tag      string
	kind     TagKind
	children []Element
}

func NewBuilder(tag string, kind TagKind) *Builder {
	return &Builder{
		tag:  tag,
		kind: kind,
	}
}

func (b *Builder) Child(elem Element) *Builder {
	b.children = append(b.children, elem)
	return b
}

func (b *Builder) Children(elems []Element) *Builder {
	b.children = append(b.children, elems...)
	return b
}

func (b *Builder) Render(ctx *Context) error {
	if err := ctx.WriteString("<" + b.tag); err != nil {
		return err
	}

	switch b.kind {
	case TagKindMeta, TagKindSelfClosing, TagKindInline:
		if err := ctx.WriteString("/>"); err != nil {
			return err
		}
	case TagKindBlock:
		if err := ctx.WriteString(">"); err != nil {
			return err
		}

		if len(b.children) > 0 {
			if err := ctx.RenderChildren(b.children); err != nil {
				return err
			}
		}

		if err := ctx.WriteString("</" + b.tag + ">"); err != nil {
			return err
		}
	}

	return nil
}
