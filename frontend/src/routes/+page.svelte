<script lang="ts">
	import { ToolItem, Title } from '$lib';
	import { store } from '$lib/store/store';
</script>

<div class="w-full text-token px-6 py-3">
	<!-- <Title title="所有工具" /> -->

	<div class="flex flex-wrap">
		{#each store.toolItems as toolItem}
			<ToolItem
				icon={toolItem.icon}
				title={toolItem.title}
				description={toolItem.description}
				href={toolItem.href}
			/>
		{/each}
	</div>
</div>
