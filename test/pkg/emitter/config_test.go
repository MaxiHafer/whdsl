package emitter

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"

	"github.com/maxihafer/whdsl/pkg/emitter"
)

func TestEmitterTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

type ConfigTestSuite struct {
	suite.Suite
}

func (s *ConfigTestSuite) SetupSuite() {

}

func (s *ConfigTestSuite) TestUnmarshalJson() {
	bytes, err := os.ReadFile("../../testdata/emitter-cfg.yaml")
	s.Require().NoError(err)

	config := &emitter.Config{}
	err = yaml.Unmarshal(bytes, config)

	s.Require().NoError(err)
}
