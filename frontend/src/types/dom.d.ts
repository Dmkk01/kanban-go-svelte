declare namespace svelte.JSX {
  interface HTMLProps<T> {
    onclick_outside?: (e: CustomEvent<HTMLElement>) => void
    onlong?: (e: CustomEvent<HTMLElement> & { detail: EventTarget & {category: string} }) => void
  }
}
