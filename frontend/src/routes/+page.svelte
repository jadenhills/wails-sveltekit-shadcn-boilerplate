<script lang="ts">
  import { Squared } from "$lib/wailsjs/go/main/App";
  import { Button } from "$lib/components/ui/button";
  import { Input } from "$lib/components/ui/input";
  import { Sun, Moon } from "lucide-svelte";
  import { toggleMode } from "mode-watcher";

  let number: number = 0;
  let result: number = 0;

  function squared() {
    console.log("Calculating square of", number);
    Squared(number)
      .then((r) => (result = r))
      .catch(() => (result = NaN));
  }
</script>

<main class="p-6">
  <div class="min-h-screen flex flex-col items-center justify-center gap-2">
    <h4 class="scroll-m-20 text-xl font-semibold tracking-tight">
      The square of {number} is <span class="animate-pulse">{result}</span>
    </h4>
    <div class="flex w-full max-w-sm flex-col gap-1.5">
      <Input type="number" bind:value={number} on:change={squared} />
      <p class="text-muted-foreground text-sm">Enter a number</p>
    </div>
    <div class="">
      <Button on:click={toggleMode} variant="outline" size="icon">
        <Sun
          class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
        />
        <Moon
          class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
        />
        <span class="sr-only">Toggle theme</span>
      </Button>
    </div>
  </div>
</main>
