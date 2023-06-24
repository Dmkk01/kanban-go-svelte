type Board = {
  id: number
  user_id: number
  name: string
  emoji: string
  created_at: string
  updated_at: string
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
