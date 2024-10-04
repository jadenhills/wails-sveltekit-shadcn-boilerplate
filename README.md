# Wails SvelteKit shadcn Boilerplate

<p>
    This project uses 
    <a href="https://github.com/wailsapp/wails">Wails</a>,
    <a href="https://github.com/sveltejs/kit">SvelteKit</a>, and
    <a href="https://github.com/huntabyte/shadcn-svelte">shadcn-svelte</a>
</p>

## Quick Start

Clone the repo and install dependencies:

```bash
git clone --depth 1 --branch main https://github.com/jadenhills/wails-sveltekit-shadcn-boilerplate.git your-project
cd your-project
cd frontend
npm install
```

## Starting Development

Start the app in the `dev` environment:

```bash
wails dev
```

## Add shadcn Components

Ensure you are in the frontend directory

```bash
cd frontend
npx shadcn-svelte@latest add <component-name>
```

## Packaging for Production

To package apps for the local platform:

```bash
wails build
```
