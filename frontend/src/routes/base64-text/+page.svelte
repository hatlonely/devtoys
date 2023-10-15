<script lang="ts">
	import 'material-symbols';
	import { Base64Decode, Base64Encode } from '$lib/wailsjs/go/devtoys/App';
	import { clipboard } from '@skeletonlabs/skeleton';
	import { Title, RadioGroup } from '$lib';
	import '@fontsource/roboto-mono';

	let op = 'encode';
	let mode = 'std';
	let align = true;
	let input = '';
	let output = '';
	let warning = '';
	let copied = false;
	let pasted = false;
	let cleared = false;

	function calculate() {
		if (op === 'encode') {
			Base64Encode(input, mode, align).then((res) => {
				output = res;
				warning = '';
			});
		} else {
			Base64Decode(input, mode, align)
				.then((res) => {
					output = res;
					warning = '';
				})
				.catch((err) => {
					warning = err;
				});
		}
	}

	function paste() {
		navigator.clipboard.readText().then((text) => {
			input = text;
			calculate();
		});
	}

	function clear() {
		input = '';
		output = '';
	}

	function updateTextareaHeight(event: Event) {
		const target = event.target as HTMLTextAreaElement;
		target.style.height = 'auto';

		if (target.scrollHeight > window.innerHeight / 3) {
			target.style.height = `${window.innerHeight / 3}px`;
		} else {
			target.style.height = `${target.scrollHeight + 1}px`;
		}
	}

	function onOpChange() {
		input = output;
		calculate();
	}

	$: op, onOpChange();
	$: mode, calculate();
	$: align, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<Title title="Base64 文本编解/解码工具" />

	<div class="w-full text-token card p-4 space-y-4">
		<RadioGroup
			bind:group={op}
			name="op"
			title="转换模式"
			description="选择您想使用的转换模式"
			icon="sync_alt"
			items={[
				{
					label: '编码',
					value: 'encode'
				},
				{
					label: '解码',
					value: 'decode'
				}
			]}
		/>

		<RadioGroup
			bind:group={mode}
			name="mode"
			title="编码模式"
			description="常规编码/URL安全编码"
			icon="link"
			items={[
				{
					label: '常规',
					value: 'std'
				},
				{
					label: 'URL',
					value: 'url'
				}
			]}
		/>

		<RadioGroup
			bind:group={align}
			name="align"
			title="补齐模式"
			description="末尾是否使用 = 补齐"
			icon="equal"
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
	</div>

	<div class="flex flex-row-reverse">
		<button
			on:click={clear}
			on:click={() => {
				cleared = true;
				setTimeout(() => {
					cleared = false;
				}, 2000);
			}}
			type="button"
			class="btn btn-sm variant-filled-primary mx-2"
		>
			<span>清空</span>
			<span class="material-symbols-outlined">
				{cleared ? 'done' : 'delete'}
			</span>
		</button>
		<button
			on:click={paste}
			on:click={() => {
				pasted = true;
				setTimeout(() => {
					pasted = false;
				}, 2000);
			}}
			type="button"
			class="btn btn-sm variant-filled-primary mx-2"
		>
			<span>粘贴</span>
			<span class="material-symbols-outlined">
				{pasted ? 'done' : 'content_paste'}
			</span>
		</button>
	</div>

	<textarea
		bind:value={input}
		on:input={calculate}
		on:input={updateTextareaHeight}
		class="textarea app-code"
		rows="4"
		placeholder="输入要编码/解码的文本"
	/>

	{#if warning}
		<div class="alert variant-filled-error app-code">
			<span class="material-symbols-outlined">warning</span>
			<div class="alert-message" data-toc-ignore>
				<h3 class="h3" data-toc-ignore>Warning</h3>
				<p>{warning}</p>
			</div>
		</div>
	{/if}

	{#if output}
		<div class="w-full card p-4">
			<button
				type="button"
				class="btn btn-sm p-0 float-right variant-filled-primary"
				on:click={() => {
					copied = true;
					setTimeout(() => {
						copied = false;
					}, 2000);
				}}
				use:clipboard={output}
			>
				<span class="material-symbols-outlined">
					{copied ? 'done' : 'content_copy'}
				</span>
			</button>
			<pre class="app-code">{output}</pre>
		</div>
	{/if}
</div>

<style>
	.app-code {
		font-size: 0.8rem;
		font-family: 'Roboto Mono', monospace;
		overflow-wrap: anywhere;
		white-space: pre-wrap; /* Since CSS 2.1 */
		white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
		white-space: -pre-wrap; /* Opera 4-6 */
		white-space: -o-pre-wrap; /* Opera 7 */
		word-wrap: break-word; /* Internet Explorer 5.5+ */
	}

	.material-symbols-outlined {
		font-variation-settings: 'FILL' 0, 'wght' 200, 'GRAD' 0, 'opsz' 24;
	}
</style>
