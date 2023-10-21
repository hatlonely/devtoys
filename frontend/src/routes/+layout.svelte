<script lang="ts">
	import '../app.postcss';
	import '@fortawesome/fontawesome-free/css/all.css';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	import { store, activeToolItemHref } from '$lib/store/store';
	import { Icon, ThemeSelector } from '$lib';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	let title = '所有工具';

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
</script>

<svelte:window on:keydown|preventDefault={shortcut} />

<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
	<svelte:fragment slot="header">
		<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">DevToys</strong>
			</svelte:fragment>
			{title}
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
					<a href="/">
						<span><Icon name="home" class="icon" /></span>
						<span>所有工具</span>
					</a>
				</li>
				<hr />
				{#each store.toolItems as tool}
					<li>
						<a
							class:variant-filled-primary={activeToolItemHref_value == tool.href}
							href={tool.href}
							on:click={() => {
								title = tool.title;
								activeToolItemHref.set(tool.href);
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
