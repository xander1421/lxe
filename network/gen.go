package network // import "github.com/automaticserver/lxe/network"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o ./libcnifake/cni.go github.com/containernetworking/cni/libcni.CNI
