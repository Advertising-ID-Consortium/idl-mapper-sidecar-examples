package main

import (
	"github.com/Advertising-ID-Consortium/idl-mapper-sidecar-examples/tcp-simple-go-client/advanced/idlmapperclient"
	"github.com/sirupsen/logrus"
)

func main() {

	// get configured idl mapper client instance
	idlMapperClient := idlmapperclient.GetIDLMapperClient("http://localhost", 3000)

	// call health
	statusCode, err := idlMapperClient.Health()
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("IDL Mapper's '/health' endpoint responded with status code: %d", statusCode)
	}

	// call map
	envelope := "AjfowUURXDJnQmc_HNeuswelMv4ZHZQJFM8TpiUnYEyA81Vdgg"
	idlMappings, err := idlMapperClient.Map(envelope)
	if err != nil {
		logrus.Error(err)
	} else {
		logrus.Infof("IDL Mapper's '/map' endpoint for envelope '%s' responded with mappings: %s", envelope, idlMappings)
	}

}
