import {writable} from 'svelte/store';


interface IStore {
    settings: UserSettings | null
    boards: Board[]
    currentBoard: Board | null
}

const initial: IStore = {
    settings: null,
    boards: [],
    currentBoard: null
}

const store = writable<IStore>(initial);


export default store;
