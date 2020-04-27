package integration

import (
	"fmt"

	"github.com/docker/docker/client"
	"testing"
)

type testSuite interface {
	Init(*testing.T)
	Run(*testing.T)
}

func startContainer(cli *client.Client, imageName string) error {
	//ctx := context.Background()
	//
	//containerName := fmt.Sprintf("%s-%d", imageName, rand.Intn(1e6))
	//
	//resp, err := cli.ContainerCreate(ctx, &container.Config{
	//	Image: imageName,
	//	//Cmd:   []string{path.Join(dataPath, filename)},
	//	Tty: true,
	//}, &container.HostConfig{
	//	Mounts: []mount.Mount{
	//		//{
	//		//	Type:     mount.TypeBind,
	//		//	Source:   bindDir,
	//		//	Target:   dataPath,
	//		//	ReadOnly: true,
	//		//},
	//	},
	//}, nil, containerName)
	//if err != nil {
	//	return fmt.Errorf("error create a container, %w", err)
	//}
	//
	//if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
	//	return fmt.Errorf("error start a container, %w", err)
	//}
	//
	//_, err = cli.ContainerWait(ctx, resp.ID)
	//if err != nil {
	//	return fmt.Errorf("error wait a container, %w", err)
	//}
	//
	//err = cli.ContainerStop(ctx, resp.ID, nil)
	//if err != nil {
	//	return fmt.Errorf("error stop a container, %w", err)
	//}
	//
	//out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
	//	ShowStdout: true,
	//	ShowStderr: true,
	//})
	//if err != nil {
	//	return fmt.Errorf("error get container logs, %w", err)
	//}
	//
	//buf, err := ioutil.ReadAll(out)
	//if err != nil {
	//	return fmt.Errorf("error read container logs, %w", err)
	//}
	//
	//err = cli.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{
	//	RemoveVolumes: true,
	//	RemoveLinks:   false,
	//	Force:         true,
	//})
	//
	//if err != nil {
	//	return fmt.Errorf("error remove container, %w", err)
	//}
	//
	//_ = buf
	//return nil
}
