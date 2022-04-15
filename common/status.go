package common

import (
	"context"
	"fmt"
	"maikurabu-robit/types"
	"net"
	"time"

	"github.com/Craftserve/mcstatus"
	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func GetServerStatus(project string, zone string, instance string) (string, error) {
	ctx := context.Background()
	auth := option.WithCredentialsFile("./takaran-server-8141624fa778.json")

	service, err := compute.NewService(ctx, auth)
	if err != nil {
		return "", err
	}

	res, err := service.Instances.Get(
		project,
		zone,
		instance,
	).Context(ctx).Do()
	if err != nil {
		return "", err
	}

	return res.Status, nil
}

func GetMCServerStatus(ip string, port uint16) (*types.ServerStatus, error) {
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
		port,
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

func StartServer(project string, zone string, instance string) error {
	ctx := context.Background()
	auth := option.WithCredentialsFile("./takaran-server-8141624fa778.json")

	service, err := compute.NewService(ctx, auth)
	if err != nil {
		return err
	}

	// Launch server
	_, err = service.Instances.Start(
		project,
		zone,
		instance,
	).Context(ctx).Do()
	if err != nil {
		return err
	}

	// Checking server status every 10s for launching
	for {
		status, err := GetServerStatus(project, zone, instance)
		if err != nil {
			return err
		}

		if status == "RUNNING" {
			break
		}

		time.Sleep(10 * time.Second)
	}

	return nil
}

func RebootServer(project string, zone string, instance string) error {
	ctx := context.Background()
	auth := option.WithCredentialsFile("./takaran-server-8141624fa778.json")

	service, err := compute.NewService(ctx, auth)
	if err != nil {
		return err
	}

	// Stop server
	_, err = service.Instances.Stop(
		project,
		zone,
		instance,
	).Context(ctx).Do()
	if err != nil {
		return err
	}

	// Checking server status every 10s for stopping
	for {
		status, err := GetServerStatus(project, zone, instance)
		if err != nil {
			return err
		}

		if status == "TERMINATED" {
			break
		}

		time.Sleep(10 * time.Second)
	}

	// Launch server
	_, err = service.Instances.Start(
		project,
		zone,
		instance,
	).Context(ctx).Do()
	if err != nil {
		return err
	}

	// Checking server status every 10s for launching
	for {
		status, err := GetServerStatus(project, zone, instance)
		if err != nil {
			return err
		}

		if status == "RUNNING" {
			break
		}

		time.Sleep(10 * time.Second)
	}

	return nil
}
