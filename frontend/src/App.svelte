<script lang="ts">
  import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query'
  import { Router, Link, Route } from 'svelte-routing'
  import Join from './pages/Join.svelte'
  import About from './pages/About.svelte'
  import GettingStarted from './pages/GettingStarted.svelte'
  import Home from './pages/home/Home.svelte'
  import Settings from './pages/home/Settings.svelte'
  import EmojiSelector from './components/common/EmojiSelector.svelte'

  import store from './store'

  const queryClient = new QueryClient()

  export let url = ''
</script>

<QueryClientProvider client={queryClient}>
  <main class={`font-mona ${import.meta.env.MODE === 'development' && 'debug-screens'}`}>
    <Router {url}>
      <Route path="/">
        <Join />
      </Route>
      <Route path="/about">
        <About />
      </Route>
      <Route path="/getting-started">
        <GettingStarted />
      </Route>
      <Route path="/home">
        <Home />
      </Route>
      <Route path="/home/settings">
        <Settings />
      </Route>
    </Router>
    {#if $store.emojis.isOpen}
      <EmojiSelector />
    {/if}
  </main>
</QueryClientProvider>
