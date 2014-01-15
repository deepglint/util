package monitor

import (
	"deepglint/util/io/file"
	"errors"
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
	Bytes      uint64
	Packets    uint64
	Errs       uint64
	Drop       uint64
	Fifo       uint64
	Frame      uint64
	Compressed uint64
	Multicast  uint64
}

type Transmit struct {
	Bytes      uint64
	Packets    uint64
	Errs       uint64
	Drop       uint64
	Fifo       uint64
	Colls      uint64
	Carrier    uint64
	Compressed uint64
}

func GetCurrentNIC(data_file_path string, dev string) (result NIC, err error) {
	var lines []string
	lines, err = file.ReadLines(data_file_path)
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
				receive_bytes      uint64
				receive_packets    uint64
				receive_errs       uint64
				receive_drop       uint64
				receive_fifo       uint64
				receive_frame      uint64
				receive_compressed uint64
				receive_multicast  uint64

				transmit_bytes      uint64
				transmit_packets    uint64
				transmit_errs       uint64
				transmit_drop       uint64
				transmit_fifo       uint64
				transmit_colls      uint64
				transmit_carrier    uint64
				transmit_compressed uint64
			)

			receive_bytes, err = strconv.ParseUint(infos[1], 10, 64)
			receive_packets, err = strconv.ParseUint(infos[2], 10, 64)
			receive_errs, err = strconv.ParseUint(infos[3], 10, 64)
			receive_drop, err = strconv.ParseUint(infos[4], 10, 64)
			receive_fifo, err = strconv.ParseUint(infos[5], 10, 64)
			receive_frame, err = strconv.ParseUint(infos[6], 10, 64)
			receive_compressed, err = strconv.ParseUint(infos[7], 10, 64)
			receive_multicast, err = strconv.ParseUint(infos[8], 10, 64)

			transmit_bytes, err = strconv.ParseUint(infos[9], 10, 64)
			transmit_packets, err = strconv.ParseUint(infos[10], 10, 64)
			transmit_errs, err = strconv.ParseUint(infos[11], 10, 64)
			transmit_drop, err = strconv.ParseUint(infos[12], 10, 64)
			transmit_fifo, err = strconv.ParseUint(infos[13], 10, 64)
			transmit_colls, err = strconv.ParseUint(infos[14], 10, 64)
			transmit_carrier, err = strconv.ParseUint(infos[15], 10, 64)
			transmit_compressed, err = strconv.ParseUint(infos[16], 10, 64)
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
