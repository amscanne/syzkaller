// Copyright 2017 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

package sys

import (
	_ "github.com/google/syzkaller/sys/linux"
)

type Target struct {
	OS                 string
	Arch               string
	PtrSize            uint64
	CArch              []string
	CFlags             []string
	CrossCFlags        []string
	CCompilerPrefix    string
	KernelArch         string
	KernelHeaderArch   string
	KernelCrossCompile string
}

var Targets = map[string]map[string]*Target{
	"linux": map[string]*Target{
		"amd64": {
			PtrSize:          8,
			CArch:            []string{"__x86_64__"},
			CFlags:           []string{"-m64"},
			CCompilerPrefix:  "x86_64-linux-gnu-",
			KernelArch:       "x86_64",
			KernelHeaderArch: "x86",
		},
		"386": {
			PtrSize:          4,
			CArch:            []string{"__i386__"},
			CFlags:           []string{"-m32"},
			CCompilerPrefix:  "x86_64-linux-gnu-",
			KernelArch:       "i386",
			KernelHeaderArch: "x86",
		},
		"arm64": {
			PtrSize:          8,
			CArch:            []string{"__aarch64__"},
			CCompilerPrefix:  "aarch64-linux-gnu-",
			KernelArch:       "arm64",
			KernelHeaderArch: "arm64",
		},
		"arm": {
			PtrSize:          4,
			CArch:            []string{"__arm__"},
			CFlags:           []string{"-D__LINUX_ARM_ARCH__=6"},
			CrossCFlags:      []string{"-march=armv6t2"},
			CCompilerPrefix:  "arm-linux-gnueabihf-",
			KernelArch:       "arm",
			KernelHeaderArch: "arm",
		},
		"ppc64le": {
			PtrSize:          8,
			CArch:            []string{"__ppc64__", "__PPC64__", "__powerpc64__"},
			CFlags:           []string{"-D__powerpc64__"},
			CCompilerPrefix:  "powerpc64le-linux-gnu-",
			KernelArch:       "powerpc",
			KernelHeaderArch: "powerpc",
		},
	},
}

func init() {
	for OS, archs := range Targets {
		for arch, target := range archs {
			target.OS = OS
			target.Arch = arch
			target.CrossCFlags = append(target.CrossCFlags, target.CFlags...)
		}
	}
}