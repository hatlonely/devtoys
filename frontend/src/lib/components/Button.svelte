<script lang="ts">
	import 'material-symbols';
	import { Icon } from '$lib';

	export let icon = '';
	export let iconSet = 'material-symbols';
	export let text: string;

	let done = false;
</script>

<button
	on:click
	on:click={() => {
		done = true;
		setTimeout(() => {
			done = false;
		}, 2000);
	}}
	type="button"
	class="btn btn-sm variant-soft hover:variant-filled"
>
	<span>{text}</span>
	{#if icon}
		{#if done}
			<span class="material-symbols-outlined"> done </span>
		{:else if iconSet == 'material-symbols'}
			<span class="material-symbols-outlined">{icon}</span>
		{:else if iconSet == 'custom'}
			<Icon name={icon} class="icon-md" />
		{/if}
	{/if}
</button>

<style>
	.material-symbols-outlined {
		font-variation-settings: 'FILL' 0, 'wght' 200, 'GRAD' 0, 'opsz' 24;
	}
</style>
