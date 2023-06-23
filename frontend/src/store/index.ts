import { writable } from 'svelte/store'

interface IStore {
  isSidebarOpen: boolean
  boardDrawer: {
    isOpen: boolean
    boardID: number | null
  }
  emojis: {
    isOpen: boolean
    activeKey: string | null
    keys: {
      key: string
      emoji: string
    }[]
  }
}

const initial: IStore = {
  isSidebarOpen: false,
  boardDrawer: {
    isOpen: true,
    boardID: null,
  },
  emojis: {
    isOpen: false,
    activeKey: null,
    keys: [],
  },
}

const store = writable<IStore>(initial)

export default store
