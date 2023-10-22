<script lang="ts">
	import { ConvertTextBinary } from '$lib/wailsjs/go/devtoys/App';
	import { RadioGroupItem, Textarea, TextViewer, MultiSelector } from '$lib';
	import '@fontsource/roboto-mono';

	let in_ = '';
	let to = '';
	let inType = 'raw';
	let toType = 'hex';
	let asByte = false;
	let lowercase = false;
	let warning: any = '';

	async function calculate() {
		warning = '';
		try {
			const res = await ConvertTextBinary({
				In: in_,
				InType: inType,
				ToType: toType,
				AsByte: asByte,
				LowerCase: lowercase
			});
			to = res.To;
		} catch (err) {
			warning = err;
		}
	}

	$: in_, calculate();
	$: inType, calculate();
	$: toType, calculate();
	$: asByte, calculate();
	$: lowercase, calculate();

	const labelValues = [
		{
			label: '原始',
			value: 'raw'
		},
		{
			label: '二进制',
			value: 'bin'
		},
		{
			label: '十进制',
			value: 'dec'
		},
		{
			label: '十六进制',
			value: 'hex'
		}
	];
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<RadioGroupItem
			bind:group={lowercase}
			name="lowercase"
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
			bind:group={asByte}
			name="asByte"
			icon="sync_alt"
			title="文本模式"
			description="二进制模式还是 Unicode 文本模式"
			items={[
				{
					label: '文本',
					value: false
				},
				{
					label: '二进制',
					value: true
				}
			]}
		/>

		<RadioGroupItem
			bind:group={inType}
			name="inType"
			icon="sync_alt"
			title="输入类型"
			description="输入的二进制类型"
			items={labelValues}
		/>

		<RadioGroupItem
			bind:group={toType}
			name="inType"
			icon="sync_alt"
			title="输出类型"
			description="输出的二进制类型"
			items={labelValues}
		/>
	</div>

	<div class="w-full text-token card p-4">
		<Textarea
			title="输入"
			bind:value={in_}
			on:input={calculate}
			placeholder="输入要转换的文本"
			code={true}
			{warning}
		/>
	</div>

	{#if to}
		<div class="w-full text-token card p-4">
			<TextViewer title="输出" text={to} />
		</div>
	{/if}
</div>
