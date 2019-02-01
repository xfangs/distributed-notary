package dnc

import (
	"net/http"

	"fmt"

	"os"

	"github.com/SmartMeshFoundation/distributed-notary/api"
	"github.com/SmartMeshFoundation/distributed-notary/api/userapi"
	"github.com/SmartMeshFoundation/distributed-notary/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/urfave/cli"
)

var mcploCmd = cli.Command{
	Name:      "main-chain-prepare-lock-out",
	ShortName: "mcplo",
	Usage:     "call MCPrepareLockout API of notary",
	Action:    mcPrepareLockout,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "mcname",
			Usage: "name of main chain contract which you want to lockout",
			Value: "ethereum",
		},
	},
}

func mcPrepareLockout(ctx *cli.Context) (err error) {
	scTokenInfo := getSCTokenByMCName(ctx.String("mcname"))
	if scTokenInfo == nil {
		fmt.Println("wrong mcname")
		os.Exit(-1)
	}
	if GlobalConfig.RunTime == nil {
		fmt.Println("must call pli first")
		os.Exit(-1)
	}
	url := GlobalConfig.NotaryHost + "/api/1/user/mcpreparelockout/" + scTokenInfo.SCToken.String()
	req := &userapi.MCPrepareLockoutRequest{
		BaseReq:              api.NewBaseReq(userapi.APIUserNameMCPrepareLockout),
		BaseReqWithResponse:  api.NewBaseReqWithResponse(),
		BaseReqWithSCToken:   api.NewBaseReqWithSCToken(scTokenInfo.SCToken),
		BaseReqWithSignature: api.NewBaseReqWithSignature(common.HexToAddress(GlobalConfig.EthUserAddress)),
		SecretHash:           common.HexToHash(GlobalConfig.RunTime.SecretHash),
		MCUserAddress:        common.HexToAddress(GlobalConfig.EthUserAddress),
		SCUserAddress:        common.HexToAddress(GlobalConfig.SmcUserAddress),
	}
	privateKey, err := getPrivateKey(GlobalConfig.EthUserAddress, GlobalConfig.EthUserPassword)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	req.Sign(req, privateKey)
	payload := utils.ToJSONString(req)
	var resp api.BaseResponse
	err = call(http.MethodPost, url, payload, &resp)
	if err != nil {
		fmt.Printf("call %s with payload=%s err :%s", url, payload, err.Error())
		os.Exit(-1)
	}
	fmt.Println("MCPrepareLockout SUCCESS")
	return
}