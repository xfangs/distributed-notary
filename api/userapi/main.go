package userapi

import (
	"fmt"

	"time"

	"github.com/SmartMeshFoundation/distributed-notary/api"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/nkbai/log"
)

// APIName :
type APIName string

/* #nosec */
const (
	APIAdminNamePrefix             = "Admin-"
	APIAdminNameCreatePrivateKey   = APIAdminNamePrefix + "CreatePrivateKey"   // 发起一次私钥协商
	APIAdminNameGetPrivateKeyList  = APIAdminNamePrefix + "GetPrivateKeyList"  // 私钥片列表查询
	APIAdminNameRegisterNewSCToken = APIAdminNamePrefix + "RegisterNewSCToken" // 注册一个新的侧链token
	APIAdminNameCancelNonce        = APIAdminNamePrefix + "CancelNonce"        // 用一笔小额交易销毁一个nonce

	APIUserNamePrefix            = "User-"
	APIUserNameGetNotaryList     = APIUserNamePrefix + "GetNotaryList"  // 公证人列表查询
	APIUserNameGetSCTokenList    = APIUserNamePrefix + "GetSCTokenList" // 当前支持的SCToken列表查询
	APIUserNameGetLockinStatus   = APIUserNamePrefix + "GetLockinStatus"
	APIUserNameSCPrepareLockin   = APIUserNamePrefix + "SCPrepareLockin"
	APIUserNameSCPrepareLockin2  = APIUserNamePrefix + "SCPrepareLockin2" //封装SCPrepareLockin,方便js处理
	APIUserNameGetLockoutStatus  = APIUserNamePrefix + "GetLockoutStatus"
	APIUserNameMCPrepareLockout  = APIUserNamePrefix + "MCPrepareLockout"
	APIUserNameMCPrepareLockout2 = APIUserNamePrefix + "MCPrepareLockout2" //封装MCPrepareLockout,方便js处理

	APIDebugNamePrefix            = "Debug-"
	APIDebugNameTransferToAccount = APIDebugNamePrefix + "TransferToAccount" // 给某个账户在所有链上转10eth,为了测试
	APIDebugNameGetAllLockinInfo  = APIDebugNamePrefix + "GetAllLockinInfo"
	APIDebugNameGetAllLockoutInfo = APIDebugNamePrefix + "GetAllLockoutInfo"
	APIDebugNameGetAllUtxo        = APIDebugNamePrefix + "GetAllUtxo"
)

// APIName2URLMap :
var APIName2URLMap map[string]string

func init() {
	APIName2URLMap = make(map[string]string)
	/*
		admin
	*/
	APIName2URLMap[APIAdminNameCreatePrivateKey] = "/api/1/admin/private-key"
	APIName2URLMap[APIAdminNameGetPrivateKeyList] = "/api/1/admin/private-keys"
	APIName2URLMap[APIAdminNameRegisterNewSCToken] = "/api/1/admin/sctoken"
	APIName2URLMap[APIAdminNameCancelNonce] = "/api/1/admin/cancel-nonce"
	/*
		user
	*/
	APIName2URLMap[APIUserNameGetNotaryList] = "/api/1/user/notaries"
	APIName2URLMap[APIUserNameGetSCTokenList] = "/api/1/user/sctokens"
	// lockin
	APIName2URLMap[APIUserNameGetLockinStatus] = "/api/1/user/lockin/:sctoken/:secrethash"
	APIName2URLMap[APIUserNameSCPrepareLockin] = "/api/1/user/scpreparelockin/:sctoken"
	APIName2URLMap[APIUserNameSCPrepareLockin2] = "/api/1/user/scpreparelockin2/:sctoken"
	// lockout
	APIName2URLMap[APIUserNameGetLockoutStatus] = "/api/1/user/lockout/:sctoken/:secrethash"
	APIName2URLMap[APIUserNameMCPrepareLockout] = "/api/1/user/mcpreparelockout/:sctoken"
	APIName2URLMap[APIUserNameMCPrepareLockout2] = "/api/1/user/mcpreparelockout2/:sctoken"
	/*
		debug
	*/
	APIName2URLMap[APIDebugNameGetAllLockinInfo] = "/api/1/debug/lockin"
	APIName2URLMap[APIDebugNameGetAllLockoutInfo] = "/api/1/debug/lockout"
	APIName2URLMap[APIDebugNameTransferToAccount] = "/api/1/debug/transfer-to-account/:account"
	APIName2URLMap[APIDebugNameGetAllUtxo] = "/api/1/debug/utxo"
}

// defaultAPITimeout : 默认api请求超时时间
var defaultAPITimeout = 1200 * time.Second

/*
UserAPI :
提供给用户的API
暂时把SystemRequest和NotaryRequest都放在UserAPI
*/
type UserAPI struct {
	api.BaseAPI
	IPPort string
}

// NewUserAPI :
func NewUserAPI(host string) *UserAPI {
	var userAPI UserAPI
	router, err := rest.MakeRouter(
		/*
			api about private key
		*/
		rest.Put(APIName2URLMap[APIAdminNameCreatePrivateKey], userAPI.createPrivateKey),
		/*
			user api
		*/
		rest.Get(APIName2URLMap[APIUserNameGetNotaryList], userAPI.getNotaryList),
		rest.Get(APIName2URLMap[APIUserNameGetSCTokenList], userAPI.getSCTokenList),
		// lockin
		rest.Get(APIName2URLMap[APIUserNameGetLockinStatus], userAPI.getLockinStatus),
		rest.Post(APIName2URLMap[APIUserNameSCPrepareLockin], userAPI.scPrepareLockin),
		rest.Post(APIName2URLMap[APIUserNameSCPrepareLockin2], userAPI.scPrepareLockin2),
		// lockout
		rest.Get(APIName2URLMap[APIUserNameGetLockoutStatus], userAPI.getLockoutStatus),
		rest.Post(APIName2URLMap[APIUserNameMCPrepareLockout], userAPI.mcPrepareLockout),
		rest.Post(APIName2URLMap[APIUserNameMCPrepareLockout2], userAPI.mcPrepareLockout2),
		/*
			admin api
		*/
		rest.Get(APIName2URLMap[APIAdminNameGetPrivateKeyList], userAPI.getPrivateKeyList),
		rest.Put(APIName2URLMap[APIAdminNameRegisterNewSCToken], userAPI.registerNewSCToken),
		rest.Post(APIName2URLMap[APIAdminNameCancelNonce], userAPI.cancelNonce),
		/*
			debug api
		*/
		rest.Get(APIName2URLMap[APIDebugNameGetAllLockinInfo], userAPI.getAllLockinInfo),
		rest.Get(APIName2URLMap[APIDebugNameGetAllLockoutInfo], userAPI.getAllLockoutInfo),
		rest.Get(APIName2URLMap[APIDebugNameTransferToAccount], userAPI.transferToAccount),
		rest.Get(APIName2URLMap[APIDebugNameGetAllUtxo], userAPI.getAllBTCUtxo),
	)
	if err != nil {
		log.Crit(fmt.Sprintf("maker router :%s", err))
	}
	// 跨域
	corsMiddleware := &rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		OriginValidator: func(origin string, request *rest.Request) bool {
			return true
		},
		AllowedMethods: []string{"GET", "POST", "PUT"},
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	}
	userAPI.BaseAPI = api.NewBaseAPI("UserAPI-Server", host, router, corsMiddleware)
	userAPI.Timeout = defaultAPITimeout
	userAPI.IPPort = host
	return &userAPI
}
