<script lang="ts">
  import { Router, Link, Route } from 'svelte-routing'
  import Join from './pages/Join.svelte'

  export let url = '/'
</script>

<main class={`font-mona ${import.meta.env.MODE === 'development' && 'debug-screens'}`}>
  <Router {url}>
    <Route path="/"><Join /></Route>
  </Router>
</main>
