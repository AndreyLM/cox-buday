package intermediate

import (
	"os/exec"

	customerror "github.com/andreylm/cox-buday.git/error_propagation/custom_error"
	"github.com/andreylm/cox-buday.git/error_propagation/lowlvl"
)

// Error - error
type Error struct {
	error
}

// RunJob - run job
func RunJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecutable, err := lowlvl.IsGloballyExec(jobBinPath)
	if err != nil {
		return Error{customerror.WrapError(err, "cannot run job %q: requisite binaries are not executable", id)}
	} else if isExecutable == false {
		return customerror.WrapError(nil, "job binary is not executable")
	}

	return exec.Command(jobBinPath, " --id="+id).Run()
}
