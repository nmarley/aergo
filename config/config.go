/**
 *  @file
 *  @copyright defined in aergo/LICENSE.txt
 */

package config

import (
	"runtime"

	"github.com/aergoio/aergo-lib/config"
	"github.com/aergoio/aergo/types"
)

type ServerContext struct {
	config.BaseContext
}

func NewServerContext(homePath string, configFilePath string) *ServerContext {
	serverCxt := &ServerContext{}
	serverCxt.BaseContext = config.NewBaseContext(serverCxt, homePath, configFilePath, EnvironmentPrefix)
	serverCxt.Vc.SetDefault("blockchain.zerofee", serverCxt.GetDefaultBlockchainConfig().ZeroFee)
	serverCxt.Vc.SetDefault("sql.maxdbsize", serverCxt.GetDefaultSQLConfig().MaxDbSize)
	return serverCxt
}

func (ctx *ServerContext) GetHomePath() string {
	return defaultAergoHomePath
}

func (ctx *ServerContext) GetConfigFileName() string {
	return defaultAergoConfigFileName
}

func (ctx *ServerContext) GetTemplate() string {
	return tomlConfigFileTemplate
}

func (ctx *ServerContext) GetDefaultConfig() interface{} {
	return &Config{
		BaseConfig: ctx.GetDefaultBaseConfig(),
		RPC:        ctx.GetDefaultRPCConfig(),
		P2P:        ctx.GetDefaultP2PConfig(),
		Blockchain: ctx.GetDefaultBlockchainConfig(),
		Mempool:    ctx.GetDefaultMempoolConfig(),
		Consensus:  ctx.GetDefaultConsensusConfig(),
		Monitor:    ctx.GetDefaultMonitorConfig(),
		Account:    ctx.GetDefaultAccountConfig(),
		Auth:       ctx.GetDefaultAuthConfig(),
		Polaris:    ctx.GetDefaultPolarisConfig(),
		Hardfork:   ctx.GetDefaultHardforkConfig(),
		SQL:        ctx.GetDefaultSQLConfig(),
	}
}

func (ctx *ServerContext) GetDefaultBaseConfig() BaseConfig {
	return BaseConfig{
		DataDir:        ctx.ExpandPathEnv("$HOME/data"),
		DbType:         "badgerdb",
		EnableProfile:  false,
		ProfilePort:    6060,
		EnableTestmode: false,
		Personal:       true,
		AuthDir:        ctx.ExpandPathEnv("$HOME/auth"),
	}
}

func (ctx *ServerContext) GetDefaultRPCConfig() *RPCConfig {
	return &RPCConfig{
		NetServiceAddr:  "127.0.0.1",
		NetServicePort:  7845,
		NetServiceTrace: false,
		NSKey:           "",
	}
}

func (ctx *ServerContext) GetDefaultP2PConfig() *P2PConfig {
	return &P2PConfig{
		NetProtocolAddr: "",
		NetProtocolPort: 7846,
		NPBindAddr:      "",
		NPBindPort:      -1,
		NPEnableTLS:     false,
		NPCert:          "",
		NPKey:           "",
		NPAddPeers:      nil,
		NPDiscoverPeers: true,
		NPMaxPeers:      100,
		NPPeerPool:      100,
		NPUsePolaris:    true,
		NPExposeSelf:    true,
		PeerRole:        "",
	}
}

func (ctx *ServerContext) GetDefaultAuthConfig() *AuthConfig {
	return &AuthConfig{
		EnableLocalConf: false,
	}
}

func (ctx *ServerContext) GetDefaultPolarisConfig() *PolarisConfig {
	return &PolarisConfig{
		GenesisFile:     "",
		AllowPrivate:    false,
		EnableBlacklist: true,
	}
}

func (ctx *ServerContext) GetDefaultBlockchainConfig() *BlockchainConfig {
	return &BlockchainConfig{
		MaxBlockSize:     types.DefaultMaxBlockSize,
		CoinbaseAccount:  "",
		MaxAnchorCount:   20,
		VerifierCount:    types.DefaultVerifierCnt,
		ForceResetHeight: 0,
		ZeroFee:          true, // deprecated
		StateTrace:       0,
	}
}

func (ctx *ServerContext) GetDefaultMempoolConfig() *MempoolConfig {
	return &MempoolConfig{
		ShowMetrics:    false,
		EnableFadeout:  false,
		FadeoutPeriod:  types.DefaultEvictPeriod,
		VerifierNumber: runtime.NumCPU(),
		DumpFilePath:   ctx.ExpandPathEnv("$HOME/mempool.dump"),
	}
}

func (ctx *ServerContext) GetDefaultConsensusConfig() *ConsensusConfig {
	return &ConsensusConfig{
		BlockInterval: 1,
	}
}

func (ctx *ServerContext) GetDefaultMonitorConfig() *MonitorConfig {
	return &MonitorConfig{
		ServerProtocol: "",
		ServerEndpoint: "",
	}
}

func (ctx *ServerContext) GetDefaultAccountConfig() *AccountConfig {
	return &AccountConfig{
		UnlockTimeout: 60,
	}
}

func (ctx *ServerContext) GetDefaultHardforkConfig() *HardforkConfig {
	return AllEnabledHardforkConfig
}

func (ctx *ServerContext) GetDefaultSQLConfig() *SQLConfig {
	return &SQLConfig{
		MaxDbSize: 20,
	}
}
