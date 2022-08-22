package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"web_complier/core"

	"io/ioutil"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func DockerRun(image string, code string, dest string, cmd string, langTimeout int64, memory int64) string {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		core.ZLogger.Sugar().Error("NewClientWithOpts:", err)
	}
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        image,
		Cmd:          []string{"sh", "-c", cmd},
		Tty:          false,
		AttachStderr: true,
		AttachStdout: true,
	}, &container.HostConfig{
		Resources: container.Resources{
			Memory: memory, // Minimum memory limit allowed is 6MB.
		},
	}, nil, nil, "")
	if err != nil {
		core.ZLogger.Sugar().Error("ContainerCreate:", err)
	}

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	err = tw.WriteHeader(&tar.Header{
		Name: dest,             // filename
		Mode: 0777,             // permissions
		Size: int64(len(code)), // filesize
	})
	if err != nil {
		core.ZLogger.Sugar().Error("docker copy err:", err)
	}
	tw.Write([]byte(code))
	tw.Close()

	// use &buf as argument for content in CopyToContainer
	cli.CopyToContainer(ctx, resp.ID, ".", &buf, types.CopyToContainerOptions{})

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		core.ZLogger.Sugar().Error("ContainerStart err:%v:", err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)

	timeout := time.NewTimer(time.Duration(langTimeout) * time.Second)
	select {
	case waitBody := <-statusCh:
		core.ZLogger.Sugar().Info("waitBody err:", waitBody.StatusCode)
		break
	case errC := <-errCh:
		core.ZLogger.Sugar().Errorf("ContainerWait statusCh err:%v:", errC)
	case <-timeout.C:
		core.ZLogger.Sugar().Error("execute timeout")
		cli.ContainerKill(ctx, resp.ID, "SIGKILL")
		core.ZLogger.Sugar().Error("ContainerKill")
		return "execute timeout"
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		core.ZLogger.Sugar().Errorf("ContainerLogs err:%v:", err)
	}

	defer out.Close()
	output, _ := ioutil.ReadAll(out)
	core.ZLogger.Sugar().Info("output:", string(output))
	return string(output)
}
