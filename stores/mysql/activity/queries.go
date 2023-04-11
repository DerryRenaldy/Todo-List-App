package activitystore

const (
	QueryCreateActivity        = `INSERT INTO todoApp.activities (title, email) VALUES (?, ?);`
	QueryGetSingleActivityById = `SELECT activity_id, title, email, created_at, updated_at FROM todoApp.activities WHERE activity_id=?;`
	QueryGetListActivity       = `SELECT activity_id, title, email, created_at, updated_at FROM todoApp.activities WHERE deleted_at IS NULL;`
	QueryUpdateDeletedTime     = `UPDATE todoApp.activities SET deleted_at=? WHERE activity_id=?;`
	QueryUpdateTitleById       = `UPDATE todoApp.activities SET title=? WHERE activity_id=?;`
)
