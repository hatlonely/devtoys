<script lang="ts">
	import { ConvertTextBinary } from '$lib/wailsjs/go/devtoys/App';
	import { Title, RadioGroup, Textarea, TextViewer, MultiSelector } from '$lib';
	import { fade } from 'svelte/transition';
	import '@fontsource/roboto-mono';

	let text = '';
	let to: Record<string, boolean> = {
		hex: true
	};
	let lowercase = false;
	let withoutSpace = false;
	let withoutFillZero = false;
	let warning: any = '';
	let results: any = {};

	async function calculate() {
		warning = '';
		for (const key in to) {
			if (!to[key]) {
				continue;
			}
			try {
				const res = await ConvertTextBinary({
					Text: text,
					To: key,
					LowerCase: lowercase,
					WithoutSpace: withoutSpace,
					WithoutFillZero: withoutFillZero
				});
				results[key] = res.Text;
			} catch (err) {
				warning = err;
			}
		}
	}

	$: text, calculate();
	$: lowercase, calculate();
	$: withoutSpace, calculate();
	$: withoutFillZero, calculate();

	const labelValues = [
		{
			label: '二进制',
			value: 'bin'
		},
		{
			label: '十六进制',
			value: 'hex'
		},
		{
			label: '十进制',
			value: 'dec'
		},
		{
			label: 'Base64',
			value: 'base64'
		}
	];
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<Title title="文本二进制转换工具" />

	<div class="w-full text-token card p-4">
		<RadioGroup
			bind:group={withoutSpace}
			name="withoutSpace"
			icon="sync_alt"
			title="空格"
			description="是否在结果中使用空格分隔"
			items={[
				{
					label: '使用',
					value: false
				},
				{
					label: '不用',
					value: true
				}
			]}
		/>

		<RadioGroup
			bind:group={withoutFillZero}
			name="withoutFillZero"
			icon="sync_alt"
			title="填充 0"
			description="是否在结果中用 0 填充"
			items={[
				{
					label: '填充',
					value: false
				},
				{
					label: '不填',
					value: true
				}
			]}
		/>

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

		<MultiSelector
			bind:values={to}
			icon="sync_alt"
			title="二进制类型"
			description="选择要转换的二进制类型"
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
		{#if to[labelValue.value] && results[labelValue.value]}
			<div class="w-full text-token card p-4" in:fade out:fade>
				<TextViewer title={labelValue.label} text={results[labelValue.value]} />
			</div>
		{/if}
	{/each}
</div>
