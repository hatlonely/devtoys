<script lang="ts">
	export let onclick: any;
	export let icon = '';
	export let iconSet = 'material-symbols';

	let done = false;
</script>

<button
	on:click={onclick}
	on:click={() => {
		done = true;
		setTimeout(() => {
			done = false;
		}, 2000);
	}}
	type="button"
	class="btn btn-sm variant-filled-primary mx-2"
>
	<span>清空</span>
	{#if icon}
		{#if done}
			<span class="material-symbols-outlined"> done </span>
		{:else if iconSet == 'material-symbols'}
			<span class="material-symbols-outlined">{icon}</span>
		{/if}
	{/if}

	<span class="material-symbols-outlined">
		{done ? 'done' : 'delete'}
	</span>
</button>
