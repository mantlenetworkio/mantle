package tsslib

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	bkeygen "github.com/binance-chain/tss-lib/ecdsa/keygen"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	tssconfig "github.com/mantlenetworkio/mantle/tss/common"
	common2 "github.com/mantlenetworkio/mantle/tss/node/tsslib/common"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/conversion"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/keygen"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/keysign"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/monitor"
	p2p2 "github.com/mantlenetworkio/mantle/tss/node/tsslib/p2p"
	storage2 "github.com/mantlenetworkio/mantle/tss/node/tsslib/storage"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"sort"
	"strings"
	"sync"
)

type TssServer struct {
	conf             common2.TssConfig
	logger           zerolog.Logger
	p2pCommunication *p2p2.Communication
	localNodePubKey  string
	participants     map[string][]string
	preParams        *bkeygen.LocalPreParams
	tssKeyGenLocker  *sync.Mutex
	stopChan         chan struct{}
	stateManager     storage2.LocalStateManager
	secretsManager   storage2.SecretsManager
	shamirManager    storage2.ShamirManager
	privateKey       *ecdsa.PrivateKey
	tssMetrics       *monitor.Metric
	secretsEnable    bool
	shamirEnable     bool
}

func NewTss(
	cmdBootstrapPeers string,
	waitFullConnected bool,
	p2pPort int,
	priKey *ecdsa.PrivateKey,
	storageFolder string,
	conf common2.TssConfig,
	preParamsFile string,
	externalIP string,
	secretsEnable bool,
	secretId string,
	shamirConfig tssconfig.ShamirConfig,
) (*TssServer, error) {

	pubkey := crypto.CompressPubkey(&priKey.PublicKey)
	pubkeyHex := hex.EncodeToString(pubkey)
	log.Info().Msgf("pub key is (%s) \n", pubkeyHex)

	peerId, err := conversion.GetPeerIDFromPubKey(pubkeyHex)
	if err != nil {
		log.Error().Err(err).Msg("ERROR: fail to get peer id by pub key")
	}
	log.Info().Msgf("peer id is (%s) \n", peerId)
	stateManager, err := storage2.NewFileStateMgr(storageFolder)
	if err != nil {
		return nil, errors.New("fail to create file state manager")
	}
	var secretsManager *storage2.SecretsMgr
	var shamirManager *storage2.ShamirMgr
	if shamirConfig.Enable {
		shamirManager, err = storage2.NewShamirMgr(shamirConfig)
		if err != nil {
			log.Error().Err(err).Msgf("fail to create shamir manager :%w", err)
			return nil, errors.New("fail to create shamir manager")
		}
	} else if secretsEnable {
		secretsManager, err = storage2.NewSecretsMgr(secretId)
		if err != nil {
			log.Error().Err(err).Msgf("fail to create secrets manager :%w", err)
			return nil, errors.New("fail to create secrets manager")
		}
	}
	cmdBootstrapPeerS := strings.Split(cmdBootstrapPeers, ",")
	var bootstrapPeers p2p2.AddrList
	savedPeers, err := stateManager.RetrieveP2PAddresses()
	if err == nil {
		bootstrapPeers = savedPeers
	}
	for _, peerAdd := range cmdBootstrapPeerS {
		if peerAdd != "" {
			if err = bootstrapPeers.Set(strings.TrimSpace(peerAdd)); err != nil {
				return nil, err
			}
		}
	}

	comm, err := p2p2.NewCommunication(bootstrapPeers, p2pPort, externalIP, waitFullConnected)
	if err != nil {
		return nil, fmt.Errorf("fail to create communication layer: %w", err)
	}
	// When using the keygen party it is recommended that you pre-compute the
	// "safe primes" and Paillier secret beforehand because this can take some
	// time.
	// This code will generate those parameters using a concurrency limit equal
	// to the number of available CPU cores.

	// if there is no preParams file specified by user, and no default preParams file path exists, then generate a new preParams for user.
	var preParams *bkeygen.LocalPreParams
	if len(preParamsFile) != 0 {
		preParams, err = stateManager.GetLocalPreParams(preParamsFile)
		if err != nil {
			return nil, fmt.Errorf("fail to generate pre parameters: %w", err)
		}
	} else if shamirConfig.Enable {
		preParams = shamirManager.GetOneLocalState()
		if preParams == nil {
			log.Info().Msg("start to generate pre params...")
			preParams, err = bkeygen.GeneratePreParams(conf.PreParamTimeout)
			if err != nil {
				return nil, fmt.Errorf("fail to generate pre parameters: %w", err)
			}
		}
	} else if secretsEnable {
		preParams = secretsManager.GetOneLocalState()
		if preParams == nil {
			log.Info().Msg("start to generate pre params...")
			preParams, err = bkeygen.GeneratePreParams(conf.PreParamTimeout)
			if err != nil {
				return nil, fmt.Errorf("fail to generate pre parameters: %w", err)
			}
		}
	} else {
		preParams, err = stateManager.GetOneLocalPreParams()
		if err != nil && os.IsNotExist(err) {
			log.Info().Msg("start to generate pre params...")
			preParams, err = bkeygen.GeneratePreParams(conf.PreParamTimeout)
			if err != nil {
				return nil, fmt.Errorf("fail to generate pre parameters: %w", err)
			}
		} else if err != nil {
			return nil, err
		}
	}

	if !preParams.Validate() {
		return nil, errors.New("invalid preparams")
	}

	if err := comm.Start(crypto.FromECDSA(priKey)); nil != err {
		return nil, fmt.Errorf("fail to start p2p network: %w", err)
	}
	//sn := keysign.NewSignatureNotifier(comm.GetHost())
	metrics := monitor.NewMetric()
	if conf.EnableMonitor {
		metrics.Enable()
	}
	tssServer := TssServer{
		conf:             conf,
		logger:           log.With().Str("module", "tss").Logger(),
		p2pCommunication: comm,
		localNodePubKey:  pubkeyHex,
		participants:     make(map[string][]string),
		preParams:        preParams,
		tssKeyGenLocker:  &sync.Mutex{},
		stopChan:         make(chan struct{}),
		stateManager:     stateManager,
		secretsManager:   secretsManager,
		shamirManager:    shamirManager,
		privateKey:       priKey,
		tssMetrics:       metrics,
		secretsEnable:    secretsEnable,
		shamirEnable:     shamirConfig.Enable,
	}

	return &tssServer, nil
}

func (t *TssServer) Start() error {
	log.Info().Msg("Starting the TSS servers")
	return nil
}

// Stop Tss server
func (t *TssServer) Stop() {
	close(t.stopChan)
	// stop the p2p and finish the p2p wait group
	err := t.p2pCommunication.Stop()
	if err != nil {
		t.logger.Error().Msgf("error in shutdown the p2p server")
	}
	log.Info().Msg("The Tss and p2p server has been stopped successfully")
}

func (t *TssServer) GetLocalPeerID() string {
	return t.p2pCommunication.GetLocalPeerID()
}

func (t *TssServer) requestToMsgId(request interface{}) (string, error) {
	var dat []byte
	var keys []string
	switch value := request.(type) {
	case keygen.Request:
		keys = value.Keys
	case keysign.Request:
		dat = value.Message
		keys = value.SignerPubKeys
	default:
		t.logger.Error().Msg("unknown request type")
		return "", errors.New("unknown request type")
	}
	keyAccumulation := ""
	sort.Strings(keys)
	for _, el := range keys {
		keyAccumulation += el
	}
	dat = append(dat, []byte(keyAccumulation)...)
	return common2.MsgToHashString(dat)
}

func (t *TssServer) requestCheck(request interface{}) error {
	var threshold int
	switch value := request.(type) {
	case keygen.Request:
		threshold = value.ThresHold
		if len(value.Keys) <= threshold {
			t.logger.Error().Msg("check params : pub_keys size is smaller than threshold !")
			return errors.New("check params : pub_keys size is smaller than threshold")
		}
	case keysign.Request:
		myPk, err := conversion.GetPubKeyFromPeerID(t.p2pCommunication.GetHost().ID().String())
		if err != nil {
			t.logger.Info().Msgf("fail to convert the p2p id(%s) to pubkey", t.p2pCommunication.GetHost().ID().String())
			return err
		}
		isSignMember := false
		for _, el := range value.SignerPubKeys {
			if myPk == el {
				isSignMember = true
				break
			}
		}
		if !isSignMember {
			t.logger.Info().Msgf("we(%s) are not the active signer", t.p2pCommunication.GetHost().ID().String())
			return errors.New("not active signer")
		}

	default:
		t.logger.Error().Msg("unknown request type")
		return errors.New("unknown request type")
	}
	return nil

}

func (t *TssServer) ExportPeerAddress() map[string]string {
	ret := make(map[string]string)
	rs := t.p2pCommunication.ExportPeerAddress()
	for k, v := range rs {
		ret[k.String()] = v.String()
	}
	return ret
}

func (t *TssServer) CheckPubKeys(signerPubKeys []string, threshold int) ([]peer.ID, error) {
	if len(signerPubKeys) < threshold {
		t.logger.Error().Msg("the keys size is smaller than threshold, please commit right params!")
		return nil, errors.New("keys size is smaller than threshold number! ")
	}

	peersID, err := conversion.GetPeerIDsFromPubKeys(signerPubKeys)
	if err != nil {
		return nil, errors.New("fail to convert the public key to peer ID")
	}
	return peersID, nil

}
