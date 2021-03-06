package dnc

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"strings"

	"os"

	"encoding/json"

	"path/filepath"

	"io/ioutil"

	"math/big"

	"github.com/SmartMeshFoundation/distributed-notary/api"
	"github.com/SmartMeshFoundation/distributed-notary/api/userapi"
	"github.com/SmartMeshFoundation/distributed-notary/models"
	"github.com/SmartMeshFoundation/distributed-notary/service"
	"github.com/SmartMeshFoundation/distributed-notary/utils"
	"github.com/btcsuite/btcutil"
	"github.com/urfave/cli"
)

type runTime struct {
	MCName              string              `json:"mc_name"`
	Secret              string              `json:"secret"`
	SecretHash          string              `json:"secret_hash"`
	BtcLockScript       []byte              `json:"btc_lock_script"`
	BtcUserAddressBytes []byte              `json:"btc_user_address_bytes"`
	BtcTXHash           string              `json:"btc_tx_hash"`
	BtcExpiration       *big.Int            `json:"btc_expiration"`
	BtcAmount           btcutil.Amount      `json:"btc_amount"`
	LockoutInfo         *models.LockoutInfo `json:"lockout_info"`
}
type dncConfig struct {
	NotaryHost string `json:"notary_host"`
	Keystore   string `json:"keystore"`

	BtcRPCUser         string `json:"btc_rpc_user"`
	BtcRPCPass         string `json:"btc_rpc_pass"`
	BtcRPCCertFilePath string `json:"btc_rpc_cert_file_path"`
	BtcRPCEndpoint     string `json:"btc_rpc_endpoint"`
	BtcUserAddress     string `json:"btc_user_address"`

	BtcWalletRPCCertFilePath string `json:"btc_wallet_rpc_cert_file_path"`
	BtcWalletRPCEndpoint     string `json:"btc_wallet_rpc_endpoint"`

	EthUserAddress  string `json:"eth_user_address"`
	EthUserPassword string `json:"eth_user_password"`
	EthRPCEndpoint  string `json:"eth_rpc_endpoint"`

	SmcUserAddress  string `json:"smc_user_address"`
	SmcUserPassword string `json:"smc_user_password"`
	SmcRPCEndpoint  string `json:"smc_rpc_endpoint"`

	SCTokenList []service.ScTokenInfoToResponse `json:"sc_token_list"`

	RunTime *runTime `json:"run_time"`
}

// GlobalConfig :
var GlobalConfig *dncConfig

// DefaultConfig :
var DefaultConfig = &dncConfig{
	NotaryHost: "http://transport01.smartmesh.cn:8032",
	Keystore:   "../../testdata/keystore",

	BtcRPCUser:         "wuhan",
	BtcRPCPass:         "wuhan",
	BtcRPCCertFilePath: filepath.Join(os.Getenv("HOME"), ".btcd/rpc.cert"),
	BtcRPCEndpoint:     "192.168.124.13:18556",
	BtcUserAddress:     "SgEQfVdPqBS65jpSNLoddAa9kCouqqxGrY",

	BtcWalletRPCEndpoint:     "192.168.124.13:18554",
	BtcWalletRPCCertFilePath: filepath.Join(os.Getenv("HOME"), ".btcwallet/rpc.cert"),

	EthUserAddress:  "0x201b20123b3c489b47fde27ce5b451a0fa55fd60",
	EthUserPassword: "123",
	EthRPCEndpoint:  "http://106.52.171.12:18003",

	SmcUserAddress:  "0x201b20123b3c489b47fde27ce5b451a0fa55fd60",
	SmcUserPassword: "123",
	SmcRPCEndpoint:  "http://106.52.171.12:17004",
}

//var configDir = path.Join(".dnc-client")c
var configFile = filepath.Join("dnc.json")

func init() {
	var err error
	// dir
	//if !utils.Exists(configDir) {
	//	err = os.MkdirAll(configDir, os.ModePerm)
	//	if err != nil {
	//		err = fmt.Errorf("configDir:%s doesn't exist and cannot create %v", configDir, err)
	//		return
	//	}
	//}
	// config file
	var contents []byte
	// #nosec
	contents, err = ioutil.ReadFile(configFile)
	if err != nil || len(contents) == 0 {
		GlobalConfig = DefaultConfig
		updateConfigFile()
		return
	}
	GlobalConfig = new(dncConfig)
	err = json.Unmarshal(contents, GlobalConfig)
	if err != nil {
		fmt.Printf("use default config instead of wrong dnc_config in file : %s\n", configFile)
		GlobalConfig = DefaultConfig
		return
	}
}

var configCmd = cli.Command{
	Name:      "config",
	ShortName: "c",
	Usage:     "manage all config of dnc",
	Action:    configManage,
	Subcommands: cli.Commands{
		cli.Command{
			Name:      "list",
			ShortName: "l",
			Usage:     "list all config",
			Action:    listAllConfig,
		},
		cli.Command{
			Name:   "reset",
			Usage:  "reset all config",
			Action: resetAllConfig,
		},
		cli.Command{
			Name:   "refresh",
			Usage:  "refresh globalConfig.SCTokenList from notary",
			Action: refreshSCTokenList,
		},
	},
}

func configManage(ctx *cli.Context) error {
	for _, param := range ctx.Args() {
		s := strings.Split(param, "=")
		if len(s) != 2 {
			fmt.Printf("wrong param : %s\n", param)
			os.Exit(-1)
		}
		switch s[0] {
		case "nh", "notary-host":
			GlobalConfig.NotaryHost = s[1]
		case "keystore":
			GlobalConfig.Keystore = s[1]

		case "eua", "eth-user-address":
			GlobalConfig.EthUserAddress = s[1]
		case "eup", "eth-user-password":
			GlobalConfig.EthUserPassword = s[1]
		case "eth", "eth-rpc-endpoint":
			GlobalConfig.EthRPCEndpoint = s[1]

		case "sua", "smc-user-address":
			GlobalConfig.SmcUserAddress = s[1]
		case "sup", "smc-user-password":
			GlobalConfig.SmcUserPassword = s[1]
		case "smc", "smc-rpc-endpoint":
			GlobalConfig.SmcRPCEndpoint = s[1]
		default:
			fmt.Printf("wrong param : %s\n", param)
			os.Exit(-1)
		}
	}
	updateConfigFile()
	fmt.Println(utils.ToJSONStringFormat(GlobalConfig))
	return nil
}

func listAllConfig(ctx *cli.Context) error {
	fmt.Println(utils.ToJSONStringFormat(GlobalConfig))
	return nil
}

//ListAllConfig :
func ListAllConfig() string {
	return utils.ToJSONStringFormat(GlobalConfig)
}
func resetAllConfig(ctx *cli.Context) error {
	GlobalConfig = DefaultConfig
	updateConfigFile()
	return nil
}

func refreshSCTokenList(ctx *cli.Context) (err error) {
	return RefreshSCTokenList()
}

// RefreshSCTokenList :
func RefreshSCTokenList() (err error) {
	if GlobalConfig.NotaryHost == "" {
		err = fmt.Errorf("must set globalConfig.NotaryHost first")
		fmt.Println(err)
		return
	}
	var resp api.BaseResponse
	url := GlobalConfig.NotaryHost + userapi.APIName2URLMap[userapi.APIUserNameGetSCTokenList]
	err = call(http.MethodGet, url, "", &resp)
	if err != nil {
		err = fmt.Errorf("call %s err : %s", url, err.Error())
		fmt.Println(err)
		return
	}
	GlobalConfig.SCTokenList = []service.ScTokenInfoToResponse{}
	err = resp.ParseData(&GlobalConfig.SCTokenList)
	if err != nil {
		err = fmt.Errorf("parse data err : %s", err.Error())
		fmt.Println(err)
	}
	updateConfigFile()
	fmt.Println(utils.ToJSONStringFormat(GlobalConfig))
	return
}

func updateConfigFile() {
	err := ioutil.WriteFile(configFile, []byte(utils.ToJSONStringFormat(GlobalConfig)), 0777)
	if err != nil {
		panic(err)
	}
}

func call(method string, url string, payload string, responseTo api.Response) (err error) {
	var reqBody io.Reader
	if payload == "" {
		reqBody = nil
	} else {
		reqBody = strings.NewReader(payload)
	}
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	defer func() {
		if req.Body != nil {
			/* #nosec */
			req.Body.Close()
		}
		if resp != nil && resp.Body != nil {
			/* #nosec */
			resp.Body.Close()
		}
	}()
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http request err : status code = %d", resp.StatusCode)
	}
	var buf [4096 * 1024]byte
	n := 0
	n, err = resp.Body.Read(buf[:])
	if err != nil && err.Error() == "EOF" {
		err = nil
	}
	respBody := buf[:n]
	if responseTo == nil {
		responseTo = new(api.BaseResponse)
	}
	err = json.Unmarshal(respBody, responseTo)
	if err != nil {
		return
	}
	if responseTo.GetErrorCode() != api.ErrorCodeSuccess {
		err = errors.New(responseTo.GetErrorMsg())
	}
	return
}
