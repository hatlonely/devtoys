<script lang="ts">
	import { GenerateSSHKey } from '$lib/wailsjs/go/devtoys/App';
	import { CodeViewer, RadioGroupItem } from '$lib';
	import '@fontsource/roboto-mono';

	let type_ = 'rsa';
	let warning: any = '';
	let privateKey = '';
	let publicKey = '';
	let result: any;

	//     type SSHKeyGeneratorRes struct {
	// 	PrivateKey string
	// 	PublicKey  string
	// }

	async function calculate() {
		warning = '';
		try {
			let res = await GenerateSSHKey({
				Type: type_
			});
			result = res;
			console.log(res);
		} catch (err) {
			warning = err;
		}
	}

	$: type_, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<RadioGroupItem
			bind:group={type_}
			name="enableDigit"
			title="使用数字"
			description="数字包括（0123456789）"
			icon="sync_alt"
			items={[
				{
					label: 'RSA',
					value: 'rsa'
				},
				{
					label: 'ECDSA',
					value: 'ecdsa'
				},
				{
					label: 'ED25519',
					value: 'ed25519'
				}
			]}
		/>
	</div>

	{#if result}
		<div class="w-full text-token card p-4">
			<CodeViewer title="输出" withLineNumbers={false} value={result?.PrivateKey} />
		</div>

		<div class="w-full text-token card p-4">
			<CodeViewer title="输出" withLineNumbers={false} value={result?.PublicKey} />
		</div>
	{/if}
</div>
