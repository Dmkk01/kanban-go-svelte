import { API_URL, authHeader, authHeaderWithJSON, handleResponse } from '.'

const getColumn = async (column_id: number) => {
  return await fetch(`${API_URL}/column/${column_id}`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<ColumnBoard>(res))
}

const getTasks = async (column_id: number) => {
  return await fetch(`${API_URL}/column/${column_id}/task`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<Task[]>(res))
}

const ColumnAPI = {
  getColumn,
  getTasks,
}

export default ColumnAPI
