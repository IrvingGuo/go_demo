package entity

import (
	"fmt"

	"github.com/xuri/excelize/v2"
	"gorm.io/datatypes"
)

type Assignment struct {
	Base
	UserId         uint           `json:"userId" orm:"comment:user id"`
	ProgramId      uint           `json:"programId" orm:"comment:program id"`
	DeptId         uint           `json:"deptId" orm:"comment:department id"`
	Allocation     float32        `json:"allocation" orm:"comment:allocation" sql:"type:decimal(4,2);"`
	AllocationTime datatypes.Date `json:"allocationTime" orm:"comment:allocation time"`
	Approval       bool           `json:"approval" orm:"comment:approved or not"`
	Status         uint           `json:"status" orm:"comment:assignment status"`
}

type AssignmentWithCn struct {
	Assignment
	Cn string `json:"cn"`
}

func (a *Assignment) String() string {
	return fmt.Sprintf("[%+v, %+v, %+v, %+v, %+v, %+v, %+v, %+v]",
		a.Id, a.UserId, a.ProgramId, a.DeptId, a.Allocation, a.AllocationTime, a.Approval, a.Status)
}

func (assignment *Assignment) Save() error {
	return db.Save(assignment).Error
}

func (assignment *Assignment) Delete() error {
	return db.Delete(assignment).Error
}

func SaveInBatches(assignments []Assignment) error {
	return db.CreateInBatches(assignments, 100).Error
}

func UpdateAssignmentsInBatches(assignments []Assignment) (err error) {
	for _, assignment := range assignments {
		if assignment.Id <= 0 {
			return ErrUpdateId
		}
		if err = db.Save(assignment).Error; err != nil {
			return
		}
	}
	return
}

func UpdateAssignmentsStatus(statusPayloads []StatusPayload) (err error) {
	for _, statusPayload := range statusPayloads {
		db.Model(Assignment{}).Where("id = ?", statusPayload.Id).Update("status", statusPayload.Status)
	}
	return
}

func PrintAssignments(assignments []Assignment) {
	for _, assignment := range assignments {
		log.Info(assignment.String())
	}
}

func FindAllAssignments() (assignments []Assignment, err error) {
	err = db.Find(&assignments).Error
	return
}

func FindAssignmentById(id uint) (assignment Assignment, err error) {
	err = db.First(&assignment, "id = ?", id).Error
	return
}

func FindAssignmentsByProgramIds(programIds []uint) (assignments []Assignment, err error) {
	err = db.Find(&assignments, "program_id IN ?", programIds).Error
	return
}

func FindAssignmentsByUserIds(userIds []uint) (assignments []Assignment, err error) {
	err = db.Find(&assignments, "user_id IN ?", userIds).Error
	return
}

func FindAssignmentsByDeptIds(deptIds []uint) (assignments []Assignment, err error) {
	err = db.Find(&assignments, "dept_id IN ?", deptIds).Error
	return
}

func DeleteAssignmentsByIds(assignmentIds []uint) (err error) {
	err = db.Delete(Assignment{}, "id IN ?", assignmentIds).Error
	return
}

func DeleteAssignmentsByUserId(userId uint) (err error) {
	err = db.Delete(Assignment{}, "user_id = ?", userId).Error
	return
}

func DeleteAssignmentsByProgId(progId uint) (err error) {
	err = db.Delete(Assignment{}, "program_id = ?", progId).Error
	return
}

func FindAssignmentsWithCnByProgramIds(programIds []uint) (assignmentWithCn []AssignmentWithCn, err error) {
	err = db.
		Table("assignments").
		Select("assignments.id, assignments.user_id, assignments.program_id, assignments.allocation, assignments.allocation_time, assignments.approval, assignments.approval_time, assignments.status, users.cn").
		Where("assignments.program_id in ?", programIds).
		Joins("LEFT JOIN users ON assignments.user_id = users.id").
		Find(&assignmentWithCn).Error
	return
}

type MasterResPlanItem struct {
	Division     string
	Department   string
	Group        string
	Name         string
	Location     string
	ResourceType string
	Program      string
	ProgramType  string
	Year         int
	Month        int
	Allocation   float32
}

func (m MasterResPlanItem) ConvertToInterfaceArr() (arr []interface{}) {
	return []interface{}{
		m.Division,
		m.Department,
		m.Group,
		m.Name,
		m.Location,
		m.ResourceType,
		m.Program,
		m.ProgramType,
		m.Year,
		m.Month,
		m.Allocation,
	}
}

func FindMasterResPlan() (masterResPlanItems []MasterResPlanItem, err error) {
	err = db.Table("master_res_plan").Find(&masterResPlanItems).Error
	return
}

func GetMasterResPlanColumns(styleId int) []interface{} {
	return []interface{}{
		excelize.Cell{Value: "Division", StyleID: styleId},
		excelize.Cell{Value: "Department", StyleID: styleId},
		excelize.Cell{Value: "Group", StyleID: styleId},
		excelize.Cell{Value: "Name", StyleID: styleId},
		excelize.Cell{Value: "Location", StyleID: styleId},
		excelize.Cell{Value: "Resource Type", StyleID: styleId},
		excelize.Cell{Value: "Program", StyleID: styleId},
		excelize.Cell{Value: "Program Type", StyleID: styleId},
		excelize.Cell{Value: "Year", StyleID: styleId},
		excelize.Cell{Value: "Month", StyleID: styleId},
		excelize.Cell{Value: "Allocation", StyleID: styleId},
		excelize.Cell{Value: "Cost Center", StyleID: styleId}, // required by langya
		excelize.Cell{Value: "Cost", StyleID: styleId},        // required by langya
	}
}

func GetMasterResPlanColumnsWidth() []int {
	return []int{
		10, // Division
		30, // Department
		30, // Group
		15, // Name
		10, // Location
		15, // ResourceType
		25, // Program
		20, // ProgramType
		8,  // Year
		8,  // Month
		10, // Allocation
	}
}
