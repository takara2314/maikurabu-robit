package processes

import (
	"context"
	"time"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/option"
)

func StartServer() error {
	ctx := context.Background()
	auth := option.WithCredentialsFile("./takaran-server-8141624fa778.json")

	service, err := compute.NewService(ctx, auth)
	if err != nil {
		return err
	}

	// 起動処理
	_, err = service.Instances.Start(
		"takaran-server",
		"asia-northeast2-c",
		"minecraft",
	).Context(ctx).Do()
	if err != nil {
		return err
	}

	// 起動を確認するまで、10秒ごとに状態を監視
	for {
		status, err := CheckServer()
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
