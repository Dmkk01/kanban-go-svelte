declare namespace svelte.JSX {
  interface HTMLProps<T> {
    onclick_outside?: (e: CustomEvent<HTMLElement>) => void
  }
}
