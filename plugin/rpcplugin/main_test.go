package rpcplugin

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/mattermost/platform/plugin/plugintest"
)

func TestMain(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	plugin := filepath.Join(dir, "plugin.exe")
	compileGo(t, `
		package main

		import (
			"github.com/mattermost/platform/plugin"
			"github.com/mattermost/platform/plugin/rpcplugin"
		)

		type MyPlugin struct {}

		func (p *MyPlugin) OnActivate(api plugin.API) error {
			return nil
		}

		func (p *MyPlugin) OnDeactivate() error {
			return nil
		}

		func main() {
			rpcplugin.Main(&MyPlugin{})
		}
	`, plugin)

	ctx, cancel := context.WithCancel(context.Background())
	p, ipc, err := NewProcess(ctx, plugin)
	require.NoError(t, err)
	defer p.Wait()

	muxer := NewMuxer(ipc, false)
	defer muxer.Close()

	defer cancel()

	var api plugintest.API

	hooks, err := ConnectMain(muxer)
	require.NoError(t, err)
	assert.NoError(t, hooks.OnActivate(&api))
	assert.NoError(t, hooks.OnDeactivate())
}
