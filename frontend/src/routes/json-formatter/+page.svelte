<script lang="ts">
	import { Textarea, CodeViewer } from '$lib';

	let text = '';
	let json = '';
	let warning: any = '';

	async function calculate() {
		if (text === '') {
			json = '';
			return;
		}

		warning = '';
		try {
			const obj = JSON.parse(text);
			json = JSON.stringify(obj, null, 2);
		} catch (err) {
			warning = err;
		}
	}

	$: text, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<Textarea
			title="输入"
			bind:value={text}
			on:input={calculate}
			placeholder="输入要格式化的 JSON 字符串"
			code={true}
			{warning}
		/>
	</div>

	{#if json}
		<div class="w-full text-token card p-4">
			<CodeViewer title="输出" value={json} />
		</div>
	{/if}
</div>
