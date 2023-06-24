package gosl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkParseFileToStruct(b *testing.B) {
	type config struct {
		Host string `koanf:"host"`
		Port string `koanf:"port"`
	}

	_ = os.MkdirAll("./test", 0o755)

	_ = os.WriteFile("./test/file.json", []byte(`{
	"host": "https://my-server.com/api/v1",
	"port": "3000"
}`), 0o755)

	for i := 0; i < b.N; i++ {
		_, _ = ParseFileToStruct("./test/file.json", &config{})
	}

	_ = os.RemoveAll("./test")
}

func BenchmarkParseFileWithEnvToStruct(b *testing.B) {
	type config struct {
		URL      string `koanf:"url"`
		AuthType string `koanf:"auth_type"`
		Token    string `koanf:"token"`
	}

	_ = os.MkdirAll("./test", 0o755)

	_ = os.WriteFile("./test/file.json", []byte(`{
	"url": "https://my-server.com/api/v1",
	"auth_type": "Bearer",
	"token": "{{ MY_CONFIG_TOKEN }}"
}`), 0o755)

	for i := 0; i < b.N; i++ {
		_, _ = ParseFileWithEnvToStruct("./test/file.json", "MY_CONFIG", &config{})
	}

	_ = os.RemoveAll("./test")
}

func TestParseFileToStruct(t *testing.T) {
	type config struct {
		Host string `koanf:"host"`
		Port string `koanf:"port"`
	}

	_, err := ParseFileToStruct("", &config{})
	require.Error(t, err)

	_ = os.MkdirAll("./test", 0o755)

	_, err = ParseFileToStruct("./test/file.unknown", &config{})
	require.Error(t, err)

	_, err = ParseFileToStruct("https://example.com/file.json", &config{})
	require.Error(t, err)

	_, err = ParseFileToStruct("example.com/file.json", &config{})
	require.Error(t, err)

	_, err = ParseFileToStruct("https://github.com/koddr/gosl/blob/main/.github/dependabot.yml", &config{})
	require.Error(t, err)

	_, err = ParseFileToStruct("./test/not-found-file.json", &config{})
	require.Error(t, err)

	_, err = ParseFileToStruct("./test", &config{})
	require.Error(t, err)

	_ = os.WriteFile("./test/file.json", []byte(`{
	"host": "https://my-server.com/api/v1",
	"port": "3000"
}`), 0o755)

	cfgJson, err := ParseFileToStruct("./test/file.json", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgJson.Host, "https://my-server.com/api/v1")

	_ = os.WriteFile("./test/file.yml", []byte(`host: https://my-server.com/api/v1
port: '3000'`), 0o755)

	cfgYaml, err := ParseFileToStruct("./test/file.yml", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgYaml.Host, "https://my-server.com/api/v1")

	_ = os.WriteFile("./test/file.toml", []byte(`host = "https://my-server.com/api/v1"
port = "3000"`), 0o755)

	cfgToml, err := ParseFileToStruct("./test/file.toml", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgToml.Host, "https://my-server.com/api/v1")

	_ = os.WriteFile("./test/file.tf", []byte(`"host" = "https://my-server.com/api/v1"
"port" = "3000"`), 0o755)

	cfgHcl, err := ParseFileToStruct("./test/file.tf", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgHcl.Host, "https://my-server.com/api/v1")

	g := GenericUtility[config, any]{} // tests for method

	_, err = g.ParseFileToStruct("./test/file.unknown", &config{})
	require.Error(t, err)

	_, err = g.ParseFileToStruct("https://example.com/file.json", &config{})
	require.Error(t, err)

	_, err = g.ParseFileToStruct("example.com/file.json", &config{})
	require.Error(t, err)

	_, err = g.ParseFileToStruct("https://github.com/koddr/gosl/blob/main/.github/dependabot.yml", &config{})
	require.Error(t, err)

	_, err = g.ParseFileToStruct("./test/not-found-file.json", &config{})
	require.Error(t, err)

	_, err = g.ParseFileToStruct("./test", &config{})
	require.Error(t, err)

	cfgJson, err = g.ParseFileToStruct("./test/file.json", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgJson.Host, "https://my-server.com/api/v1")

	cfgYaml, err = g.ParseFileToStruct("./test/file.yml", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgYaml.Host, "https://my-server.com/api/v1")

	cfgToml, err = g.ParseFileToStruct("./test/file.toml", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgToml.Host, "https://my-server.com/api/v1")

	cfgHcl, err = g.ParseFileToStruct("./test/file.tf", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgHcl.Host, "https://my-server.com/api/v1")

	_ = os.RemoveAll("./test")
}

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

	_, err = ParseFileWithEnvToStruct("example.com/file.json", "", &config{})
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

	_, err = ParseFileWithEnvToStruct("./test/file.tf", "", &config{})
	require.Error(t, err)

	g := GenericUtility[config, any]{} // tests for method

	_, err = g.ParseFileWithEnvToStruct("./test/file.unknown", "", &config{})
	require.Error(t, err)

	_, err = g.ParseFileWithEnvToStruct("https://example.com/file.json", "", &config{})
	require.Error(t, err)

	_, err = g.ParseFileWithEnvToStruct("example.com/file.json", "", &config{})
	require.Error(t, err)

	_, err = g.ParseFileWithEnvToStruct("https://github.com/koddr/gosl/blob/main/.github/dependabot.yml", "", &config{})
	require.Error(t, err)

	_, err = g.ParseFileWithEnvToStruct("./test/not-found-file.json", "", &config{})
	require.Error(t, err)

	_, err = g.ParseFileWithEnvToStruct("./test", "", &config{})
	require.Error(t, err)

	cfgJson, err = g.ParseFileWithEnvToStruct("./test/file.json", "MY_CONFIG", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgJson.URL, "https://my-server.com/api/v1")

	cfgYaml, err = g.ParseFileWithEnvToStruct("./test/file.yml", "MY_CONFIG", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgYaml.URL, "https://my-server.com/api/v1")

	cfgToml, err = g.ParseFileWithEnvToStruct("./test/file.toml", "MY_CONFIG", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgToml.URL, "https://my-server.com/api/v1")

	cfgHcl, err = g.ParseFileWithEnvToStruct("./test/file.tf", "MY_CONFIG", &config{})
	assert.NoError(t, err)
	assert.EqualValues(t, cfgHcl.URL, "https://my-server.com/api/v1")

	_, err = g.ParseFileWithEnvToStruct("./test/file.tf", "", &config{})
	require.Error(t, err)

	_ = os.RemoveAll("./test")
}
