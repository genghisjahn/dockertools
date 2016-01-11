package docker

import "fmt"

type ContainerNotFoundError struct {
	s string
}

func (e *ContainerNotFoundError) Error() string {
	return e.s
}

func NewContainerNotFound(name string) *ContainerNotFoundError {
	return &ContainerNotFoundError{fmt.Sprintf("Container %s does not exist", name)}
}

type ContainerInfo struct {
	AppArmorProfile string        `json:"AppArmorProfile"`
	Args            []interface{} `json:"Args"`
	Config          struct {
		AttachStderr bool        `json:"AttachStderr"`
		AttachStdin  bool        `json:"AttachStdin"`
		AttachStdout bool        `json:"AttachStdout"`
		Cmd          []string    `json:"Cmd"`
		Domainname   string      `json:"Domainname"`
		Entrypoint   interface{} `json:"Entrypoint"`
		Env          []string    `json:"Env"`
		ExposedPorts struct {
			Eight000_tcp struct{} `json:"8000/tcp"`
		} `json:"ExposedPorts"`
		Hostname   string      `json:"Hostname"`
		Image      string      `json:"Image"`
		Labels     struct{}    `json:"Labels"`
		OnBuild    interface{} `json:"OnBuild"`
		OpenStdin  bool        `json:"OpenStdin"`
		StdinOnce  bool        `json:"StdinOnce"`
		StopSignal string      `json:"StopSignal"`
		Tty        bool        `json:"Tty"`
		User       string      `json:"User"`
		Volumes    interface{} `json:"Volumes"`
		WorkingDir string      `json:"WorkingDir"`
	} `json:"Config"`
	Created     string      `json:"Created"`
	Driver      string      `json:"Driver"`
	ExecDriver  string      `json:"ExecDriver"`
	ExecIDs     interface{} `json:"ExecIDs"`
	GraphDriver struct {
		Data interface{} `json:"Data"`
		Name string      `json:"Name"`
	} `json:"GraphDriver"`
	HostConfig struct {
		Binds           []string      `json:"Binds"`
		BlkioWeight     int           `json:"BlkioWeight"`
		CapAdd          interface{}   `json:"CapAdd"`
		CapDrop         interface{}   `json:"CapDrop"`
		CgroupParent    string        `json:"CgroupParent"`
		ConsoleSize     []int         `json:"ConsoleSize"`
		ContainerIDFile string        `json:"ContainerIDFile"`
		CPUPeriod       int           `json:"CpuPeriod"`
		CPUQuota        int           `json:"CpuQuota"`
		CPUShares       int           `json:"CpuShares"`
		CpusetCpus      string        `json:"CpusetCpus"`
		CpusetMems      string        `json:"CpusetMems"`
		Devices         []interface{} `json:"Devices"`
		DNS             []interface{} `json:"Dns"`
		DNSOptions      []interface{} `json:"DnsOptions"`
		DNSSearch       []interface{} `json:"DnsSearch"`
		ExtraHosts      interface{}   `json:"ExtraHosts"`
		GroupAdd        interface{}   `json:"GroupAdd"`
		IpcMode         string        `json:"IpcMode"`
		KernelMemory    int           `json:"KernelMemory"`
		Links           []string      `json:"Links"`
		LogConfig       struct {
			Config struct{} `json:"Config"`
			Type   string   `json:"Type"`
		} `json:"LogConfig"`
		LxcConf           []interface{} `json:"LxcConf"`
		Memory            int           `json:"Memory"`
		MemoryReservation int           `json:"MemoryReservation"`
		MemorySwap        int           `json:"MemorySwap"`
		MemorySwappiness  int           `json:"MemorySwappiness"`
		NetworkMode       string        `json:"NetworkMode"`
		OomKillDisable    bool          `json:"OomKillDisable"`
		PidMode           string        `json:"PidMode"`
		PortBindings      struct {
			Eight000_tcp []struct {
				HostIP   string `json:"HostIp"`
				HostPort string `json:"HostPort"`
			} `json:"8000/tcp"`
		} `json:"PortBindings"`
		Privileged      bool `json:"Privileged"`
		PublishAllPorts bool `json:"PublishAllPorts"`
		ReadonlyRootfs  bool `json:"ReadonlyRootfs"`
		RestartPolicy   struct {
			MaximumRetryCount int    `json:"MaximumRetryCount"`
			Name              string `json:"Name"`
		} `json:"RestartPolicy"`
		SecurityOpt  interface{} `json:"SecurityOpt"`
		UTSMode      string      `json:"UTSMode"`
		Ulimits      interface{} `json:"Ulimits"`
		VolumeDriver string      `json:"VolumeDriver"`
		VolumesFrom  interface{} `json:"VolumesFrom"`
	} `json:"HostConfig"`
	HostnamePath string `json:"HostnamePath"`
	HostsPath    string `json:"HostsPath"`
	ID           string `json:"Id"`
	Image        string `json:"Image"`
	LogPath      string `json:"LogPath"`
	MountLabel   string `json:"MountLabel"`
	Mounts       []struct {
		Destination string `json:"Destination"`
		Mode        string `json:"Mode"`
		RW          bool   `json:"RW"`
		Source      string `json:"Source"`
	} `json:"Mounts"`
	Name            string `json:"Name"`
	NetworkSettings struct {
		Bridge                 string `json:"Bridge"`
		EndpointID             string `json:"EndpointID"`
		Gateway                string `json:"Gateway"`
		GlobalIPv6Address      string `json:"GlobalIPv6Address"`
		GlobalIPv6PrefixLen    int    `json:"GlobalIPv6PrefixLen"`
		HairpinMode            bool   `json:"HairpinMode"`
		IPAddress              string `json:"IPAddress"`
		IPPrefixLen            int    `json:"IPPrefixLen"`
		IPv6Gateway            string `json:"IPv6Gateway"`
		LinkLocalIPv6Address   string `json:"LinkLocalIPv6Address"`
		LinkLocalIPv6PrefixLen int    `json:"LinkLocalIPv6PrefixLen"`
		MacAddress             string `json:"MacAddress"`
		Networks               struct {
			Bridge struct {
				EndpointID          string `json:"EndpointID"`
				Gateway             string `json:"Gateway"`
				GlobalIPv6Address   string `json:"GlobalIPv6Address"`
				GlobalIPv6PrefixLen int    `json:"GlobalIPv6PrefixLen"`
				IPAddress           string `json:"IPAddress"`
				IPPrefixLen         int    `json:"IPPrefixLen"`
				IPv6Gateway         string `json:"IPv6Gateway"`
				MacAddress          string `json:"MacAddress"`
			} `json:"bridge"`
		} `json:"Networks"`
		Ports struct {
			Eight000_tcp []struct {
				HostIP   string `json:"HostIp"`
				HostPort string `json:"HostPort"`
			} `json:"8000/tcp"`
		} `json:"Ports"`
		SandboxID              string      `json:"SandboxID"`
		SandboxKey             string      `json:"SandboxKey"`
		SecondaryIPAddresses   interface{} `json:"SecondaryIPAddresses"`
		SecondaryIPv6Addresses interface{} `json:"SecondaryIPv6Addresses"`
	} `json:"NetworkSettings"`
	Path           string `json:"Path"`
	ProcessLabel   string `json:"ProcessLabel"`
	ResolvConfPath string `json:"ResolvConfPath"`
	RestartCount   int    `json:"RestartCount"`
	State          struct {
		Dead       bool   `json:"Dead"`
		Error      string `json:"Error"`
		ExitCode   int    `json:"ExitCode"`
		FinishedAt string `json:"FinishedAt"`
		OOMKilled  bool   `json:"OOMKilled"`
		Paused     bool   `json:"Paused"`
		Pid        int    `json:"Pid"`
		Restarting bool   `json:"Restarting"`
		Running    bool   `json:"Running"`
		StartedAt  string `json:"StartedAt"`
		Status     string `json:"Status"`
	} `json:"State"`
}
