package configs

import (
	"bcos_lottery_demo/contract"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"reflect"
	"runtime"
)

var Conf = Config{}

type (
	Config struct {
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		LogPath         string `yaml:"logPath"`
		logLevel        string `yaml:"logLevel"`
		ContractAddress string `yaml:"contractAddress"`
		IsGm            bool   `yaml:"isGm"`
		ConfigFile      string `yaml:"configFile"`
		Client          *contract.LotteryRecordSession
	}
)

func init() {
	fmt.Println("**********start load config")
	dataBytes, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		os.Exit(2)
	}
	if err := yaml.Unmarshal(dataBytes, &Conf); err != nil {
		fmt.Println("read config error", err)
		os.Exit(2)
	}
	if reflect.DeepEqual(Conf, &Config{}) {
		fmt.Println("get empty config,please check")
		os.Exit(2)
	}
	// 获取 sdk配置
	if err = Conf.initSdkClient(); err != nil {
		fmt.Printf(" init  bcos client error===%v===\n", err)
		os.Exit(2)
	}
}

func (cf *Config) String() string {
	b, err := json.Marshal(*cf)
	if err != nil {
		return fmt.Sprintf("%+v", *cf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *cf)
	}
	return out.String()
}

func (cf *Config) initSdkClient() error {
	var filePath string
	confFile := cf.ConfigFile
	dir, _ := os.Getwd()
	if runtime.GOOS == "linux" {
		filePath = fmt.Sprintf("%s/%s/%s", dir, "configs", confFile)
	} else {
		filePath = fmt.Sprintf("%s\\%s\\%s", dir, "configs", confFile)
	}
	configs, err := conf.ParseConfigFile(filePath)
	if err != nil {
		log.Panicf("parse config error: %v", err)
		return err
	}
	client, err := client.Dial(&configs[0])
	if err != nil {
		return fmt.Errorf("failed to initialize FISCO BCOS client: %v", err)
	}
	contractAddress := common.HexToAddress(cf.ContractAddress)
	instance, err := contract.NewLotteryRecord(contractAddress, client)
	if err != nil {
		log.Panicf("NewHelloWorld error: %v", err)
		return err
	}
	clientSession := &contract.LotteryRecordSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
	cf.Client = clientSession
	return nil
}
