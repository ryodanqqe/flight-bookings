package query

const UpdateUserQuery = `
	UPDATE users
	SET email = $2, password = #3, phone = $4 
	WHERE id = $1;
`
const DeleteUserQuery = `
	DELETE FROM users
	WHERE id = $1
`
