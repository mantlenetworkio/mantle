package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	bkeygen "github.com/binance-chain/tss-lib/ecdsa/keygen"
	"github.com/rs/zerolog/log"
	"time"
)

const CtxTimeout = 30 * time.Second

type SecretsManager interface {
	Save() error
	PutKeyFile(state KeygenLocalState) error
	GetKeyFile(pubKey string) (KeygenLocalState, error)
}

type SecretsMgr struct {
	secretId string
	keys     map[string]KeygenLocalState
	client   *secretsmanager.Client
}

func NewSecretsMgr(secretId string) (*SecretsMgr, error) {
	log.Debug().Msg("start to load default config for secrets manager")
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Error().Err(err).Msg("fail to load config for secrets manager")
		return nil, errors.New("fail to new secrets manager instance")
	}
	client := secretsmanager.NewFromConfig(cfg)
	params := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretId),
	}
	result, err := client.GetSecretValue(ctx, params)
	if err != nil {

		log.Error().Err(err).Msgf("fail to get secret value form aws secrets manager :%w", err)
		return nil, err
	}
	var keys map[string]KeygenLocalState
	if err := json.Unmarshal(result.SecretBinary, &keys); err != nil {
		log.Error().Err(err).Msgf("fail to unmarshal data to map :%w", err)
		keys = map[string]KeygenLocalState{}
	}

	return &SecretsMgr{
		secretId: secretId,
		keys:     keys,
		client:   client,
	}, nil

}

func (sm *SecretsMgr) PutKeyFile(stat KeygenLocalState) error {
	_, ok := sm.keys[stat.PubKey]
	if !ok {
		sm.keys[stat.PubKey] = stat
	}
	return nil
}

func (sm *SecretsMgr) Save() error {
	if len(sm.keys) == 0 {
		log.Warn().Msg("secrets storage size is 0 ,don't need to put to aws secrets manager")
		return nil
	}
	buf, err := json.Marshal(sm.keys)
	if err != nil {
		log.Error().Err(err).Msgf("fail to marshal secrets keys map to json: %w", err)
		return err
	}
	params := &secretsmanager.UpdateSecretInput{
		SecretId:     aws.String(sm.secretId),
		SecretBinary: buf,
	}
	output, err := sm.client.UpdateSecret(context.TODO(), params)
	if err != nil {
		log.Error().Err(err).Msgf("fail to put data to aws's secrets manager : %w", err)
		return err
	}
	log.Info().Msgf("put data to aws's secrets manager success, version is:%s", output.VersionId)
	return nil
}

func (sm *SecretsMgr) GetKeyFile(pubKey string) (KeygenLocalState, error) {
	value, ok := sm.keys[pubKey]
	if !ok {
		log.Warn().Msgf("can not find keygenlocalstate from storage by this pubKey (%s)", pubKey)
		return KeygenLocalState{}, nil
	}
	return value, nil
}

func (sm *SecretsMgr) GetOneLocalState() *bkeygen.LocalPreParams {
	var preParams *bkeygen.LocalPreParams
	if len(sm.keys) > 0 {
		for _, value := range sm.keys {
			preParams = &value.LocalData.LocalPreParams
			break
		}
	} else {
		return nil
	}
	return preParams
}
