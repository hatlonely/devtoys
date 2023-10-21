<script lang="ts">
	import { GeneratePassword } from '$lib/wailsjs/go/devtoys/App';
	import { RadioGroup, TextViewer, Input, NumberInputItem } from '$lib';

	let length = 16;
	let enableDigit = true;
	let enableUpper = true;
	let enableLower = true;
	let enableSpecific = false;
	let specific = '!@#$%^&*()_+-=';
	let characterSet = '';
	let warning: any = '';
	let result = '';

	async function calculate() {
		warning = '';
		try {
			let res = await GeneratePassword({
				Length: length,
				EnableDigit: enableDigit,
				EnableUpper: enableUpper,
				EnableLower: enableLower,
				EnableSpecific: enableSpecific,
				Specific: specific,
				CharacterSet: characterSet
			});
			result = res.Text;
		} catch (err) {
			warning = err;
		}
	}

	function calculateCharacterSet() {
		characterSet = '';
		if (enableDigit) {
			characterSet += '0123456789';
		}
		if (enableUpper) {
			characterSet += 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
		}
		if (enableLower) {
			characterSet += 'abcdefghijklmnopqrstuvwxyz';
		}
		if (enableSpecific) {
			characterSet += specific;
		}
	}

	$: enableDigit, calculateCharacterSet();
	$: enableUpper, calculateCharacterSet();
	$: enableLower, calculateCharacterSet();
	$: enableSpecific, calculateCharacterSet();
	$: specific, calculateCharacterSet();
	$: characterSet, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<NumberInputItem
			bind:value={length}
			title="密码长度"
			description="密码长度范围（1-128）"
			icon="sync_alt"
		/>

		<RadioGroup
			bind:group={enableDigit}
			name="enableDigit"
			title="使用数字"
			description="数字包括（0123456789）"
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

		<RadioGroup
			bind:group={enableUpper}
			name="enableUpper"
			title="使用大写字母"
			description="大写字母包括（ABCDEFGHIJKLMNOPQRSTUVWXYZ）"
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

		<RadioGroup
			bind:group={enableLower}
			name="enableLower"
			title="使用小写字母"
			description="小写字母包括（abcdefghijklmnopqrstuvwxyz）"
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

		<RadioGroup
			bind:group={enableSpecific}
			name="enableSpecific"
			title="使用特殊字符"
			description="特殊字符包括（!@#$%^&*()_+-=）"
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

		{#if enableSpecific}
			<Input
				bind:value={specific}
				title="自定义特殊字符"
				placeholder="输入特殊字符集合"
				code={true}
				{warning}
			/>
		{/if}
	</div>

	<div class="w-full text-token card p-4">
		<Input
			bind:value={characterSet}
			on:input={calculate}
			on:clear={calculate}
			title="字符集合"
			placeholder="输入字符集合"
			code={true}
			{warning}
			extraButton={[
				{
					icon: 'submit',
					iconSet: 'custom',
					text: '生成',
					onClick: () => {
						calculate();
					}
				}
			]}
		/>
	</div>

	{#if result}
		<div class="w-full text-token card p-4">
			<TextViewer title="输出" text={result} />
		</div>
	{/if}
</div>
