package main

import "fmt"

type Store interface {
	Open()
}

type StorageType int

const (
	StorageTypeDisk StorageType = 1 << iota
	StorageTypeTemp
	StorageTypeMemory
)

type DiskStorage struct{}

var _ Store = &DiskStorage{}

func (ms *DiskStorage) Open() {}

func newDiskStorage() Store {
	return &DiskStorage{}
}

type TempStorage struct{}

var _ Store = &TempStorage{}

func (ms *TempStorage) Open() {}

func newTempStorage() Store {
	return &TempStorage{}
}

type MemoryStorage struct{}

var _ Store = &MemoryStorage{}

func (ms *MemoryStorage) Open() {}

func newMemoryStorage() Store {
	return &MemoryStorage{}
}

func NewStore(t StorageType) Store {
	switch t {
	case StorageTypeDisk:
		return newDiskStorage()
	case StorageTypeMemory:
		return newMemoryStorage()
	case StorageTypeTemp:
		return newTempStorage()
	}
	return nil
}

type Interviewer interface {
	TakeInterview()
}

type Developer struct{}

var _ Interviewer = &Developer{}

func (i *Developer) TakeInterview() {
	fmt.Println("Asking about design patterns!")
}

type CommunityExecutive struct{}

var _ Interviewer = &CommunityExecutive{}

func (i *CommunityExecutive) TakeInterview() {
	fmt.Println("Asking about community building")
}

type HiringManager interface {
	MakeInterviewer() Interviewer
}

type DevelopmentManager struct{}

var _ HiringManager = &DevelopmentManager{}

func (dm *DevelopmentManager) MakeInterviewer() Interviewer {
	return &Developer{}
}

type MarketingManager struct{}

var _ HiringManager = &MarketingManager{}

func (mm *MarketingManager) MakeInterviewer() Interviewer {
	return &CommunityExecutive{}
}

type Title int

const (
	TitleDeveloper Title = iota
	TitleCommunityExecutive
)

func NewInterviewer(title Title) Interviewer {
	switch title {
	case TitleDeveloper:
		dm := DevelopmentManager{}
		return dm.MakeInterviewer()
	case TitleCommunityExecutive:
		mm := MarketingManager{}
		return mm.MakeInterviewer()
	}
	return nil
}

func main() {
	developer := NewInterviewer(TitleDeveloper)
	developer.TakeInterview() // 输出: Asking about design patterns

	communityExecutive := NewInterviewer(TitleCommunityExecutive)
	communityExecutive.TakeInterview() // 输出: Asking about community building
}
