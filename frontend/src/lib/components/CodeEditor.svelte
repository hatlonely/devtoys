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
	import Button from './Button.svelte';
	import CodeViewer from './CodeViewer.svelte';
	import Warning from './Warning.svelte';
	import { createEventDispatcher } from 'svelte';
	import { FileDropzone } from '@skeletonlabs/skeleton';

	export let title = '';
	export let value = '';
	export let syntax = '';
	export let withLineNumbers = false;

	export let warning = '';

	export let extraButton: any = [];

	export let enableUpload = false;
	export let uploadMessage = '点击或者拖拽文件';
	export let uploadAccept = 'text/*';
	export let asByte = false;

	let upload = false;
	let files: FileList;
	let uploadDone = false;

	const dispatch = createEventDispatcher();

	function paste() {
		navigator.clipboard.readText().then((text) => {
			value = text;
		});
	}

	function clear() {
		value = '';
		uploadDone = false;
		dispatch('clear', {});
	}

	function toggleUpload() {
		upload = !upload;
	}

	function readFile(file: File): Promise<string> {
		return new Promise<string>((resolve, reject) => {
			const reader = new FileReader();
			reader.onload = () => {
				resolve(reader.result as string);
			};
			reader.onerror = () => {
				reject(reader.error);
			};
			reader.readAsText(file);
		});
	}

	async function onUpload(): Promise<void> {
		warning = '';
		if (!files || files.length === 0) {
			return;
		}

		const file = files[0];
		const acceptTypes = uploadAccept.split(',');
		if (!acceptTypes.some((type) => file.type.match(type.trim()))) {
			warning = `不支持的文件类型：${file.type}`;
			return;
		}

		const content = await readFile(file);
		if (file.type.startsWith('text/')) {
			value = content;
			asByte = false;
		} else {
			value = btoa(content);
			asByte = true;
		}

		uploadDone = true;
		dispatch('upload', { value });
	}

	$: files, onUpload;
</script>

<div class="space-y-2">
	<div class="flex">
		<div class="flex flex-auto">
			<span class="font-bold devtoys-align-center">{title}</span>
		</div>
		<div class="flex flex-right space-x-2">
			<div class="flex flex-right space-x-2">
				{#each extraButton as button}
					<Button
						on:click={button.onClick}
						icon={button.icon}
						iconSet={button.iconSet}
						text={button.text}
					/>
				{/each}

				{#if enableUpload}
					<Button on:click={toggleUpload} icon="upload" text="文件" />
				{/if}
				<Button on:click={paste} icon="content_paste" text="粘贴" />
				<Button on:click={clear} icon="clear" text="清空" />
			</div>
		</div>
	</div>

	<div>
		{#if enableUpload && upload}
			{#if !uploadDone}
				<FileDropzone name="files" accept={uploadAccept} bind:files on:change={onUpload}>
					<svelte:fragment slot="lead">
						<span class="material-symbols-outlined text-6xl"> upload_file </span>
					</svelte:fragment>
					<svelte:fragment slot="message">{uploadMessage}</svelte:fragment>
					<svelte:fragment slot="meta">支持的文件类型：{uploadAccept}</svelte:fragment>
				</FileDropzone>
			{:else if files[0].type.startsWith('text/')}
				<CodeViewer title={files[0].name} {value} />
			{:else if files[0].type.startsWith('image/')}
				<img src={URL.createObjectURL(files[0])} alt={files[0].name} />
			{:else if files[0].type === 'application/x-x509-ca-cert'}
				<CodeViewer title={files[0].name} {value} />
			{:else}
				<CodeViewer value={files[0].name} />
			{/if}
		{:else}
			<CodeJar
				class="hljs devtoys-code p-4 rounded-xl disabled"
				{withLineNumbers}
				{syntax}
				{highlight}
				bind:value
			/>
		{/if}
	</div>
	<Warning bind:message={warning} />
</div>

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
