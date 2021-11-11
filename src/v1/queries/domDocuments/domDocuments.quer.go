package domDocuments

const attr = `id, name, type, creation_date, update_date, creation_user, update_user`
const table = `store.dom_documents`
const globalSelect = `SELECT ` + attr + ` FROM ` + table

const GET_ALL_DOCUMENTS = globalSelect
