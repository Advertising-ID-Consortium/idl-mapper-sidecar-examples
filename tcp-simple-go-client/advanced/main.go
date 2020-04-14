package main

import (
	"io/ioutil"

	"github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tcp-simple-go-client/advanced/idlmapperclient"
	"github.com/sirupsen/logrus"
)

func main() {

	// get configured idl mapper client instance
	idlMapperClient := idlmapperclient.GetIDLMapperClient("http://localhost", 3000)

	// call health
	resp, err := idlMapperClient.Health()
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("IDL Mapper's '/health' endpoint responded with status: %s", string(resp.Status))
	}

	// call map
	envelope := "AjfowUURXDJnQmc_HNeuswelMv4ZHZQJFM8TpiUnYEyA81Vdgg"
	resp, err = idlMapperClient.Map(envelope)
	if err != nil {
		logrus.Error(err)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Error(err)
		} else {
			logrus.Infof("IDL Mapper's '/map' endpoint for envelope '%s' responded with mappings: %s", envelope, string(body))
		}
	}

}
