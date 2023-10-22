<script lang="ts">
	import { CodeEditor, CodeViewer, RadioGroupItem, NumberInputItem } from '$lib';
	import { format } from 'sql-formatter';

	let sql = '';
	let formattedSQL = '';
	let keywordCase: any = 'upper';
	let indentStyle: any = 'standard';
	let indent = 2;
	let warning: any = '';

	async function calculate() {
		if (sql === '') {
			formattedSQL = '';
			return;
		}

		formattedSQL = format(sql, {
			language: 'mysql',
			keywordCase: keywordCase,
			tabWidth: indent
		});
	}

	$: sql, calculate();
	$: indent, calculate();
	$: keywordCase, calculate();
	$: indentStyle, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<RadioGroupItem
			bind:group={keywordCase}
			name="keywordCase"
			title="关键字大小写"
			description="关键字大小写"
			icon="sync_alt"
			items={[
				{ value: 'upper', label: '大写' },
				{ value: 'lower', label: '小写' },
				{ value: 'preserve', label: '保留' }
			]}
		/>

		<!-- // 'standard' | 'tabularLeft' | 'tabularRight'; -->
		<RadioGroupItem
			bind:group={indentStyle}
			name="indentStyle"
			title="缩进样式"
			description="缩进样式"
			icon="sync_alt"
			items={[
				{ value: 'standard', label: '标准' },
				{ value: 'tabularLeft', label: '左对齐' },
				{ value: 'tabularRight', label: '右对齐' }
			]}
		/>

		<NumberInputItem
			bind:value={indent}
			title="缩进空格"
			description="缩进空格数量"
			icon="sync_alt"
		/>
	</div>

	<div class="w-full text-token card p-4">
		<CodeEditor title="输入" bind:value={sql} on:input={calculate} {warning} />
	</div>

	{#if formattedSQL}
		<div class="w-full text-token card p-4">
			<CodeViewer title="输出" value={formattedSQL} />
		</div>
	{/if}
</div>
