package bs

import (
	"github.com/gouniverse/hb"
)

type Breadcrumb struct {
	URL   string
	Name  string
	Class string
	Icon  string
}

func Breadcrumbs(breadcrumbs []Breadcrumb) *hb.Tag {
	ul := hb.NewOL().Class("breadcrumb").Style("margin: 0px;")

	for index, breadcrumb := range breadcrumbs {
		content := hb.NewSpan().HTML(breadcrumb.Name)
		if breadcrumb.URL != "" {
			content = hb.NewHyperlink().Attr("href", breadcrumb.URL)
			if breadcrumb.Icon != "" {
				content.Child(hb.NewSpan().Style("margin-right:4px;").HTML(breadcrumb.Icon))
			}
			content.HTML(breadcrumb.Name)
		}

		active := ""
		if index+1 == len(breadcrumbs) {
			active = "active"
		}

		li := hb.NewLI().Class("breadcrumb-item").Class(active).Child(content)
		ul.Child(li)
	}

	nav := hb.NewNav().Attr("aria-label", "breadcrumb").Child(ul)
	return nav
}
