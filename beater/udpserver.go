// Copyright 2019-2022 F5 Inc.
// See LICENSE.txt for the Apache-2 License.

package beater

import (
	"fmt"
	"net"
	"os"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/threatstack/gelfbeat/config"
)

func listen(receive chan string, config config.Config) {
	sAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", config.Listen, config.Port))
	if err != nil {
		logp.Error(err)
		os.Exit(1)
	}
	sConn, err := net.ListenUDP("udp", sAddr)
	if err != nil {
		logp.Error(err)
		os.Exit(1)
	}
	defer sConn.Close()

	buf := make([]byte, 8192)

	var c = make(chan []byte, 2)
	var nn = make(chan []byte)
	var ac = make(chan GELFMessage, 2)

	go printer(c, nn)
	go buildMessage(c, ac)
	i := 0
	for {
		i++
		n, err := sConn.Read(buf)
		// process msg
		go msgBuildWorker(c, ac, nn, i)
		nn <- buf[0:n]
		// process result
		data := <-nn

		if string(data) != "" {
			receive <- string(data)
		}

		if err != nil {
			logp.Error(err)
		}
	}

}
