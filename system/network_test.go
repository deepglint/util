package system

import (
	"testing"
)

func TestNetworkModel(t *testing.T) {
	var network_info NetworkInfoModel
	network_info.NICs = make(map[string]NICModel)

	receive0 := ReceiveModel{0, 0, 0, 0, 0, 0, 0, 0}
	transmit0 := TransmitModel{0, 0, 0, 0, 0, 0, 0, 0}
	nic0 := NICModel{receive0, transmit0}
	network_info.NICs["eth0"] = nic0

	receive1 := ReceiveModel{532947, 8192, 0, 0, 0, 0, 0, 0}
	transmit1 := TransmitModel{532947, 8192, 0, 0, 0, 0, 0, 0}
	nic1 := NICModel{receive1, transmit1}
	network_info.NICs["lo"] = nic1

	receive2 := ReceiveModel{2471883108, 4850064, 0, 0, 0, 0, 0, 0}
	transmit2 := TransmitModel{13016753, 53151, 0, 0, 0, 0, 0, 0}
	nic2 := NICModel{receive2, transmit2}
	network_info.NICs["wlan0"] = nic2
}
