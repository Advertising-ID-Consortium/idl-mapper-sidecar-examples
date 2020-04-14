package idlmapperclient

import (
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

func (idlMapperClient IDLMapperClient) Health() (resp *http.Response, err error) {
	return idlMapperClient.httpGet("/health")
}

func (idlMapperClient IDLMapperClient) Map(envelope string) (resp *http.Response, err error) {
	return idlMapperClient.httpGet("/map?env=" + envelope)
}

func (idlMapperClient IDLMapperClient) httpGet(endpoint string) (resp *http.Response, err error) {
	return http.Get(idlMapperClient.host + ":" + strconv.Itoa(idlMapperClient.port) + endpoint)
}
