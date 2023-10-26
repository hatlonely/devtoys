<script lang="ts">
	import { AES } from '$lib/wailsjs/go/devtoys/App';
	import { TextViewer, RadioGroupItem, Textarea, TextInputItem } from '$lib';

	let func = 'encrypt';
	let text = '';
	let iv = '';
	let base64Text = false;
	let key = '';
	let base64Key = false;
	let type_ = 'AES128';
	let encryptionMode = 'cbc';
	let warning: any = '';
	let result = '';

	async function calculate() {
		warning = '';
		try {
			let res = await AES({
				Function: func,
				Text: text,
				IV: iv,
				Base64Text: base64Text,
				Key: key,
				Base64Key: base64Key,
				Type: type_,
				EncryptionMode: encryptionMode
			});
			result = res.Text;
		} catch (err) {
			warning = err;
		}
	}

	$: text, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<RadioGroupItem
			bind:group={func}
			name="func"
			title="加解密"
			description="加密/解密"
			icon="sync_alt"
			items={[
				{
					label: '加密',
					value: 'encrypt'
				},
				{
					label: '解密',
					value: 'decrypt'
				}
			]}
		/>

		<RadioGroupItem
			bind:group={base64Text}
			name="base64Text"
			title="Base64文本"
			description="是否使用 base64 编码文本"
			icon="sync_alt"
			items={[
				{
					label: '使用',
					value: true
				},
				{
					label: '不用',
					value: false
				}
			]}
		/>

		<RadioGroupItem
			bind:group={base64Key}
			name="base64Key"
			title="Base64秘钥"
			description="是否使用 base64 编码秘钥"
			icon="sync_alt"
			items={[
				{
					label: '使用',
					value: true
				},
				{
					label: '不用',
					value: false
				}
			]}
		/>

		<RadioGroupItem
			bind:group={type_}
			name="type_"
			title="AES 类型"
			description="AES 类型"
			icon="sync_alt"
			items={[
				{
					label: 'AES128',
					value: 'AES128'
				},
				{
					label: 'AES192',
					value: 'AES192'
				},
				{
					label: 'AES256',
					value: 'AES256'
				}
			]}
		/>

		<RadioGroupItem
			bind:group={encryptionMode}
			name="encryptionMode"
			title="加密模式"
			description="加密模式"
			icon="sync_alt"
			items={[
				{
					label: 'CBC',
					value: 'cbc'
				},
				{
					label: 'CFB',
					value: 'cfb'
				},
				{
					label: 'CTR',
					value: 'ctr'
				},
				{
					label: 'OFB',
					value: 'ofb'
				}
			]}
		/>
	</div>

	<div class="w-full text-token card p-4">
		<Textarea
			title="输入"
			bind:value={text}
			on:input={calculate}
			placeholder="输入要加密的文本"
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
