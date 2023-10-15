<script lang="ts">
	import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';

	export let group: any;
	export let name: string;
	export let title: string;
	export let description: string;
	export let items: any;
	export let icon: string;
	export let iconSet = 'material-symbols';
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
	<span>
		<RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
			{#each items as item}
				<RadioItem bind:group {name} value={item.value}>{item.label}</RadioItem>
			{/each}
		</RadioGroup>
	</span>
</div>
