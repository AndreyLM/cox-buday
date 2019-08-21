package lowlvl

import (
	"os"

	customerror "github.com/andreylm/cox-buday.git/error_propagation/custom_error"
)

// LowLevelErr - module specific error
type LowLevelErr struct {
	error
}

// IsGloballyExec - some func
func IsGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelErr{(customerror.WrapError(err, err.Error()))}
	}
	return info.Mode().Perm()&0100 == 0100, nil
}
