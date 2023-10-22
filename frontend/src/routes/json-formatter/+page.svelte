<script lang="ts">
	import { CodeEditor, CodeViewer, RadioGroupItem, NumberInputItem } from '$lib';

	let text = '';
	let json = '';
	let indent = 2;
	let warning: any = '';

	async function calculate() {
		if (text === '') {
			json = '';
			return;
		}

		warning = '';
		json = '';

		const lines = text.split(/\r?\n/);
		lines.forEach((line) => {
			if (line == '') {
				return;
			}
			try {
				const obj = JSON.parse(line);
				json += JSON.stringify(obj, null, indent) + '\n';
			} catch (err) {
				json += line + '\n';
				console.log(json);
				warning = err;
			}
		});
	}

	$: text, calculate();
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
		<CodeEditor title="输入" bind:value={text} on:input={calculate} {warning} />
	</div>

	{#if json}
		<div class="w-full text-token card p-4">
			<CodeViewer title="输出" value={json} />
		</div>
	{/if}
</div>
