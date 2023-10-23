<script lang="ts">
	import { CodeEditor, CodeViewer, RadioGroupItem, Warning } from '$lib';

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
			const res = await window.fetch(
				url + '?' + (query ? new URLSearchParams(JSON.parse(query)) : ''),
				{
					method: method,
					headers: header ? JSON.parse(header) : {},
					body: method != 'GET' ? body : undefined
				}
			);

			resStatus = res.status;
			resHeader = JSON.stringify(res.headers, null, 2);
			resBody = await res.text();

			console.log(resStatus);
			console.log(resHeader);
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
