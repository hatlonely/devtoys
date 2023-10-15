<script lang="ts">
	import { clipboard } from '@skeletonlabs/skeleton';
	import 'material-symbols';
	import Button from './Button.svelte';

	export let text: string;

	let copied = false;
</script>

{#if text}
	<div class="space-y-2">
		<div>
			<button
				type="button"
				class="btn btn-sm mx-2 variant-soft hover:variant-filled"
				on:click={() => {
					copied = true;
					setTimeout(() => {
						copied = false;
					}, 2000);
				}}
				use:clipboard={text}
			>
				<span>复制</span>
				<span class="material-symbols-outlined">
					{copied ? 'done' : 'content_copy'}
				</span>
			</button>
		</div>
		<div class="w-full card p-4">
			<pre class="dt-code">{text}</pre>
		</div>
	</div>
{/if}

<style>
	.dt-code {
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
