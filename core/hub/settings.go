package hub

import (
	"context"
	"errors"

	hubpb "github.com/tron-us/go-btfs-common/protos/hub"
	nodepb "github.com/tron-us/go-btfs-common/protos/node"
	"github.com/tron-us/go-btfs-common/utils/grpc"
)

func GetHostSettings(ctx context.Context, addr, peerId string) (*nodepb.Node_Settings, error) {
	// get from remote
	ns := new(nodepb.Node_Settings)
	err := grpc.HubQueryClient(addr).WithContext(ctx, func(ctx context.Context, client hubpb.HubQueryServiceClient) error {
		req := new(hubpb.SettingsReq)
		req.Id = peerId
		resp, err := client.GetSettings(ctx, req)
		if err != nil {
			return err
		}
		if resp.Code != hubpb.ResponseCode_SUCCESS {
			return errors.New(resp.Message)
		}
		ns.StorageTimeMin = uint64(resp.SettingsData.StorageTimeMin)
		ns.StoragePriceAsk = uint64(resp.SettingsData.StoragePriceAsk)
		ns.BandwidthLimit = resp.SettingsData.BandwidthLimit
		ns.BandwidthPriceAsk = uint64(resp.SettingsData.BandwidthPriceAsk)
		ns.CollateralStake = uint64(resp.SettingsData.CollateralStake)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return ns, nil
}