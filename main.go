package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/awalterschulze/gographviz"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	words, err := readWords()
	if err != nil {
		return err
	}

	return generateDotfile(generateTrie(words))
}

// ファイルから入力
func readWords() ([]string, error) {
	f, err := os.Open("data.txt")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	words := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		words = append(words, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

// Trie木を生成
func generateTrie(words []string) *Node {
	t := newNode("", make(map[rune]*Node), false)
	for _, w := range words {
		t.insert(w)
	}
	return t
}

var id = 0

type Node struct {
	ID       int
	Key      string
	Children map[rune]*Node
	End      bool
}

func newNode(k string, c map[rune]*Node, e bool) *Node {
	id++
	return &Node{
		ID:       id,
		Key:      k,
		Children: c,
		End:      e,
	}
}

func (n *Node) getLabel() string {
	if n.Key == "" {
		return "\"0\""
	}
	return fmt.Sprintf("\"%s\"", n.Key)
}

func (n *Node) getShape() string {
	if n.End {
		return "\"doublecircle\""
	}
	return "\"circle\""
}

func (n *Node) insert(w string) error {
	runes := []rune(w)
	currentNode := n
	for i, r := range runes {
		if nextNode, ok := currentNode.Children[r]; ok {
			currentNode = nextNode
		} else {
			currentNode.Children[r] = newNode(string(r), make(map[rune]*Node), false)
			currentNode = currentNode.Children[r]
		}

		// 終端にチェック
		if i == len(runes)-1 {
			currentNode.End = true
		}
	}
	return nil
}

var fontSize = "35"
var g = gographviz.NewGraph()

func dfs(n *Node) error {
	for _, v := range n.Children {
		if err := g.AddNode("G", strconv.Itoa(n.ID), map[string]string{"label": n.getLabel(), "shape": n.getShape(), "fontsize": fontSize}); err != nil {
			return err
		}
		if err := g.AddNode("G", strconv.Itoa(v.ID), map[string]string{"label": v.getLabel(), "shape": v.getShape(), "fontsize": fontSize}); err != nil {
			return err
		}
		if err := g.AddEdge(strconv.Itoa(n.ID), strconv.Itoa(v.ID), true, nil); err != nil {
			return err
		}
		dfs(v)
	}
	return nil
}

// Trie木からdotファイルを生成
func generateDotfile(trie *Node) error {

	g.SetName("G")
	g.SetDir(true)
	if err := dfs(trie); err != nil {
		return err
	}

	f, err := os.Create("trie.dot")
	if err != nil {
		return err
	}
	f.WriteString(g.String())
	return nil
}
