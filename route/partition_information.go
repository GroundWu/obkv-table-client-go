package route

import (
	"strconv"

	"github.com/pkg/errors"

	"github.com/oceanbase/obkv-table-client-go/log"
	"github.com/oceanbase/obkv-table-client-go/table"
)

type obPartLevel int

const (
	PartLevelUnknown obPartLevel = -1
	PartLevelZero    obPartLevel = 0
	PartLevelOne     obPartLevel = 1
	PartLevelTwo     obPartLevel = 2
)

type obPartitionInfo struct {
	level           obPartLevel
	firstPartDesc   obPartDesc
	subPartDesc     obPartDesc
	partColumns     []*obColumn
	partTabletIdMap map[int64]int64
	partNameIdMap   map[string]int64
}

func (p *obPartitionInfo) SubPartDesc() obPartDesc {
	return p.subPartDesc
}

func (p *obPartitionInfo) FirstPartDesc() obPartDesc {
	return p.firstPartDesc
}

func (p *obPartitionInfo) GetTabletId(partId int64) (int64, error) {
	if p.partTabletIdMap == nil {
		log.Warn("partTabletIdMap is nil")
		return 0, errors.New("partTabletIdMap is nil")
	}
	return p.partTabletIdMap[partId], nil
}

func (p *obPartitionInfo) Level() obPartLevel {
	return p.level
}

func (p *obPartitionInfo) setRowKeyElement(rowKeyElement *table.ObRowKeyElement) {
	if p.firstPartDesc != nil {
		p.firstPartDesc.setRowKeyElement(rowKeyElement)
	}
	if p.subPartDesc != nil {
		p.subPartDesc.setRowKeyElement(rowKeyElement)
	}
}

func (p *obPartitionInfo) String() string {
	// partColumns to string
	var partColumnsStr string
	partColumnsStr = partColumnsStr + "["
	for i := 0; i < len(p.partColumns); i++ {
		if i > 0 {
			partColumnsStr += ", "
		}
		partColumnsStr += p.partColumns[i].String()
	}
	partColumnsStr += "]"

	// partTabletIdMap to string
	var partTabletIdMapStr string
	var i = 0
	partTabletIdMapStr = partTabletIdMapStr + "{"
	for k, v := range p.partTabletIdMap {
		if i > 0 {
			partTabletIdMapStr += ", "
		}
		i++
		partTabletIdMapStr += "m[" + strconv.Itoa(int(k)) + "]=" + strconv.Itoa(int(v))
	}
	partTabletIdMapStr += "}"

	// partNameIdMap to string
	var partNameIdMapStr string
	i = 0
	partNameIdMapStr = partNameIdMapStr + "{"
	for k, v := range p.partNameIdMap {
		if i > 0 {
			partNameIdMapStr += ", "
		}
		i++
		partNameIdMapStr += "m[" + k + "]=" + strconv.Itoa(int(v))
	}
	partNameIdMapStr += "}"

	// firstPartDesc to string
	var firstPartDescStr string
	if p.firstPartDesc != nil {
		firstPartDescStr = p.firstPartDesc.String()
	} else {
		firstPartDescStr = "nil"
	}

	// subPartDesc to string
	var subPartDescStr string
	if p.subPartDesc != nil {
		subPartDescStr = p.firstPartDesc.String()
	} else {
		subPartDescStr = "nil"
	}

	return "obPartitionInfo{" +
		"level:" + strconv.Itoa(int(p.level)) + ", " +
		"firstPartDesc:" + firstPartDescStr + ", " +
		"subPartDesc:" + subPartDescStr + ", " +
		"partColumns:" + partColumnsStr + ", " +
		"partTabletIdMap:" + partTabletIdMapStr + ", " +
		"partNameIdMap:" + partNameIdMapStr +
		"}"
}
