package service

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	libnode "github.com/hashicorp/consul/integration/consul-container/libs/node"
	"github.com/hashicorp/consul/integration/consul-container/libs/utils"
)

// connectContainer
type connectContainer struct {
	ctx       context.Context
	container testcontainers.Container
	ip        string
	port      int
	req       testcontainers.ContainerRequest
}

func (g connectContainer) GetName() string {
	name, err := g.container.Name(g.ctx)
	if err != nil {
		return ""
	}
	return name
}

func (g connectContainer) GetAddr() (string, int) {
	return g.ip, g.port
}

func (g connectContainer) Start() error {
	if g.container == nil {
		return fmt.Errorf("container has not been initialized")
	}
	return g.container.Start(context.Background())
}

// Terminate attempts to terminate the container. On failure, an error will be
// returned and the reaper process (RYUK) will handle cleanup.
func (c connectContainer) Terminate() error {
	if c.container == nil {
		return nil
	}

	err := c.container.StopLogProducer()

	if err1 := c.container.Terminate(c.ctx); err == nil {
		err = err1
	}

	c.container = nil

	return err
}

func NewConnectService(ctx context.Context, name string, serviceName string, serviceBindPort int, node libnode.Node) (Service, error) {
	// TODO (dans): add the dc name into the cluster
	containerName := fmt.Sprintf("dc1-service-connect-%s", name)

	envoyVersion := getEnvoyVersion()
	buildargs := map[string]*string{
		"ENVOY_VERSION": utils.StringPointer(envoyVersion),
	}

	dockerfileCtx, err := getDevContainerDockerfile()
	if err != nil {
		return nil, err
	}
	dockerfileCtx.BuildArgs = buildargs

	nodeIP, _ := node.GetAddr()

	req := testcontainers.ContainerRequest{
		FromDockerfile: dockerfileCtx,
		WaitingFor:     wait.ForLog("").WithStartupTimeout(10 * time.Second),
		AutoRemove:     false,
		Name:           containerName,
		Cmd: []string{
			"consul", "connect", "envoy",
			"-sidecar-for", serviceName,
			"-service", name,
			"-admin-bind", "0.0.0.0:19000",
			"-grpc-addr", fmt.Sprintf("%s:8502", nodeIP),
			"-http-addr", fmt.Sprintf("%s:8500", nodeIP),
			"--",
			"--log-level", "debug"},
		Env: map[string]string{"CONSUL_HTTP_ADDR": fmt.Sprintf("%s:%d", nodeIP, 8500)},
		ExposedPorts: []string{
			fmt.Sprintf("%d/tcp", serviceBindPort), // Envoy Listener
			"19000/tcp",                            // Envoy Admin Port
		},
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}
	ip, err := container.ContainerIP(ctx)
	if err != nil {
		return nil, err
	}
	mappedPort, err := container.MappedPort(ctx, nat.Port(fmt.Sprintf("%d", serviceBindPort)))
	if err != nil {
		return nil, err
	}

	if err := container.StartLogProducer(ctx); err != nil {
		return nil, err
	}
	container.FollowOutput(&ServiceLogConsumer{
		Prefix: containerName,
	})

	// Register the termination function the node so the containers can stop together
	terminate := func() error {
		return container.Terminate(context.Background())
	}
	node.RegisterTermination(terminate)

	return &connectContainer{container: container, ip: ip, port: mappedPort.Int()}, nil
}
