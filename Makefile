gen:
	go run main.go
	dot -T png trie.dot -o trie.png

gen2:
	go run main.go | dot -T png -o tgaks.png
