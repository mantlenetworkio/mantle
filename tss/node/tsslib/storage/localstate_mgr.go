package storage

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/conversion"
	"github.com/mantlenetworkio/mantle/tss/node/tsslib/p2p"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/binance-chain/tss-lib/ecdsa/keygen"
	"github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const PreParams = "pre_params"

type KeygenLocalState struct {
	PubKey          string                    `json:"pub_key"`
	LocalData       keygen.LocalPartySaveData `json:"local_data"`
	ParticipantKeys []string                  `json:"participant_keys"` // the paticipant of last key gen
	LocalPartyKey   string                    `json:"local_party_key"`
	Threshold       int                       `json:"threshold"`
}

type LocalStateManager interface {
	SaveLocalState(state KeygenLocalState) error
	GetLocalState(pubKey string) (KeygenLocalState, error)
	SaveAddressBook(addressBook map[peer.ID]p2p.AddrList) error
	RetrieveP2PAddresses() (p2p.AddrList, error)
	SavePreParams(preParams *keygen.LocalPreParams) error
	GetLocalPreParams(string) (*keygen.LocalPreParams, error)
}

type FileStateMgr struct {
	folder    string
	logger    zerolog.Logger
	writeLock *sync.RWMutex
}

func NewFileStateMgr(folder string) (*FileStateMgr, error) {
	if len(folder) > 0 {
		_, err := os.Stat(folder)
		if err != nil && os.IsNotExist(err) {
			if err := os.MkdirAll(folder, os.ModePerm); err != nil {
				return nil, err
			}
		}
	}
	return &FileStateMgr{
		folder:    folder,
		logger:    log.With().Str("module", "storage").Logger(),
		writeLock: &sync.RWMutex{},
	}, nil
}

func (fsm *FileStateMgr) getFilePathName(pubKey string) (string, error) {
	ret, err := conversion.CheckKeyOnCurve(pubKey)
	if err != nil {
		return "", err
	}
	if !ret {
		return "", errors.New("invalid pubkey for file name")
	}

	localFileName := fmt.Sprintf("localstate-%s.json", pubKey)
	if len(fsm.folder) > 0 {
		return filepath.Join(fsm.folder, localFileName), nil
	}
	return localFileName, nil
}

func (fsm *FileStateMgr) getOneFilePathName() (string, error) {
	var pattern = "localstate*.json"
	if len(fsm.folder) > 0 {
		pattern = filepath.Join(fsm.folder, pattern)
	}
	files, err := filepath.Glob(pattern)
	if err != nil {
		return "", err
	}
	if len(files) > 0 {
		return files[0], nil
	}
	return "", nil

}

func (fsm *FileStateMgr) SaveLocalState(state KeygenLocalState) error {
	buf, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("fail to marshal KeygenLocalState to json: %w", err)
	}
	filePathName, err := fsm.getFilePathName(state.PubKey)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filePathName, buf, 0o655)
}

func (fsm *FileStateMgr) GetLocalState(pubKey string) (KeygenLocalState, error) {
	if len(pubKey) == 0 {
		return KeygenLocalState{}, errors.New("pub key is empty")
	}
	filePathName, err := fsm.getFilePathName(pubKey)
	if err != nil {
		return KeygenLocalState{}, err
	}
	if _, err := os.Stat(filePathName); os.IsNotExist(err) {
		return KeygenLocalState{}, err
	}

	buf, err := ioutil.ReadFile(filePathName)
	if err != nil {
		return KeygenLocalState{}, fmt.Errorf("file to read from file(%s): %w", filePathName, err)
	}
	var localState KeygenLocalState
	if err := json.Unmarshal(buf, &localState); nil != err {
		return KeygenLocalState{}, fmt.Errorf("fail to unmarshal KeygenLocalState: %w", err)
	}
	return localState, nil
}

func (fsm *FileStateMgr) GetOneLocalPreParams() (*keygen.LocalPreParams, error) {
	filePathName, err := fsm.getOneFilePathName()
	if err != nil {
		return nil, err
	}
	if len(filePathName) == 0 {
		return nil, os.ErrNotExist
	}
	if _, err := os.Stat(filePathName); os.IsNotExist(err) {
		return nil, err
	}
	buf, err := ioutil.ReadFile(filePathName)
	if err != nil {
		return nil, fmt.Errorf("file to read from file(%s): %w", filePathName, err)
	}
	var localState KeygenLocalState
	if err := json.Unmarshal(buf, &localState); nil != err {
		return nil, fmt.Errorf("fail to unmarshal KeygenLocalState: %w", err)
	}
	return &localState.LocalData.LocalPreParams, nil

}

func (fsm *FileStateMgr) SaveAddressBook(address map[peer.ID]p2p.AddrList) error {
	if len(fsm.folder) < 1 {
		return errors.New("base file path is invalid")
	}
	filePathName := filepath.Join(fsm.folder, "address_book.seed")
	var buf bytes.Buffer

	for peer, addrs := range address {
		for _, addr := range addrs {
			// we do not save the loopback addr
			if strings.Contains(addr.String(), "127.0.0.1") {
				continue
			}
			record := addr.String() + "/p2p/" + peer.String() + "\n"
			_, err := buf.WriteString(record)
			if err != nil {
				return errors.New("fail to write the record to buffer")
			}
		}
	}
	fsm.writeLock.Lock()
	defer fsm.writeLock.Unlock()
	return ioutil.WriteFile(filePathName, buf.Bytes(), 0o655)
}

func (fsm *FileStateMgr) RetrieveP2PAddresses() (p2p.AddrList, error) {
	if len(fsm.folder) < 1 {
		return nil, errors.New("base file path is invalid")
	}
	filePathName := filepath.Join(fsm.folder, "address_book.seed")

	_, err := os.Stat(filePathName)
	if err != nil {
		return nil, err
	}
	fsm.writeLock.RLock()
	input, err := ioutil.ReadFile(filePathName)
	if err != nil {
		fsm.writeLock.RUnlock()
		return nil, err
	}
	fsm.writeLock.RUnlock()
	data := strings.Split(string(input), "\n")
	var peerAddresses []ma.Multiaddr
	for _, el := range data {
		// we skip the empty entry
		if len(el) == 0 {
			continue
		}
		addr, err := ma.NewMultiaddr(el)
		if err != nil {
			return nil, fmt.Errorf("invalid address in address book %w", err)
		}
		peerAddresses = append(peerAddresses, addr)
	}
	return peerAddresses, nil
}

func (fsm *FileStateMgr) SavePreParams(preParams *keygen.LocalPreParams) error {
	buf, err := json.Marshal(preParams)
	if err != nil {
		return fmt.Errorf("fail to marshal keygen local preparams to json: %w", err)
	}
	localFileName := fmt.Sprintf("%s.json", PreParams)
	var filePathName = localFileName
	if len(fsm.folder) > 0 {
		filePathName = filepath.Join(fsm.folder, localFileName)
	}
	return ioutil.WriteFile(filePathName, buf, 0o655)
}

func (fsm *FileStateMgr) GetLocalPreParams(filePath string) (*keygen.LocalPreParams, error) {
	if len(filePath) == 0 {
		filePath = filepath.Join(fsm.folder, fmt.Sprintf("%s.json", PreParams))
	}
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	buf = bytes.Trim(buf, "\n")
	var preParam keygen.LocalPreParams
	if err = json.Unmarshal(buf, &preParam); err != nil {
		log.Error().Err(err).Msg("fail to unmarshal file content to LocalPreParams")
		return nil, err
	}
	return &preParam, nil

}
