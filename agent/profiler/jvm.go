package profiler

import (
	"bytes"
	"fmt"
	"github.com/VerizonMedia/kubectl-flame/agent/details"
	"github.com/VerizonMedia/kubectl-flame/agent/utils"
	"os"
	"os/exec"
	"path"
	"strconv"
)

const (
	profilerDir = "/tmp/async-profiler"
	fileName    = "/tmp/flamegraph.html"
	profilerSh  = profilerDir + "/profiler.sh"
)

type JvmProfiler struct{}

func (j *JvmProfiler) SetUp(job *details.ProfilingJob) error {
	targetFs, err := utils.GetTargetFileSystemLocation(job.ContainerID)
	if err != nil {
		return err
	}

	err = os.RemoveAll("/tmp")
	if err != nil {
		return err
	}

	err = os.Symlink(path.Join(targetFs, "tmp"), "/tmp")
	if err != nil {
		return err
	}

	return j.copyProfilerToTempDir()
}

func (j *JvmProfiler) Invoke(job *details.ProfilingJob) error {
	fmt.Printf("JvmProfiler invoke with job: %+v\n", job)
	pid, err := utils.FindProcessId(job)
	if err != nil {
		return err
	}

	duration := strconv.Itoa(int(job.Duration.Seconds()))
	event := string(job.Event)
	fmt.Printf("trying exec.Command with fileName = %s, event = %s, pid = %s\n", fileName, event, pid)
	cmd := exec.Command(profilerSh, "-d", duration, "-f", fileName, "-e", event, pid) // todo: can consider to add more command here :) to provide support for k8s too; the error from this cmd when it run !!

	fmt.Printf("finish exec.Command with profilersh: %s\n", profilerSh)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	fmt.Printf("cmd.Run with err :%+v\n", err)

	if err != nil {
		fmt.Println("ignore error")
		return nil
		return err
	}

	return utils.PublishFlameGraph(fileName)
}

func (j *JvmProfiler) copyProfilerToTempDir() error {
	cmd := exec.Command("cp", "-r", "/app/async-profiler", "/tmp")
	return cmd.Run()
}
