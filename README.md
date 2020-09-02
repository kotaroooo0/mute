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

$ git clone https://github.com/kotaroooo0/mute.git
$ cd mute
$ vi data.txt # create data text for trie. see: data.txt.sample
$ make

# TODO: Ideal
$ go get github.com/kotaroooo0/mute
$ mute -f data.txt -o .
```
