package main

import (
	"fmt"
	"log"
	"main/block"
	"main/pulls"
	"os"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
)

var (
	path = "./pullipfs/"
)
var (
	pathdown = "./downipfs/"
)

func main() {
	os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)

	var i int
	var url string
	var key string
	var hash string
	fmt.Print("Type a number: ")
	fmt.Scan(&i)
	switch {
	case i == 1:
		fmt.Print("Url: ")
		fmt.Scan(&url)
		fmt.Print("Key: ")
		fmt.Scan(&key)
		pulls.Pullrepo(url)
		encrypt(path, key)
		time.Sleep(time.Millisecond * 15)
		ipfspull(path)
	case i == 2:
		os.OpenFile(pathdown, os.O_RDONLY|os.O_CREATE, 0666)
		fmt.Print("Hash: ")
		fmt.Scan(&hash)
		fmt.Print("Key: ")
		fmt.Scan(&key)
		ipfsdown(hash, path)
		time.Sleep(time.Millisecond * 15)
		decryption(path, key)
	}

}
func ipfsdown(hash string, paths string) {
	sh := shell.NewShell("localhost:5001")
	sh.Get(hash, paths)
}
func ipfspull(paths string) {
	sh := shell.NewShell("localhost:5001")

	cid, err := sh.AddDir(paths)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("https://ipfs.io/ipfs/%s", cid)
}
func encrypt(paths string, key string) {
	block.Blockencrypt(paths, key)
}
func decryption(paths string, key string) {
	block.Blockdecryption(paths, key)
}
