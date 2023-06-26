type Board = {
  id: number
  user_id: number
  name: string
  emoji: string
  created_at: string
  updated_at: string
}

type BoardFull = {
  board: Board
  columns: ColumnBoardFull[]
}

type ColumnBoardFull = {
  column: ColumnBoard
  tasks: Task[]
}

type ColumnBoard = {
  id: number
  board_id: number
  created_at: string
  updated_at: string
  emoji: string
  name: string
  position: number
}

type TaskCreate = {
  title: string
  description: string
  time_needed: number
  due_date: string
  position: number
  sub_tasks: SubTaskCreate[]
  links: LinkTaskCreate[]
}

type Task = {
  id: number
  column_id: number
  board_id: number
  title: string
  description: string
  time_needed: number
  due_date: string
  position: number
  created_at: string
  updated_at: string
  sub_tasks: SubTask[]
  links: LinkTask[]
  tags: Tag[]
}

type SubTaskCreate = {
  completed: boolean
  title: string
}

type SubTask = {
  id: number
  task_id: number
  title: string
  completed: boolean
  created_at: string
  updated_at: string
}

type LinkTaskCreate = {
  title: string
  url: string
  emoji: string
}

type LinkTask = {
  id: number
  task_id: number
  title: string
  url: string
  emoji: string
  created_at: string
  updated_at: string
}

type Tag = {
  board_id: number
  id: number
  title: string
  color: string
  created_at: string
  updated_at: string
}

type TagType = {
  id: number
  title: string
  color: string
  isNew?: boolean
}
