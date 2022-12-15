package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"io"
	"math/rand"
	"strings"
	"time"
	"web_complier/core"
)

func DockerRun(image string, code string, dest string, cmd string, langTimeout int64, memory int64, ext string) string {
	/*t0 := time.Now().UnixMicro()
	t := t0*/
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	/*fmt.Println("initializes a new client cost:", time.Now().UnixMicro()-t0)
	t0 = time.Now().UnixMicro()*/
	if err != nil {
		core.ZLogger.Sugar().Error("NewClientWithOpts:", err)
	}

	optionFilters := filters.NewArgs()
	optionFilters.Add("name", image)
	//optionFilters.Add("health", "starting")
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "container",
		Filters: optionFilters,
		Limit:   1,
	})
	//fmt.Println("get container list:", containers)
	/*fmt.Println("get container list:", len(containers), " cost:", time.Now().UnixMicro()-t0)
	t0 = time.Now().UnixMicro()*/
	if err != nil {
		core.ZLogger.Sugar().Error("docker list err:", err)
	}
	var containerID string
	if len(containers) > 0 {
		/*for _, t := range containers {
			fmt.Println(t.SizeRw)
		}*/
		filtersContainer := containers[0]
		//fmt.Println(filtersContainer)
		containerID = containers[0].ID
		if filtersContainer.State == "exited" {
			if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
				core.ZLogger.Sugar().Error("ContainerStart err:%v:", err)
			}
		}
		/*fmt.Println("range containers end", time.Now().UnixMicro()-t0)
		t0 = time.Now().UnixMicro()*/
	} else {
		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image:        image,
			AttachStderr: true,
			AttachStdout: true,
			Tty:          true,
		}, &container.HostConfig{
			Resources: container.Resources{
				Memory: memory, // Minimum memory limit allowed is 6MB.
			},
		}, nil, nil, fmt.Sprintf("%s", image)) //并发创建容器会报错，但可以保证只有一个容器

		if err != nil {
			core.ZLogger.Sugar().Error("ContainerCreate:", err)
		}
		//fmt.Println("create new container cost:", time.Now().UnixMicro()-t0)
		containerID = resp.ID
		if err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
			core.ZLogger.Sugar().Error("ContainerStart err:%v:", err)
		}
		//t0 = time.Now().UnixMicro()
	}

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	rand.Seed(time.Now().UnixMicro())
	filename := fmt.Sprintf("%stest%d", dest, rand.Uint32())
	fname := fmt.Sprintf("%s.%s", filename, ext)
	//fmt.Println(fname)
	err = tw.WriteHeader(&tar.Header{
		Name: fname,            // filename
		Mode: 0777,             // permissions
		Size: int64(len(code)), // filesize
	})
	if err != nil {
		core.ZLogger.Sugar().Error("docker copy err:", err)
	}
	tw.Write([]byte(code))
	tw.Close()

	/*fmt.Println("write buffer cost:", time.Now().UnixMicro()-t0)
	t0 = time.Now().UnixMicro()*/

	// use &buf as argument for content in CopyToContainer
	cli.CopyToContainer(ctx, containerID, ".", &buf, types.CopyToContainerOptions{})

	/*fmt.Println("CopyToContainer cost:", time.Now().UnixMicro()-t0)
	t0 = time.Now().UnixMicro()*/
	logFile := fmt.Sprintf("%s.log", filename)
	cmd = strings.ReplaceAll(cmd, "filename", filename)
	cmd = fmt.Sprintf("timeout %d %s > %s", langTimeout, cmd, logFile)
	fmt.Println(cmd)
	res, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd: []string{"sh", "-c", cmd},
	})
	if err != nil {
		removeFile([]string{fname}, cli, ctx, containerID)
		core.ZLogger.Sugar().Error("docker exec create err:", err)
	}
	/*fmt.Println([]string{"sh", "-c", cmd})
	fmt.Println(res, err)*/

	c := make(chan int64)
	var resSting string
	go func() {
		t1 := time.Now().UnixMicro()
		if err := cli.ContainerExecStart(ctx, res.ID, types.ExecStartCheck{Detach: false, Tty: false}); err != nil {
			removeFile([]string{fname, logFile}, cli, ctx, containerID)
			core.ZLogger.Sugar().Errorf("ContainerExecStart %d err:%v:", res.ID, err)
		}
		for tryTimes := 5; tryTimes > 0; tryTimes-- {
			fmt.Println("tryTimes:", tryTimes)
			readIO, stat, err := cli.CopyFromContainer(ctx, containerID, logFile)
			fmt.Println("read:", logFile, readIO, stat, err)
			if err != nil {
				core.ZLogger.Sugar().Info("CopyFromContainer err", err)
			}
			if stat.Size > 0 {
				var resBytes []byte
				tr := tar.NewReader(readIO)

				for {
					_, err := tr.Next()
					if err == io.EOF {
						break // End of archive
					}
					if err != nil {
						core.ZLogger.Sugar().Info("CopyFromContainer err", err)
					}
					//fmt.Printf("Contents of %s:\n", hdr.Name)
					//io.Copy(os.Stdout, tr)
					resBytes, err = io.ReadAll(tr)
					//fmt.Println(resBytes)
					if err != nil {
						core.ZLogger.Sugar().Info("CopyFromContainer err", err)
					}
				}
				resSting = string(resBytes)
				//fmt.Println(resSting)
				break
			}
			//time.Sleep(time.Duration(1) * time.Microsecond)
		}
		timeCost := time.Now().UnixMicro() - t1
		c <- timeCost
	}()

	timeout := time.NewTimer(time.Duration(langTimeout) * time.Second)
	timeoutFlag := false
	select {
	case timeCost := <-c:
		fmt.Println("ContainerExecStart cost:", timeCost)
		//core.ZLogger.Sugar().Info("ContainerExecStart %d cost:%v:", res.ID, timeCost)
		//time.Sleep(time.Duration(150) * time.Millisecond)
		break
	case <-timeout.C:
		timeoutFlag = true
		fmt.Println("exec timeout")
		//core.ZLogger.Sugar().Error("execute timeout")
		return "execute timeout"
	}

	removeFile([]string{fname, logFile}, cli, ctx, containerID)
	if timeoutFlag {
		resSting = resSting + "\n execute timeout"
	}
	//fmt.Println("All action cost:", time.Now().UnixMicro()-t)
	//core.ZLogger.Sugar().Info("output:", string(resBytes))
	return resSting
}

func removeFile(files []string, cli *client.Client, ctx context.Context, containerID string) {
	cmd := fmt.Sprintf("rm %s", strings.Join(files, " "))
	//fmt.Println(cmd)
	res, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd: []string{"sh", "-c", cmd},
	})
	if err != nil {
		fmt.Println(err)
		core.ZLogger.Sugar().Info("docker exec cmd ", cmd, " err:", err)
	}
	if err := cli.ContainerExecStart(ctx, res.ID, types.ExecStartCheck{Detach: true, Tty: false}); err != nil {
		core.ZLogger.Sugar().Info("ContainerExecStart %d err:%v:", res.ID, err)
	}
}
