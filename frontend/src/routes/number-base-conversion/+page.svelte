<script lang="ts">
	import { ConvertNumberBase } from '$lib/wailsjs/go/devtoys/App';
	import { Title, RadioGroupItem, Textarea, TextViewer, MultiSelector, TextInputItem } from '$lib';
	import '@fontsource/roboto-mono';

	let text = '';
	let inBase = '0123456789';
	let toBases: Record<string, boolean> = {
		'0123456789ABCDEF': true
	};
	let lowercase = false;
	let warning: any = '';
	let results: any = {};

	async function convertNumberBase(inBase: string, toBase: string) {
		if (lowercase) {
			inBase = inBase.toLowerCase();
			toBase = toBase.toLowerCase();
		} else {
			inBase = inBase.toUpperCase();
			toBase = toBase.toUpperCase();
		}

		return await ConvertNumberBase({
			Number: text,
			LowerCase: lowercase,
			InBaseCharacters: inBase,
			ToBaseCharacters: toBase
		});
	}

	async function calculate() {
		warning = '';
		for (const toBase in toBases) {
			if (!toBases[toBase]) {
				continue;
			}
			try {
				const res = await convertNumberBase(inBase, toBase);
				results[toBase] = res.Number;
			} catch (err) {
				warning = err;
			}
		}
	}

	$: text, calculate();
	$: lowercase, calculate();
	$: inBase, calculate();

	const labelValues = [
		{
			label: '二进制',
			value: '01'
		},
		{
			label: '八进制',
			value: '01234567'
		},
		{
			label: '十进制',
			value: '0123456789'
		},
		{
			label: '十六进制',
			value: '0123456789ABCDEF'
		}
	];
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<!-- <Title title="进制转换工具" /> -->
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

		<MultiSelector
			bind:values={toBases}
			icon="sync_alt"
			title="输出进制类型"
			description="选择输出的进制类型"
			{labelValues}
			on:select={(e) => {
				if (e.detail.selected) {
					calculate();
				} else {
					results[e.detail.value] = '';
				}
			}}
		/>
	</div>

	<div class="w-full text-token card p-4">
		<TextInputItem
			bind:value={text}
			on:input={calculate}
			placeholder="输入要转换的数字"
			code={true}
			{warning}
		/>
	</div>

	{#each labelValues as labelValue}
		{#if toBases[labelValue.value] && results[labelValue.value]}
			<div class="w-full text-token card p-4">
				<TextViewer title={labelValue.label} text={results[labelValue.value]} />
			</div>
		{/if}
	{/each}
</div>
