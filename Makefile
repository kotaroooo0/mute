gen:
	go run main.go
	dot -T png trie.dot -o trie.png
