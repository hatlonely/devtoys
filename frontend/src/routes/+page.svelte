<script lang="ts">
	import { Icon, ToolItem, Title } from '$lib';

	let toolItems = [
		{
			icon: 'base64',
			title: 'Base64 编解码',
			description: 'Base64 编解码工具',
			href: '/base64-text'
		},
		{
			icon: 'timestamp',
			title: 'Unix 时间戳',
			description: 'Unix 时间戳工具',
			href: '/unix-timestamp'
		},
		{
			icon: 'escape',
			title: '字符串转义',
			description: '字符串、HTML、URL、URL Query 转义工具',
			href: '/escape'
		},
		{
			icon: 'text',
			title: '文本二进制转换',
			description: '查看文本的二进制、八进制、十六进制标识',
			href: '/text-binary-conversion'
		},
		{
			icon: 'number',
			title: '数值进制转换',
			description: '数值二进制、八进制、十进制、十六进制转换',
			href: '/number-base-conversion'
		},
		{
			icon: 'hash',
			title: '哈希散列',
			description: 'MD5、SHA1、SHA256、SHA512、SHA3、RIPEMD160、HMAC 哈希散列工具',
			href: '/hash'
		},
		{
			icon: 'string',
			title: '字符串转换',
			description: '字符串大小写转换、驼峰转换、下划线转换、中划线转换',
			href: '/string-conversion'
		}
	];
</script>

<div class="w-full text-token px-6 py-3">
	<Title title="所有工具" />

	<div class="flex flex-wrap">
		{#each toolItems as toolItem}
			<ToolItem
				icon={toolItem.icon}
				title={toolItem.title}
				description={toolItem.description}
				href={toolItem.href}
			/>
		{/each}
	</div>
</div>
