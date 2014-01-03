package system

import (
	"errors"
	"github.com/liutong19890905/util/io"
	"regexp"
	"strconv"
)

type NetworkInfoModel struct {
	NICs map[string]NICModel
}

type NICModel struct {
	Receive  ReceiveModel
	Transmit TransmitModel
}

type ReceiveModel struct {
	Bytes      int
	Packets    int
	Errs       int
	Drop       int
	Fifo       int
	Frame      int
	Compressed int
	Multicast  int
}

type TransmitModel struct {
	Bytes      int
	Packets    int
	Errs       int
	Drop       int
	Fifo       int
	Colls      int
	Carrier    int
	Compressed int
}

func GetCurrentNetworkInfo() (network_info NetworkInfoModel, err error) {
	var lines []string
	lines, err = io.ReadLines("/proc/net/dev")
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`[^ |^:]+`)
	network_info.NICs = make(map[string]NICModel)

	for i := 2; i < len(lines); i++ {
		infos := reg.FindAllString(lines[i], -1)
		// fmt.Println(len(infos))
		if len(infos) != 17 {
			err = errors.New("There has something wrong with /proc/net/dev")
			return
		}
		var (
			receive_bytes      int
			receive_packets    int
			receive_errs       int
			receive_drop       int
			receive_fifo       int
			receive_frame      int
			receive_compressed int
			receive_multicast  int

			transmit_bytes      int
			transmit_packets    int
			transmit_errs       int
			transmit_drop       int
			transmit_fifo       int
			transmit_colls      int
			transmit_carrier    int
			transmit_compressed int
		)

		receive_bytes, err = strconv.Atoi(infos[1])
		receive_packets, err = strconv.Atoi(infos[2])
		receive_errs, err = strconv.Atoi(infos[3])
		receive_drop, err = strconv.Atoi(infos[4])
		receive_fifo, err = strconv.Atoi(infos[5])
		receive_frame, err = strconv.Atoi(infos[6])
		receive_compressed, err = strconv.Atoi(infos[7])
		receive_multicast, err = strconv.Atoi(infos[8])

		transmit_bytes, err = strconv.Atoi(infos[9])
		transmit_packets, err = strconv.Atoi(infos[10])
		transmit_errs, err = strconv.Atoi(infos[11])
		transmit_drop, err = strconv.Atoi(infos[12])
		transmit_fifo, err = strconv.Atoi(infos[13])
		transmit_colls, err = strconv.Atoi(infos[14])
		transmit_carrier, err = strconv.Atoi(infos[15])
		transmit_compressed, err = strconv.Atoi(infos[16])
		if err != nil {
			return
		}
		receive := ReceiveModel{receive_bytes, receive_packets, receive_errs, receive_drop,
			receive_fifo, receive_frame, receive_compressed, receive_multicast}
		transmit := TransmitModel{transmit_bytes, transmit_packets, transmit_errs,
			transmit_drop, transmit_fifo, transmit_colls, transmit_carrier, transmit_compressed}
		nic := NICModel{receive, transmit}
		network_info.NICs[infos[0]] = nic
	}
	return
}
