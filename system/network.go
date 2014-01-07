package system

import (
	"errors"
	"github.com/liutong19890905/util/io"
	"regexp"
	"strconv"
	"time"
)

type CurrentNetworkState struct {
	NICs map[string]NIC
}

type NIC struct {
	Receive  Receive
	Transmit Transmit
}

type Receive struct {
	Bytes      int
	Packets    int
	Errs       int
	Drop       int
	Fifo       int
	Frame      int
	Compressed int
	Multicast  int
}

type Transmit struct {
	Bytes      int
	Packets    int
	Errs       int
	Drop       int
	Fifo       int
	Colls      int
	Carrier    int
	Compressed int
}

type IntervalNetworkState struct {
	ReceiveBytes  float64
	TransmitBytes float64
}

func GetIntervalNetworkState(dev string, test_time int, data_file_path string) (result IntervalNetworkState, err error) {
	network_info_before, err := GetCurrentNetworkState(data_file_path)
	if err != nil {
		return
	}
	timer := time.NewTimer(time.Duration(test_time) * time.Millisecond)
	<-timer.C
	network_info_after, err := GetCurrentNetworkState(data_file_path)
	if err != nil {
		return
	}
	result.ReceiveBytes = (float64(network_info_after.NICs[dev].Receive.Bytes) - float64(network_info_before.NICs[dev].Receive.Bytes)) / float64(test_time) * 1000.0
	result.TransmitBytes = (float64(network_info_after.NICs[dev].Transmit.Bytes) - float64(network_info_before.NICs[dev].Transmit.Bytes)) / float64(test_time) * 1000.0
	return
}

func GetCurrentNetworkState(data_file_path string) (network_info CurrentNetworkState, err error) {
	var lines []string
	lines, err = io.ReadLines(data_file_path)
	if err != nil {
		return
	}
	reg := regexp.MustCompile(`[^ |^:]+`)
	network_info.NICs = make(map[string]NIC)

	for i := 2; i < len(lines); i++ {
		infos := reg.FindAllString(lines[i], -1)
		if len(infos) != 17 {
			err = errors.New("There has something wrong with " + data_file_path)
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
		receive := Receive{receive_bytes, receive_packets, receive_errs, receive_drop,
			receive_fifo, receive_frame, receive_compressed, receive_multicast}
		transmit := Transmit{transmit_bytes, transmit_packets, transmit_errs,
			transmit_drop, transmit_fifo, transmit_colls, transmit_carrier, transmit_compressed}
		nic := NIC{receive, transmit}
		network_info.NICs[infos[0]] = nic
	}
	return
}
