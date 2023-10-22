<script context="module" lang="ts">
	import hljs from 'highlight.js/lib/core';
	import javascript from 'highlight.js/lib/languages/javascript';
	import json from 'highlight.js/lib/languages/json';
	import cpp from 'highlight.js/lib/languages/cpp';
	import python from 'highlight.js/lib/languages/python';
	import java from 'highlight.js/lib/languages/java';
	import bash from 'highlight.js/lib/languages/bash';
	import css from 'highlight.js/lib/languages/css';
	import markdown from 'highlight.js/lib/languages/markdown';
	import nginx from 'highlight.js/lib/languages/nginx';
	import php from 'highlight.js/lib/languages/php';
	import sql from 'highlight.js/lib/languages/sql';
	import xml from 'highlight.js/lib/languages/xml';
	import yaml from 'highlight.js/lib/languages/yaml';
	import typescript from 'highlight.js/lib/languages/typescript';
	import go from 'highlight.js/lib/languages/go';
	import kotlin from 'highlight.js/lib/languages/kotlin';
	import rust from 'highlight.js/lib/languages/rust';
	import swift from 'highlight.js/lib/languages/swift';
	import scala from 'highlight.js/lib/languages/scala';
	import ruby from 'highlight.js/lib/languages/ruby';
	import perl from 'highlight.js/lib/languages/perl';
	import lua from 'highlight.js/lib/languages/lua';
	import 'highlight.js/styles/monokai.css';

	hljs.registerLanguage('javascript', javascript);
	hljs.registerLanguage('json', json);
	hljs.registerLanguage('cpp', cpp);
	hljs.registerLanguage('python', python);
	hljs.registerLanguage('java', java);
	hljs.registerLanguage('bash', bash);
	hljs.registerLanguage('css', css);
	hljs.registerLanguage('markdown', markdown);
	hljs.registerLanguage('nginx', nginx);
	hljs.registerLanguage('php', php);
	hljs.registerLanguage('sql', sql);
	hljs.registerLanguage('xml', xml);
	hljs.registerLanguage('yaml', yaml);
	hljs.registerLanguage('typescript', typescript);
	hljs.registerLanguage('go', go);
	hljs.registerLanguage('kotlin', kotlin);
	hljs.registerLanguage('rust', rust);
	hljs.registerLanguage('swift', swift);
	hljs.registerLanguage('scala', scala);
	hljs.registerLanguage('ruby', ruby);
	hljs.registerLanguage('perl', perl);
	hljs.registerLanguage('lua', lua);

	const highlight = (code: any, syntax: any) => {
		if (syntax) {
			return hljs.highlight(code, {
				language: syntax
			}).value;
		} else {
			return hljs.highlightAuto(code).value;
		}
	};
</script>

<script lang="ts">
	import { CodeJar } from '@novacbn/svelte-codejar';
	import { clipboard } from '@skeletonlabs/skeleton';

	export let title = '';
	export let value = '';
	export let syntax = '';
	export let withLineNumbers = true;
	export let disabled = false;

	let copied = false;
	let codeJar: any;
</script>

{#if value}
	<div class="space-y-2">
		<div class="flex">
			<div class="flex flex-auto">
				<span class="font-bold devtoys-align-center">{title}</span>
			</div>
			<div class="flex flex-right space-x-2">
				<button
					type="button"
					class="btn btn-sm mx-2 variant-soft hover:variant-filled"
					on:click={() => {
						copied = true;
						setTimeout(() => {
							copied = false;
						}, 2000);
					}}
					use:clipboard={value}
				>
					<span>复制</span>
					<span class="material-symbols-outlined">
						{copied ? 'done' : 'content_copy'}
					</span>
				</button>
			</div>
		</div>
		<CodeJar
			class="hljs devtoys-code p-4 rounded-xl disabled"
			{withLineNumbers}
			{syntax}
			{highlight}
			{value}
			bind:this={codeJar}
		/>
	</div>
{/if}

<style>
	.devtoys-align-center {
		display: inline-flex;
		align-items: center;
		vertical-align: bottom;
		height: 100%;
	}

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
