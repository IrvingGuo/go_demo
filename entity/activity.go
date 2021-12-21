package entity

type Activity struct {
	Base
	Name   string `json:"name" orm:"comment:program name"`
	Status bool   `json:"status" orm:"comment:program status"`
}

func (activity *Activity) Save() error {
	return db.Save(activity).Error
}

func FindAllActivities() (activities []Activity, err error) {
	err = db.Find(&activities).Error
	return
}

func DeleteActivityById(activityId uint) error {
	return db.Delete(Activity{}, "id = ?", activityId).Error
}
