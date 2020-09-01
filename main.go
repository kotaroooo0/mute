package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

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
	replace := []string{"@", " ", "　", "　", "(", "（", ")", "）", ",", "-", "/"}
	for s.Scan() {
		t := s.Text()
		for _, r := range replace {
			t = strings.Replace(t, r, "", -1)
		}
		words = append(words, t)
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

// Trie木を生成
func generateTrie(words []string) *Node {
	t := newNode("", make(map[rune]*Node))
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
}

func newNode(k string, c map[rune]*Node) *Node {
	id++
	return &Node{
		ID:       id,
		Key:      k,
		Children: c,
	}
}

func (n *Node) getLabel() string {
	if n.Key == "" {
		return "ROOT"
	}
	return n.Key
}

func (n *Node) insert(w string) error {
	runes := []rune(w)
	currentNode := n
	for _, r := range runes {
		if nextNode, ok := currentNode.Children[r]; ok {
			currentNode = nextNode
		} else {
			currentNode.Children[r] = newNode(string(r), make(map[rune]*Node))
			currentNode = currentNode.Children[r]
		}
	}
	return nil
}

var g = gographviz.NewGraph()

func dfs(n *Node) error {
	for _, v := range n.Children {
		if err := g.AddNode("G", strconv.Itoa(n.ID), map[string]string{"label": n.getLabel(), "fontsize": "35"}); err != nil {
			return err
		}
		if err := g.AddNode("G", strconv.Itoa(v.ID), map[string]string{"label": v.getLabel(), "fontsize": "35"}); err != nil {
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
