package topic

import (
	"app/fragment"
	//"app/fragment/cache"
	//"github.com/alecthomas/log4go"
	"html/template"
)

type NodeInfo struct {
}

func (this *NodeInfo) Render(args *fragment.RenderArgs) template.HTML {
	return ""
}

func (this *NodeInfo) PrepareArgs(args *fragment.RenderArgs) {

}
