package action

import (
	"context"
	"encoding/json"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/volume"
	"github.com/richbai90/xfer-bin/pkg/dockerclient"
	"github.com/spf13/cobra"
)

type Inspect struct {
	CreatedAt  string
	Driver     string
	Labels     map[string]string
	Mountpoint string
	Name       string
	Options    map[string]string
	Scope      string
}

func Restore(src *string, dest *string) func(cmd *cobra.Command, args []string) {
	ctx := context.Background()
	return func(cmd *cobra.Command, args []string) {
		info, err := os.Stat(*src)
		if err != nil {
			return
		}
		if info.IsDir() {
			// TODO: copy directory
		} else {
			// Read in volume metadata
			dkr, err := dockerclient.Client()
			inspect, err := readSource(*src)
			if err != nil {
				return
			}
			// Create the volume
			v, err := dkr.VolumeCreate(ctx, volume.VolumeCreateBody{
				Driver:     inspect.Driver,
				DriverOpts: inspect.Options,
				Labels:     inspect.Labels,
				Name:       inspect.Name,
			})

			if err != nil { return }

			// Keep the container alive
			containerConfig := container.Config{
				Cmd: []string{
					"tail",
					"-f",
					"/dev/null",
				},
			}

			hostConfig := container.HostConfig{
				Mounts: []mount.Mount{
					{
						Type:   "bind",
						Source: inspect.Name,
						Target: "/restore",
					},
				},
				AutoRemove: true,
			}

			ctr, err := dkr.ContainerCreate(ctx, &containerConfig, &hostConfig, nil, nil, "restore")

			if err != nil {
				return
			}
			// Restore the backed up tar file to the volume
			oscmd := exec.Command("tar", "-zxvf", "/backup.tar.gz")
			oscmd.Dir = "/restore"
			oscmd.Run()
			// cleanup
			dkr.ContainerKill(ctx, ctr.ID, "SIGKILL")
			dkr.VolumeRemove(ctx, v.Name, true)
		}
	}
}

func readSource(src string) (Inspect, error) {
	inspectRaw := make([]byte, 1024)
	f, err := os.Open(src)
	if err != nil {
		return Inspect{}, err
	}
	_, err = f.Read(inspectRaw)
	if err != nil {
		return Inspect{}, err
	}
	inspect := Inspect{}
	err = json.Unmarshal(inspectRaw, &inspect)
	if err != nil {
		return Inspect{}, err
	}

	return inspect, nil
}
