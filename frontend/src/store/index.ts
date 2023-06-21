import { writable } from 'svelte/store'

interface IStore {
  isSidebarOpen: boolean
}

const initial: IStore = {
  isSidebarOpen: false,
}

const store = writable<IStore>(initial)

export default store
