<script lang="ts">
	import { CodeEditor, CodeViewer, RadioGroupItem, Warning } from '$lib';
	import { DoHttp } from '$lib/wailsjs/go/devtoys/App';

	let url = '';
	let method = 'GET';
	let header = '';
	let query = '';
	let body = '';

	let headerWarning = '';
	let queryWarning = '';

	let resStatus = 0;
	let resHeader = '';
	let resBody = '';

	let warning = '';

	async function fetch() {
		console.log('fetch');
		try {
			const res = await DoHttp({
				Method: method,
				URL: url,
				Query: query ? JSON.parse(query) : undefined,
				Headers: header ? JSON.parse(header) : undefined,
				Body: body
			});

			resStatus = res.Status;
			resHeader = JSON.stringify(res.Headers, null, 2);
			resBody = res.Body;
		} catch (err) {
			console.log(err);
		}
	}

	const extraButton = [
		{
			icon: 'submit',
			iconSet: 'custom',
			text: '发送',
			onClick: () => {
				fetch();
			}
		}
	];
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<div class="w-full text-token card p-4">
		<RadioGroupItem
			bind:group={method}
			name="method"
			icon="sync_alt"
			title="HTTP 方法"
			description="HTTP 请求方法"
			items={[
				{
					label: 'GET',
					value: 'GET'
				},
				{
					label: 'POST',
					value: 'POST'
				},
				{
					label: 'PUT',
					value: 'PUT'
				},
				{
					label: 'DELETE',
					value: 'DELETE'
				},
				{
					label: 'PATCH',
					value: 'PATCH'
				},
				{
					label: 'HEAD',
					value: 'HEAD'
				},
				{
					label: 'OPTIONS',
					value: 'OPTIONS'
				}
			]}
		/>

		<div class="w-full text-token card p-4">
			<CodeEditor title="Request URL" bind:value={url} warning={headerWarning} {extraButton} />
		</div>
		<div class="w-full text-token card p-4">
			<CodeEditor title="Request Query" bind:value={query} warning={headerWarning} {extraButton} />
		</div>
		<div class="w-full text-token card p-4">
			<CodeEditor title="Request Header" bind:value={header} warning={queryWarning} {extraButton} />
		</div>
		<div class="w-full text-token card p-4">
			<CodeEditor title="Request Body" bind:value={body} warning={headerWarning} {extraButton} />
		</div>
	</div>

	{#if resStatus != 0}
		<div class="w-full text-token card p-4">
			{resStatus}

			<div class="w-full text-token card p-4">
				<CodeViewer title="Response Header" value={resHeader} />
			</div>
			<div class="w-full text-token card p-4">
				<CodeViewer title="Response Body" value={resBody} />
			</div>

			{#if warning}
				<Warning message={warning} />
			{/if}
		</div>
	{/if}
</div>
