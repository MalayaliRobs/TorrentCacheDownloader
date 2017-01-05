// Takes P2P blocklist text format in stdin, and outputs the packed format
// from the iplist package.
package main

import (
	"bufio"
	"os"

	"github.com/anacrolix/missinggo"
	"github.com/anacrolix/missinggo/args"

	"github.com/anacrolix/torrent/iplist"
)

func main() {
	args.Parse()
	l, err := iplist.NewFromReader(os.Stdin)
	if err != nil {
		missinggo.Fatal(err)
	}
	wb := bufio.NewWriter(os.Stdout)
	defer wb.Flush()
	err = l.WritePacked(wb)
	if err != nil {
		missinggo.Fatal(err)
	}
}
