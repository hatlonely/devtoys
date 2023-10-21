<script lang="ts">
	import Button from './Button.svelte';
	import Warning from './Warning.svelte';
	import { createEventDispatcher } from 'svelte';
	import { FileDropzone } from '@skeletonlabs/skeleton';
	import TextViewer from './TextViewer.svelte';

	export let value: string;
	export let row = 4;
	export let placeholder = '';
	export let code = false;
	export let warning = '';
	export let title = '';
	export let extraButton: any = [];

	export let enableUpload = false;
	export let uploadMessage = '点击或者拖拽文件';
	export let uploadAccept = 'text/*';

	let upload = false;
	let files: FileList;
	let uploadDone = false;

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

		value = await readFile(file);
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
				<TextViewer text={value} />
			{:else if files[0].type.startsWith('image/')}
				<img src={URL.createObjectURL(files[0])} alt={files[0].name} />
			{:else if files[0].type === 'application/x-x509-ca-cert'}
				<TextViewer text={value} />
			{:else}
				<TextViewer text={files[0].name} />
			{/if}
		{:else}
			<textarea
				bind:value
				on:input
				on:input={updateTextareaHeight}
				class="textarea {code ? 'devtoys-code' : ''}"
				spellcheck="false"
				rows={row}
				{placeholder}
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
</style>
