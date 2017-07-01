# Gojekyll
[![Build Status](https://travis-ci.org/osteele/gojekyll.svg?branch=master)](https://travis-ci.org/osteele/gojekyll)
[![Go Report Card](https://goreportcard.com/badge/github.com/osteele/gojekyll)](https://goreportcard.com/report/github.com/osteele/gojekyll)

Gojekyll is an incomplete implementation of the [Jekyll](https://jekyllrb.com) static site generator, in the [Go](https://golang.org) programming language.

Missing features:

- Themes, page tags, excerpts, plugins (except for `avatar`), and pagination
- Site variables: `pages`, `static_files`, `html_pages`, `html_files`, `documents`, and `tags`
- Jekyll's `group_by_exp`, `pop`, `shift`, `cgi_escape`, `uri_escape`, `scssify`, and `smartify` filters
- Jekyll's `include_relative`, `post_url`, `gist`, and `highlight` tags
- The [Go Liquid template engine](https://github.com/osteele/gojekyll) is also missing some tags and filters.
- Data files must be YAML; CSV and JSON are not supported.
- `{% highlight %}` uses Pygments. There's no way to tell it to use Rouge.
- `<div markdown=1>` doesn't work. I think this is a limitation of the *blackfriday* Markdown processor.
- Parse errors aren't reported very nicely.

Other differences from Jekyll:

- `serve` generates pages on the fly; it doesn't write to the file system.
- Files are cached to `/tmp/gojekyll-${USER}`.
- No `.sass-cache`; therefore, no `sass: cache: false` option.
- Server live reload is always on.
- The server reloads the `_config.yml` file too.

## Installation

1. [Install go](https://golang.org/doc/install#install). On macOS running Homebrew, `brew install go` is easier.
2. `go get -u osteele/gojekyll/cmd/gojekyll`
3. You need to install Pygments `pip install Pygments` to use the `{% highlight %}` tag.

## Usage

```bash
gojekyll -s path/to/site build                # builds into ./_site
gojekyll -s path/to/site serve                # serves from memory, w/ live reload
gojekyll help
gojekyll help build
```

## Status

- [ ] Content
  - [x] Front Matter
  - [ ] Posts
    - [x] Categories
    - [ ] Tags
    - [x] Drafts
    - [x] Future
    - [x] Related
  - [x] Static Files
  - [x] Variables
  - [x] Collections
  - [ ] Data Files
    - [ ] CSV
    - [ ] JSON
    - [x] YAML
  - [ ] Assets
    - [ ] Coffeescript
    - [x] Sass/SCSS
- [ ] Customization
  - [x] Templates
    - [ ] Jekyll filters
      - [ ] `group_by_exp` `pop` `shift` `cgi_escape` `uri_escape` `scssify` `smartify`
      - [x] everything else
    - [ ] Jekyll tags
      - [x] `include`
      - [ ] `include_relative`
      - [x] `link`
      - [x] `post_url`
      - [ ] `gist`
      - [x] `highlight`
  - [x] Includes
      - [x] `include` parameters
      - [x] `include` variables (e.g. `{% include {{ expr }} %}`)
  - [x] Permalinks
  - [ ] Pagination
  - [ ] Plugins
    - [x] `jekyll-avatar`
    - [ ] `jekyll-coffeescript`
    - [x] `jekyll-live-reload` (always on)
    - [ ] `jekyll-paginate`
  - [ ] Themes
  - [x] Layouts
- [x] Server
  - [x] Directory watch
- [ ] Commands
  - [x] `build`
    - [x] `--source`, `--destination`, `--drafts`, `--future`, `--unpublished`
    - [ ] `--config`, `--baseurl`, `--lsi`, etc.
  - [x] `clean`
  - [ ] `doctor`
  - [x] `help`
  - [ ] `import`
  - [ ] `new`
  - [ ] `new-theme`
  - [x] `serve`
    - [x] `--open-uri`
    - [ ] `--detach`, `--host`, `--port`, etc.
- [ ] Windows

## Contributing

Install package dependencies and development tools:

```bash
make setup
```

### Testing

```bash
make test
make lint
gojekyll  -s path/to/site render index.md      # render a file to stdout
gojekyll  -s path/to/site render /             # render a URL to stdout
gojekyll  -s path/to/site variables /          # print a file or URL's variables
./scripts/coverage && go tool cover -html=coverage.out
```

`./scripts/gojekyll` is an alternative to the `gojekyll` executable, that uses `go run` each time it's invoked.

### Profiling

```bash
gojekyll -s path/to/site profile
go tool pprof gojekyll gojekyll.prof
```

## Credits

Gojekyll uses these libraries:

| Package | Author | Description |
| --- | --- | --- |
| [github.com/jaschaephraim/lrserver](https://github.com/jaschaephraim/lrserver) | Jascha Ephraim | Live Reload server |
| [github.com/osteele/liquid](https://github.com/osteele/liquid) | Oliver Steele | Liquid processor |
| [github.com/pkg/browser](https://github.com/pkg/browser) | [pkg](https://github.com/pkg) | The `serve -o` option to open the site in the browser |
| [github.com/russross/blackfriday](https://github.com/russross/blackfriday) | Russ Ross | Markdown processor |
| [github.com/sass/libsass](https://github.com/sass/libsass) | Listed [here](https://https://github.com/sass/libsass) | C port of the Ruby SASS compiler |
| [github.com/wellington/go-libsass](https://github.com/wellington/go-libsass) | Drew Wells | Go bindings to libsass |
| [gopkg.in/alecthomas/kingpin.v2](https://github.com/alecthomas/kingpin)  | Alec Thomas | command line and flag parser |
| [gopkg.in/yaml.v2](https://github.com/go-yaml) | Canonical | YAML support |

In addition to being totally and obviously inspired by the Jekyll, Jekyll's solid documentation was indispensible. Many of the filter test cases are taken directly from the Jekyll documentation, and the [Jekyll docs](https://jekyllrb.com/docs/home/) were always open in at least one tab.

The help text was taken from the output of `jekyll help`.

The gopher image in the test directory is from [Wikimedia Commons](https://commons.wikimedia.org/wiki/File:Gophercolor.jpg). It is used under the [Creative Commons Attribution-Share Alike 3.0 Unported license](https://creativecommons.org/licenses/by-sa/3.0/deed.en).

## Related

[Hugo](https://gohugo.io) isn't Jekyll-compatible (-), but actually works (+++).

[Jekyll](https://jekyllrb.com), of course.

## License

MIT

## Alternate Naming Possibilities

* "Gekyll". (Hard or soft "g"? See [gif](https://en.wikipedia.org/wiki/GIF#Pronunciation_of_GIF).)
* "Gekko"
