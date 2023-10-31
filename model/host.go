package model

import (
	pb "github.com/redamancy2319/pseudo-nezha-agent/proto"
)

type HostState struct {
	CPU            float64
	MemUsed        uint64
	SwapUsed       uint64
	DiskUsed       uint64
	NetInTransfer  uint64
	NetOutTransfer uint64
	NetInSpeed     uint64
	NetOutSpeed    uint64
	Uptime         uint64
	Load1          float64
	Load5          float64
	Load15         float64
	TcpConnCount   uint64
	UdpConnCount   uint64
	ProcessCount   uint64
}

func (s *HostState) PB() *pb.State {
	return &pb.State{
		Cpu:            s.CPU,
		MemUsed:        s.MemUsed,
		SwapUsed:       s.SwapUsed,
		DiskUsed:       s.DiskUsed,
		NetInTransfer:  s.NetInTransfer,
		NetOutTransfer: s.NetOutTransfer,
		NetInSpeed:     s.NetInSpeed,
		NetOutSpeed:    s.NetOutSpeed,
		Uptime:         s.Uptime,
		Load1:          s.Load1,
		Load5:          s.Load5,
		Load15:         s.Load15,
		TcpConnCount:   s.TcpConnCount,
		UdpConnCount:   s.UdpConnCount,
		ProcessCount:   s.ProcessCount,
	}
}

type Host struct {
	Platform        string
	PlatformVersion string
	CPU             []string
	MemTotal        uint64
	DiskTotal       uint64
	SwapTotal       uint64
	Arch            string
	Virtualization  string
	BootTime        uint64
	IP              string `json:"-"`
	CountryCode     string
	Version         string
}

func (h *Host) PB() *pb.Host {
	return &pb.Host{
		Platform:        h.Platform,
		PlatformVersion: h.PlatformVersion,
		Cpu:             h.CPU,
		MemTotal:        h.MemTotal,
		DiskTotal:       h.DiskTotal,
		SwapTotal:       h.SwapTotal,
		Arch:            h.Arch,
		Virtualization:  h.Virtualization,
		BootTime:        h.BootTime,
		Ip:              h.IP,
		CountryCode:     h.CountryCode,
		Version:         h.Version,
	}
}

type PseudoParam struct {
	PseudoPlatform        string // 系统名
	PseudoPlatformVersion string // 系统版本
	PseudoArch            string // 系统架构
	PseudoVirt            string // 系统虚拟化
	PseudoCPUModel        string // CPU型号
	PseudoCPUCoreCount    string // CPU核数
	PseudoMemTotal        uint64 // 内存大小
	PseudoDiskTotal       uint64 // 硬盘大小
	PseudoSwapTotal       uint64 // 交换大小
	PseudoIP              string // 客户端IP
	PseudoLoc             string // 客户端位置
	PseudoBootTime        uint64 // 系统启动时间
	PseudoVersion         string // Agent版本
}
