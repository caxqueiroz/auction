package auctionrunner_test

import (
	"time"

	"github.com/cloudfoundry-incubator/auction/auctiontypes"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	. "github.com/onsi/gomega"
)

func BuildLRPStartAuction(processGuid string, instanceGuid string, index int, stack string, memoryMB int, diskMB int) models.LRPStartAuction {
	return models.LRPStartAuction{
		DesiredLRP: models.DesiredLRP{
			ProcessGuid: processGuid,
			MemoryMB:    memoryMB,
			DiskMB:      diskMB,
			Stack:       stack,
		},
		InstanceGuid: instanceGuid,
		Index:        index,
	}
}

func BuildLRPStopAuction(processGuid string, index int) models.LRPStopAuction {
	return models.LRPStopAuction{
		ProcessGuid: processGuid,
		Index:       index,
	}
}

func BuildStartAuction(start models.LRPStartAuction, queueTime time.Time) auctiontypes.StartAuction {
	return auctiontypes.StartAuction{
		LRPStartAuction: start,
		QueueTime:       queueTime,
	}
}

func BuildStopAuction(stop models.LRPStopAuction, queueTime time.Time) auctiontypes.StopAuction {
	return auctiontypes.StopAuction{
		LRPStopAuction: stop,
		QueueTime:      queueTime,
	}
}

func BuildCellState(memoryMB int, diskMB int, containers int, lrps []auctiontypes.LRP) auctiontypes.CellState {
	totalResources := auctiontypes.Resources{
		MemoryMB:   memoryMB,
		DiskMB:     diskMB,
		Containers: containers,
	}

	availableResources := totalResources
	for _, lrp := range lrps {
		availableResources.MemoryMB -= lrp.MemoryMB
		availableResources.DiskMB -= lrp.DiskMB
		availableResources.Containers -= 1
	}

	Ω(availableResources.MemoryMB).Should(BeNumerically(">=", 0), "Check your math!")
	Ω(availableResources.DiskMB).Should(BeNumerically(">=", 0), "Check your math!")
	Ω(availableResources.Containers).Should(BeNumerically(">=", 0), "Check your math!")

	return auctiontypes.CellState{
		Stack:              "lucid64",
		AvailableResources: availableResources,
		TotalResources:     totalResources,
		LRPs:               lrps,
	}
}