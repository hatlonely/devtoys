<script lang="ts">
	import { Icon } from '$lib';

	export let icon: string;
	export let title: string;
	export let description: string;
</script>

<div class="card w-40 p-8">
	<span><span><Icon name={icon} class="icon-lg" /></span></span>

	<span>{title}</span>
	<span>{description}</span>
</div>
