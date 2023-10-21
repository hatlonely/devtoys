<script lang="ts">
	import 'material-symbols';

	export let title: string = '';
	export let headers: string[] = [];
	export let names: string[] = [];
	export let rows: any[] = [];
	export let fixed: boolean = false;
</script>

<table class="table table-hover {fixed ? 'table-fixed' : 'table-auto'}">
	<thead>
		{#if title}
			<tr>
				<th colspan={names.length}>
					{title}
				</th>
			</tr>
		{/if}
		{#if headers}
			<tr>
				{#each headers as label}
					<th>{label}</th>
				{/each}
			</tr>
		{/if}
	</thead>
	<tbody>
		{#each rows as row, i}
			<tr>
				{#each names as name}
					<td>
						{#if row['_meta']?.[name]?.code}
							<pre><code>{row[name]}</code></pre>
						{:else}
							<span>{row[name]}</span>
						{/if}
					</td>
				{/each}
			</tr>
		{/each}
	</tbody>
</table>
