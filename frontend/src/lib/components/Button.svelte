<script lang="ts">
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
	class="btn btn-sm variant-filled-primary mx-2"
>
	<span>{text}</span>
	{#if icon}
		{#if done}
			<span class="material-symbols-outlined"> done </span>
		{:else if iconSet == 'material-symbols'}
			<span class="material-symbols-outlined">{icon}</span>
		{/if}
	{/if}
</button>
