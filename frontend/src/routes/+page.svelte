<script lang="ts">
	import { ToolItem } from '$lib';
	import { store, activeToolItemHref, title } from '$lib/store/store';

	function setTitleAndActiveToolItem(t: string, href: string) {
		title.set(t);
		activeToolItemHref.set(href);
	}
</script>

<div class="w-full text-token px-6 py-3">
	<div class="flex flex-wrap">
		{#each store.toolItems as toolItem}
			<ToolItem
				on:click={() => {
					setTitleAndActiveToolItem(toolItem.title, toolItem.href);
				}}
				icon={toolItem.icon}
				title={toolItem.title}
				description={toolItem.description}
				href={toolItem.href}
			/>
		{/each}
	</div>
</div>
