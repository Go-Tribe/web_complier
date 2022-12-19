package docker

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
	c "web_complier/configs"
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
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "container",
		Filters: optionFilters,
		Limit:   1,
	})
	//fmt.Println("get container list:", len(containers), " cost:", time.Now().UnixMicro()-t0)
	/*t0 = time.Now().UnixMicro()*/
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
		bindsString := fmt.Sprintf("%s:%s", c.Config.StaticBasePath, dest)
		//fmt.Println("bindsString", bindsString)
		resp, err := cli.ContainerCreate(ctx, &container.Config{
			Image:        image,
			AttachStderr: true,
			AttachStdout: true,
			Tty:          true,
		}, &container.HostConfig{
			Binds: []string{bindsString},
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

	rand.Seed(time.Now().UnixMicro())
	filename := fmt.Sprintf("test_%d", rand.Uint32())
	fname := fmt.Sprintf("%s/%s.%s", c.Config.StaticBasePath, filename, ext)
	//fmt.Println(fname)

	err = os.WriteFile(fname, []byte(code), 0777)
	if err != nil {
		core.ZLogger.Sugar().Error("write file err:", err)
	}
	/*fmt.Println("write buffer cost:", time.Now().UnixMicro()-t0)
	t0 = time.Now().UnixMicro()*/

	// use &buf as argument for content in CopyToContainer
	//cli.CopyToContainer(ctx, containerID, ".", &buf, types.CopyToContainerOptions{})

	/*fmt.Println("CopyToContainer cost:", time.Now().UnixMicro()-t0)
	t0 = time.Now().UnixMicro()*/
	cmd = strings.ReplaceAll(cmd, "filename", dest+"/"+filename)
	cmd = fmt.Sprintf("timeout %d %s > %s/%s.log", langTimeout, cmd, dest, filename)
	//fmt.Println(cmd)
	res, err := cli.ContainerExecCreate(ctx, containerID, types.ExecConfig{
		Cmd: []string{"sh", "-c", cmd},
	})
	if err != nil {
		removeFile(filename, ext)
		core.ZLogger.Sugar().Error("docker exec create err:", err)
	}
	/*fmt.Println([]string{"sh", "-c", cmd})
	fmt.Println(res, err)*/

	chanC := make(chan int64)
	var resSting string
	go func() {
		t1 := time.Now().UnixMicro()
		if err := cli.ContainerExecStart(ctx, res.ID, types.ExecStartCheck{Detach: false, Tty: false}); err != nil {
			removeFile(filename, ext)
			core.ZLogger.Sugar().Errorf("ContainerExecStart %d err:%v:", res.ID, err)
		}

		logFile := fmt.Sprintf("%s/%s.log", c.Config.StaticBasePath, filename)
		for tryTimes := 200; tryTimes > 0; tryTimes-- {
			time.Sleep(time.Duration(20) * time.Millisecond)
			dir, err := os.Stat(logFile)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				core.ZLogger.Sugar().Error("open log file err:", err)
			}
			if dir.Size() > 0 {
				content, err := os.ReadFile(logFile)
				if err != nil {
					core.ZLogger.Sugar().Info("read log file err", err)
				}
				resSting = string(content)
				break
			}
		}
		timeCost := time.Now().UnixMicro() - t1
		chanC <- timeCost
	}()

	timeout := time.NewTimer(time.Duration(langTimeout) * time.Second)
	timeoutFlag := false
	select {
	case <-chanC:
		break
	case <-timeout.C:
		timeoutFlag = true
		fmt.Println("exec timeout")
		//core.ZLogger.Sugar().Error("execute timeout")
		return "execute timeout"
	}
	removeFile(filename, ext)
	if timeoutFlag {
		resSting = resSting + "\n execute timeout"
	}
	//fmt.Println("All action cost:", time.Now().UnixMicro()-t)
	//core.ZLogger.Sugar().Info("output:", string(resBytes))
	return resSting
}

func removeFile(filename string, ext string) {
	pattern := fmt.Sprintf("%s/%s*", c.Config.StaticBasePath, filename)
	files, err := filepath.Glob(pattern)
	if err != nil {
		core.ZLogger.Sugar().Info("err:", err)
	}
	for _, f := range files {
		fmt.Println(f)
		if err := os.Remove(f); err != nil {
			core.ZLogger.Sugar().Info("err:", err)
		}
	}
}
