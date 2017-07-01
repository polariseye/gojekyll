package tags

import (
	"fmt"
	"io"

	"github.com/osteele/gojekyll/config"
	"github.com/osteele/liquid"
	"github.com/osteele/liquid/chunks"
)

// A LinkTagHandler given an include tag file name returns a URL.
type LinkTagHandler func(string) (string, bool)

// AddJekyllTags adds the Jekyll tags to the Liquid engine.
func AddJekyllTags(e liquid.Engine, c config.Config, lh LinkTagHandler) {
	tc := tagContext{c, lh}
	e.DefineTag("link", tc.linkTag)
	e.DefineTag("include", tc.includeTag)

	// TODO unimplemented
	e.DefineTag("post_url", MakeUnimplementedTag())
	e.DefineStartTag("highlight", highlightTag)
}

type tagContext struct {
	config config.Config
	lh     LinkTagHandler
}

// MakeUnimplementedTag creates a tag definition that prints a warning the first
// time it's rendered, and otherwise does nothing.
func MakeUnimplementedTag() liquid.TagDefinition {
	warned := false
	return func(_ io.Writer, ctx chunks.RenderContext) error {
		if !warned {
			fmt.Printf("The %q tag has not been implemented. It is being ignored.\n", ctx.TagName())
			warned = true
		}
		return nil
	}
}

func (tc tagContext) linkTag(w io.Writer, ctx chunks.RenderContext) error {
	filename := ctx.TagArgs()
	url, found := tc.lh(filename)
	if !found {
		return fmt.Errorf("missing link filename: %s", filename)
	}
	_, err := w.Write([]byte(url))
	return err
}