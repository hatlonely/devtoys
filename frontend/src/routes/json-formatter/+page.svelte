<script lang="ts">
	import { Textarea, TextViewer, SingleSelector } from '$lib';

	let text = '';
	let json = '';
	let warning: any = '';

	async function calculate() {
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
			bind:value={text}
			on:input={calculate}
			placeholder="输入要转换的文本"
			code={true}
			{warning}
		/>
	</div>

	{#if json}
		<div class="w-full text-token card p-4">
			<TextViewer title="输出" text={json} />
		</div>
	{/if}
</div>
