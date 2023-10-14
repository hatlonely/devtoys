<script lang="ts">
	import 'material-symbols';
	import { AppShell } from '@skeletonlabs/skeleton';
	import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';
	import { CodeBlock } from '@skeletonlabs/skeleton';
	import { Base64Decode, Base64Encode } from '$lib/wailsjs/go/apps/Base64TextApp';

	import { CodeJar } from '@novacbn/svelte-codejar';
	import hljs from 'highlight.js/lib/core';
	const highlight = (code, syntax) => hljs.highlightAuto(code).value;

	import { HighlightAuto, LineNumbers } from 'svelte-highlight';
	import github from 'svelte-highlight/styles/monokai';

	let op = 'encode';
	let mode = 'std';
	let align = true;
	let input = '';
	let output = '';

	function calculate() {
		if (op === 'encode') {
			Base64Encode(input, mode, align).then((res) => {
				output = res;
			});
		} else {
			Base64Decode(input, mode, align)
				.then((res) => {
					output = res;
				})
				.catch((err) => {
					output = err;
				});
		}
	}

	function paste() {
		navigator.clipboard.readText().then((text) => {
			input = text;
			calculate();
		});
	}

	function updateTextareaHeight(event: Event) {
		const target = event.target as HTMLTextAreaElement;
		target.style.height = 'auto';
		target.style.height = `${target.scrollHeight + 1}px`;
	}

	$: op, calculate();
	$: mode, calculate();
	$: align, calculate();
</script>

<svelte:head>
	{@html github}
</svelte:head>

<AppShell>
	<svelte:fragment>
		<div class="card p-4 w-full text-token space-y-4 mt-2">
			<h2 class="h2 my-4">Base64 文本编解/解码工具</h2>

			<div class="w-full text-token card p-4 space-y-4">
				<div class="w-full text-token p-2 flex flex-row">
					<span class="badge-icon p-4 variant-soft-secondary mt-2">
						<span class="material-symbols-outlined">sync_alt</span>
					</span>
					<span class="flex-auto ml-5">
						<dt class="font-bold">转换模式</dt>
						<dd class="text-sm opacity-50">选择您想使用的转换模式</dd>
					</span>
					<span>
						<RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
							<RadioItem bind:group={op} name="op" value="encode">编码</RadioItem>
							<RadioItem bind:group={op} name="op" value="decode">解码</RadioItem>
						</RadioGroup>
					</span>
				</div>

				<div class="w-full text-token p-2 flex flex-row">
					<span class="badge-icon p-4 variant-soft-secondary mt-2">
						<span class="material-symbols-outlined">link</span>
					</span>
					<span class="flex-auto ml-5">
						<dt class="font-bold">编码模式</dt>
						<dd class="text-sm opacity-50">常规编码/URL安全编码</dd>
					</span>
					<span>
						<RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
							<RadioItem bind:group={mode} name="mode" value="std">常规</RadioItem>
							<RadioItem bind:group={mode} name="mode" value="url">URL</RadioItem>
						</RadioGroup>
					</span>
				</div>

				<div class="w-full text-token p-2 flex flex-row">
					<span class="badge-icon p-4 variant-soft-secondary mt-2">
						<span class="material-symbols-outlined">equal</span>
					</span>
					<span class="flex-auto ml-5">
						<dt class="font-bold">补齐模式</dt>
						<dd class="text-sm opacity-50">末尾是否使用 = 补齐</dd>
					</span>
					<span>
						<RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
							<RadioItem bind:group={align} name="align" value={true}>开启</RadioItem>
							<RadioItem bind:group={align} name="align" value={false}>关闭</RadioItem>
						</RadioGroup>
					</span>
				</div>
			</div>

			<div class="flex flex-row-reverse">
				<button type="button" class="btn btn-sm variant-filled mx-2">
					<span>清空</span> <span class="material-symbols-outlined">delete</span>
				</button>
				<button type="button" on:click={paste} class="btn btn-sm variant-filled mx-2">
					<span>粘贴</span> <span class="material-symbols-outlined">content_paste</span>
				</button>
			</div>
			<!-- <textarea
				bind:value={input}
				on:input={(e) => {
					calculate();
					updateTextareaHeight(e);
				}}
				class="textarea h-auto resize-none"
				placeholder="输入要编码/解码的文本"
				rows="4"
			/> -->

			<div class="card shadow-none">
				<CodeJar
					class="hljs"
					{highlight}
					bind:value={input}
					on:change={calculate}
					withLineNumbers={true}
				/>
			</div>

			<CodeBlock bind:code={output} lineNumbers={true} />
			<HighlightAuto bind:code={output} let:highlighted>
				<LineNumbers {highlighted} wrapLines style="break-normal" />
			</HighlightAuto>
		</div>
	</svelte:fragment>
</AppShell>
