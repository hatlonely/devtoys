<script lang="ts">
	import '../app.postcss';
	import '@fortawesome/fontawesome-free/css/all.css';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	import { store, activeToolItemHref, title, setting } from '$lib/store/store';
	import { Icon, ThemeSelector } from '$lib';
	import {
		WindowFullscreen,
		WindowIsFullscreen,
		WindowUnfullscreen
	} from '$lib/wailsjs/runtime/runtime';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	function shortcut(e: KeyboardEvent) {
		switch (e.code) {
			// 按下 F11 全屏显示
			case 'F11':
				WindowIsFullscreen().then((isFullscreen) =>
					isFullscreen ? WindowUnfullscreen() : WindowFullscreen()
				);
				break;
		}
	}

	setting.subscribe((s) => {
		document.body.setAttribute('data-theme', s?.App?.Theme);
	});

	function setTitleAndActiveToolItem(t: string, href: string) {
		title.set(t);
		activeToolItemHref.set(href);
	}
</script>

<svelte:window on:keydown={shortcut} />

<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
	<svelte:fragment slot="header">
		<AppBar gridColumns="grid-cols-3" slotDefault="place-self-center" slotTrail="place-content-end">
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">DevToys</strong>
			</svelte:fragment>
			{$title}
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
						class:variant-filled-primary={$activeToolItemHref === '/'}
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
							class:variant-filled-primary={$activeToolItemHref === tool.href}
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
