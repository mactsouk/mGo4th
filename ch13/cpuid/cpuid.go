package main

import (
	"fmt"
	"strings"

	. "github.com/klauspost/cpuid/v2"
)

func main() {
	// Print basic CPU information:
	fmt.Println("Name:", CPU.BrandName)
	fmt.Println("PhysicalCores:", CPU.PhysicalCores)
	fmt.Println("LogicalCores:", CPU.LogicalCores)
	fmt.Println("ThreadsPerCore:", CPU.ThreadsPerCore)

	fmt.Println("Family", CPU.Family, "Model:", CPU.Model, "Vendor ID:", CPU.VendorID)
	fmt.Println("Features:", strings.Join(CPU.FeatureSet(), ","))

	fmt.Println("Cacheline bytes:", CPU.CacheLine)
	fmt.Println("L1 Data Cache:", CPU.Cache.L1D, "bytes")
	fmt.Println("L1 Instruction Cache:", CPU.Cache.L1I, "bytes")
	fmt.Println("L2 Cache:", CPU.Cache.L2, "bytes")
	fmt.Println("L3 Cache:", CPU.Cache.L3, "bytes")
	fmt.Println("Frequency", CPU.Hz, "hz")

	// Test if we have these specific features:
	// if CPU.Supports(SSE, SSE2) {
	// 	fmt.Println("We have Streaming SIMD 2 Extensions")
	// }
}
