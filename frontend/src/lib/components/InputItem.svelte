<script lang="ts">
	import 'material-symbols';

	export let value: any;
	export let title: string;
	export let description: string;
	export let code = false;
	export let placeholder = '';
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
		<input
			bind:value
			on:input
			class="input {code ? 'devtoys-code' : ''}"
			spellcheck="false"
			{placeholder}
		/>
	</span>
</div>

<style>
	.material-symbols-outlined {
		font-variation-settings: 'FILL' 0, 'wght' 200, 'GRAD' 0, 'opsz' 24;
	}
</style>
