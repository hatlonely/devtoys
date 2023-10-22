<script lang="ts">
	import { Escape } from '$lib/wailsjs/go/devtoys/App';
	import { Title, RadioGroupItem as RadioGroup2, Textarea, TextViewer } from '$lib';
	import { fade } from 'svelte/transition';
	import '@fontsource/roboto-mono';

	let mode = 'unescape';
	let type_ = 'string';
	let text = '';
	let output = '';
	let warning = '';

	function calculate() {
		Escape({
			text: text,
			mode,
			type: type_
		})
			.then((res: any) => {
				output = res.Text;
				warning = '';
			})
			.catch((err) => {
				warning = err;
			});
	}

	function onModeChange() {
		text = output;
		calculate();
	}

	$: mode, onModeChange();
	$: type_, calculate();
	$: text, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<RadioGroup2
			bind:group={mode}
			name="mode"
			title="转义模式"
			description="选择您想使用的转义模式"
			icon="sync_alt"
			items={[
				{
					label: '转义',
					value: 'escape'
				},
				{
					label: '反转义',
					value: 'unescape'
				}
			]}
		/>

		<RadioGroup2
			bind:group={type_}
			name="type_"
			title="转义类型"
			description="常规编码/URL安全编码"
			icon="link"
			items={[
				{
					label: 'HTML',
					value: 'html'
				},
				{
					label: 'URL',
					value: 'url'
				},
				{
					label: 'URLQuery',
					value: 'urlquery'
				},
				{
					label: 'String',
					value: 'string'
				}
			]}
		/>
	</div>

	<div class="w-full text-token card p-4">
		<Textarea
			bind:value={text}
			on:input={calculate}
			placeholder="输入要转义/反转义的文本"
			code={true}
			{warning}
		/>
	</div>

	{#if output}
		<div class="w-full text-token card p-4" in:fade out:fade>
			<TextViewer bind:text={output} />
		</div>
	{/if}
</div>
