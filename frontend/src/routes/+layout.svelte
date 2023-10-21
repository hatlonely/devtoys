<script lang="ts">
	import '../app.postcss';
	import '@fortawesome/fontawesome-free/css/all.css';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	import { store, activeToolItemHref, title } from '$lib/store/store';
	import { Icon, ThemeSelector } from '$lib';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	function shortcut(e: KeyboardEvent) {
		console.log(e.code);
		switch (e.code) {
			// 按下 F11 全屏显示
			case 'F11':
				document.fullscreenElement
					? document.exitFullscreen()
					: document.documentElement.requestFullscreen();
				break;
		}
	}

	let activeToolItemHref_value = '';
	activeToolItemHref.subscribe((href) => {
		activeToolItemHref_value = href;
	});

	let title_value = '';
	title.subscribe((t) => {
		title_value = t;
	});

	function setTitleAndActiveToolItem(t: string, href: string) {
		title.set(t);
		activeToolItemHref.set(href);
	}
</script>

<svelte:window on:keydown|preventDefault={shortcut} />

<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
	<svelte:fragment slot="header">
		<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">DevToys</strong>
			</svelte:fragment>
			{title_value}
			<svelte:fragment slot="trail">
				<ThemeSelector />

				<a href="https://github.com/hatlonely/devtoys" target="_blank" rel="noreferrer">
					<i class="fa-brands fa-github" />
				</a>
			</svelte:fragment>
		</AppBar>
	</svelte:fragment>
	<slot />

	<svelte:fragment slot="sidebarLeft">
		<nav class="list-nav">
			<ul>
				<li>
					<a
						href="/"
						class:variant-filled-primary={activeToolItemHref_value === '/'}
						on:click={() => {
							setTitleAndActiveToolItem('所有工具', '/');
						}}
					>
						<span><Icon name="home" class="icon" /></span>
						<span>所有工具</span>
					</a>
				</li>
				<hr />
				{#each store.toolItems as tool}
					<li>
						<a
							class:variant-filled-primary={activeToolItemHref_value === tool.href}
							href={tool.href}
							on:click={() => {
								setTitleAndActiveToolItem(tool.title, tool.href);
							}}
						>
							<span><Icon name={tool.icon} class="icon" /></span>
							<span>{tool.title}</span>
						</a>
					</li>
				{/each}
			</ul>
		</nav>
	</svelte:fragment>
</AppShell>
