<script lang="ts">
	import { GenerateSSHKey } from '$lib/wailsjs/go/devtoys/App';
	import { CodeViewer, RadioGroupItem } from '$lib';
	import '@fontsource/roboto-mono';

	let type_ = 'rsa';
	let warning: any = '';
	let result: any;

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
			title="加密算法"
			description="生成 sshkey 的加密算法"
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
