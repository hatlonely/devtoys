<script lang="ts">
	import 'material-symbols';
	import { createEventDispatcher } from 'svelte';

	export let title: string;
	export let description: string;
	export let labelValues: any[];
	export let value = '';
	export let icon: string;
	export let iconSet = 'material-symbols';

	const dispatch = createEventDispatcher();

	function select(newValue: string): void {
		const oldValue = value;
		value = newValue;
		dispatch('select', {
			oldValue,
			newValue
		});
	}
</script>

<div class="w-full text-token p-2 flex flex-row">
	<span class="badge-icon p-4 variant-soft-secondary mt-2">
		{#if icon && iconSet == 'material-symbols'}
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
				class="btn btn-sm text-xs {value == labelValue.value
					? 'variant-filled-primary'
					: 'variant-soft'} hover:variant-soft-primary"
				on:click={() => {
					select(labelValue.value);
				}}
			>
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
