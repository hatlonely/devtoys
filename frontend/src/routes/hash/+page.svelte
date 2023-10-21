<script lang="ts">
	import { Hash } from '$lib/wailsjs/go/devtoys/App';
	import { Title, RadioGroup, Textarea, TextViewer, Input } from '$lib';
	import '@fontsource/roboto-mono';

	let text = '';
	let type_ = 'md5';
	let encode = 'hex';
	let key = '';
	let hmac = false;
	let warning: any = '';
	let result = '';

	async function calculate() {
		warning = '';
		try {
			const res = await Hash({
				Text: text,
				Type: type_,
				Hmac: hmac,
				Encode: encode,
				Key: key
			});
			result = res.Text;
		} catch (err) {
			warning = err;
		}
	}

	$: text, calculate();
	$: key, calculate();
	$: hmac, calculate();
	$: type_, calculate();
	$: encode, calculate();

	const labelValues = [
		{
			label: 'MD5',
			value: 'md5'
		},
		{
			label: 'SHA1',
			value: 'sha1'
		},
		{
			label: 'SHA256',
			value: 'sha256'
		},
		{
			label: 'SHA224',
			value: 'sha224'
		},
		{
			label: 'SHA512',
			value: 'sha512'
		}
	];
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<!-- <Title title="哈希散列工具" /> -->

	<div class="w-full text-token card p-4">
		<RadioGroup
			bind:group={encode}
			name="encode"
			icon="sync_alt"
			title="编码类型"
			description="输出结果的编码类型"
			items={[
				{
					label: '十六进制',
					value: 'hex'
				},
				{
					label: 'Base64',
					value: 'base64'
				}
			]}
		/>

		<RadioGroup
			bind:group={type_}
			name="type_"
			icon="sync_alt"
			title="哈希算法"
			description="哈希算法类型"
			items={labelValues}
		/>

		<RadioGroup
			bind:group={hmac}
			name="hamc"
			icon="sync_alt"
			title="HMAC"
			description="HMAC 模式，需要输入密码"
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

		{#if hmac}
			<Input
				bind:value={key}
				on:input={calculate}
				on:clear={calculate}
				title="HMAC 密码"
				placeholder="输入密码"
				code={true}
				{warning}
			/>
		{/if}
	</div>

	<div class="w-full text-token card p-4">
		<Textarea
			bind:value={text}
			on:input={calculate}
			placeholder="输入要哈希的文本"
			code={true}
			enableUpload={true}
			{warning}
		/>
	</div>

	{#each labelValues as labelValue}
		{#if result && type_ == labelValue.value}
			<div class="w-full text-token card p-4">
				<TextViewer title={labelValue.label} text={result} />
			</div>
		{/if}
	{/each}
</div>
