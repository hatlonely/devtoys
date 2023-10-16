<script lang="ts">
	import 'material-symbols';
	import { createEventDispatcher } from 'svelte';

	export let title: string;
	export let description: string;
	export let labelValues: any[];
	export let values: Record<string, boolean> = {};
	export let icon: string;
	export let iconSet = 'material-symbols';

	const dispatch = createEventDispatcher();

	function filter(key: string): void {
		values[key] = !values[key];
		dispatch('select', {
			value: key,
			selected: values[key]
		});
	}
</script>

<div class="w-full text-token p-2 flex flex-row">
	<span class="badge-icon p-4 variant-soft-secondary mt-2">
		{#if iconSet == 'material-symbols'}
			<span class="material-symbols-outlined">{icon}</span>
		{/if}
	</span>
	<span class="flex-auto ml-5">
		<dt class="font-bold">{title}</dt>
		<dd class="text-sm opacity-50">{description}</dd>
	</span>
	<div class="flex items-center space-x-2">
		{#each labelValues as labelValue}
			<button
				class="chip {values[labelValue.value]
					? 'variant-filled'
					: 'variant-soft'} hover:variant-soft-secondary"
				on:click={() => {
					filter(labelValue.value);
				}}
			>
				{#if values[labelValue.value]}<span class="material-symbols-outlined">done</span>{/if}
				<span class="capitalize">{labelValue.label}</span>
			</button>
		{/each}
	</div>
</div>

<style>
	.material-symbols-outlined {
		font-variation-settings: 'FILL' 0, 'wght' 200, 'GRAD' 0, 'opsz' 12;
	}
</style>
