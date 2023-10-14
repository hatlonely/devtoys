<script lang="ts">
	import '../app.postcss';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';

	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	import '@fortawesome/fontawesome-free/css/all.css';
</script>

<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
	<svelte:fragment slot="header">
		<AppBar>
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">DevToys</strong>
			</svelte:fragment>
			<svelte:fragment slot="trail">
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
						<span class="material-symbols-outlined"> home </span>
						<span>所有工具</span>
					</a>
				</li>
				<li><a href="/base64-text">Base64 文本</a></li>
			</ul>
		</nav>
	</svelte:fragment>
</AppShell>
