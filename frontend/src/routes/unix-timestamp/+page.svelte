<script lang="ts">
	import 'material-symbols';
	import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';
	import { AnalystTimeInfo } from '$lib/wailsjs/go/devtoys/App';
	import { clipboard } from '@skeletonlabs/skeleton';
	import '@fontsource/roboto-mono';

	let input = '';
	let info = {
		Timestamp: 0,
		TimestampMill: 0,
		LocalTime: '',
		UTCTime: '',
		Relative: ''
	};
	let copied = {
		Timestamp: false,
		TimestampMill: false,
		LocalTime: false,
		UTCTime: false,
		Relative: false
	};
	let warning = '';
	let pasted = false;
	let cleared = false;

	function calculate() {
		AnalystTimeInfo(input)
			.then((res) => {
				console.log(res);
				info = res;
				warning = '';
			})
			.catch((err) => {
				warning = err;
			});
	}

	function paste() {
		navigator.clipboard.readText().then((text) => {
			input = text;
			calculate();
		});
	}

	function clear() {
		input = '';
		calculate();
	}

	let display = [
		{ field: 'Timestamp', label: '当前时间戳' },
		{ field: 'TimestampMill', label: '当前时间戳（毫秒）' },
		{ field: 'LocalTime', label: '本地时间' },
		{ field: 'UTCTime', label: 'UTC 时间' },
		{ field: 'Relative', label: '相对时间' }
	];
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<h2 class="h2 my-4">Unix 时间戳</h2>

	<div class="w-full text-token card p-4 space-y-4">
		<div class="w-full text-token p-2 space-x-10 flex flex-row">
			<span class="flex flex-row">
				<span class="badge-icon variant-soft-secondary mt-2">
					<span class="material-symbols-outlined">schedule</span>
				</span>
				<span class="flex-auto ml-5 w-fit">
					<dt class="font-bold">时间</dt>
					<dd class="text-sm opacity-50">任意的合法的时间格式均可</dd>
				</span>
			</span>

			<span class="flex-auto">
				<input
					bind:value={input}
					on:input={calculate}
					class="input app-code"
					placeholder="输入时间"
				/>
			</span>

			<span class="flex-right">
				<button
					on:click={paste}
					on:click={() => {
						pasted = true;
						setTimeout(() => {
							pasted = false;
						}, 2000);
					}}
					type="button"
					class="btn btn-sm variant-filled mx-2"
				>
					<span>粘贴</span>
					<span class="material-symbols-outlined">
						{pasted ? 'done' : 'content_paste'}
					</span>
				</button>

				<button
					on:click={clear}
					on:click={() => {
						cleared = true;
						setTimeout(() => {
							cleared = false;
						}, 2000);
					}}
					type="button"
					class="btn btn-sm variant-filled mx-2"
				>
					<span>当前时间</span>
					<span class="material-symbols-outlined">
						{cleared ? 'done' : 'update'}
					</span>
				</button>
			</span>
		</div>
	</div>

	{#if warning}
		<div class="alert variant-filled-error app-code">
			<span class="material-symbols-outlined">warning</span>
			<div class="alert-message" data-toc-ignore>
				<h3 class="h3" data-toc-ignore>Warning</h3>
				<p>{warning}</p>
			</div>
		</div>
	{/if}

	<div class="w-full text-token card p-4 space-y-4">
		<div class="w-full p-4">
			{#each display as v}
				<div class="flex flex-row">
					<span class="w-1/3">{v.label}</span>
					<span class="flex-auto app-code">{info[v.field] ? info[v.field] : ''}</span>
					<button
						type="button"
						class="btn btn-sm"
						on:click={() => {
							copied[v.field] = true;
							setTimeout(() => {
								copied[v.field] = false;
							}, 2000);
						}}
						use:clipboard={info[v.field]}
					>
						<span class="material-symbols-outlined">
							{copied[v.field] ? 'done' : 'content_copy'}
						</span>
					</button>
				</div>
			{/each}
		</div>
	</div>
</div>

<style>
	.app-code {
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
