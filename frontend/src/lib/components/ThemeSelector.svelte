<script lang="ts">
	import Icon from '../icons/Icon.svelte';
	import { setting } from '$lib/store/store';
	import { SetSetting } from '$lib/wailsjs/go/devtoys/App';
	import { popup, LightSwitch } from '@skeletonlabs/skeleton';
	import type { PopupSettings } from '@skeletonlabs/skeleton';

	let currentTheme = document.body.getAttribute('data-theme');

	const popupFeatured: PopupSettings = {
		event: 'click',
		target: 'theme',
		placement: 'bottom'
	};

	const themes = [
		{ type: 'skeleton', name: 'Skeleton', icon: '💀' },
		{ type: 'wintry', name: 'Wintry', icon: '🌨️' },
		{ type: 'modern', name: 'Modern', icon: '🤖' },
		{ type: 'rocket', name: 'Rocket', icon: '🚀' },
		{ type: 'seafoam', name: 'Seafoam', icon: '🧜‍♀️' },
		{ type: 'vintage', name: 'Vintage', icon: '📺' },
		{ type: 'sahara', name: 'Sahara', icon: '🏜️' },
		{ type: 'hamlindigo', name: 'Hamlindigo', icon: '👔' },
		{ type: 'gold-nouveau', name: 'Gold Nouveau', icon: '💫' },
		{ type: 'crimson', name: 'Crimson', icon: '⭕' }
	];

	function setTheme(theme: string) {
		document.body.setAttribute('data-theme', theme);
		currentTheme = theme;
		setting.update((s) => {
			s.App.Theme = theme;
			SetSetting(s);
			return s;
		});
	}
</script>

<button use:popup={popupFeatured}>
	<Icon name="theme" />
</button>

<div class="card p-4 w-60 shadow-xl" data-popup="theme">
	<div class="space-y-4">
		<section class="flex justify-between items-center">
			<h6 class="h6">Mode</h6>
			<LightSwitch />
		</section>
		<hr />
		<nav class="list-nav p-4 -m-4 max-h-64 lg:max-h-[500px] overflow-y-auto">
			<ul>
				{#each themes as { icon, name, type }}
					<li>
						<button
							on:click={() => setTheme(type)}
							class="option w-full h-full"
							type="button"
							name="theme"
							class:bg-primary-active-token={currentTheme === type}
						>
							<span>{icon}</span>
							<span class="flex-auto text-left">{name}</span>
						</button>
					</li>
				{/each}
			</ul>
		</nav>
	</div>
</div>
