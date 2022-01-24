package version_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xmlking/grpc-starter-kit/internal/version"
)

func TestGetBuildInfo(t *testing.T) {
	assert := assert.New(t)

	expected := ""
	actual := ""

	t.Log(version.GetBuildInfo())
	t.Logf("build_info:%s", version.GetBuildInfo().PrettyString())
	t.Log(version.GetSoftwareBOM())

	assert.Equal(actual, expected)
}

func ExampleGetBuildInfo() {
	fmt.Println(version.GetBuildInfo())
	fmt.Println(version.GetSoftwareBOM())
	//fmt.Printf("build_info:%s", version.GetBuildInfo().PrettyString())

	// Output:
	//{"tag":"v0.1.0-74-g5f16c94-dirty","commit":"","branch":"develop","state":"dirty","build_time":"","go_version":"go1.17.6","compiler":"gc","platform":"darwin/amd64"}
	//not built in module mode
}
