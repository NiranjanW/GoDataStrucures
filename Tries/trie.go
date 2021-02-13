package main

import "fmt"

//Allphabet is the number of possible characters in the trie
const Alphabetsize = 26

//Node
type Node struct {
	children [25]*Node
	isEnd    bool
}

//Trie
type Trie struct {
	root *Node
}

//InitTrie will create a new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[int(charIndex)] = &Node{}
		}
		currentNode = currentNode.children[int(charIndex)]
	}
	currentNode.isEnd = true
}

// Search
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root
	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[int(charIndex)]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	myTrie := InitTrie()
	myTrie.Insert("aragon")
	myTrie.Search("aragon")
	toAdd := []string{
		"aragon",
		"orgon",
		"oregon",
		"oreo",
	}
	for _, v := range toAdd {
		myTrie.Insert(v)
	}
	fmt.Println(myTrie.Search("aragon"))
}
