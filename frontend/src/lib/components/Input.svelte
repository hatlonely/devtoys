<script lang="ts">
	import Button from './Button.svelte';
	import Warning from './Warning.svelte';
	import { createEventDispatcher } from 'svelte';

	export let value: string;
	export let placeholder = '';
	export let code = false;
	export let warning = '';
	export let title = '';

	const dispatch = createEventDispatcher();

	function paste() {
		navigator.clipboard.readText().then((text) => {
			value = text;
		});
	}

	function clear() {
		value = '';
		dispatch('clear', {});
	}

	function enter(e: any) {
		if (e.key === 'Enter' || e.keyCode === 13) {
			dispatch('enter', {});
		}
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
		<input
			bind:value
			on:input
			on:keyup={enter}
			class="input {code ? 'devtoys-code' : ''}"
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
