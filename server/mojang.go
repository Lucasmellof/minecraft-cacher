package server

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/lucasmellof/mojangcache/model"
	"io/ioutil"
	"net/http"
)

const API_URL = "https://api.mojang.com"
const SESSION_URL = "https://sessionserver.mojang.com"

func UsernameToUuid(username string) model.MojangResponse {
	URL := API_URL + "/users/profiles/minecraft/" + username
	return mojangGet(URL)
}

func UsernamesToUuids(jsonData []byte) model.MojangResponse {
	URL := API_URL + "/profiles/minecraft"
	return mojangPost(URL, jsonData)
}

func UuidToUsername(uuid string) model.MojangResponse {
	URL := API_URL + "/user/profile/" + uuid
	fmt.Println(URL)
	return mojangGet(URL)
}

func UuidToProfile(uuid string, unsigned string) model.MojangResponse {
	URL := SESSION_URL + "/session/minecraft/profile/" + uuid + "?unsigned=" + unsigned
	return mojangGet(URL)
}

func mojangGet(URL string) model.MojangResponse {
	cache := HasCache(URL)
	if cache.HasCache {
		return cache.Response
	}
	response, err := http.Get(URL)
	return SaveCache(URL, mojangResponse(response, err)).Response
}

func mojangPost(URL string, jsonData []byte) model.MojangResponse {
	hash := hex.EncodeToString(md5.New().Sum(jsonData))
	key := URL + hash

	cache := HasCache(key)
	if cache.HasCache {
		return cache.Response
	}

	req, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))
	return SaveCache(key, mojangResponse(req, err)).Response

}

func mojangResponse(resp *http.Response, err error) model.MojangResponse {
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	return model.MojangResponse{Code: resp.StatusCode, Json: string(b)}
}
