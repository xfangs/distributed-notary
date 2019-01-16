package userapi

import (
	"github.com/SmartMeshFoundation/distributed-notary/api"
	"github.com/ant0ine/go-json-rest/rest"
)

// GetPrivateKeyListRequest :
type GetPrivateKeyListRequest struct {
	api.BaseRequest
}

/*
getPrivateKeyList 私钥列表查询
*/
func (ua *UserAPI) getPrivateKeyList(w rest.ResponseWriter, r *rest.Request) {
	req := &GetPrivateKeyListRequest{
		BaseRequest: api.NewBaseRequest(APIAdminNameGetPrivateKeyList),
	}
	api.Return(w, ua.SendToServiceAndWaitResponse(req))
}

// CreatePrivateKeyRequest :
type CreatePrivateKeyRequest struct {
	api.BaseRequest
}

/*
CreatePrivateKey 该接口仅仅是发启一次过程,不参与实际协商过程,由SystemService处理
发起一次私钥协商过程,生成一组私钥片
*/
func (ua *UserAPI) createPrivateKey(w rest.ResponseWriter, r *rest.Request) {
	req := &CreatePrivateKeyRequest{
		BaseRequest: api.NewBaseRequest(APIAdminNameCreatePrivateKey),
	}
	api.Return(w, ua.SendToServiceAndWaitResponse(req))
}

// RegisterSCTokenRequest :
type RegisterSCTokenRequest struct {
	api.BaseRequest
	MainChainName string `json:"main_chain_name,omitempty"` // 主链名,目前仅支持以太坊
	PrivateKeyID  string `json:"private_key_id,omitempty"`  // 部署合约使用的私钥ID
}

/*
registerNewSCToken :
	注册一个新的侧链token地址,dnotary将完成以下工作:
	1.部署一个主链合约,一个侧链合约
*/
func (ua *UserAPI) registerNewSCToken(w rest.ResponseWriter, r *rest.Request) {
	req := &RegisterSCTokenRequest{
		BaseRequest: api.NewBaseRequest(APIAdminNameRegisterNewSCToken),
	}
	err := r.DecodeJsonPayload(req)
	if err != nil {
		api.Return(w, api.NewFailResponse(req.RequestID, api.ErrorCodeParamsWrong))
		return
	}
	if req.MainChainName == "" {
		api.Return(w, api.NewFailResponse(req.RequestID, api.ErrorCodeParamsWrong, "main_chain_name can not be null"))
		return
	}
	if req.PrivateKeyID == "" {
		api.Return(w, api.NewFailResponse(req.RequestID, api.ErrorCodeParamsWrong, "private_key_id can not be null"))
		return
	}
	api.Return(w, ua.SendToServiceAndWaitResponse(req))
}
