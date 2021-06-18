package processes

import (
	"fmt"
	"maikurabu-robit/types"
	"net"

	"github.com/Craftserve/mcstatus"
)

func GetServerStatus(ip string, port int) (*types.ServerStatus, error) {
	address, err := net.ResolveTCPAddr(
		"tcp4",
		fmt.Sprintf("%s:%d", ip, port),
	)
	if err != nil {
		return nil, err
	}

	apiStatus, ping, err := mcstatus.CheckStatusNew(
		address,
		"localhost",
		25565,
	)
	if err != nil {
		return nil, err
	}

	status := types.ServerStatus{
		Version: apiStatus.GameVersion,
		Player:  apiStatus.Players,
		Max:     apiStatus.Slots,
		Players: apiStatus.PlayersSample,
		Ping:    ping,
		Icon:    apiStatus.Favicon,
	}

	return &status, nil
}
