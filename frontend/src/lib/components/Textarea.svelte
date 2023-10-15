<script lang="ts">
	export let value: string;
	export let row = 4;
	export let placeholder = '';
	export let code = false;

	function updateTextareaHeight(event: Event) {
		const target = event.target as HTMLTextAreaElement;
		target.style.height = 'auto';

		if (target.scrollHeight > window.innerHeight / 3) {
			target.style.height = `${window.innerHeight / 3}px`;
		} else {
			target.style.height = `${target.scrollHeight + 1}px`;
		}
	}
</script>

<textarea
	bind:value
	on:input
	on:input={updateTextareaHeight}
	class="textarea {code ? 'devtoys-code' : ''}}"
	rows={row}
	{placeholder}
/>

<style>
	.devtoys-code {
		font-size: 0.8rem;
		font-family: 'Roboto Mono', monospace;
		overflow-wrap: anywhere;
		white-space: pre-wrap; /* Since CSS 2.1 */
		white-space: -moz-pre-wrap; /* Mozilla, since 1999 */
		white-space: -pre-wrap; /* Opera 4-6 */
		white-space: -o-pre-wrap; /* Opera 7 */
		word-wrap: break-word; /* Internet Explorer 5.5+ */
	}
</style>
