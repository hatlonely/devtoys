<script lang="ts">
	import { ConvertNumberBase } from '$lib/wailsjs/go/devtoys/App';
	import { Title, RadioGroup, Textarea, TextViewer, MultiSelector, SingleSelector } from '$lib';
	import { fade } from 'svelte/transition';
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

	async function onSelectInBase(e: any) {
		try {
			const res = await convertNumberBase(e.detail.oldValue, e.detail.newValue);
			console.log(res);
			text = res.Number;
		} catch (err) {
			warning = err;
		}
	}

	$: text, calculate();
	$: lowercase, calculate();

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
	<Title title="进制转换工具" />
	<div class="w-full text-token card p-4">
		<RadioGroup
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

		<SingleSelector
			bind:value={inBase}
			icon="sync_alt"
			title="输入进制类型"
			description="选择输入的二进制类型"
			{labelValues}
			on:select={onSelectInBase}
		/>

		<MultiSelector
			bind:values={toBases}
			icon="sync_alt"
			title="输出进制类型"
			description="选择输出的二进制类型"
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
		<Textarea
			bind:value={text}
			on:input={calculate}
			placeholder="输入要编码/解码的文本"
			code={true}
			{warning}
		/>
	</div>

	{#each labelValues as labelValue}
		{#if toBases[labelValue.value] && results[labelValue.value]}
			<div class="w-full text-token card p-4" in:fade out:fade>
				<TextViewer title={labelValue.label} text={results[labelValue.value]} />
			</div>
		{/if}
	{/each}
</div>
