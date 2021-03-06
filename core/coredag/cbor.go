package coredag

import (
	"io"
	"io/ioutil"

	ipld "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-ipld-format"
	ipldcbor "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-ipld-cbor"
)

func cborJSONParser(r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error) {
	nd, err := ipldcbor.FromJSON(r, mhType, mhLen)
	if err != nil {
		return nil, err
	}

	return []ipld.Node{nd}, nil
}

func cborRawParser(r io.Reader, mhType uint64, mhLen int) ([]ipld.Node, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	nd, err := ipldcbor.Decode(data, mhType, mhLen)
	if err != nil {
		return nil, err
	}

	return []ipld.Node{nd}, nil
}
