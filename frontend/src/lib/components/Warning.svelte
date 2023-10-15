<script lang="ts">
	export let message: string;
	import 'material-symbols';
</script>

{#if message}
	<div class="alert variant-filled-error devtoys-code space-x-8">
		<div class="ml-5">
			<span class="material-symbols-outlined">warning</span>
		</div>
		<div class="alert-message space-y-0">
			<h3 class="h3">错误</h3>
			<p>{message}</p>
		</div>
	</div>
{/if}

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

	.material-symbols-outlined {
		font-variation-settings: 'FILL' 0, 'wght' 200, 'GRAD' 0, 'opsz' 24;
	}
</style>
