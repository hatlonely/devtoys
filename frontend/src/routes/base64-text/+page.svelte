<script lang="ts">
	import { Base64Text } from '$lib/wailsjs/go/devtoys/App';
	import { Title, RadioGroupItem, Textarea, TextViewer } from '$lib';
	import { fade } from 'svelte/transition';
	import '@fontsource/roboto-mono';

	let mode = 'encode';
	let type_ = 'std';
	let padding = true;
	let input = '';
	let output = '';
	let warning = '';

	function calculate() {
		Base64Text({
			Text: input,
			Mode: mode,
			Type: type_,
			Padding: padding
		})
			.then((res) => {
				output = res.Text;
				warning = '';
			})
			.catch((err) => {
				warning = err;
			});
	}

	function onOpChange() {
		input = output;
		calculate();
	}

	$: mode, onOpChange();
	$: type_, calculate();
	$: padding, calculate();
	$: input, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<RadioGroupItem
			bind:group={mode}
			name="mode"
			title="转换模式"
			description="选择您想使用的转换模式"
			icon="sync_alt"
			items={[
				{
					label: '编码',
					value: 'encode'
				},
				{
					label: '解码',
					value: 'decode'
				}
			]}
		/>

		<RadioGroupItem
			bind:group={type_}
			name="type_"
			title="编码模式"
			description="常规编码/URL安全编码"
			icon="link"
			items={[
				{
					label: '常规',
					value: 'std'
				},
				{
					label: 'URL',
					value: 'url'
				}
			]}
		/>

		<RadioGroupItem
			bind:group={padding}
			name="padding"
			title="补齐模式"
			description="末尾是否使用 = 补齐"
			icon="equal"
			items={[
				{
					label: '开启',
					value: true
				},
				{
					label: '关闭',
					value: false
				}
			]}
		/>
	</div>

	<div class="w-full text-token card p-4">
		<Textarea
			title="输入"
			bind:value={input}
			on:input={calculate}
			placeholder="输入要编码/解码的文本"
			enableUpload={true}
			uploadAccept="text/*,application/x-x509-ca-cert,image/svg+xml"
			code={true}
			{warning}
		/>
	</div>

	{#if output}
		<div class="w-full text-token card p-4" in:fade out:fade>
			<TextViewer title="输出" bind:text={output} />
		</div>
	{/if}
</div>
