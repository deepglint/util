package system

import (
	"errors"
	"github.com/liutong19890905/util/io"
	"regexp"
	"strconv"
	"strings"
)

type NIC struct {
	Receive       Receive
	Transmit      Transmit
	ReceiveBytes  float64
	TransmitBytes float64
}

type Receive struct {
	Bytes      int64
	Packets    int64
	Errs       int64
	Drop       int64
	Fifo       int64
	Frame      int64
	Compressed int64
	Multicast  int64
}

type Transmit struct {
	Bytes      int64
	Packets    int64
	Errs       int64
	Drop       int64
	Fifo       int64
	Colls      int64
	Carrier    int64
	Compressed int64
}

func GetCurrentNIC(data_file_path string, dev string) (result NIC, err error) {
	var lines []string
	lines, err = io.ReadLines(data_file_path)
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`[^ |^:]+`)
	for i := 2; i < len(lines); i++ {
		infos := reg.FindAllString(lines[i], -1)
		if len(infos) != 17 {
			err = errors.New("There has something wrong with " + data_file_path)
			return
		}
		if strings.EqualFold(infos[0], dev) {
			var (
				receive_bytes      int64
				receive_packets    int64
				receive_errs       int64
				receive_drop       int64
				receive_fifo       int64
				receive_frame      int64
				receive_compressed int64
				receive_multicast  int64

				transmit_bytes      int64
				transmit_packets    int64
				transmit_errs       int64
				transmit_drop       int64
				transmit_fifo       int64
				transmit_colls      int64
				transmit_carrier    int64
				transmit_compressed int64
			)

			receive_bytes, err = strconv.ParseInt(infos[1], 10, 64)
			receive_packets, err = strconv.ParseInt(infos[2], 10, 64)
			receive_errs, err = strconv.ParseInt(infos[3], 10, 64)
			receive_drop, err = strconv.ParseInt(infos[4], 10, 64)
			receive_fifo, err = strconv.ParseInt(infos[5], 10, 64)
			receive_frame, err = strconv.ParseInt(infos[6], 10, 64)
			receive_compressed, err = strconv.ParseInt(infos[7], 10, 64)
			receive_multicast, err = strconv.ParseInt(infos[8], 10, 64)

			transmit_bytes, err = strconv.ParseInt(infos[9], 10, 64)
			transmit_packets, err = strconv.ParseInt(infos[10], 10, 64)
			transmit_errs, err = strconv.ParseInt(infos[11], 10, 64)
			transmit_drop, err = strconv.ParseInt(infos[12], 10, 64)
			transmit_fifo, err = strconv.ParseInt(infos[13], 10, 64)
			transmit_colls, err = strconv.ParseInt(infos[14], 10, 64)
			transmit_carrier, err = strconv.ParseInt(infos[15], 10, 64)
			transmit_compressed, err = strconv.ParseInt(infos[16], 10, 64)
			if err != nil {
				return
			}
			receive := Receive{receive_bytes, receive_packets, receive_errs, receive_drop,
				receive_fifo, receive_frame, receive_compressed, receive_multicast}
			transmit := Transmit{transmit_bytes, transmit_packets, transmit_errs,
				transmit_drop, transmit_fifo, transmit_colls, transmit_carrier, transmit_compressed}
			result = NIC{receive, transmit, 0, 0}
		}
	}
	return
}
