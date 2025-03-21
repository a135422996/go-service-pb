// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: gorpc.proto

package gosysclient

import (
	"context"

	"github.com/a135422996/go-service-pb/gorpc/pb/gorpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AccountDeviceListReq         = gorpc.AccountDeviceListReq
	AccountDeviceListResp        = gorpc.AccountDeviceListResp
	AccountDeviceVo              = gorpc.AccountDeviceVo
	AccountListReq               = gorpc.AccountListReq
	AccountListResp              = gorpc.AccountListResp
	AccountVo                    = gorpc.AccountVo
	AreaParamListReq             = gorpc.AreaParamListReq
	AreaParamListResp            = gorpc.AreaParamListResp
	AreaParamVo                  = gorpc.AreaParamVo
	BaseReq                      = gorpc.BaseReq
	CommonResult                 = gorpc.CommonResult
	ConfigGetReq                 = gorpc.ConfigGetReq
	DeleteIdResp                 = gorpc.DeleteIdResp
	DeviceListReq                = gorpc.DeviceListReq
	DeviceListRes                = gorpc.DeviceListRes
	DeviceVo                     = gorpc.DeviceVo
	DictGetReq                   = gorpc.DictGetReq
	DictListReq                  = gorpc.DictListReq
	DictListResp                 = gorpc.DictListResp
	DictVo                       = gorpc.DictVo
	Empty                        = gorpc.Empty
	GameListReq                  = gorpc.GameListReq
	GameListRes                  = gorpc.GameListRes
	GameReq                      = gorpc.GameReq
	GameVo                       = gorpc.GameVo
	GlobalParamListReq           = gorpc.GlobalParamListReq
	GlobalParamListResp          = gorpc.GlobalParamListResp
	GlobalParamVo                = gorpc.GlobalParamVo
	I18NVo                       = gorpc.I18NVo
	Ids                          = gorpc.Ids
	Int64SelectController        = gorpc.Int64SelectController
	IntSelectController          = gorpc.IntSelectController
	NationalLanguageListReq      = gorpc.NationalLanguageListReq
	NationalLanguageListResp     = gorpc.NationalLanguageListResp
	NationalLanguageVo           = gorpc.NationalLanguageVo
	Page                         = gorpc.Page
	Page_OrderItem               = gorpc.Page_OrderItem
	ProductActionListReq         = gorpc.ProductActionListReq
	ProductActionListRes         = gorpc.ProductActionListRes
	ProductActionReq             = gorpc.ProductActionReq
	ProductActionVo              = gorpc.ProductActionVo
	ProductListReq               = gorpc.ProductListReq
	ProductListResp              = gorpc.ProductListResp
	ProductParamListReq          = gorpc.ProductParamListReq
	ProductParamListResp         = gorpc.ProductParamListResp
	ProductParamVo               = gorpc.ProductParamVo
	ProductVo                    = gorpc.ProductVo
	RelGameRankListReq           = gorpc.RelGameRankListReq
	RelGameRankListRes           = gorpc.RelGameRankListRes
	RelGameRankReq               = gorpc.RelGameRankReq
	RelGameRankVo                = gorpc.RelGameRankVo
	RelGameVpnGroupListReq       = gorpc.RelGameVpnGroupListReq
	RelGameVpnGroupListRes       = gorpc.RelGameVpnGroupListRes
	RelGameVpnGroupReq           = gorpc.RelGameVpnGroupReq
	RelGameVpnGroupVo            = gorpc.RelGameVpnGroupVo
	Request                      = gorpc.Request
	ResourceDelReq               = gorpc.ResourceDelReq
	Response                     = gorpc.Response
	StringSelectController       = gorpc.StringSelectController
	TipDelReq                    = gorpc.TipDelReq
	TipKey                       = gorpc.TipKey
	TipListReq                   = gorpc.TipListReq
	TipListResp                  = gorpc.TipListResp
	TipVo                        = gorpc.TipVo
	TransactionOperation         = gorpc.TransactionOperation
	TransactionReq               = gorpc.TransactionReq
	UpgradePkgGetReq             = gorpc.UpgradePkgGetReq
	UpgradePkgListReq            = gorpc.UpgradePkgListReq
	UpgradePkgListResp           = gorpc.UpgradePkgListResp
	UpgradePkgPlanDetailListReq  = gorpc.UpgradePkgPlanDetailListReq
	UpgradePkgPlanDetailListResp = gorpc.UpgradePkgPlanDetailListResp
	UpgradePkgPlanDetailVo       = gorpc.UpgradePkgPlanDetailVo
	UpgradePkgPlanListReq        = gorpc.UpgradePkgPlanListReq
	UpgradePkgPlanListResp       = gorpc.UpgradePkgPlanListResp
	UpgradePkgPlanVo             = gorpc.UpgradePkgPlanVo
	UpgradePkgVo                 = gorpc.UpgradePkgVo
	UpgradeResConfigGetReq       = gorpc.UpgradeResConfigGetReq
	UpgradeResConfigListReq      = gorpc.UpgradeResConfigListReq
	UpgradeResConfigListResp     = gorpc.UpgradeResConfigListResp
	UpgradeResConfigVo           = gorpc.UpgradeResConfigVo
	UpgradeResGetReq             = gorpc.UpgradeResGetReq
	UpgradeResListReq            = gorpc.UpgradeResListReq
	UpgradeResListResp           = gorpc.UpgradeResListResp
	UpgradeResPlanDetailListReq  = gorpc.UpgradeResPlanDetailListReq
	UpgradeResPlanDetailListResp = gorpc.UpgradeResPlanDetailListResp
	UpgradeResPlanDetailVo       = gorpc.UpgradeResPlanDetailVo
	UpgradeResPlanListReq        = gorpc.UpgradeResPlanListReq
	UpgradeResPlanListResp       = gorpc.UpgradeResPlanListResp
	UpgradeResPlanVo             = gorpc.UpgradeResPlanVo
	UpgradeResVo                 = gorpc.UpgradeResVo
	UserAppConfigGetReq          = gorpc.UserAppConfigGetReq
	UserAppConfigListReq         = gorpc.UserAppConfigListReq
	UserAppConfigListRes         = gorpc.UserAppConfigListRes
	UserAppConfigVo              = gorpc.UserAppConfigVo
	VpnGroupAddReq               = gorpc.VpnGroupAddReq
	VpnGroupGetReq               = gorpc.VpnGroupGetReq
	VpnGroupListReq              = gorpc.VpnGroupListReq
	VpnGroupListResp             = gorpc.VpnGroupListResp
	VpnGroupTagListReq           = gorpc.VpnGroupTagListReq
	VpnGroupTagListResp          = gorpc.VpnGroupTagListResp
	VpnGroupTagVo                = gorpc.VpnGroupTagVo
	VpnGroupVo                   = gorpc.VpnGroupVo
	VpnLocationExVo              = gorpc.VpnLocationExVo
	VpnLocationGetReq            = gorpc.VpnLocationGetReq
	VpnLocationListReq           = gorpc.VpnLocationListReq
	VpnLocationListResp          = gorpc.VpnLocationListResp
	VpnLocationServerCountReq    = gorpc.VpnLocationServerCountReq
	VpnLocationServerCountResp   = gorpc.VpnLocationServerCountResp
	VpnLocationServerCountVo     = gorpc.VpnLocationServerCountVo
	VpnLocationVo                = gorpc.VpnLocationVo
	VpnServerAddReq              = gorpc.VpnServerAddReq
	VpnServerListReq             = gorpc.VpnServerListReq
	VpnServerListResp            = gorpc.VpnServerListResp
	VpnServerLocationListReq     = gorpc.VpnServerLocationListReq
	VpnServerLocationListResp    = gorpc.VpnServerLocationListResp
	VpnServerLocationVo          = gorpc.VpnServerLocationVo
	VpnServerVo                  = gorpc.VpnServerVo
	VpnTagGetReq                 = gorpc.VpnTagGetReq
	VpnTagListReq                = gorpc.VpnTagListReq
	VpnTagListResp               = gorpc.VpnTagListResp
	VpnTagVo                     = gorpc.VpnTagVo

	GoSysClient interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		// 登录
		Login(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*Response, error)
	}

	defaultGoSysClient struct {
		cli zrpc.Client
	}
)

func NewGoSysClient(cli zrpc.Client) GoSysClient {
	return &defaultGoSysClient{
		cli: cli,
	}
}

func (m *defaultGoSysClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := gorpc.NewGoSysClientClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

// 登录
func (m *defaultGoSysClient) Login(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*Response, error) {
	client := gorpc.NewGoSysClientClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}
