<script lang="ts">
	import { ConvertString } from '$lib/wailsjs/go/devtoys/App';
	import { Title, Textarea, TextViewer, SingleSelector } from '$lib';
	import '@fontsource/roboto-mono';

	let text = '';
	let type_ = 'camel';
	let warning: any = '';
	let result = '';

	async function calculate() {
		warning = '';
		try {
			let res = await ConvertString({
				Text: text,
				Type: type_
			});
			result = res.Text;
		} catch (err) {
			warning = err;
		}
	}

	$: text, calculate();

	const labelValues = [
		{ label: '单词首字母大写', value: 'title' },
		{ label: '句子首字母大写', value: 'sentence' },
		{ label: '小写', value: 'lower' },
		{ label: '大写', value: 'upper' },
		{ label: '驼峰', value: 'camel' },
		{ label: '帕斯卡', value: 'pascal' },
		{ label: '下划线', value: 'snake' },
		{ label: '中划线', value: 'kebab' },
		{ label: '下划线大写', value: 'snakeAllCaps' },
		{ label: '中划线大写', value: 'kebabAllCaps' },
		{ label: '中划线首字母大写', value: 'train' }
	];
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<!-- <Title title="字符串转换" /> -->
	<div class="w-full text-token card p-4">
		<SingleSelector
			bind:value={type_}
			icon="sync_alt"
			title="输出类型"
			description="命名风格"
			{labelValues}
			on:select={calculate}
		/>
	</div>

	<div class="w-full text-token card p-4">
		<Textarea
			bind:value={text}
			on:input={calculate}
			placeholder="输入要转换的文本"
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
