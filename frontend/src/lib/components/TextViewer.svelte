<script lang="ts">
	import { clipboard } from '@skeletonlabs/skeleton';
	import 'material-symbols';

	export let text: string;

	let copied = false;
</script>

{#if text}
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
			use:clipboard={text}
		>
			<span class="material-symbols-outlined">
				{copied ? 'done' : 'content_copy'}
			</span>
		</button>
		<pre class="dt-code">{text}</pre>
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
