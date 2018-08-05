package menu
import (
	"errors"
)
type Node struct {
	Id	int 
	Pid 	int
	Name 	string
	Controller 	string
	Method 		string	
	Childer		[]Node
}
type menu struct {
	node []Node
}

func NewMenu() *menu {
	return &menu{}
}

func (this *menu) Put (data []Node) error {
	if (len(data)<=0) {
		return errors.New("Data cannot be empty")
	}
	
	this.node = append(this.node,data...)
	return nil
}

func (this *menu) Tree() []Node {
	tree := recursion(this.node,0)
	return tree
}

func recursion(data []Node,pid int) []Node {
	tree := make([]Node,0)
	count := len(data)
	if (count<1) {
		return tree
	}
        for _,v := range data {
		if (v.Pid == pid) {
			v.Childer = append(v.Childer,recursion(data,v.Id)...) 
			tree = append(tree,v)		
		}
	} 
	return tree;
}
