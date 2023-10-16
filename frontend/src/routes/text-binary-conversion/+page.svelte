<script lang="ts">
	import { Base64Decode, Base64Encode, ConvertTextBinary } from '$lib/wailsjs/go/devtoys/App';
	import { Title, RadioGroup, Textarea, TextViewer } from '$lib';
	import '@fontsource/roboto-mono';

	let text = '';
	let to = '';
	let lowercase = false;
	let withoutSpace = false;
	let withoutFillZero = false;
	let warning = '';

	let results = {
		bin: '',
		hex: '',
		dec: '',
		base64: ''
	};

	function calculate() {
		ConvertTextBinary({
			Text: text,
			To: 'bin',
			LowerCase: lowercase,
			WithoutSpace: withoutSpace,
			WithoutFillZero: withoutFillZero
		})
			.then((res) => {
				warning = '';
				console.log(res);
				results.bin = res.Text;
			})
			.catch((err) => {
				warning = err;
			});
	}

	$: text, calculate();
	$: lowercase, calculate();
	$: withoutSpace, calculate();
	$: withoutFillZero, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<Title title="Base64 文本编解/解码工具" />

	<div class="w-full text-token card p-4">
		<RadioGroup
			bind:group={withoutSpace}
			name="withoutSpace"
			title="空格"
			description="是否在结果中使用空格分隔"
			icon="sync_alt"
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
			title="填充 0"
			description="是否在结果中用 0 填充"
			icon="sync_alt"
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
			title="大小写"
			description="是否使用小写字符编码（仅十六进制有效）"
			icon="sync_alt"
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

	{#if results.bin}
		<div class="w-full text-token card p-4">
			<TextViewer title="二进制" bind:text={results.bin} />
		</div>
	{/if}
</div>
