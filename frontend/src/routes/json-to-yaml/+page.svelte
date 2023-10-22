<script lang="ts">
	import { CodeEditor, CodeViewer, RadioGroupItem, NumberInputItem } from '$lib';
	import YAML from 'yaml';

	let json = '';
	let yaml = '';
	let indent = 2;
	let warning: any = '';

	async function calculate() {
		if (json === '') {
			yaml = '';
			return;
		}

		try {
			yaml = YAML.stringify(JSON.parse(json), {
				indent
			});
			warning = '';
		} catch (err) {
			warning = err;
		}
	}

	$: json, calculate();
	$: indent, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<NumberInputItem
			bind:value={indent}
			title="缩进空格"
			description="缩进空格数量"
			icon="sync_alt"
		/>
	</div>

	<div class="w-full text-token card p-4">
		<CodeEditor title="输入" bind:value={json} on:input={calculate} {warning} />
	</div>

	{#if yaml}
		<div class="w-full text-token card p-4">
			<CodeViewer title="输出" value={yaml} />
		</div>
	{/if}
</div>
