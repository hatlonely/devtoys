<script lang="ts">
	import { Base64Decode, Base64Encode } from '$lib/wailsjs/go/devtoys/App';
	import { Title, RadioGroup, Button, Textarea, Warning, TextViewer } from '$lib';
	import '@fontsource/roboto-mono';

	let mode = 'encode';
	let type_ = 'std';
	let align = true;
	let input = '';
	let output = '';
	let warning = '';

	function calculate() {
		if (mode === 'encode') {
			Base64Encode(input, type_, align).then((res) => {
				output = res;
				warning = '';
			});
		} else {
			Base64Decode(input, type_, align)
				.then((res) => {
					output = res;
					warning = '';
				})
				.catch((err) => {
					warning = err;
				});
		}
	}

	function onOpChange() {
		input = output;
		calculate();
	}

	$: mode, onOpChange();
	$: type_, calculate();
	$: align, calculate();
	$: input, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<Title title="Base64 文本编解/解码工具" />

	<div class="w-full text-token card p-4">
		<RadioGroup
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

		<RadioGroup
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

		<RadioGroup
			bind:group={align}
			name="align"
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
			bind:value={input}
			on:input={calculate}
			placeholder="输入要编码/解码的文本"
			code={true}
		/>

		<Warning bind:message={warning} />
	</div>

	<div class="w-full text-token card p-4">
		<TextViewer bind:text={output} />
	</div>
</div>
