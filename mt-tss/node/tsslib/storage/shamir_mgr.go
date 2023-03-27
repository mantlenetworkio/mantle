package storage

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	sssas "github.com/SSSaaS/sssa-golang"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	bkeygen "github.com/bnb-chain/tss-lib/ecdsa/keygen"
	"github.com/google/uuid"
	nodeconfig "github.com/mantlenetworkio/mantle/mt-tss/common"
	"github.com/rs/zerolog/log"
)

const (
	n = 5
	k = 4
)

type ShamirManager interface {
	PutKeyFile(state KeygenLocalState) error
	GetKeyFile(pubKey, localPartyKey string) (KeygenLocalState, error)
}

type (
	ShamirMgr struct {
		shamirConfig nodeconfig.ShamirConfig
		keys         map[string]KeygenLocalState
	}

	Share struct {
		DexCipher   []byte `json:"dex_cipher"`
		ShareCipher []byte `json:"share_cipher"`
	}
)

func NewShamirMgr(config nodeconfig.ShamirConfig) (*ShamirMgr, error) {
	log.Debug().Msg("create shamir instance ")
	var keys = map[string]KeygenLocalState{}
	return &ShamirMgr{
		shamirConfig: config,
		keys:         keys,
	}, nil

}

func (sh *ShamirMgr) PutKeyFile(stat KeygenLocalState) error {
	log.Info().Msg("start to storage new keygen ")
	_, ok := sh.keys[stat.PubKey]
	if !ok {
		sh.keys[stat.PubKey] = stat
		err := sh.SaveEncrypt(stat)
		if err != nil {
			log.Error().Err(err).Msg("put key file failed")
			return err
		}
	}
	return nil
}

func (sh *ShamirMgr) GetKeyFile(pubKey, localPartyKey string) (KeygenLocalState, error) {
	value, ok := sh.keys[pubKey]
	if !ok {
		log.Warn().Msgf("can not find keygenlocalstate from memory storage by this pubKey (%s),need to get from aws", pubKey)
		keygen, err := sh.GetDecrypt(pubKey, localPartyKey)
		if err != nil {
			log.Error().Err(err).Msg("failed to get keygen local state from aws")
			return KeygenLocalState{}, err
		}
		var keygenlocalstate KeygenLocalState
		if err := json.Unmarshal([]byte(keygen), &keygenlocalstate); err != nil {
			log.Error().Err(err).Msgf("fail to unmarshal data to map :%w", err)
			return KeygenLocalState{}, err
		}
		//缓存在内存中
		sh.keys[pubKey] = keygenlocalstate

		return keygenlocalstate, nil
	}
	return value, nil
}

func (sh *ShamirMgr) SaveEncrypt(stat KeygenLocalState) error {
	log.Info().Msg("start to save keygen to aws")
	bytes, err := json.Marshal(stat)
	if err != nil {
		log.Error().Err(err).Msg("keygen localstate json marshal failed ")
		return err
	}
	//1-通过Shamir密钥分片算法将私钥分成k/n份
	log.Info().Msg("1-start to use shamir to split keygen")
	shares, err := sssas.Create(k, n, string(bytes))
	if err != nil {
		log.Error().Err(err).Msg("shamir create failed! ")
		return err
	}

	//2-将分片好的私钥进行异或，异或的key可以随机生成
	log.Info().Msg("2-start to xor the shares.")
	var shares_xor []string
	if len(sh.shamirConfig.Xor) == 0 {
		log.Error().Msg("Shamir.Xor value is empty ,please set it! ")
		return errors.New("Shamir.Xor value is empty ,please set it! ")
	}
	for i := 0; i < n; i++ {
		xor := ByXOR(shares[i], sh.shamirConfig.Xor)
		shares_xor = append(shares_xor, xor)
	}
	//3-生成aes密文秘钥
	log.Info().Msg("3-start to generate data key")
	dataKey, err := sh.GenerateDataKey()
	if err != nil {
		return err
	}

	//4-将异或后的分片私钥通过aes加密
	log.Info().Msg("4-start to encrypt shares")
	var shares_xor_encrypted [][]byte
	for i := 0; i < n; i++ {
		cipher, err := AesEncrypt([]byte(shares_xor[i]), dataKey.Plaintext)
		if err != nil {
			return err
		}
		var share = Share{
			DexCipher:   dataKey.CiphertextBlob,
			ShareCipher: cipher,
		}
		share_bytes, err := json.Marshal(share)
		if err != nil {
			log.Error().Err(err).Msg("share struct json marshal failed ")
			return err
		}
		shares_xor_encrypted = append(shares_xor_encrypted, share_bytes)
	}

	//6-将加密的分片私钥密文存入S3和SM，分片索引值奇数存S3,偶数存SM。

	var shares_s3 [][]byte
	var shares_sm [][]byte

	var buckets = strings.Split(sh.shamirConfig.S3.Buckets, ",")
	var secrets = strings.Split(sh.shamirConfig.Sm.SecretIds, ",")
	sort.Strings(buckets)
	sort.Strings(secrets)
	if len(buckets) == 0 && len(secrets) == 0 {
		log.Error().Msg("S3 config doesn't have buckets value,Sm config doesn't have secrets value! ")
		return errors.New("S3 config doesn't have buckets value,Sm config doesn't have secrets value! ")
	} else if len(buckets) == 0 {
		shares_sm = shares_xor_encrypted
	} else if len(secrets) == 0 {
		shares_s3 = shares_xor_encrypted
	} else {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				shares_s3 = append(shares_s3, shares_xor_encrypted[i])
			} else {
				shares_sm = append(shares_sm, shares_xor_encrypted[i])
			}
		}
	}

	if shares_s3_len := len(shares_s3); shares_s3_len > 0 {
		buckets_len := len(buckets)
		log.Info().Msg("6-start to new s3 uploader ")
		s3_uploader, err := sh.NewS3Uploader()
		if err != nil {
			return err
		}
		log.Info().Msg("6-start to uploader ")

		for i := 0; i < shares_s3_len; i++ {
			var filename = stat.PubKey + ":" + stat.LocalPartyKey + ":" + strconv.Itoa(i)
			err = uploadToS3(buckets[i%buckets_len], filename, shares_s3[i], *s3_uploader)
			if err != nil {
				return err
			}
		}
	}

	if shares_sm_len := len(shares_sm); shares_sm_len > 0 {
		secrets_len := len(secrets)
		log.Info().Msg("6-start to new s3 secrets manager ")

		svc, err := sh.NewSecretsManager()
		if err != nil {
			return err
		}
		log.Info().Msg("6-start to put value to sm ")

		for i := 0; i < shares_sm_len; i++ {
			var key = stat.PubKey + ":" + strconv.Itoa(i)
			err := putSecretValue(secrets[i%secrets_len], key, shares_sm[i], *svc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (sh *ShamirMgr) GetDecrypt(pubKey, localPartyKey string) (string, error) {
	var buckets = strings.Split(sh.shamirConfig.S3.Buckets, ",")
	var secrets = strings.Split(sh.shamirConfig.Sm.SecretIds, ",")
	sort.Strings(buckets)
	sort.Strings(secrets)
	var sm_len, s3_len int
	if len(buckets) == 0 && len(secrets) == 0 {
		log.Error().Msg("S3 config doesn't have buckets value,Sm config doesn't have secrets value! ")
		return "", errors.New("S3 config doesn't have buckets value,Sm config doesn't have secrets value! ")
	} else if len(buckets) == 0 {
		sm_len = n
	} else if len(secrets) == 0 {
		s3_len = n
	} else {
		s3_len = n/2 + n%2
		sm_len = n / 2
	}

	//1-从SM、S3获取分片私钥密文
	var shares_xor_encrypted [][]byte
	if s3_len > 0 {
		buckets_len := len(buckets)
		s3_downloader, err := sh.NewS3Downloader()
		if err != nil {
			return "", err
		}
		for i := 0; i < s3_len; i++ {
			var filename = pubKey + ":" + localPartyKey + ":" + strconv.Itoa(i)
			cipher, _ := getFromS3(buckets[i%buckets_len], filename, *s3_downloader)
			if cipher != nil {
				shares_xor_encrypted = append(shares_xor_encrypted, cipher)
			}
		}
	}
	if sm_len > 0 {
		secrets_len := len(secrets)
		svc, err := sh.NewSecretsManager()
		if err != nil {
			return "", err
		}
		for i := 0; i < sm_len; i++ {
			var key = pubKey + ":" + strconv.Itoa(i)
			value, _ := getSecretValue(secrets[i%secrets_len], key, *svc)
			if value != nil {
				shares_xor_encrypted = append(shares_xor_encrypted, value)
			}
		}
	}
	if len(shares_xor_encrypted) < k {
		return "", errors.New("shares number " + strconv.Itoa(len(shares_xor_encrypted)) + " is smaller than shamir threshold " + strconv.Itoa(k))
	}

	//2-将分片私钥通过AES解密
	var shares_xor []string
	var plainDex []byte
	for i := 0; i < len(shares_xor_encrypted); i++ {
		var share Share
		if err := json.Unmarshal(shares_xor_encrypted[i], &share); err != nil {
			log.Error().Err(err).Msgf("shamir json unmarshal share index %d struct failed", i)
			return "", err
		}

		//2-1 获取dex密文秘钥，通过kms解密得到dex明文
		if plainDex == nil {
			result, err := sh.kmsDecrypt(share.DexCipher)
			if err != nil {
				return "", err
			}
			plainDex = result
		}
		//2-2 通过Aes解密
		shareCipher, err := AesDecrypt(share.ShareCipher, plainDex)
		if err != nil {
			return "", err
		}
		shares_xor = append(shares_xor, string(shareCipher))
	}

	//3-将解密后的分片私钥进行异或，异或的key与分片过程中使用的key一致
	var shares_recover []string
	for i := 0; i < n; i++ {
		xor := ByXOR(shares_xor[i], sh.shamirConfig.Xor)
		shares_recover = append(shares_recover, xor)
	}

	//4-将解密后的N个分片私钥通过Shamir算法恢复成原有私钥
	privateKey_recover, err := sssas.Combine(shares_recover)
	if err != nil {
		log.Error().Err(err).Msg("shamir combine failed!")
		return "", err
	}
	return privateKey_recover, nil
}

func ByXOR(message, keywords string) string {
	messageLen := len(message)
	keywordsLen := len(keywords)
	result := ""

	for i := 0; i < messageLen; i++ {
		result += string(message[i] ^ keywords[i%keywordsLen])
	}
	return result
}

func (sh *ShamirMgr) kmsEncrypt(plaintext []byte) ([]byte, error) {
	sess, err := NewSession(sh.shamirConfig.Kms.Region, sh.shamirConfig.Kms.Aksk.Id, sh.shamirConfig.Kms.Aksk.Secret)
	if err != nil {
		return nil, err
	}
	// Create KMS service client
	svc := kms.New(sess)

	// Encrypt the data
	result, err := svc.Encrypt(&kms.EncryptInput{
		KeyId:     aws.String(sh.shamirConfig.Kms.KeyId),
		Plaintext: plaintext,
	})

	if err != nil {
		log.Error().Err(err).Msg("kms fail to encrypt data ")
		return nil, err
	}

	return result.CiphertextBlob, nil
}

func (sh *ShamirMgr) kmsDecrypt(ciphertext []byte) ([]byte, error) {
	sess, err := NewSession(sh.shamirConfig.Kms.Region, sh.shamirConfig.Kms.Aksk.Id, sh.shamirConfig.Kms.Aksk.Secret)
	if err != nil {
		return nil, err
	}
	// Create KMS service client
	svc := kms.New(sess)
	// Decrypt the data

	result, err := svc.Decrypt(&kms.DecryptInput{
		KeyId:          aws.String(sh.shamirConfig.Kms.KeyId),
		CiphertextBlob: ciphertext,
	})
	if err != nil {
		log.Error().Err(err).Msg("kms decrypt failed ")
		return nil, err
	}
	return result.Plaintext, nil
}

func (sh *ShamirMgr) NewSecretsManager() (*secretsmanager.SecretsManager, error) {
	sess, err := NewSession(sh.shamirConfig.Sm.Region, sh.shamirConfig.Sm.Aksk.Id, sh.shamirConfig.Sm.Aksk.Secret)
	if err != nil {
		return nil, err
	}
	svc := secretsmanager.New(sess)
	return svc, nil
}

func putSecretValue(secretId, key string, value []byte, svc secretsmanager.SecretsManager) error {

	uuid, err := createUUID(key)
	if err != nil {
		return err
	}

	input := &secretsmanager.PutSecretValueInput{
		SecretId:           aws.String(secretId),
		SecretBinary:       value,
		ClientRequestToken: aws.String(uuid),
	}

	_, err = svc.PutSecretValue(input)
	if err != nil {
		log.Error().Err(err).Msgf("put secret failed, secretId %s  ", secretId)
		return err
	}
	log.Info().Msgf("Put secret successful!")
	return nil
}

func getSecretValue(secretId, key string, svc secretsmanager.SecretsManager) ([]byte, error) {

	uuid, err := createUUID(key)
	if err != nil {
		return nil, err
	}

	input := &secretsmanager.GetSecretValueInput{
		SecretId:  aws.String(secretId),
		VersionId: aws.String(uuid),
	}

	result, err := svc.GetSecretValue(input)
	if err != nil {
		log.Error().Err(err).Msgf("get secret value failed, (%s)", err)
		return nil, err
	}
	log.Info().Msgf("%s get secretvalue successful!", key)
	return result.SecretBinary, nil
}

func (sh *ShamirMgr) NewS3Downloader() (*s3manager.Downloader, error) {
	sess, err := NewSession(sh.shamirConfig.S3.Region, sh.shamirConfig.S3.Aksk.Id, sh.shamirConfig.S3.Aksk.Secret)
	if err != nil {
		return nil, err
	}

	downloader := s3manager.NewDownloader(sess)
	return downloader, nil
}

func (sh *ShamirMgr) NewS3Uploader() (*s3manager.Uploader, error) {
	sess, err := NewSession(sh.shamirConfig.S3.Region, sh.shamirConfig.S3.Aksk.Id, sh.shamirConfig.S3.Aksk.Secret)
	if err != nil {
		return nil, err
	}

	uploader := s3manager.NewUploader(sess)
	return uploader, nil
}

func uploadToS3(bucket, filename string, file []byte, uploader s3manager.Uploader) error {

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   bytes.NewReader(file),
	})

	if err != nil {
		log.Error().Err(err).Msgf("Unable to upload %q to %q, %v", filename, bucket, err)
		return err
	}
	log.Info().Msgf("Successfully uploaded %q to %q\n", filename, bucket)
	return nil
}

func getFromS3(bucket, filename string, downloader s3manager.Downloader) ([]byte, error) {

	var buf []byte
	file := aws.NewWriteAtBuffer(buf)

	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	})
	if err != nil {
		log.Error().Err(err).Msgf("Ubable to download file %q, %v", filename, err)
		return nil, err
	}
	log.Info().Msgf("Downloaded", filename, numBytes, "bytes")
	return file.Bytes(), nil

}

func (sh *ShamirMgr) GetOneLocalState() *bkeygen.LocalPreParams {
	var preParams *bkeygen.LocalPreParams
	if len(sh.keys) > 0 {
		for _, value := range sh.keys {
			preParams = &value.LocalData.LocalPreParams
			break
		}
	} else {
		return nil
	}
	return preParams
}

func (sh *ShamirMgr) GenerateDataKey() (*kms.GenerateDataKeyOutput, error) {
	sess, err := NewSession(sh.shamirConfig.Kms.Region, sh.shamirConfig.Kms.Aksk.Id, sh.shamirConfig.Kms.Aksk.Secret)
	if err != nil {
		return nil, err
	}
	// Create KMS service client
	svc := kms.New(sess)
	result, err := svc.GenerateDataKey(&kms.GenerateDataKeyInput{
		KeyId:   aws.String(sh.shamirConfig.Kms.KeyId),
		KeySpec: aws.String("AES_256"),
	})
	if err != nil {
		log.Error().Err(err).Msg("fail to generate data key")
		return nil, err
	}
	return result, nil
}

func AesEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error().Err(err).Msg("fail to use aes to new cipher ")
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Error().Err(err).Msg("fail to use aes to new GCM ")
		return nil, err
	}
	nonce := make([]byte, 12)
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func AesDecrypt(chipertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Error().Err(err).Msg("fail to use aes to new cipher ")
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Error().Err(err).Msg("fail to use aes to new GCM ")
		return nil, err
	}
	nonce := make([]byte, 12)
	plaintext, err := aesgcm.Open(nil, nonce, chipertext, nil)
	if err != nil {
		log.Error().Err(err).Msg("fail to use aes decrypt data")
		return nil, err
	}
	return plaintext, nil
}

func NewSession(region, id, secret string) (*session.Session, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(id, secret, ""),
		},
	})

	if err != nil {
		log.Error().Err(err).Msg("failt to new aws session")
		return nil, err
	}
	return sess, nil
}

func createUUID(key string) (string, error) {
	uuidBytes := make([]byte, 16)
	keyBytes := []byte(key)
	copy(uuidBytes, keyBytes[len(keyBytes)-16:len(keyBytes)])
	uuid, err := uuid.FromBytes(uuidBytes)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return uuid.String(), nil

}
