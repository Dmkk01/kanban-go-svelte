import { API_URL, authHeaderWithJSON, handleResponse } from '.'

const getTask = async (task_id: number) => {
  return await fetch(`${API_URL}/task/${task_id}`, {
    method: 'GET',
    headers: authHeaderWithJSON(),
  }).then(async (res) => handleResponse<Task>(res))
}

const taskUpdatePosition = async (data: { task_id: number; position: number; column_id: number }[]) => {
  return await fetch(`${API_URL}/task/position`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify(data),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const createTask = async (column_id: number, data: TaskCreate) => {
  return await fetch(`${API_URL}/column/${column_id}/task`, {
    method: 'POST',
    headers: authHeaderWithJSON(),
    body: JSON.stringify(data),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const updateSubtask = async (task_id: number, subtask_id: number, completed: boolean, title: string) => {
  return await fetch(`${API_URL}/task/${task_id}/subtask/${subtask_id}`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ completed, title, id: subtask_id }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const updateSubtaskCompleted = async (subtask_id: number, completed: boolean) => {
  return await fetch(`${API_URL}/subtask/${subtask_id}/complete`, {
    method: 'PUT',
    headers: authHeaderWithJSON(),
    body: JSON.stringify({ completed }),
  }).then(async (res) => handleResponse<StatusResponse>(res))
}

const TaskAPI = {
  getTask,
  taskUpdatePosition,
  createTask,
  updateSubtask,
  updateSubtaskCompleted
}

export default TaskAPI
