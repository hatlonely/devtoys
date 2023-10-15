<script lang="ts">
	import 'material-symbols';
	import { AnalystTimeInfo } from '$lib/wailsjs/go/devtoys/App';
	import '@fontsource/roboto-mono';
	import { Title, Input, DataTable } from '$lib';

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

	const labels = {
		Timestamp: '当前时间戳',
		TimestampMill: '当前时间戳（毫秒）',
		LocalTime: '本地时间',
		UTCTime: 'UTC 时间',
		Relative: '相对时间'
	};

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

	<div class="w-full text-token">
		<DataTable
			fixed={true}
			keys={['label', 'value']}
			copyable={{ value: true }}
			alignRight={{ label: true }}
			rows={['Timestamp', 'TimestampMill', 'LocalTime', 'UTCTime', 'Relative'].map((v) => {
				return {
					label: labels[v],
					value: info[v]
				};
			})}
		/>
	</div>
</div>
