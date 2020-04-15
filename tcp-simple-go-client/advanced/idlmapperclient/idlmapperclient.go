package idlmapperclient

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

type IDLMapperClient struct {
	host string
	port int
}

func GetIDLMapperClient(host string, port int) IDLMapperClient {
	return IDLMapperClient{
		host: host,
		port: port,
	}
}

// returns status code returned by /health endpoint
func (idlMapperClient IDLMapperClient) Health() (statusCode int, err error) {
	resp, err := idlMapperClient.httpGet("/health")
	if err != nil {
		return 0, err
	} else {
		return resp.StatusCode, nil
	}
}

// returns SeatIDs and associated IdentityLinks in map format
func (idlMapperClient IDLMapperClient) Map(envelope string) (idlMappings map[string]string, err error) {
	resp, err := idlMapperClient.httpGet("/map?env=" + envelope)
	if err != nil {
		return map[string]string{}, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return map[string]string{}, err
		} else {
			var idlMappings map[string]string
			if err := json.Unmarshal(body, &idlMappings); err != nil {
				return map[string]string{}, err
			}
			return idlMappings, nil
		}
	}
}

func (idlMapperClient IDLMapperClient) httpGet(endpoint string) (resp *http.Response, err error) {
	return http.Get(idlMapperClient.host + ":" + strconv.Itoa(idlMapperClient.port) + endpoint)
}
