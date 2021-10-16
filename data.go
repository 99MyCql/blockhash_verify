package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const (
	btcBlockApi = "https://chain.api.btc.com/v3/block/"
)

type blockData map[string]interface{}

func (b blockData) getBlockHeader() *blockHeader {
	return &blockHeader{
		version:       uint32toBytes4(uint32(b["version"].(float64))),
		prevBlockHash: str2Bytes32(b["prev_block_hash"].(string), 16),
		merkleRoot:    str2Bytes32(b["mrkl_root"].(string), 16),
		time:          uint32toBytes4(uint32(b["timestamp"].(float64))),
		bits:          uint32toBytes4(uint32(b["bits"].(float64))),
		nonce:         uint32toBytes4(uint32(b["nonce"].(float64))),
	}
}

func (b blockData) getBlockHash() [32]byte {
	return str2Bytes32(b["hash"].(string), 16)
}

// getBlock 请求相关API获取区块信息
func getBlock(height uint) (*blockData, error) {
	// 请求API，获取区块信息
	url := fmt.Sprintf(btcBlockApi+"%d", height)
	rsp, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrap(err, "请求btc.com接口出错")
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "读取响应数据出错")
	}
	// fmt.Println(string(body))

	// json反序列化
	var m map[string]interface{}
	if err := json.Unmarshal(body, &m); err != nil {
		return nil, errors.Wrap(err, "反序列化出错")
	}

	if m["err_no"].(float64) != 0 {
		fmt.Println(m)
		return nil, errors.New(m["message"].(string))
	}

	var data blockData = m["data"].(map[string]interface{})
	return &data, nil
}
