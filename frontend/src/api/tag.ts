import { API_URL, authHeader, authHeaderWithJSON, handleResponse } from '.'

const createTag = async (board_id: number, color: string, title: string) => {
  return await fetch(`${API_URL}/board/${board_id}/tag`, {
    method: 'POST',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ color, title, board_id }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const updateTag = async (board_id: number, tag_id: number, color: string, title: string) => {
  return await fetch(`${API_URL}/tag/${tag_id}`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ color, title, board_id }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const deleteTag = async (tag_id: number) => {
  return await fetch(`${API_URL}/tag/${tag_id}`, {
    method: 'DELETE',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const TagAPI = {
  createTag,
  updateTag,
  deleteTag,
}

export default TagAPI
