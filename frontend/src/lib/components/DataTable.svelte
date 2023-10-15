<script lang="ts">
	import 'material-symbols';
	import { clipboard } from '@skeletonlabs/skeleton';

	export let names: string[] = [];
	export let keys: string[] = [];
	export let rows: any[] = [];
	export let copyable: any = {};
	export let fixed: boolean = false;
	export let alignRight: any = {};

	let copied = {};
</script>

<table class="table table-hover {fixed ? 'table-fixed' : 'table-auto'}">
	<thead>
		<tr>
			{#each names as name}
				<th>{name}</th>
			{/each}
		</tr>
	</thead>
	<tbody>
		{#each rows as row, i}
			<tr>
				{#each keys as key}
					<td class={alignRight[key] ? 'text-right' : ''}>
						{#if copyable[key]}
							<button
								type="button"
								class="btn btn-sm variant-soft hover:variant-filled"
								on:click={() => {
									copied[key + i] = true;
									setTimeout(() => {
										copied[key + i] = false;
									}, 2000);
								}}
								use:clipboard={row[key]}
							>
								<span>{row[key]}</span>
								{#if copied[key + i]}
									<span class="material-symbols-outlined align-baseline"> done </span>
								{/if}
							</button>
						{:else}
							<span class="align-middle">{row[key]}</span>
						{/if}
					</td>
				{/each}
			</tr>
		{/each}
	</tbody>
</table>
