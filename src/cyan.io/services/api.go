package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"cyan.io/utils"
)

func SignedPostRequest(localVarPostBody interface{}) {
	bodyBuf := &bytes.Buffer{}
	json.NewEncoder(bodyBuf).Encode(localVarPostBody)

	log.Printf("Body Order is %+s", bodyBuf.String())

	SignedGetRequest([]byte(bodyBuf.String()))
}

func SignedGetRequest(message []byte) {
	//reqParams := "timestamp=" + strconv.FormatInt(timeNow, 10)
	// reqParams := "recvWindow=" + strconv.Itoa(recvWindow) + "&timestamp=" + strconv.FormatInt(timeNow, 10)
	//reqParams := "timestamp=" + strconv.FormatInt(timeNow, 10) + "&recvWindow=" + strconv.Itoa(recvWindow)
	//log.Println("Request params: " + reqParams)

	cfg.AddDefaultHeader("Authorization", utils.Authorization)

	// message := []byte(reqParams)
	hash := hmac.New(sha256.New, utils.Secret)
	hash.Write(message)
	cfg.AddDefaultHeader("Signature", hex.EncodeToString(hash.Sum(nil)))
	log.Printf("F payload: %s", string(message))
	log.Println("Signature is: " + hex.EncodeToString(hash.Sum(nil)))
	// cossBalance := GetAccountBalances(ctx, apiClient, int32(recvWindow), timeNow)
	// log.Println(cossBalance)
}

func PrintHTTPResponse(httpResponse *http.Response, err error) {
	log.Printf("HTTP resquest failed with errors %v\n", err)
	log.Printf("HTTP Response: %v\n", httpResponse)
	log.Printf("HTTP Response: %d\n", httpResponse.StatusCode)
}

func ComposeHeaderForGetRequest(timeNow int64) {
	reqParams := "timestamp=" + strconv.FormatInt(timeNow, 10)
	SignedGetRequest([]byte(reqParams))
}
