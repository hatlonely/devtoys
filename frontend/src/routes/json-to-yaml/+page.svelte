<script lang="ts">
	import { CodeEditor, CodeViewer, RadioGroupItem, NumberInputItem } from '$lib';
	import YAML from 'yaml';

	let in_ = '';
	let to = '';
	let indent = 2;
	let warning: any = '';

	async function calculate() {
		if (in_ === '') {
			to = '';
			return;
		}

		if (in_.match(/^[{\[]/)) {
			try {
				to = YAML.stringify(JSON.parse(in_), {
					indent
				});
				warning = '';
			} catch (err) {
				warning = err;
			}
		} else {
			try {
				to = JSON.stringify(YAML.parse(in_), null, indent);
				warning = '';
			} catch (err) {
				warning = err;
			}
		}
	}

	$: in_, calculate();
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
		<CodeEditor title="输入" bind:value={in_} on:input={calculate} {warning} />
	</div>

	{#if to}
		<div class="w-full text-token card p-4">
			<CodeViewer title="输出" value={to} />
		</div>
	{/if}
</div>
