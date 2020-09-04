# mute

## overview

Visualization of prefix match using trie tree.

```
# data.txt
keynote
keycase
king
kingdom
macbook
```

<img src="https://user-images.githubusercontent.com/31947384/91886968-126bf980-ecc5-11ea-8209-3b7131a6c3f2.png" width="160">

## setup

```sh
$ brew install graphviz
$ go get github.com/kotaroooo0/mute

# create data text for trie. see: data.txt.sample
$ vi data.txt
$ mute -s data.txt | dot -T png -o sample.png

# only create .dot
$ mute -s data.txt -o sample.dot
```
