<script lang="ts">
	import '../app.postcss';
	import '@fortawesome/fontawesome-free/css/all.css';
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';

	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	import { Icon, ThemeSelector } from '$lib';

	const toolList = [
		{
			href: '/base64-text',
			icon: 'base64',
			title: 'Base64 文本编解码'
		},
		{
			href: '/unix-timestamp',
			icon: 'timestamp',
			title: 'Unix 时间戳'
		},
		{
			href: '/escape',
			icon: 'escape',
			title: '字符串转义'
		},
		{
			href: '/text-binary-conversion',
			icon: 'text',
			title: '文本二进制转换'
		},
		{
			href: '/number-base-conversion',
			icon: 'number',
			title: '数值进制转换'
		},
		{
			href: '/hash',
			icon: 'hash',
			title: '哈希散列'
		},
		{
			href: '/string-conversion',
			icon: 'string',
			title: '字符串转换'
		},
		{
			href: '/ssl-certificate',
			icon: 'https',
			title: 'https 证书'
		}
	];
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
				<hr />
				{#each toolList as tool}
					<li>
						<a href={tool.href}>
							<span><Icon name={tool.icon} class="icon" /></span>
							<span>{tool.title}</span>
						</a>
					</li>
				{/each}
			</ul>
		</nav>
	</svelte:fragment>
</AppShell>
