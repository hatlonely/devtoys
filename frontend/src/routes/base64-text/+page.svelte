<script lang="ts">
	import { AppShell } from '@skeletonlabs/skeleton';
	import 'material-symbols';
	import { RadioGroup, RadioItem } from '@skeletonlabs/skeleton';

	let op = 'encode';
	let mode = 'std';
	let align = 'true';

	const op_desc_map: Map<string, string> = new Map([
		['encode', '编码'],
		['decode', '解码']
	]);
</script>

<AppShell>
	<svelte:fragment>
		<div class="card p-4 w-full text-token space-y-4">
			<h2 class="h2 my-4">Base64 文本编解/解码工具</h2>

			<div class="w-full text-token card p-4 space-y-4">
				<div class="w-full text-token p-2 md:flex">
					<span class="badge-icon p-4 variant-soft-secondary mt-2">
						<span class="material-symbols-outlined">sync_alt</span>
					</span>
					<span class="flex-auto ml-5">
						<dt class="font-bold">转换模式</dt>
						<dd class="text-sm opacity-50">选择您想使用的转换模式</dd>
					</span>
					<span>
						<RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
							<RadioItem bind:group={op} name="op" value="encode">编码</RadioItem>
							<RadioItem bind:group={op} name="op" value="decode">解码</RadioItem>
						</RadioGroup>
					</span>
				</div>

				<div class="w-full text-token p-2 md:flex">
					<span class="badge-icon p-4 variant-soft-secondary mt-2">
						<span class="material-symbols-outlined">link</span>
					</span>
					<span class="flex-auto ml-5">
						<dt class="font-bold">编码模式</dt>
						<dd class="text-sm opacity-50">常规编码/URL安全编码</dd>
					</span>
					<span>
						<RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
							<RadioItem bind:group={mode} name="mode" value="std">常规</RadioItem>
							<RadioItem bind:group={mode} name="mode" value="url">URL</RadioItem>
						</RadioGroup>
					</span>
				</div>

				<div class="w-full text-token p-2 md:flex">
					<span class="badge-icon p-4 variant-soft-secondary mt-2">
						<span class="material-symbols-outlined">equal</span>
					</span>
					<span class="flex-auto ml-5">
						<dt class="font-bold">补齐模式</dt>
						<dd class="text-sm opacity-50">末尾是否使用 = 补齐</dd>
					</span>
					<span>
						<RadioGroup active="variant-filled-primary" hover="hover:variant-soft-primary">
							<RadioItem bind:group={align} name="align" value="true">开启</RadioItem>
							<RadioItem bind:group={align} name="align" value="false">关闭</RadioItem>
						</RadioGroup>
					</span>
				</div>
			</div>

			<label class="label">
				<textarea
					class="textarea"
					rows="4"
					placeholder="Lorem ipsum dolor sit amet consectetur adipisicing elit."
				/>
			</label>

			<label class="label">
				<textarea
					class="textarea"
					rows="4"
					placeholder="Lorem ipsum dolor sit amet consectetur adipisicing elit."
				/>
			</label>
		</div>
	</svelte:fragment>
</AppShell>
