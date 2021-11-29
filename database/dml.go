package database

const InsertInfo = `INSERT INTO info (curr_to_curr, raw_display) VALUES (?, ?);`
const UpdateInfo = `UPDATE info SET raw_display = ? WHERE curr_to_curr = ?;`
const CheckHasInfo = `Select count(*) from info where curr_to_curr = ?;`
const GetFromDBInfo = `Select raw_display from info where curr_to_curr = ?`
