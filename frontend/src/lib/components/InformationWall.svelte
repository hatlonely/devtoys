<script lang="ts">
	import 'material-symbols';
	import { clipboard } from '@skeletonlabs/skeleton';
	export let labelValues: any[] = [];
	let copied = {};
</script>

<div class="flex flex-wrap space-y-1">
	{#each labelValues as labelValue, i}
		<div class="sm:w-full md:w-full lg:w-1/2 xl:w-1/3 2xl:w-1/4">
			<div class="flex flex-row space-x-3">
				<div class="align-sub">
					<span class="text-gray-500 align-sub">{labelValue.label}</span>
				</div>

				<button
					type="button"
					class="btn btn-sm variant-soft hover:variant-filled"
					on:click={() => {
						copied[i] = true;
						setTimeout(() => {
							copied[i] = false;
						}, 2000);
					}}
					use:clipboard={labelValue.value}
				>
					<span>{labelValue.value}</span>
					{#if copied[i]}
						<span class="material-symbols-outlined align-baseline"> done </span>
					{/if}
				</button>
			</div>
		</div>
	{/each}
</div>
