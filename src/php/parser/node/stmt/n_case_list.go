package stmt

import (
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/position"
	"github.com/VKCOM/noverify/src/php/parser/walker"
)

// CaseList node
type CaseList struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Cases        []node.Node
}

// NewCaseList node constructor
func NewCaseList(Cases []node.Node) *CaseList {
	return &CaseList{
		FreeFloating: nil,
		Cases:        Cases,
	}
}

// SetPosition sets node position
func (n *CaseList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *CaseList) GetPosition() *position.Position {
	return n.Position
}

func (n *CaseList) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *CaseList) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Cases != nil {
		for _, nn := range n.Cases {
			if nn != nil {
				nn.Walk(v)
			}
		}
	}

	v.LeaveNode(n)
}
