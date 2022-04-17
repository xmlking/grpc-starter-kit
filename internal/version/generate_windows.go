//go:build windows

package version

//go:generate cmd /c "git describe --tags --dirty --always > version.txt"
//go:generate cmd /c "git branch --show-current > branch.txt"
