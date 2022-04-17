//go:build !windows

package version

//go:generate sh -c "git describe --tags --dirty --always > version.txt"
//go:generate sh -c "git branch --show-current > branch.txt"
