package commands

import (
	"github.com/spf13/cobra"

	"github.com/33cn/chain33/rpc/jsonclient"
	rpctypes "github.com/33cn/chain33/rpc/types"
)

func NetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "net",
		Short: "Net operation",
		Args:  cobra.MinimumNArgs(1),
	}

	cmd.AddCommand(
		GetPeerInfoCmd(),
		IsClockSyncCmd(),
		IsSyncCmd(),
		GetNetInfoCmd(),
		GetFatalFailureCmd(),
		GetTimeStausCmd(),
	)

	return cmd
}

// get peers connected info
func GetPeerInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "peer_info",
		Short: "Get remote peer nodes",
		Run:   peerInfo,
	}
	return cmd
}

func peerInfo(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res rpctypes.PeerList
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetPeerInfo", nil, &res)
	ctx.Run()
}

// get ntp clock sync status
func IsClockSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is_clock_sync",
		Short: "Get ntp clock synchronization status",
		Run:   isClockSync,
	}
	return cmd
}

func isClockSync(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res bool
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.IsNtpClockSync", nil, &res)
	ctx.Run()
}

// get local db sync status
func IsSyncCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is_sync",
		Short: "Get blockchain synchronization status",
		Run:   isSync,
	}
	return cmd
}

func isSync(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res bool
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.IsSync", nil, &res)
	ctx.Run()
}

// get net info
func GetNetInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "Get net information",
		Run:   netInfo,
	}
	return cmd
}

func netInfo(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res rpctypes.NodeNetinfo
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetNetInfo", nil, &res)
	ctx.Run()
}

// get FatalFailure
func GetFatalFailureCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fault",
		Short: "Get system fault",
		Run:   fatalFailure,
	}
	return cmd
}

func fatalFailure(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res int64
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetFatalFailure", nil, &res)
	ctx.Run()
}

// get time status
func GetTimeStausCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "time",
		Short: "Get time status",
		Run:   timestatus,
	}
	return cmd
}

func timestatus(cmd *cobra.Command, args []string) {
	rpcLaddr, _ := cmd.Flags().GetString("rpc_laddr")
	var res rpctypes.TimeStatus
	ctx := jsonclient.NewRpcCtx(rpcLaddr, "Chain33.GetTimeStatus", nil, &res)
	ctx.Run()
}
