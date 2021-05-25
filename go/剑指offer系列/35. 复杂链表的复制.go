package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodes := make([]*Node, 0)
	nodes2 := make([]*Node, 0)
	for head != nil {
		nodes = append(nodes, head)
		nodes2 = append(nodes2, &Node{Val:head.Val})
		head = head.Next
	}
	indexes := make([]int, len(nodes))
	for i:=0; i< len(nodes); i++ {
		random := nodes[i].Random
		if random == nil {
			indexes[i] = -1
			continue
		}
		for j :=0; j< len(nodes); j++ {
			if random == nodes[j] {
				indexes[i] = j
				break
			}
		}
	}
	for k,v :=range nodes2 {
		if k> 0 {
			nodes2[k-1].Next = v
		}
		if indexes[k] > -1 {
			v.Random = nodes2[indexes[k]]
		}
	}
	return nodes2[0]
}
