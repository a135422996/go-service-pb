// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.1
// Source: gorpc.proto

package gouserclient

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

	GoUserClient interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
		// 登录
		Login(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*Response, error)
		// 用户app配置管理（ring）
		UserAppConfigAdd(ctx context.Context, in *UserAppConfigVo, opts ...grpc.CallOption) (*UserAppConfigVo, error)
		UserAppConfigUpdate(ctx context.Context, in *UserAppConfigVo, opts ...grpc.CallOption) (*UserAppConfigVo, error)
		UserAppConfigDel(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error)
		UserAppConfigGet(ctx context.Context, in *UserAppConfigGetReq, opts ...grpc.CallOption) (*UserAppConfigVo, error)
		UserAppConfigList(ctx context.Context, in *UserAppConfigListReq, opts ...grpc.CallOption) (*UserAppConfigListRes, error)
		// 设备管理
		DeviceAdd(ctx context.Context, in *DeviceVo, opts ...grpc.CallOption) (*DeviceVo, error)
		DeviceUpdate(ctx context.Context, in *DeviceVo, opts ...grpc.CallOption) (*DeviceVo, error)
		DeviceDel(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error)
		DeviceGet(ctx context.Context, in *DeviceVo, opts ...grpc.CallOption) (*DeviceVo, error)
		DeviceList(ctx context.Context, in *DeviceListReq, opts ...grpc.CallOption) (*DeviceListRes, error)
		// Account
		AccountAdd(ctx context.Context, in *AccountVo, opts ...grpc.CallOption) (*AccountVo, error)
		AccountUpdate(ctx context.Context, in *AccountVo, opts ...grpc.CallOption) (*AccountVo, error)
		AccountDel(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error)
		AccountGet(ctx context.Context, in *AccountVo, opts ...grpc.CallOption) (*AccountVo, error)
		AccountList(ctx context.Context, in *AccountListReq, opts ...grpc.CallOption) (*AccountListResp, error)
		// AccountDevice
		AccountDeviceAdd(ctx context.Context, in *AccountDeviceVo, opts ...grpc.CallOption) (*AccountDeviceVo, error)
		AccountDeviceUpdate(ctx context.Context, in *AccountDeviceVo, opts ...grpc.CallOption) (*AccountDeviceVo, error)
		AccountDeviceDel(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error)
		AccountDeviceGet(ctx context.Context, in *AccountDeviceVo, opts ...grpc.CallOption) (*AccountDeviceVo, error)
		AccountDeviceList(ctx context.Context, in *AccountDeviceListReq, opts ...grpc.CallOption) (*AccountDeviceListResp, error)
	}

	defaultGoUserClient struct {
		cli zrpc.Client
	}
)

func NewGoUserClient(cli zrpc.Client) GoUserClient {
	return &defaultGoUserClient{
		cli: cli,
	}
}

func (m *defaultGoUserClient) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}

// 登录
func (m *defaultGoUserClient) Login(ctx context.Context, in *BaseReq, opts ...grpc.CallOption) (*Response, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 用户app配置管理（ring）
func (m *defaultGoUserClient) UserAppConfigAdd(ctx context.Context, in *UserAppConfigVo, opts ...grpc.CallOption) (*UserAppConfigVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.UserAppConfigAdd(ctx, in, opts...)
}

func (m *defaultGoUserClient) UserAppConfigUpdate(ctx context.Context, in *UserAppConfigVo, opts ...grpc.CallOption) (*UserAppConfigVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.UserAppConfigUpdate(ctx, in, opts...)
}

func (m *defaultGoUserClient) UserAppConfigDel(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.UserAppConfigDel(ctx, in, opts...)
}

func (m *defaultGoUserClient) UserAppConfigGet(ctx context.Context, in *UserAppConfigGetReq, opts ...grpc.CallOption) (*UserAppConfigVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.UserAppConfigGet(ctx, in, opts...)
}

func (m *defaultGoUserClient) UserAppConfigList(ctx context.Context, in *UserAppConfigListReq, opts ...grpc.CallOption) (*UserAppConfigListRes, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.UserAppConfigList(ctx, in, opts...)
}

// 设备管理
func (m *defaultGoUserClient) DeviceAdd(ctx context.Context, in *DeviceVo, opts ...grpc.CallOption) (*DeviceVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.DeviceAdd(ctx, in, opts...)
}

func (m *defaultGoUserClient) DeviceUpdate(ctx context.Context, in *DeviceVo, opts ...grpc.CallOption) (*DeviceVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.DeviceUpdate(ctx, in, opts...)
}

func (m *defaultGoUserClient) DeviceDel(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.DeviceDel(ctx, in, opts...)
}

func (m *defaultGoUserClient) DeviceGet(ctx context.Context, in *DeviceVo, opts ...grpc.CallOption) (*DeviceVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.DeviceGet(ctx, in, opts...)
}

func (m *defaultGoUserClient) DeviceList(ctx context.Context, in *DeviceListReq, opts ...grpc.CallOption) (*DeviceListRes, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.DeviceList(ctx, in, opts...)
}

// Account
func (m *defaultGoUserClient) AccountAdd(ctx context.Context, in *AccountVo, opts ...grpc.CallOption) (*AccountVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountAdd(ctx, in, opts...)
}

func (m *defaultGoUserClient) AccountUpdate(ctx context.Context, in *AccountVo, opts ...grpc.CallOption) (*AccountVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountUpdate(ctx, in, opts...)
}

func (m *defaultGoUserClient) AccountDel(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountDel(ctx, in, opts...)
}

func (m *defaultGoUserClient) AccountGet(ctx context.Context, in *AccountVo, opts ...grpc.CallOption) (*AccountVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountGet(ctx, in, opts...)
}

func (m *defaultGoUserClient) AccountList(ctx context.Context, in *AccountListReq, opts ...grpc.CallOption) (*AccountListResp, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountList(ctx, in, opts...)
}

// AccountDevice
func (m *defaultGoUserClient) AccountDeviceAdd(ctx context.Context, in *AccountDeviceVo, opts ...grpc.CallOption) (*AccountDeviceVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountDeviceAdd(ctx, in, opts...)
}

func (m *defaultGoUserClient) AccountDeviceUpdate(ctx context.Context, in *AccountDeviceVo, opts ...grpc.CallOption) (*AccountDeviceVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountDeviceUpdate(ctx, in, opts...)
}

func (m *defaultGoUserClient) AccountDeviceDel(ctx context.Context, in *Ids, opts ...grpc.CallOption) (*Empty, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountDeviceDel(ctx, in, opts...)
}

func (m *defaultGoUserClient) AccountDeviceGet(ctx context.Context, in *AccountDeviceVo, opts ...grpc.CallOption) (*AccountDeviceVo, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountDeviceGet(ctx, in, opts...)
}

func (m *defaultGoUserClient) AccountDeviceList(ctx context.Context, in *AccountDeviceListReq, opts ...grpc.CallOption) (*AccountDeviceListResp, error) {
	client := gorpc.NewGoUserClientClient(m.cli.Conn())
	return client.AccountDeviceList(ctx, in, opts...)
}
