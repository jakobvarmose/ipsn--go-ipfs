package main

import (
	"fmt"
	"net"
	"os"

	logging "github.com/ipsn/go-ipfs/gxlibs/ipfs/Qmbi1CTJsbnBZjCEgc2otwu8cUFPsGpzWXG7edVCLZ7Gvk/go-log"
)

var log = logging.Logger("seccat")

func exit(format string, vals ...interface{}) {
	if format != "" {
		fmt.Fprintf(os.Stderr, "seccat: error: "+format+"\n", vals...)
	}
	Usage()
	os.Exit(1)
}

func out(format string, vals ...interface{}) {
	if verbose {
		fmt.Fprintf(os.Stderr, "seccat: "+format+"\n", vals...)
	}
}

type logConn struct {
	net.Conn
	n string
}

func (r *logConn) Read(buf []byte) (int, error) {
	n, err := r.Conn.Read(buf)
	if n > 0 {
		log.Debugf("%s read: %v", r.n, buf)
	}
	return n, err
}

func (r *logConn) Write(buf []byte) (int, error) {
	log.Debugf("%s write: %v", r.n, buf)
	return r.Conn.Write(buf)
}

func (r *logConn) Close() error {
	return r.Conn.Close()
}
