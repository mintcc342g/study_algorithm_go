package node

// 연결 리스트 예제1
type Node struct {
	data     string
	nextNode *Node
}

func NewNode(data string) *Node {
	return &Node{
		data: data,
	}
}

func (n *Node) addNextNode(nextNode *Node) {
	n.nextNode = nextNode
}

type LinkedList struct {
	firstNode *Node
}

func NewLinkedList(firstNode *Node) *LinkedList {
	return &LinkedList{
		firstNode: firstNode,
	}
}

func (l *LinkedList) read(idx int) string {
	currentNode := l.firstNode
	currentIdx := 0

	for currentIdx < idx { // 찾고 있는 인덱스에 도달할 때까지 노드의 링크를 계속 따라들어감.
		currentNode = currentNode.nextNode
		currentIdx++

		if currentNode == nil { // 만약 현재 노드가 nil 이라면 리스트 끝에 도달했다는 거고,
			return "" // 찾고자 하는 인덱스의 노드를 못 찾았다는 뜻이므로 제로값 반환
		}
	}

	// currentIdx == idx 일 때 for문을 빠져 나왔을테니까
	// 찾은 것이므로 데이터 리턴
	return currentNode.data
}
