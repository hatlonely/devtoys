<script lang="ts">
	import { ConvertNumberBase } from '$lib/wailsjs/go/devtoys/App';
	import { RadioGroupItem, TextViewer, TextInputItem } from '$lib';
	import '@fontsource/roboto-mono';

	let text = '';
	let inBase = 10;
	let toBase = 16;
	let inBaseCharacters = '';
	let toBaseCharacters = '';
	let lowercase = false;
	let warning: any = '';
	let result = '';

	async function calculate() {
		if (inBase == 0 && inBaseCharacters == '') {
			return;
		}
		if (toBase == 0 && toBaseCharacters == '') {
			return;
		}

		warning = '';
		if (inBase != 0) {
			inBaseCharacters = '';
		}
		if (toBase != 0) {
			toBaseCharacters = '';
		}

		try {
			const res = await ConvertNumberBase({
				Number: text,
				LowerCase: lowercase,
				InBase: inBase,
				ToBase: toBase,
				InBaseCharacters: inBaseCharacters,
				ToBaseCharacters: toBaseCharacters
			});
			result = res.Number;
		} catch (err) {
			warning = err;
		}
	}

	$: text, calculate();
	$: lowercase, calculate();
	$: inBase, calculate();
	$: toBase, calculate();
	$: inBaseCharacters, calculate();
	$: toBaseCharacters, calculate();

	const labelValues = [
		{
			label: '二进制',
			value: 2
		},
		{
			label: '八进制',
			value: 8
		},
		{
			label: '十进制',
			value: 10
		},
		{
			label: '十六进制',
			value: 16
		},
		{
			label: '自定义',
			value: 0
		}
	];
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<RadioGroupItem
			bind:group={lowercase}
			name="mode"
			icon="sync_alt"
			title="大小写"
			description="是否使用小写字符编码（仅十六进制有效）"
			items={[
				{
					label: '大写',
					value: false
				},
				{
					label: '小写',
					value: true
				}
			]}
		/>

		<RadioGroupItem
			bind:group={inBase}
			name="mode"
			icon="sync_alt"
			title="输入进制类型"
			description="选择输入的进制类型"
			items={labelValues}
		/>

		{#if inBase === 0}
			<TextInputItem
				title="输入进制字符集"
				bind:value={inBaseCharacters}
				on:input={calculate}
				placeholder="输入进制字符集"
				code={true}
			/>
		{/if}

		<RadioGroupItem
			bind:group={toBase}
			name="mode"
			icon="sync_alt"
			title="输出进制类型"
			description="选择输出的进制类型"
			items={labelValues}
		/>

		{#if toBase === 0}
			<TextInputItem
				title="输出进制字符集"
				bind:value={toBaseCharacters}
				on:input={calculate}
				placeholder="输出进制字符集"
				code={true}
			/>
		{/if}
	</div>

	<div class="w-full text-token card p-4">
		<TextInputItem
			title="输入"
			bind:value={text}
			on:input={calculate}
			placeholder="输入要转换的数字"
			code={true}
			{warning}
		/>
	</div>

	{#if result}
		<div class="w-full text-token card p-4">
			<TextViewer title="输出" text={result} />
		</div>
	{/if}
</div>
