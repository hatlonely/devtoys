<script lang="ts">
	import Button from './Button.svelte';
	import Warning from './Warning.svelte';
	import { createEventDispatcher } from 'svelte';

	export let value: string;
	export let row = 4;
	export let placeholder = '';
	export let code = false;
	export let warning = '';
	export let title = '';

	const dispatch = createEventDispatcher();

	function updateTextareaHeight(event: Event) {
		const target = event.target as HTMLTextAreaElement;
		target.style.height = 'auto';

		if (target.scrollHeight > window.innerHeight / 3) {
			target.style.height = `${window.innerHeight / 3}px`;
		} else {
			target.style.height = `${target.scrollHeight + 1}px`;
		}
	}

	function paste() {
		navigator.clipboard.readText().then((text) => {
			value = text;
		});
	}

	function clear() {
		value = '';
		dispatch('clear', {});
	}
</script>

<div class="space-y-2">
	<div class="flex">
		<div class="flex flex-auto">
			<span class="font-bold align-sub">{title}</span>
		</div>
		<div class="flex flex-right space-x-2">
			<Button on:click={paste} icon="content_paste" text="粘贴" />
			<Button on:click={clear} icon="delete" text="清空" />
		</div>
	</div>
	<div>
		<textarea
			bind:value
			on:input
			on:input={updateTextareaHeight}
			class="textarea {code ? 'devtoys-code' : ''}"
			spellcheck="false"
			rows={row}
			{placeholder}
		/>
	</div>
	<Warning bind:message={warning} />
</div>

<style>
	.devtoys-code {
		font-size: 1rem;
		font-family: 'Roboto Mono', monospace;
		overflow-wrap: anywhere;
		white-space: pre-wrap; /* Since CSS 2.1 */
		white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
		white-space: -pre-wrap; /* Opera 4-6 */
		white-space: -o-pre-wrap; /* Opera 7 */
		word-wrap: break-word; /* Internet Explorer 5.5+ */
	}
</style>
