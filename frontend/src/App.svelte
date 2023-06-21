<script lang="ts">
  import { QueryClient, QueryClientProvider } from '@sveltestack/svelte-query'
  import { Router, Link, Route } from 'svelte-routing'
  import Join from './pages/Join.svelte'
  import About from './pages/About.svelte'
  import GettingStarted from './pages/GettingStarted.svelte'
  import Home from './pages/Home.svelte'

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
    </Router>
  </main>
</QueryClientProvider>
