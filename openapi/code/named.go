package code

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

var reservedWords = []string{
	"break", "default", "func", "interface", "select", "case", "defer", "go",
	"map", "struct", "chan", "else", "goto", "switch", "const", "fallthrough",
	"if", "range", "type", "continue", "for", "import", "return", "var",
	"append", "bool", "byte", "iota", "len", "make", "new", "package",
}

// Named holds common methods for identifying and describing things
type Named struct {
	Name        string
	Description string
}

func (n *Named) IsNameReserved() bool {
	for _, v := range reservedWords {
		if n.CamelName() == v {
			return true
		}
	}
	return false
}

func (n *Named) isNamePlural() bool {
	if n.Name == "" {
		return false
	}
	return n.Name[len(n.Name)-1] == 's'
}

// simplified ruleset for singularizing multiples
var singularizers = []*regexTransform{
	// branches -> branch
	newRegexTransform("(s|ss|sh|ch|x|z)es$", "$1"),

	// policies -> policy
	newRegexTransform("([bcdfghjklmnpqrstvwxz])ies$", "${1}y"),

	// permissions -> permission
	newRegexTransform("([a-z])s$", "$1"),
}

var singularExceptions = map[string]string{
	"dbfs":       "dbfs",
	"warehouses": "warehouse",
	"databricks": "databricks",
}

type regexTransform struct {
	Search  *regexp.Regexp
	Replace string
}

func newRegexTransform(search, replace string) *regexTransform {
	return &regexTransform{
		Search:  regexp.MustCompile(search),
		Replace: replace,
	}
}

func (t *regexTransform) Transform(src string) string {
	return t.Search.ReplaceAllString(src, t.Replace)
}

func (n *Named) Singular() *Named {
	if !n.isNamePlural() {
		return n
	}
	exception, ok := singularExceptions[strings.ToLower(n.Name)]
	if ok {
		return &Named{
			Name:        exception,
			Description: n.Description,
		}
	}
	for _, t := range singularizers {
		after := t.Transform(n.Name)
		if after == n.Name {
			continue
		}
		return &Named{
			Name:        after,
			Description: n.Description,
		}
	}
	return n
}

// emulate positive lookaheads from JVM regex:
// (?<=[a-z])(?=[A-Z])|(?<=[A-Z])(?=[A-Z][a-z])|([-_\s])
// and convert all words to lower case
func (n *Named) splitASCII() (w []string) {
	var current []rune
	nameLen := len(n.Name)
	var last, this, next, lookahead bool
	// we do assume here that all named entities are strictly ASCII
	for i := 0; i < nameLen; i++ {
		r := rune(n.Name[i])
		if r == '$' {
			// we're naming language literals, $ is usually not allowed
			continue
		}
		this = unicode.IsUpper(r)
		r = unicode.ToLower(r)
		next = false
		lookahead = i+1 < nameLen
		if lookahead {
			next = unicode.IsUpper(rune(n.Name[i+1]))
		}
		split, before, after := false, false, true
		if last && this && !next && lookahead {
			// (?<=[A-Z])(?=[A-Z][a-z])
			split = true
			before = false
			after = true
		}
		if !this && next {
			// (?<=[a-z])(?=[A-Z])
			split = true
			before = true
			after = false
		}
		if r == '-' || r == '_' || r == ' ' {
			// ([-_\s])
			split = true
			before = false
			after = false
		}
		if before {
			current = append(current, r)
		}
		if split && len(current) > 0 {
			w = append(w, string(current))
			current = []rune{}
		}
		if after {
			current = append(current, r)
		}
		last = this
	}
	if len(current) > 0 {
		w = append(w, string(current))
	}
	return w
}

// PascalName creates NamesLikesThis
func (n *Named) PascalName() string {
	var sb strings.Builder
	for _, w := range n.splitASCII() {
		sb.WriteString(strings.Title(w))
	}
	return sb.String()
}

// TitleName creates Names Likes This
func (n *Named) TitleName() string {
	return strings.Title(strings.Join(n.splitASCII(), " "))
}

// CamelName creates namesLikesThis
func (n *Named) CamelName() string {
	cc := n.PascalName()
	return strings.ToLower(cc[0:1]) + cc[1:]
}

// SnakeName creates names_like_this
func (n *Named) SnakeName() string {
	return strings.Join(n.splitASCII(), "_")
}

// ConstantName creates NAMES_LIKE_THIS
func (n *Named) ConstantName() string {
	return strings.ToUpper(n.SnakeName())
}

// KebabName creates names-like-this
func (n *Named) KebabName() string {
	return strings.Join(n.splitASCII(), "-")
}

// AbbrName returns `nlt` for `namesLikeThis`
func (n *Named) AbbrName() string {
	var abbr []byte
	for _, v := range n.splitASCII() {
		abbr = append(abbr, v[0])
	}
	return string(abbr)
}

func (n *Named) HasComment() bool {
	return n.Description != ""
}

func (n *Named) sentences() []string {
	if n.Description == "" {
		return []string{}
	}
	norm := whitespace.ReplaceAllString(n.Description, " ")
	trimmed := strings.TrimSpace(norm)
	return strings.Split(trimmed, ". ")
}

// Summary gets the first sentence from the description. Always ends in a dot.
func (n *Named) Summary() string {
	sentences := n.sentences()
	if len(sentences) > 0 {
		return strings.TrimSuffix(sentences[0], ".") + "."
	}
	return ""
}

// match markdown links, ignoring new lines
var markdownLink = regexp.MustCompile(`\[([^\]]+)\]\(([^\)]+)\)`)

func (n *Named) DescriptionWithoutSummary() string {
	sentences := n.sentences()
	if len(sentences) > 1 {
		return strings.Join(sentences[1:], ". ")
	}
	return ""
}

// Comment formats description into language-specific comment multi-line strings
func (n *Named) Comment(prefix string, maxLen int) string {
	if n.Description == "" {
		return ""
	}
	trimmed := strings.TrimSpace(n.Description)
	// collect links, which later be sorted
	links := map[string]string{}
	// safe to iterate and update, as match slice is a snapshot
	for _, m := range markdownLink.FindAllStringSubmatch(trimmed, -1) {
		label := strings.TrimSpace(m[1])
		link := strings.TrimSpace(m[2])
		if !strings.HasPrefix(link, "http") {
			// this condition is here until OpenAPI spec normalizes all links
			continue
		}
		// simplify logic by overriding links in case of duplicates.
		// this will also lead us to alphabetically sort links in the bottom,
		// instead of always following the order they appear in the comment.
		// luckily, this doesn't happen often.
		links[label] = link
		// replace [test](url) with [text]
		trimmed = strings.ReplaceAll(trimmed, m[0], fmt.Sprintf("[%s]", label))
	}
	linksInBottom := []string{}
	for k, v := range links {
		link := fmt.Sprintf("[%s]: %s", k, v)
		linksInBottom = append(linksInBottom, link)
	}
	sort.Strings(linksInBottom)
	description := strings.ReplaceAll(trimmed, "\n\n", " __BLANK__ ")
	var lines []string
	currentLine := strings.Builder{}
	for _, v := range whitespace.Split(description, -1) {
		if v == "__BLANK__" {
			lines = append(lines, currentLine.String())
			lines = append(lines, "")
			currentLine.Reset()
			continue
		}
		if len(prefix)+currentLine.Len()+len(v)+1 > maxLen {
			lines = append(lines, currentLine.String())
			currentLine.Reset()
		}
		if currentLine.Len() > 0 {
			currentLine.WriteRune(' ')
		}
		currentLine.WriteString(v)
	}
	if currentLine.Len() > 0 {
		lines = append(lines, currentLine.String())
		currentLine.Reset()
	}
	if len(linksInBottom) > 0 {
		lines = append(append(lines, ""), linksInBottom...)
	}
	return strings.TrimLeft(prefix, "\t ") + strings.Join(lines, "\n"+prefix)
}
