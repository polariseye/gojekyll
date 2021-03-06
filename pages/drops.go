package pages

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"

	"github.com/osteele/gojekyll/templates"
	"github.com/osteele/gojekyll/utils"
)

// ToLiquid is part of the liquid.Drop interface.
func (d *StaticFile) ToLiquid() interface{} {
	return map[string]interface{}{
		"name":          d.relpath,
		"basename":      utils.TrimExt(d.relpath),
		"path":          d.Permalink(),
		"modified_time": d.fileModTime,
		"extname":       d.OutputExt(),
	}
}

func (f *file) ToLiquid() interface{} {
	var (
		relpath = "/" + filepath.ToSlash(f.relpath)
		base    = path.Base(relpath)
		ext     = path.Ext(relpath)
	)

	return templates.MergeVariableMaps(f.frontMatter, map[string]interface{}{
		"path":          relpath,
		"modified_time": f.fileModTime,
		"name":          base,
		"basename":      utils.TrimExt(base),
		"extname":       ext,
	})
}

// ToLiquid is in the liquid.Drop interface.
func (p *page) ToLiquid() interface{} {
	var (
		relpath = p.relpath
		ext     = filepath.Ext(relpath)
		root    = utils.TrimExt(p.relpath)
		base    = filepath.Base(root)
		content = p.maybeContent(true)
	)
	data := map[string]interface{}{
		"content": content,
		"excerpt": p.excerpt(),
		"path":    relpath,
		"url":     p.Permalink(),
		// TODO output

		// not documented, but present in both collection and non-collection pages
		"permalink": p.Permalink(),

		// TODO only in non-collection pages:
		"dir":  "/" + path.Dir(relpath),
		"name": path.Base(relpath),
		// TODO next previous

		// TODO Documented as present in all pages, but de facto only defined for collection pages
		"id": base,
		// "title": base, // TODO capitalize
		// TODO excerpt category? categories tags
		// TODO slug
		"categories": p.Categories(),
		"tags":       p.Tags(),

		// TODO Only present in collection pages https://jekyllrb.com/docs/collections/#documents
		"relative_path": filepath.ToSlash(p.site.RelativePath(p.filename)),
		// TODO collection(name)

		// TODO undocumented; only present in collection pages:
		"ext": ext,
	}
	for k, v := range p.frontMatter {
		switch k {
		// doc implies these aren't present, but they appear to be present in a collection page:
		// case "layout", "published":
		case "permalink":
		// omit this, in order to use the value above
		default:
			data[k] = v
		}
	}
	return data
}

func (p *page) maybeContent(fallback bool) []byte {
	p.Lock()
	defer p.Unlock()
	if p.content != nil {
		return *p.content
	}
	if fallback {
		return p.raw
	}
	return nil
}

func (p *page) excerpt() []byte {
	if ei, ok := p.frontMatter["excerpt"]; ok {
		return []byte(fmt.Sprint(ei))
	}
	content := p.maybeContent(true)
	pos := bytes.Index(content, []byte(p.site.Config().ExcerptSeparator))
	if pos >= 0 {
		content = content[:pos]
	}
	return content
}

// MarshalYAML is part of the yaml.Marshaler interface
// The variables subcommand uses this.
func (f *file) MarshalYAML() (interface{}, error) {
	return f.ToLiquid(), nil
}

// MarshalYAML is part of the yaml.Marshaler interface
// The variables subcommand uses this.
func (p *page) MarshalYAML() (interface{}, error) {
	return p.ToLiquid(), nil
}
