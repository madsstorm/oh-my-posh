package segments

import (
	"path/filepath"
	"testing"

	"github.com/jandedobbeleer/oh-my-posh/src/runtime/mock"

	"github.com/stretchr/testify/assert"
)

func TestGetNodePackageVersion(t *testing.T) {
	cases := []struct {
		Case        string
		PackageJSON string
		Version     string
		ShouldFail  bool
		NoFiles     bool
	}{
		{Case: "14.1.5", Version: "14.1.5", PackageJSON: "{ \"name\": \"nx\",\"version\": \"14.1.5\"}"},
		{Case: "14.0.0", Version: "14.0.0", PackageJSON: "{ \"name\": \"nx\",\"version\": \"14.0.0\"}"},
		{Case: "no files", NoFiles: true, ShouldFail: true},
		{Case: "bad data", ShouldFail: true, PackageJSON: "bad data"},
	}

	for _, tc := range cases {
		var env = new(mock.Environment)
		env.On("Pwd").Return("posh")
		path := filepath.Join("posh", "node_modules", "nx")
		env.On("HasFilesInDir", path, "package.json").Return(!tc.NoFiles)
		env.On("FileContent", filepath.Join(path, "package.json")).Return(tc.PackageJSON)

		got, err := getNodePackageVersion(env, "nx")

		if tc.ShouldFail {
			assert.Error(t, err, tc.Case)
			return
		}

		assert.Nil(t, err, tc.Case)
		assert.Equal(t, tc.Version, got, tc.Case)
	}
}
