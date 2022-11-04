package cache_file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/mantlenetworkio/mantle/subsidy/types"
)

type PayerStateFileWriter struct {
	homeDir       string
	cacheDir      string
	cacheFilename string
}

func NewPayerStateFileWriter(homeDir, cacheDir, cacheFilename string) *PayerStateFileWriter {
	cache := path.Join(homeDir, cacheDir)
	filePath := path.Join(cache, cacheFilename)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if _, err := os.Stat(cache); os.IsNotExist(err) {
			// Create the home folder
			if err = os.MkdirAll(cache, os.ModePerm); err != nil {
				panic(err)
			}
		}
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		file.Close()
	}
	return &PayerStateFileWriter{
		homeDir:       homeDir,
		cacheDir:      cacheDir,
		cacheFilename: cacheFilename,
	}
}

func (w *PayerStateFileWriter) Write(data *types.PayerState) error {
	cacheDataWriteBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	cacheDir := path.Join(w.homeDir, w.cacheDir)
	filename := path.Join(cacheDir, w.cacheFilename)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// And the home folder doesn't exist
		//if _, err := os.Stat(w.homeDir); os.IsNotExist(err) {
		//	// Create the home folder
		//	if err = os.MkdirAll(w.homeDir, os.ModePerm); err != nil {
		//		return err
		//	}
		//}
		// Create the home config folder
		if _, err := os.Stat(cacheDir); os.IsNotExist(err) {
			// Create the home folder
			if err = os.MkdirAll(cacheDir, os.ModePerm); err != nil {
				return err
			}
		}
		// Then create the file...
		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		if _, err = file.Write(cacheDataWriteBytes); err != nil {
			return err
		}

	} else {
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err = file.Write(cacheDataWriteBytes); err != nil {
			return err
		}
	}
	return nil
}

func (w *PayerStateFileWriter) LoadCache() *types.PayerState {
	// If the file exists, the initial height is the latest_height in the file
	filename := path.Join(w.homeDir, w.cacheDir, w.cacheFilename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("read cache file err:%+v ", err))
	}
	cacheData := types.PayerState{}
	if len(content) == 0 {
		return &cacheData
	}
	if err = json.Unmarshal(content, &cacheData); err != nil {
		panic(fmt.Sprintf("read cache file unmarshal err:%+v ", err))
	}
	return &cacheData
}
