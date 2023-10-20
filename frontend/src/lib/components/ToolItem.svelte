<script lang="ts">
	import { Icon } from '$lib';

	export let icon: string;
	export let title: string;
	export let description: string;
	export let href: string;
</script>

<div class="w-48 h-72 card m-4 p-8 content-center">
	<a {href}>
		<div class="flex justify-center"><Icon name={icon} class="icon-xl" /></div>

		<div class="mt-8">
			<span>
				<dt class="font-bold text-center">{title}</dt>
				<dd class="text-sm opacity-50 text-center mt-3">{description}</dd>
			</span>
		</div>
	</a>
</div>
