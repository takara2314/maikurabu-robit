package processes

import (
	"context"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func CheckServer() (string, error) {
	ctx := context.Background()
	auth := option.WithCredentialsFile("./takaran-server-8141624fa778.json")

	service, err := compute.NewService(ctx, auth)
	if err != nil {
		return "", err
	}

	res, err := service.Instances.Get(
		"takaran-server",
		"asia-northeast2-c",
		"minecraft-v2",
	).Context(ctx).Do()
	if err != nil {
		return "", err
	}

	return res.Status, nil
}
