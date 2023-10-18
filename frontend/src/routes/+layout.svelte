<script lang="ts">
	import '../app.postcss';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';

	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	import '@fortawesome/fontawesome-free/css/all.css';
	import { Icon, ThemeSelector } from '$lib';

	document.body.setAttribute('data-theme', 'crimson');
</script>

<AppShell slotSidebarLeft="bg-surface-500/5 w-56 p-4">
	<svelte:fragment slot="header">
		<AppBar>
			<svelte:fragment slot="lead">
				<strong class="text-xl uppercase">DevToys</strong>
			</svelte:fragment>
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
				<li>
					<a href="/base64-text">
						<span><Icon name="delete" class="icon" /></span>
						<span>Base64 编解码</span>
					</a>
				</li>
				<li>
					<a href="/unix-timestamp">
						<span><Icon name="timestamp" class="icon" /></span>
						<span>Unix 时间戳</span>
					</a>
				</li>
				<li>
					<a href="/escape">
						<span><Icon name="timestamp" class="icon" /></span>
						<span>转义</span>
					</a>
				</li>
				<li>
					<a href="/text-binary-conversion">
						<span><Icon name="timestamp" class="icon" /></span>
						<span>文本二进制转换</span>
					</a>
				</li>
				<li>
					<a href="/number-base-conversion">
						<span><Icon name="timestamp" class="icon" /></span>
						<span>进制转换</span>
					</a>
				</li>
				<li>
					<a href="/hash">
						<span><Icon name="timestamp" class="icon" /></span>
						<span>哈希散列</span>
					</a>
				</li>
			</ul>
		</nav>
	</svelte:fragment>
</AppShell>
