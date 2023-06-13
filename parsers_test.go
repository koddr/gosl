package gosl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFileWithEnvToStruct(t *testing.T) {
	type config struct {
		URL      string `koanf:"url"`
		AuthType string `koanf:"auth_type"`
		Token    string `koanf:"token"`
	}

	_, err := ParseFileWithEnvToStruct("", "", &config{})
	require.Error(t, err)

	_ = os.MkdirAll("./test", 0o755)

	_, err = ParseFileWithEnvToStruct("./test/file.unknown", "", &config{})
	require.Error(t, err)

	_, err = ParseFileWithEnvToStruct("https://example.com/file.json", "", &config{})
	require.Error(t, err)

	_, err = ParseFileWithEnvToStruct("https://github.com/koddr/gosl/blob/main/.github/dependabot.yml", "", &config{})
	require.Error(t, err)

	_, err = ParseFileWithEnvToStruct("./test/not-found-file.json", "", &config{})
	require.Error(t, err)

	_, err = ParseFileWithEnvToStruct("./test", "", &config{})
	require.Error(t, err)

	_ = os.WriteFile("./test/file.json", []byte(`{
	"url": "https://my-server.com/api/v1",
	"auth_type": "Bearer",
	"token": "{{ MY_CONFIG_TOKEN }}"
}`), 0o755)

	cfgJson, err := ParseFileWithEnvToStruct("./test/file.json", "MY_CONFIG", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgJson.URL, "https://my-server.com/api/v1")

	_ = os.WriteFile("./test/file.yml", []byte(`url: https://my-server.com/api/v1
auth_type: Bearer
token: '{{ MY_CONFIG_TOKEN }}'`), 0o755)

	cfgYaml, err := ParseFileWithEnvToStruct("./test/file.yml", "MY_CONFIG", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgYaml.URL, "https://my-server.com/api/v1")

	_ = os.WriteFile("./test/file.toml", []byte(`url = "https://my-server.com/api/v1"
auth_type = "Bearer"
token = "{{ MY_CONFIG_TOKEN }}"`), 0o755)

	cfgToml, err := ParseFileWithEnvToStruct("./test/file.toml", "MY_CONFIG", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgToml.URL, "https://my-server.com/api/v1")

	_ = os.WriteFile("./test/file.tf", []byte(`"url" = "https://my-server.com/api/v1"
"auth_type" = "Bearer"
"token" = "{{ MY_CONFIG_TOKEN }}"`), 0o755)

	cfgHcl, err := ParseFileWithEnvToStruct("./test/file.tf", "MY_CONFIG", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgHcl.URL, "https://my-server.com/api/v1")

	_ = os.RemoveAll("./test")
}
