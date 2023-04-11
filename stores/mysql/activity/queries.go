package activitystore

const (
	QueryCreateActivity        = `INSERT INTO todoApp.activities (title, email) VALUES (?, ?);`
	QueryGetSingleActivityById = `SELECT * FROM todoApp.activities WHERE activity_id=?;`
)
