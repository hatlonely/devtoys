<script lang="ts">
	import { Icon, ToolItem, Title } from '$lib';

	let toolItems = [
		{
			icon: 'home',
			title: 'Base64 编解码工具',
			description: 'Base64 编解码工具'
		},
		{
			icon: 'home',
			title: 'Unix 时间戳工具',
			description: 'Unix 时间戳工具'
		},
		{
			icon: 'home',
			title: '转义工具',
			description: '转义工具'
		},
		{
			icon: 'home',
			title: '文本二进制转换工具',
			description: '文本二进制转换工具'
		}
	];
</script>

<div class="w-full text-token px-6 py-3">
	<Title title="所有工具" />

	<div class="flex flex-row space-x-10">
		{#each toolItems as toolItem}
			<ToolItem icon={toolItem.icon} title={toolItem.title} description={toolItem.description} />
		{/each}
	</div>
</div>
