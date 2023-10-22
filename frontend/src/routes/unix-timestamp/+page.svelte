<script lang="ts">
	import 'material-symbols';
	import '@fontsource/roboto-mono';
	import { UnixTimestamp } from '$lib/wailsjs/go/devtoys/App';
	import { Title, TextInputItem, InformationWall } from '$lib';

	let text = '';
	let warning: any = '';
	let info = {};

	async function calculate() {
		warning = '';
		try {
			let res = await UnixTimestamp({
				Time: text
			});
			info = res;
		} catch (err) {
			warning = err;
		}
	}

	const labelValues = [
		{ value: 'Timestamp', label: '当前时间戳' },
		{ value: 'TimestampMill', label: '当前时间戳（毫秒）' },
		{ value: 'Relative', label: '相对时间' },
		{ value: 'LocalTime', label: '本地时间' },
		{ value: 'UTCTime', label: 'UTC 时间' },
		{ value: 'ANSIC', label: 'ANSIC' },
		{ value: 'UnixDate', label: 'UnixDate' },
		{ value: 'RubyDate', label: 'RubyDate' },
		{ value: 'RFC822', label: 'RFC822' },
		{ value: 'RFC822Z', label: 'RFC822Z' },
		{ value: 'RFC850', label: 'RFC850' },
		{ value: 'RFC1123', label: 'RFC1123' },
		{ value: 'RFC1123Z', label: 'RFC1123Z' },
		{ value: 'RFC3339', label: 'RFC3339' },
		{ value: 'RFC3339Nano', label: 'RFC3339Nano' },
		{ value: 'Kitchen', label: 'Kitchen' },
		{ value: 'Stamp', label: 'Stamp' }
	];

	$: text, calculate();
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<!-- <Title title="Unix 时间戳" /> -->

	<div class="w-full text-token card p-4">
		<TextInputItem
			bind:value={text}
			on:input={calculate}
			on:clear={calculate}
			title="时间"
			placeholder="输入任意时间格式"
			code={true}
			{warning}
		/>
	</div>

	<div class="w-full card p-4 text-token">
		<InformationWall
			labelValues={labelValues.map((v) => {
				return {
					label: v.label,
					value: info[v.value]
				};
			})}
		/>
	</div>
</div>
