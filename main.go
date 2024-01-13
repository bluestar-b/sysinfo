package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	for range time.Tick(800 * time.Millisecond) {
		cpuInfo, _ := cpu.Percent(0, true)
		ramInfo, _ := mem.VirtualMemory()

		overallCPU := average(cpuInfo)

		loadStr := formatLoad(cpuInfo)

		fmt.Printf("\033[1m\033[2K\rCPU: \033[33;1m%.2f%%\033[0m, \033[1mLoad:%s, \033[1mRAM: \033[33;1m%.2f\033[0mGB, \033[1mFree: \033[33;1m%.2f\033[0mGB",
			overallCPU, loadStr, ramInfo.UsedPercent, float64(ramInfo.Free)/1024/1024/1024)
	}
}

func average(nums []float64) float64 {
	total := 0.0
	for _, num := range nums {
		total += num
	}
	return total / float64(len(nums))
}

func formatLoad(cpuInfo []float64) string {
	loadStr := ""
	for _, load := range cpuInfo {
		loadStr += fmt.Sprintf(" \033[0m[\033[33;1m%.2f%%\033[0m]", load)
	}
	return loadStr
}
