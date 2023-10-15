<script lang="ts">
	import 'material-symbols';
	import { AnalystTimeInfo } from '$lib/wailsjs/go/devtoys/App';
	import { clipboard } from '@skeletonlabs/skeleton';
	import '@fontsource/roboto-mono';
	import { Title, Input } from '$lib';

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

	function calculate() {
		AnalystTimeInfo(input)
			.then((res) => {
				info = res;
				warning = '';
			})
			.catch((err) => {
				warning = err;
			});
	}

	let display = [
		{ field: 'Timestamp', label: '当前时间戳' },
		{ field: 'TimestampMill', label: '当前时间戳（毫秒）' },
		{ field: 'LocalTime', label: '本地时间' },
		{ field: 'UTCTime', label: 'UTC 时间' },
		{ field: 'Relative', label: '相对时间' }
	];

	$: input, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<Title title="Unix 时间戳" />

	<div class="w-full text-token card p-4">
		<Input
			bind:value={input}
			on:input={calculate}
			on:clear={calculate}
			title="时间"
			placeholder="输入任意时间格式"
			code={true}
			{warning}
		/>
	</div>

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
		font-size: 1.2rem;
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
